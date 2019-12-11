package main

import (
	"context"
	"log"
	"math/rand"
	"time"

	"simple-pub-sub/cmd/internal/config"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	log.Printf("context: %v, cancel: %v", ctx, cancel)

	conn, channel, connErr := config.RedisPubConn()
	if connErr != nil {
		log.Fatalf("Cannot establish connection. Exiting.")
	}

	defer conn.Close()

	for {
		num := rand.Int31()

		conn.Do("PUBLISH", channel, num)
		log.Printf("Wrote :%v to channel\n", num)
		time.Sleep(time.Second)
	}
}
