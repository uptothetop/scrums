// Package with testing utils
package testutils

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert" // Optional assertion library
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// MakeRequest creates and executes HTTP requests
func MakeRequest(api string, method string, payload interface{}) (*httptest.ResponseRecorder, *http.Request) {
	requestBody, _ := json.Marshal(payload)                              // Marshal payload to JSON
	req, _ := http.NewRequest(method, api, bytes.NewBuffer(requestBody)) // Create new request
	rr := httptest.NewRecorder()                                         // Create a new ResponseRecorder
	return rr, req                                                       // Return both for usage in tests
}

// SetupDB initializes an in-memory database and migrates schemas
func SetupDB(t *testing.T, models ...interface{}) *gorm.DB {
	db, _ := gorm.Open(sqlite.Open("file::memory:?cache=shared"), &gorm.Config{})
	db.AutoMigrate(models...)

	// Cleanup function to close the database connection
	t.Cleanup(func() {
		sqlDB, _ := db.DB()
		sqlDB.Close()
	})

	return db
}

// Assert function to compare expected and actual values
func Assert(t *testing.T, expected, actual interface{}, msgAndArgs ...interface{}) {
	assert.Equal(t, expected, actual, msgAndArgs...)
}
