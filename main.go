package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/routes"
)

func main() {
	

    router := gin.Default()

    routes.SetupPaymentRoutes(router)
    routes.SetupUserRoutes(router)
    routes.SetupAccountRoutes(router)
    routes.SetupTransactionRoutes(router)
    

    fmt.Println("Server is running on port 8080")
    err := router.Run("localhost:8080")
    if err != nil {
        log.Fatal("Error starting the server:", err)
    }
}