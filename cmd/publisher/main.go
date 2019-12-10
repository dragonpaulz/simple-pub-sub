package publisher

import (
	_ "log"

	_ "github.com/gomodule/redigo/redis"
)

func main() {
	// redisAddress := "localhost:6060"

	// ctx, cancel := context.WithCancel(context.Background())

	// conn, err := redis.Dial("tcp", redisAddress)
	// if err != nil {
	// 	log.Printf("Error while dialing: %v\n", err)
	// 	return
	// }

	// defer conn.Close()
}
