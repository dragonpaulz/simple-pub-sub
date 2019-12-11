package main

import (
	"log"

	"simple-pub-sub/cmd/internal/config"
	"simple-pub-sub/cmd/subscriber/internal/receive"
)

func main() {
	psc, cErr := config.RedisSubConn()
	if cErr != nil {
		log.Printf("Error connecting to redis, %v\n", cErr)
	}

	defer psc.Conn.Close()

	done := make(chan error, 1)

	receive.Receive(done, psc)

}
