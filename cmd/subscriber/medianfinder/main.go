package main

import (
	"fmt"
	"log"
	"os"
	"simple-pub-sub/cmd/internal/config"
	"simple-pub-sub/cmd/subscriber/internal/receive"
	"simple-pub-sub/cmd/subscriber/medianfinder/median"
	"time"
)

func main() {
	psc, bc, cErr := config.RedisSubConn(os.Args[1])
	if cErr != nil {
		log.Printf("Error connecting to redis, %v\n", cErr)
	}

	defer psc.Conn.Close()

	done := make(chan error, 1)
	newNum := make(chan int, 10)
	waiting := make([]int, 0)

	go receive.Receive(done, newNum, psc)

	ticker := time.NewTicker(time.Duration(bc.SumWindowSeconds) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case n := <-newNum:
			fmt.Println("Received: ", n)
			waiting = append(waiting, n)
		case t := <-ticker.C:
			received := make([]int, len(waiting))
			copy(received, waiting)
			waiting = make([]int, 0)
			m := median.Find(received)
			fmt.Println("Median ", m, " at time ", t)
		}
	}
}
