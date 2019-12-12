package main

import (
	"log"
	"simple-pub-sub/cmd/internal/config"
	"simple-pub-sub/cmd/subscriber/internal/receive"
	"simple-pub-sub/cmd/subscriber/medianfinder/median"
)

func main() {
	psc, bc, cErr := config.RedisSubConn()
	if cErr != nil {
		log.Printf("Error connecting to redis, %v\n", cErr)
	}

	defer psc.Conn.Close()

	done := make(chan error, 1)
	newNum := make(chan int, 10)

	var waiting := make([]int,0)

	go receive.Receive(done, psc)

	ticker := time.NewTicker(time.Duration(bc.SumWindowSeconds) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case n := <-newNum:
			fmt.Println("Received: ", n)
			waiting = append(waiting, n)
		case t := <-ticker.C:
			fmt.Println("Sum: ", sum, " at time ", t)
			received := make([]int, len(waiting))
			copy(received, waiting)
			waiting = make([]int,0)
			median.Find(received)
		}
	}
}
