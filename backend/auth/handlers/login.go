// Login handler for the auth service
package handlers

import (
	"encoding/json"
	"net/http"
	"time"

	"scrums/m/v2/models"

	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// Set JWT time-to-live
const JWT_TTL = 15 * time.Minute

// Creds contract
type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// JWT Claims
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// Handles the new user register endpoint of the auth service.
func RegisterHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var creds Credentials

		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, "Invalid Request", http.StatusBadRequest)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(creds.Password), bcrypt.DefaultCost)
		if err != nil {
			http.Error(w, "Internal Server Error", http.StatusInternalServerError)
		}

		user := models.User{
			Username:     creds.Username,
			PasswordHash: string(hashedPassword),
		}
		if err := db.Create(&user).Error; err != nil {
			http.Error(w, "Error creating user", http.StatusInternalServerError)
			return
		}
		w.Write([]byte("User registered successfully"))
	}
}

// Handles the user login endpoint of the auth service.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// Handles the JWT-refresh endpoint of the auth service.
func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// Handles the JWT Verification endpoint of the auth service.
func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}
