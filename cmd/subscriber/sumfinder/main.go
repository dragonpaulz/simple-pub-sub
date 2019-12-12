package main

import (
	"fmt"
	"log"
	"time"

	"simple-pub-sub/cmd/internal/config"
	"simple-pub-sub/cmd/subscriber/internal/receive"
)

func main() {
	psc, bc, cErr := config.RedisSubConn()
	if cErr != nil {
		log.Printf("Error connecting to redis, %v\n", cErr)
	}

	defer psc.Conn.Close()

	done := make(chan error, 1)
	newNum := make(chan int, 10)

	var sum int64 = 0

	go receive.Receive(done, newNum, psc)

	ticker := time.NewTicker(time.Duration(bc.SumWindowSeconds) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case n := <-newNum:
			fmt.Println("Received: ", n)
			sum += int64(n)
		case t := <-ticker.C:
			fmt.Println("Sum: ", sum, " at time ", t)
			sum = 0
		}
	}
}
