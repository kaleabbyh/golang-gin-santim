package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/config"
	"github.com/kaleabbyh/golang-santim/routes"
	_ "github.com/lib/pq"
)

func main() {
	db, err := config.ConnectDB()
    if err != nil {
        log.Fatal("Error connecting to the database:", err)
    }


    router := gin.Default()
    routes.SetupRoutes(router)
    routes.SetupUserRoutes(router, db)

    fmt.Println("Server is running on port 8080")
    err = router.Run("localhost:8080")
    if err != nil {
        log.Fatal("Error starting the server:", err)
    }
}