package main

import (
	"errors"
	"go-http-error-adapter/internal/httpadapter"
	"log"
	"net/http"
)

func main() {
	router := httpadapter.NewServeMuxErr(func(hfe httpadapter.HandlerFuncErr) http.HandlerFunc {
		return func(w http.ResponseWriter, r *http.Request) {
			if err := hfe(w, r); err != nil {
				log.Printf("Encountered err: %s", err)
				w.WriteHeader(http.StatusInternalServerError)
			}
		}
	})

	router.HandleFunc("GET /health_check", func(w http.ResponseWriter, r *http.Request) {})

	router.HandleFuncErr("GET /error_check", func(w http.ResponseWriter, r *http.Request) error {
		return errors.New("Error check")
	})

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	log.Println("listening on localhost:8080")
	if err := server.ListenAndServe(); err != nil {
		log.Fatalf("cannot start server: %s", err)
	}
}
