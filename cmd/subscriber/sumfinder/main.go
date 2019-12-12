package main

import (
	"fmt"
	"log"
	"strconv"
	"time"

	"simple-pub-sub/cmd/internal/config"
	"simple-pub-sub/cmd/subscriber/internal/receive"

	"github.com/gomodule/redigo/redis"
)

func main() {
	psc, bc, cErr := config.RedisSubConn()
	if cErr != nil {
		log.Printf("Error connecting to redis, %v\n", cErr)
	}

	defer psc.Conn.Close()

	done := make(chan error, 1)
	newNum := make(chan int64, 1)

	var sum int64 = 0
	// var mux sync.Mutex

	// onReceiveMsg := func(m redis.Message, sum *int) {
	// 	val, _ := strconv.Atoi(string(m.Data))
	// 	sum += int64(val)
	// }()

	receive.Receive(done, newNum, psc, onReceiveMsg)

	ticker := time.NewTicker(time.Duration(bc.SumWindowSeconds) * time.Second)
	defer ticker.Stop()

	for {
		select {
		case n := <-newNum:
			sum += n
		case t := <-ticker.C:
			fmt.Println("Sum: ", sum, " at time ", t)
			sum = 0
		}
	}
}

func onReceiveMsg(m redis.Message, num chan int64) {
	val, _ := strconv.Atoi(string(m.Data))
	num <- int64(val)
}
