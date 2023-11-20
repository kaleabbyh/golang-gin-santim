package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func GetAllTransactions(c *gin.Context){
   
	var transactions []Transaction
    if err := db.Find(&transactions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
        return
    }

    c.JSON(http.StatusOK, transactions)
	

}