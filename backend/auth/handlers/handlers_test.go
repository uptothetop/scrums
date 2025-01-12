package handlers

import (
	"encoding/json"
	"net/http"
	"scrums/auth/m/v2/models"
	"testing"
	"testutils"

	"golang.org/x/crypto/bcrypt"

	"gorm.io/gorm"
)

func setupDB(t *testing.T) *gorm.DB {
	db := testutils.SetupDB(t, &models.User{})

	// Insert a new record
	testUser := "testuser"
	testPass := "testpass"
	password, _ := bcrypt.GenerateFromPassword([]byte(testPass), bcrypt.DefaultCost)
	db.Create(&models.User{Username: testUser, PasswordHash: string(password)})

	return db
}

func TestLoginHandlers(t *testing.T) {
	db := setupDB(t)

	t.Run("LoginHandler", func(t *testing.T) {
		t.Run("Success", func(t *testing.T) {
			rr, req := testutils.MakeRequest("/login", "POST", map[string]string{
				"username": "testuser",
				"password": "testpass",
			})
			handler := http.HandlerFunc(LoginHandler(db))
			handler.ServeHTTP(rr, req)

			testutils.Assert(t, http.StatusOK, rr.Code, "Expected successful login")

			// Assert specific responses
		})

		t.Run("Failure", func(t *testing.T) {
			rr, req := testutils.MakeRequest("/login", "POST", map[string]string{
				"username": "wronguser",
				"password": "wrongpass",
			})
			handler := http.HandlerFunc(LoginHandler(db))
			handler.ServeHTTP(rr, req)

			testutils.Assert(t, http.StatusUnauthorized, rr.Code, "Expected unauthorized for bad creds")
		})
	})
}

func TestRegisterHandler(t *testing.T) {
	db := testutils.SetupDB(t, &models.User{})

	t.Run("Successful Registration", func(t *testing.T) {
		rr, req := testutils.MakeRequest("/register", "POST", map[string]string{
			"username": "newuser",
			"password": "newpass",
		})
		handler := http.HandlerFunc(RegisterHandler(db))
		handler.ServeHTTP(rr, req)

		testutils.Assert(t, http.StatusOK, rr.Code, "Expected successful registration")
	})
}

func TestRefreshHandler(t *testing.T) {
	// Simulate JWT creation (assuming generateJWT is available in scope):
	token, _ := generateJWT("testuser")

	t.Run("Valid Token Refresh", func(t *testing.T) {
		rr, req := testutils.MakeRequest("/refresh", "POST", nil)
		req.Header.Set("Authorization", "Bearer "+token) // Including token in header

		handler := http.HandlerFunc(RefreshHandler)
		handler.ServeHTTP(rr, req)

		testutils.Assert(t, http.StatusOK, rr.Code, "Expected successful token refresh")

		var response map[string]string
		json.NewDecoder(rr.Body).Decode(&response)
		if _, exists := response["token"]; !exists {
			t.Error("Expected token in response, got none")
		}
	})

	t.Run("Invalid Token Refresh", func(t *testing.T) {
		rr, req := testutils.MakeRequest("/refresh", "POST", nil)
		req.Header.Set("Authorization", "Bearer invalidtoken")

		handler := http.HandlerFunc(RefreshHandler)
		handler.ServeHTTP(rr, req)

		testutils.Assert(t, http.StatusUnauthorized, rr.Code, "Expected unauthorized for bad token")
	})
}

func TestVerifyJWT(t *testing.T) {
	validToken, _ := generateJWT("testuser")

	t.Run("Valid Token Verification", func(t *testing.T) {
		rr, req := testutils.MakeRequest("/verify", "POST", nil)
		req.Header.Set("Authorization", "Bearer "+validToken)

		handler := http.HandlerFunc(VerifyHandler)
		handler.ServeHTTP(rr, req)

		testutils.Assert(t, http.StatusOK, rr.Code, "Expected valid token verification")
	})

	t.Run("Invalid Token Verification", func(t *testing.T) {
		rr, req := testutils.MakeRequest("/verify", "POST", nil)
		req.Header.Set("Authorization", "Bearer badtoken")

		handler := http.HandlerFunc(VerifyHandler)
		handler.ServeHTTP(rr, req)

		testutils.Assert(t, http.StatusUnauthorized, rr.Code, "Expected unauthorized for invalid token")
	})
}
