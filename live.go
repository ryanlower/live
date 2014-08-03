package main

import (
	"fmt"
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

const (
	// Name of redis channel for hits subscription
	redisChannel = "hits"
)

func main() {
	connection := connect()
	listen(connection)
}

func connect() redis.Conn {
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	connection, err := redis.Dial("tcp", redisAddress)
	if err != nil {
		log.Panic("Eror connecting to redis...", err)
	}

	if len(redisPassword) > 0 {
		err := connection.Send("AUTH", redisPassword)
		if err != nil {
			log.Panic("Eror authenticating...", err)
		}
	}

	return connection
}

func listen(connection redis.Conn) {
	pubSubConnection := redis.PubSubConn{connection}

	pubSubConnection.Subscribe(redisChannel)
	fmt.Println("Listening to " + redisChannel + " channel...")

	for {
		reply, message := pubSubConnection.Receive().(redis.Message)
    if message {
      logHit(reply)
    }
	}
}

func logHit(hitMessage redis.Message)  {
  fmt.Println(hitMessage.Data)
}
