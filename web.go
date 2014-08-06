package main

import (
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

var templates = template.Must(template.ParseFiles("web/home.html"))
var assets = template.Must(template.ParseFiles("web/assets/home.js"))

var connections = make(map[*websocket.Conn]bool)

func serve() {
	port := os.Getenv("PORT")

	http.HandleFunc("/assets/", assetHandler)
	http.HandleFunc("/websocket", webSocketHandler)
	http.HandleFunc("/", webHandler)

	go receiveHits()

	log.Print("Listening to serve on port " + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panic(err)
	}
}

func receiveHits() {
	log.Print("Receiving hits...")

	for {
		hit := <-hits

		log.Print("Received hit, sending to connections...")

		// Send hit to socket connections
		for conn := range connections {
			conn.WriteJSON(hit)
		}
	}
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "home.html", r.Host)
}

func webSocketHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Panic(err)
	}

	// Add connection to connections map
	connections[conn] = true

	// TODO, close sockets and remove from connections map
	// conn.Close()
}

func assetHandler(w http.ResponseWriter, r *http.Request) {
	assets.ExecuteTemplate(w, "home.js", r.Host)
}
