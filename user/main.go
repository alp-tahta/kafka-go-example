package main

import (
	"log"
	"net/http"
	"os"
	"time"
	"user/handlers"
	"user/messaging"
	"user/service"
)

func main() {
	broker := os.Getenv("KAFKA_BROKER")
	log.Println(broker)
	time.Sleep(10 * time.Second)
	log.Println("HIIIIIIIIIIIIII")
	port := ":8080"
	msgBrokes := []string{broker}

	msgClient, err := messaging.NewKafkaClient(msgBrokes, nil)
	if err != nil {
		log.Println(err)
	}
	service := service.NewService(msgClient)
	handler := handlers.NewHandler(service)

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
