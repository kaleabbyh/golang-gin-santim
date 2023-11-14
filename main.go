package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	// "github.com/kaleabbyh/golang-santim/routes"
)


type payment struct {
    ID     string  `json:"id"`
    Name  string  `json:"name"`
    Reason string  `json:"reason"`
    Price  float64 `json:"price"`
}

// payments slice to seed record payment data.
var payments = []payment{
    {ID: "1", Name: "Blues", Reason: "payed for rent", Price: 56.99},
    {ID: "2", Name: "Jeru", Reason: "payed for rent", Price: 17.99},
    {ID: "3", Name: "Sarah", Reason: "payed for rent", Price: 39.99},
}

// getpayments responds with the list of all payments as JSON.
func getpayments(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, payments)
}

// postAlbums adds an album from JSON received in the request body.
func postPayments(c *gin.Context) {
    var newAlbum payment
    if err := c.BindJSON(&newAlbum); err != nil {
        return
    }

    // Add the new album to the slice.
    payments = append(payments, newAlbum)
    c.IndentedJSON(http.StatusCreated, newAlbum)
}

func getPaymentByID(c *gin.Context) {
    id := c.Param("id")

    // Loop over the list of albums, looking for
    // an album whose ID value matches the parameter.
    for _, a := range payments {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "album not found"})
}

func main() {
    router := gin.Default()
    router.GET("/payments", getpayments)
	router.GET("/payments/:id", getPaymentByID)
	router.POST("/payments", postPayments)

	fmt.Println("Server is running on port 8000")
    err :=router.Run("localhost:8080")
		if err != nil {
		log.Fatal("Error starting the server:", err)
	}
}
