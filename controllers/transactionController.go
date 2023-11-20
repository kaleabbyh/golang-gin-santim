package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/utils"
)

func GetAllTransactions(c *gin.Context){
    _, role, error := utils.GetUserIdFromToken(c)

	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	if role != "admin" && role != "superadmin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized to search for account"})
		return
	}
	var transactions []Transaction
    if err := db.Find(&transactions).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve transactions"})
        return
    }

    c.JSON(http.StatusOK, transactions)
	

}