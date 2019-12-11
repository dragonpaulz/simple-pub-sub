package main

import (
	"log"
	"simple-pub-sub/cmd/internal/config"
	"simple-pub-sub/cmd/subscriber/internal/receive"
)

func main() {
	config := config.ReadConfig()

	_, psc, cErr := config.RedisConnection()
	if cErr != nil {
		log.Printf("Error connecting to redis, %v\n", cErr)
	}

	done := make(chan error, 1)

	receive.Receive(done, psc)
}
