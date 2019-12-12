package receive

import (
	"log"
	"strconv"

	"github.com/gomodule/redigo/redis"
)

// Receiver is an attempt to make the code testable.
type Receiver interface {
	Receive() interface{}
}

// Receive will read from subscription queue
func Receive(
	done chan error,
	newNum chan int64,
	rconn Receiver,
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
			val, _ := strconv.Atoi(string(n.Data))
			newNum <- int64(val)
		}
	}
}
