package main

import (
	"fmt"

	"github.com/garyburd/redigo/redis"
)

const (
	// TODO, read from env
	redisAddress = "localhost:6379"
	// Name of redis channel for hits subscription
	redisChannel = "hits"
)

func main() {
	connection := connect()
	listen(connection)
}

func connect() redis.Conn {
	// TODO:
	// - handle connection errors
	// - add authentication handling
	connection, _ := redis.Dial("tcp", redisAddress)
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
