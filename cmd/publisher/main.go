package publisher

import (
	"fmt"
	_ "log"

	_ "github.com/gomodule/redigo/redis"
)

func main() {
	fmt.Printf("%v\n", 0)
	// redisAddress := "localhost:6060"

	// ctx, cancel := context.WithCancel(context.Background())

	// conn, err := redis.Dial("tcp", redisAddress)
	// if err != nil {
	// 	log.Printf("Error while dialing: %v\n", err)
	// 	return
	// }

	// defer conn.Close()
}
