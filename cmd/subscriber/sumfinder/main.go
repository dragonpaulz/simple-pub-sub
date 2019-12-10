package sumfinder

import (
	_ "context"
	_ "log"

	_ "github.com/gomodule/redigo/redis"
)

func main() {
	// channel := "arcticwolf"
	// redisAddress := "localhost:6060"

	// ctx, cancel := context.WithCancel(context.Background())

	// conn, err := redis.Dial("tcp", redisAddress)
	// if err != nil {
	// 	log.Printf("Error while dialing: %v\n", err)
	// 	return
	// }

	// defer conn.Close()
	// rconn := redis.PubSubConn{Conn: conn}
	// rconn.Subscribe(channel)
}
