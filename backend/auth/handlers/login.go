// Login handler for the auth service
package handlers

import (
	"encoding/json"
	"net/http"
	"os"
	"time"

	"scrums/auth/m/v2/models"
	"scrums/auth/m/v2/utils"

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
		response := map[string]string{"status": "ok"}
		var creds Credentials

		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			response["status"] = "Unauthorized"
			utils.SendJson(w, response, http.StatusUnauthorized)
			return
		}

		hashedPassword, err := bcrypt.GenerateFromPassword(
			[]byte(creds.Password), bcrypt.DefaultCost)
		if err != nil {
			response["status"] = "Unauthorized"
			utils.SendJson(w, response, http.StatusUnauthorized)
			return
		}

		user := models.User{
			Username:     creds.Username,
			PasswordHash: string(hashedPassword),
		}
		if err := db.Create(&user).Error; err != nil {
			response["status"] = "Unauthorized"
			utils.SendJson(w, response, http.StatusUnauthorized)
			return
		}

		utils.SendJson(w, response, http.StatusOK)
	}
}

// Handles the user login endpoint of the auth service.
func LoginHandler(db *gorm.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// Get user creds from the request body
		var creds Credentials
		err := json.NewDecoder(r.Body).Decode(&creds)
		if err != nil {
			http.Error(w, "Invalid request", http.StatusBadRequest)
			return
		}

		// Checking if the user exists
		var user models.User
		if err := db.Where("username = ?", creds.Username).First(&user).Error; err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Checking if the password is correct
		if err := bcrypt.CompareHashAndPassword(
			[]byte(user.PasswordHash),
			[]byte(creds.Password),
		); err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Generate JWT Token
		tokenString, err := generateJWT(user.Username)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		result := map[string]string{"token": tokenString}
		utils.SendJson(w, result, http.StatusOK)
	}
}

// Handles the JWT-refresh endpoint of the auth service.
func RefreshHandler(w http.ResponseWriter, r *http.Request) {
	currentToken := r.Header.Get("Authorization")[7:]
	token, err := jwt.ParseWithClaims(currentToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil || !token.Valid {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	claims, ok := token.Claims.(*Claims)
	if !ok {
		http.Error(w, "Invalid token", http.StatusUnauthorized)
		return
	}

	// Generate new token
	tokenString, err := generateJWT(claims.Username)
	if err != nil {
		http.Error(w, "Error generating token", http.StatusInternalServerError)
		return
	}

	result := map[string]string{"token": tokenString}
	utils.SendJson(w, result, http.StatusOK)
}

// Handles the JWT Verification endpoint of the auth service.
func VerifyHandler(w http.ResponseWriter, r *http.Request) {
	tokenString := r.Header.Get("Authorization")[7:]

	_, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		http.Error(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	result := map[string]string{
		"result": "ok",
	}
	utils.SendJson(w, result, http.StatusOK)
}

// Private functions

// Generates JWT token based on given string with given expiration.
func generateJWT(username string) (string, error) {
	expirationTime := time.Now().Add(JWT_TTL)

	claims := &Claims{
		Username: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expirationTime.Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(os.Getenv("JWT_SECRET")))
}
