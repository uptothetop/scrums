package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"scrums/m/v2/handlers"
	"scrums/m/v2/models"

	"github.com/gorilla/mux"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Fetch ENV variables
	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbName := os.Getenv("DB_NAME")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	port := os.Getenv("PORT")

	// Connect to the DB
	dsn := fmt.Sprintf("user=%s password=%s dbname=%s host=%s port=%s sslmode=disable",
		dbUser, dbPassword, dbName, dbHost, dbPort)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	// Automate DB creation/migration
	db.AutoMigrate(&models.User{})

	r := mux.NewRouter()

	// Register the routes.
	r.HandleFunc("/login", handlers.LoginHandler(db)).Methods("POST")
	r.HandleFunc("/register", handlers.RegisterHandler(db)).Methods("POST")
	r.HandleFunc("/refresh", handlers.RefreshHandler).Methods("POST")
	r.HandleFunc("/verify", handlers.VerifyHandler).Methods("POST")

	log.Printf("Starting Auth service on port :%s...", port)
	if err := http.ListenAndServe(":"+port, r); err != nil {
		log.Fatalf("Failed to start auth-service: %v", err)
	}
}
