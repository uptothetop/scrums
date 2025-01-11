package main

import (
	"fmt"
	"net/http"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "User Service is active")
}

func main() {
	http.HandleFunc("/", defaultHandler)

	fmt.Println("User Service is running on port 8080")
	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		fmt.Println("Error starting server", err)
		return
	}
}
