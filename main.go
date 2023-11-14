package main

import (
	"encoding/json"
	"net/http"
)


func getUsers(w http.ResponseWriter, r *http.Request) {
    // Handle GET /users endpoint
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

func main() {
    // Define your API routes
    http.HandleFunc("/users", getUsers)

    // Start the server
    http.ListenAndServe(":8000", nil)
}
