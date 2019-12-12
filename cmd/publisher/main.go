package main

import (
	"log"
	"math/rand"
	"time"

	"simple-pub-sub/cmd/internal/config"
)

func main() {
	conn, channel, conf, connErr := config.RedisPubConn()
	if connErr != nil {
		log.Fatalf("Cannot establish connection. Exiting.")
	}

	defer conn.Close()

	sleepTime := time.Second / time.Duration(conf.PerSecond)
	for {
		num := rand.Int31n(10)
		sign := rand.Int31n(2)

		if sign == 1 {
			num *= -1
		}

		conn.Do("PUBLISH", channel, num)
		log.Printf("Wrote :%v to channel\n", num)
		time.Sleep(sleepTime)
	}
}
