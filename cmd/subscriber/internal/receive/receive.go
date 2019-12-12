package receive

import (
	"fmt"
	"log"

	"github.com/gomodule/redigo/redis"
)

// Receive will read from subscription queue
func Receive(
	done chan error,
	newInt chan int64,
	rconn redis.PubSubConn,
	onMessage func(redis.Message, chan int64),
) {
	for {
		switch n := rconn.Receive().(type) {
		case error:
			done <- n
			log.Println("Received Error: ", n.(error))
			return
		case redis.Subscription:
			fmt.Printf("Successfully subscribed to %s\n", n.Channel)
		case redis.Message:
			go onMessage(n, newInt)
		}
	}
}
