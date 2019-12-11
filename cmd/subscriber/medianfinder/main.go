package main

import (
	"fmt"
	"log"
	"simple-pub-sub/cmd/internal/config"
	"simple-pub-sub/cmd/subscriber/internal/receive"

	"github.com/gomodule/redigo/redis"
)

func main() {
	channel := "test"
	config := config.ReadConfig()
	redisAddress := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)

	// ctx, cancel := context.WithCancel(context.Background())

	conn, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		log.Printf("Error while dialing: %v\n", err)
		return
	}

	defer conn.Close()
	rconn := redis.PubSubConn{Conn: conn}
	if err := rconn.Subscribe(channel); err != nil {
		log.Fatalf("Cannot subscribe to %v, receiving error: %v",
			channel,
			err,
		)
	}

	done := make(chan error, 1)

	receive.Receive(done, rconn)
}
