package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type payment struct {
    ID     string  `json:"id"`
    Name   string  `json:"name"`
    Reason string  `json:"reason"`
    Price  float64 `json:"price"`
}

var payments = []payment{
    {ID: "1", Name: "Blues", Reason: "payed for rent", Price: 56.99},
    {ID: "2", Name: "Jeru", Reason: "payed for rent", Price: 17.99},
    {ID: "3", Name: "Sarah", Reason: "payed for rent", Price: 39.99},
}

func Getpayments(c *gin.Context) {
    c.IndentedJSON(http.StatusOK, payments)
}

func PostPayments(c *gin.Context) {
    var newPayment payment
    if err := c.BindJSON(&newPayment); err != nil {
        return
    }

    payments = append(payments, newPayment)
    c.IndentedJSON(http.StatusCreated, newPayment)
}

func GetPaymentByID(c *gin.Context) {
    id := c.Param("id")

    for _, a := range payments {
        if a.ID == id {
            c.IndentedJSON(http.StatusOK, a)
            return
        }
    }
    c.IndentedJSON(http.StatusNotFound, gin.H{"message": "payment not found"})
}
