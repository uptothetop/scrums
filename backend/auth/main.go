package main

import (
	"log"
	"net/http"
	"os"
	"scrums/m/v2/handlers"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register the routes.
	r.HandleFunc("/login", handlers.LoginHandler).Methods("POST")

	port := os.Getenv("PORT")

	log.Printf("Starting Auth service on port :%s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start auth-service: %v", err)
	}
}
