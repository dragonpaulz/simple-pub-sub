package medianfinder

import "context"
import "log"
import "github.com/gomodule/redigo/redis"

var channel := "arcticwolf"

func main() {
	var redisAddress := "localhost:6060"

	ctx, cancel := context.WithCancel(ctx.Background())

	conn, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		log.Printf("Error while dialing: %v\n", err)
		return
	}

	defer conn.Close()
	rconn := redis.PubSubConn{Conn: conn}
	rconn.Subscribe(channel)
}
