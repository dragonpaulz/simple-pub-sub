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
	Redis RedisConfig
	Queue QueueConfig
}

// RedisConfig contains the information to connect to redis
type RedisConfig struct {
	Port, Host string
}

// QueueConfig contains the information once connected to redis
type QueueConfig struct {
	Channel string
}

// ReadConfig takes the JSON Configuration file that specifies how to connect to the channel
func ReadConfig() PubSubConfig {
	data, readErr := ioutil.ReadFile("/home/paul/go/src/simple-pub-sub/config.json")

	if readErr != nil {
		log.Printf("Error reading file: %v", readErr)
	}

	log.Println(data)

	var config PubSubConfig
	err := json.Unmarshal(data, &config)
	if err != nil {
		log.Println("Error unmarshalling configuration file: ", err)
	}
	return config
}

// RedisConnection opens a connection to a redis for PubSub
func (psc PubSubConfig) RedisConnection() (redis.Conn, redis.PubSubConn, error) {
	redisAddress := fmt.Sprintf("%s:%s", psc.Redis.Host, psc.Redis.Port)
	conn, dErr := redis.Dial("tcp", redisAddress)
	if dErr != nil {
		log.Printf("Error while dialing: %v\n", dErr)
		return nil, redis.PubSubConn{}, dErr
	}

	rconn := redis.PubSubConn{Conn: conn}
	if sErr := rconn.Subscribe(psc.Queue.Channel); sErr != nil {
		log.Printf("Cannot subscribe to %v, receiving error: %v",
			psc.Queue.Channel,
			sErr,
		)

		return nil, redis.PubSubConn{}, sErr
	}

	return conn, rconn, nil
}
