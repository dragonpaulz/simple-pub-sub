package receive

import (
	"fmt"

	"github.com/gomodule/redigo/redis"
)

// Receive will read from subscription queue
func Receive(done chan error, rconn redis.PubSubConn) {

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
