package main

import (
	"fmt"
	"log"

	"simple-pub-sub/cmd/internal/config"

	"github.com/gomodule/redigo/redis"
)

func main() {
	channel := "test"
	// redisAddress := "localhost:6060"
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

	for {
		switch n := rconn.Receive().(type) {
		case error:
			done <- n
			return
		case redis.Message:
			fmt.Printf("%v\n", string(n.Data))
			// if err := onMessage(n.Channel, n.Data); err != nil {
			// 	done <- err
			// 	return
			// }
			// case redis.Subscription:
			// 	switch n.Count {
			// 	case len(channels):
			// 		// Notify application when all channels are subscribed.
			// 		if err := onStart(); err != nil {
			// 			done <- err
			// 			return
			// 		}
			// 	case 0:
			// 		// Return from the goroutine when all channels are unsubscribed.
			// 		done <- nil
			// 		return
			// 	}
		}
	}
}
