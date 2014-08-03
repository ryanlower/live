package main

import (
	"encoding/json"
	"log"
	"os"

	"github.com/garyburd/redigo/redis"
)

const (
	// Name of redis channel for hits subscription
	redisChannel = "hits"
)

type Hit struct {
	Host string
	Path string
	Code string
}

func main() {
	connection := connect()
	go listen(connection)

	serve()
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
	log.Println("Listening to " + redisChannel + " channel...")

	for {
		reply, message := pubSubConnection.Receive().(redis.Message)
		if message {
			hit := parseMessage(reply)
			log.Print("[" + hit.Code + "] " + hit.Host + hit.Path)
		}
	}
}

func parseMessage(hitMessage redis.Message) Hit {
	var hit Hit
	json.Unmarshal(hitMessage.Data, &hit)

	return hit
}
