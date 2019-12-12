package receive

import (
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

// Receive will read from subscription queue
func Receive(
	done chan error,
	newNum chan int,
	rconn redis.PubSubConn,
) {
	for {
		switch n := rconn.Receive().(type) {
		case error:
			done <- n
			log.Println("Received Error: ", n.(error))
			return
		case redis.Subscription:
			log.Printf("Successfully subscribed to %s\n", n.Channel)
		case redis.Message:
			if val, err := strconv.Atoi(string(n.Data)); err == nil {
				newNum <- val
			} else {
				log.Printf("Non integer value received, %v", n.Data)
			}
		}
	}
}
