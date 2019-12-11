package main

import (
	"context"
	"fmt"
	"log"
	"math/rand"

	"simple-pub-sub/cmd/internal/config"

	"github.com/gomodule/redigo/redis"
)

func main() {
	config := config.ReadConfig()
	redisAddress := fmt.Sprintf("%s:%s", config.Redis.Host, config.Redis.Port)

	ctx, cancel := context.WithCancel(context.Background())

	log.Printf("context: %v, cancel: %v", ctx, cancel)

	conn, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		log.Printf("Error while dialing: %v\n", err)
		return
	}

	defer conn.Close()

	// psc := redis.PubSubConn{Conn: conn}

	// done := make(chan error, 1)

	num := rand.Int31()

	// conn.Do("PUBLISH", config.Queue.Channel, num)
	conn.Do("PUBLISH", "test", num)
	log.Printf("Wrote :%v to channel\n", num)
}
