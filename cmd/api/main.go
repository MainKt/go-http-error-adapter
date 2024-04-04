package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("GET /health_check", func(w http.ResponseWriter, r *http.Request) {})
	log.Println("listening on localhost:8080")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
