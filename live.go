package main

import (
	"fmt"
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
	// TODO:
	// - handle connection / authentication errors
	redisAddress := os.Getenv("REDIS_ADDRESS")
	redisPassword := os.Getenv("REDIS_PASSWORD")

	connection, _ := redis.Dial("tcp", redisAddress)

	if len(redisPassword) > 0 {
		connection.Send("AUTH", redisPassword)
	}

	return connection
}

func listen(connection redis.Conn) {
	pubSubConnection := redis.PubSubConn{connection}

	pubSubConnection.Subscribe(redisChannel)
	fmt.Println("Listening to " + redisChannel + " channel...")

	for {
		reply := pubSubConnection.Receive()
		fmt.Println(reply)
	}
}
