package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Service is active")
}

func main() {
	http.HandleFunc("/", defaultHandler)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
		log.Printf("WARN: PORT environment variable not set. Using default port :%s. Please consider changing to a different one.", port)
	}

	log.Printf("Service is running on port :%s", port)
	if err := http.ListenAndServe(":"+port, nil); err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
}
