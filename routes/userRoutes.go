package routes

import (
	"github.com/gorilla/mux"
	"github.com/kaleabbyh/golang-santim/controllers"
)

func NewRouter() *mux.Router {
	// Create a new instance of the router
	router := mux.NewRouter()

	// Create a new instance of the UserController
	userController := controllers.NewUserController()

	// Define your API routes
	router.HandleFunc("/users", userController.GetUsers).Methods("GET")

	return router
}