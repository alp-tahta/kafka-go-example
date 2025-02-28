package main

import (
	"log"
	"net/http"
	"time"
	"user/handlers"
)

func main() {
	port := ":8080"

	handler := handlers.NewHandler()

	router := http.NewServeMux()
	router.HandleFunc("GET /health", handler.HealthChecker)
	router.HandleFunc("POST /users", handler.CreateUser)

	s := &http.Server{
		Addr:           port,
		Handler:        router,
		ReadTimeout:    5 * time.Second,
		WriteTimeout:   5 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	log.Printf("Starting server on Port%s\n", port)
	log.Fatal(s.ListenAndServe())
}
