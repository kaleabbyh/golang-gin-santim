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
    defer db.Close()


    router := gin.Default()
    routes.SetupRoutes(router)


    fmt.Println("Server is running on port 8080")
    err = router.Run("localhost:8080")
    if err != nil {
        log.Fatal("Error starting the server:", err)
    }
}