package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func serve() {
	port := os.Getenv("PORT")

	http.HandleFunc("/", webHandler)

	log.Print("Listening to serve on port " + port + "...")
	err := http.ListenAndServe(":"+port, nil)
	if err != nil {
		log.Panic(err)
	}
}

func webHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hits are coming")
}
