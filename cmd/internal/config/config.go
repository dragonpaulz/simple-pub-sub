package config

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/gomodule/redigo/redis"
)

// PubSubConfig describes the information for connecting to redis for pubsub
type PubSubConfig struct {
	Redis  RedisConfig
	Queue  QueueConfig
	Broker BrokerConfig
}

// RedisConfig contains the information to connect to redis
type RedisConfig struct {
	Port, Host string
}

// QueueConfig contains the information once connected to redis
type QueueConfig struct {
	Channel string
}

// BrokerConfig contains configurations for the publishers and subscribers
type BrokerConfig struct {
	PerSecond, SumWindowSeconds, MedianWindowSeconds float64
}

// ReadConfig takes the JSON Configuration file that specifies how to connect to the channel
func ReadConfig(configPath string) PubSubConfig {
	data, readErr := ioutil.ReadFile(configPath)

	if readErr != nil {
		log.Printf("Error reading file: %v", readErr)
	}

	var config PubSubConfig
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Println("Error unmarshalling configuration file: ", err)
	}
	return config
}

// RedisConnection opens a connection to a redis for PubSub
func (psc PubSubConfig) redisConnection() (redis.Conn, error) {
	redisAddress := fmt.Sprintf("%s:%s", psc.Redis.Host, psc.Redis.Port)
	conn, dErr := redis.Dial("tcp", redisAddress)
	if dErr != nil {
		log.Printf("Error while dialing: %v\n", dErr)
		return nil, dErr
	}

	return conn, nil
}

// RedisSubConn will return a connection for a subscriber of a channel
func RedisSubConn(configPath string) (redis.PubSubConn, BrokerConfig, error) {
	psc := ReadConfig(configPath)
	conn, dErr := psc.redisConnection()
	if dErr != nil {
		return redis.PubSubConn{}, BrokerConfig{}, dErr
	}

	rconn := redis.PubSubConn{Conn: conn}
	if sErr := rconn.Subscribe(psc.Queue.Channel); sErr != nil {
		log.Printf("Cannot subscribe to %v, receiving error: %v",
			psc.Queue.Channel,
			sErr,
		)

		return redis.PubSubConn{}, BrokerConfig{}, sErr
	}

	return rconn, psc.Broker, nil
}

// RedisPubConn will return a connection for a publisher of a channel
func RedisPubConn(configPath string) (redis.Conn, string, BrokerConfig, error) {
	psc := ReadConfig(configPath)
	conn, err := psc.redisConnection()
	return conn, psc.Queue.Channel, psc.Broker, err
}
