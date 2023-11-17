package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/utils"
)

//Create Account by if User is LoggedIn
func CreateAccount(c *gin.Context) {
    userID := utils.GetUserIdFromToken(c)
    if userID == 0 {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
        return
    }

    var accountData Account
    if err := c.ShouldBindJSON(&accountData); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    accountData.UserID = userID
    result := db.Create(&accountData)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    var user User
    result = db.First(&user, userID)
    if result.Error != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
        return
    }

    response := AccountResponse{
        Account: accountData,
        User:    user,
        Message: "Account created successfully",
    }

    c.JSON(http.StatusCreated, response)
}

//Get Accounts by Acount Id
func GetAccountByID(c *gin.Context) {
	accountID := c.Param("id")
	var account Account
	result := db.First(&account, accountID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
		return
	}

	var user User
	result = db.First(&user, account.UserID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	response := AccountResponse{
		Account: account,
		User:    user,
	}
	c.JSON(http.StatusOK, response)
}


//Get Accounts by Acount User Id
func GetAccountByUser(c *gin.Context) {
	UserID := c.Param("UserID")
	var accounts []Account
	result := db.Where("user_id = ?", UserID).Find(&accounts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accounts"})
		return
	}

	var user User
	result = db.First(&user, UserID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}
	
	response := AccountResponses{
		Account: accounts,
		User:    user,
	}
	c.JSON(http.StatusOK, response)
}


//Get Accounts by Acount LoggedIn User
func GetAccountByLoggedInUser(c *gin.Context) {
	UserID := utils.GetUserIdFromToken(c)
	var accounts []Account
	result := db.Where("user_id = ?", UserID).Find(&accounts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accounts"})
		return
	}

	var user User
	result = db.First(&user, UserID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	response := AccountResponses{
		Account: accounts,
		User:    user,
	}
	c.JSON(http.StatusOK, response)
}


func GetAllAccounts(c *gin.Context){
	var accounts []Account
	if err := db.Find(&accounts).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve accounts"})
        return
    }

    c.JSON(http.StatusOK, accounts)
}
