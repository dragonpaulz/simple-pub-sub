package config

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

type PubSubConfig struct {
	Redis RedisConfig
	Queue QueueConfig
}

type RedisConfig struct {
	Port, Host string
}

type QueueConfig struct {
	Channel string
}

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
