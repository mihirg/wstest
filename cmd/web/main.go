package main

import (
	"log"
	"net/http"
	"ws/internal/handlers"
)

// run using go run cmd/web/*.go

func main() {
	mux := routes()

	log.Println("Starting channel listener")
	go handlers.ListenToWsChannel()
	log.Println("Starting webserver on port 8080")
	_ = http.ListenAndServe(":8080", mux)

}
