package handlers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"scrums/auth/m/v2/models"
	"testing"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func setupTestDB() *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(&models.User{})

	// Insert a new record
	testUser := "testuser"
	testPass := "testpass"
	password, _ := bcrypt.GenerateFromPassword([]byte(testPass), bcrypt.DefaultCost)
	db.Create(&models.User{Username: testUser, PasswordHash: string(password)})

	return db
}

func TestHandlers(t *testing.T) {
	db := setupTestDB()

	t.Run("RegisterHandler", func(t *testing.T) {
		t.Run("Success Registration", func(t *testing.T) {
			requestBody, _ := json.Marshal(map[string]string{
				"username": "newuser",
				"password": "newpass",
			})

			req, err := http.NewRequest(
				"POST",
				"/api/v1/register",
				bytes.NewBuffer(requestBody),
			)

			if err != nil {
				t.Fatalf("Request creation error: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(RegisterHandler(db))
			handler.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
			}
		})

		t.Run("Empty User", func(t *testing.T) {
			requestBody, _ := json.Marshal(map[string]string{
				"username": "",
				"password": "newpass",
			})

			req, err := http.NewRequest(
				"POST",
				"/api/v1/register",
				bytes.NewBuffer(requestBody),
			)

			if err != nil {
				t.Fatalf("Request creation error: %v", err)
			}

			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(RegisterHandler(db))
			handler.ServeHTTP(rr, req)

			if rr.Code != http.StatusOK {
				t.Errorf("Expected status %d, got %d", http.StatusOK, rr.Code)
			}
		})
	})

	t.Run("LoginHandler", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			requestBody, _ := json.Marshal(map[string]string{
				"username": "testuser",
				"password": "testpass",
			})

			// Create tester request
			req := httptest.NewRequest(
				"POST",
				"/api/v1/login",
				bytes.NewBuffer(requestBody),
			)

			// Make Login call
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(LoginHandler(db))
			handler.ServeHTTP(rr, req)

			// Assert the login is success
			if rr.Code != http.StatusOK {
				t.Errorf("Expected %d, got %d", http.StatusOK, rr.Code)
			}

			// Assert we got token
			var response map[string]string
			json.NewDecoder(rr.Body).Decode(&response)

			if _, exists := response["token"]; !exists {
				t.Error("Expected token in response, but got none")
			}
		})

		t.Run("Failure", func(t *testing.T) {
			requestBody, _ := json.Marshal(map[string]string{
				"username": "wronguser",
				"password": "wrongpass",
			})

			// Create tester request
			req := httptest.NewRequest(
				"POST",
				"/api/v1/login",
				bytes.NewBuffer(requestBody),
			)

			// Make Login call
			rr := httptest.NewRecorder()
			handler := http.HandlerFunc(LoginHandler(db))
			handler.ServeHTTP(rr, req)

			// Assert the login is success
			if rr.Code != http.StatusUnauthorized {
				t.Errorf("Expected %d, got %d", http.StatusUnauthorized, rr.Code)
			}

			// Assert we got token
			var response map[string]string
			json.NewDecoder(rr.Body).Decode(&response)

			if _, exists := response["token"]; exists {
				t.Error("Expected no token in response, but got it")
			}
		})
	})
}
