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

	log.Print("Listening to serve on port " + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panic(err)
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

	for {
		var hit Hit
		hit = <-hits

		conn.WriteMessage(websocket.TextMessage, []byte(hit.Code))
	}

	conn.Close()
}
