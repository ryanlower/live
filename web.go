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

func serve() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", webHandler)
	http.HandleFunc("/websocket", webSocketHandler)

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
		// TODO, send to open sockets
		log.Print(hit)
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

	// TODO, remember open sockets so we can send hits to them
	conn.WriteMessage(websocket.TextMessage, []byte("hi there"))

	conn.Close()
}
