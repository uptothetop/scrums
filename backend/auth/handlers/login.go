// Login handler for the auth service
package handlers

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

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

// Handles the user login endpoint of the auth service.
func LoginHandler(w http.ResponseWriter, r *http.Request) {
	http.Error(w, "Not Implemented", http.StatusNotImplemented)
}

// Handles the new user register endpoint of the auth service.
func RegisterHandler(w http.ResponseWriter, r *http.Request) {
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
