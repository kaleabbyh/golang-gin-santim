package controllers

import (
	"encoding/json"
	"net/http"
)

type UserController struct{}

func NewUserController() *UserController {
	return &UserController{}
}

func (uc *UserController) GetUsers(w http.ResponseWriter, r *http.Request) {
	// Retrieve users from the database or any data source
	users := []string{"John", "Jane", "Alice", "Bob"}

	// Convert users to JSON
	jsonBytes, err := json.Marshal(users)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Set response headers
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)

	// Write JSON response
	w.Write(jsonBytes)
}