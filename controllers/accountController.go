package controllers

import (
	//"fmt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// Create Account by if User is LoggedIn
func CreateAccount(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve userID"})
		return
	}
	
	var accountData Account
	if err := c.ShouldBindJSON(&accountData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	userIDValue, _ := userID.(uuid.UUID)
	accountData.CreatedBy =  userIDValue

	result := db.Create(&accountData)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var user User
	result = db.First(&user, accountData.UserID)
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

// Get Accounts by Acount Id
func GetAccountByID(c *gin.Context) {
	accountID := c.Param("id")

	var account Account
	result := db.First(&account, "id = ?", accountID)
	fmt.Println(accountID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	var user User
	result = db.First(&user, account.UserID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{ "Account": account, "User":    user })
}


// Get Accounts by Acount Id
func GetAccountByAccountNumber(c *gin.Context) {
	
	AccountNumber := c.Query("account_number")
	var account Account
	result := db.First(&account, "account_number = ?", AccountNumber)
	fmt.Println(AccountNumber)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error})
		return
	}

	var user User
	result = db.First(&user,"id=?", account.UserID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{ "Account": account, "User":    user })
}

// Get Accounts by Acount User Id
func GetAccountByUser(c *gin.Context) {
	userID, exists := c.Get("userID")

	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve userID"})
		return
	}

	var accounts []Account
	result := db.Where("user_id = ?", userID).Find(&accounts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accounts"})
		return
	}

	var user User
	result = db.First(&user, "id = ?", userID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{ "Account": accounts, "User":    user })
}


// Get Accounts by Acount LoggedIn User
func GetAccountByLoggedInUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve userID"})
		return
	}
	
	var accounts []Account
	result := db.Where("user_id = ?", userID).Find(&accounts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accounts"})
		return
	}
	fmt.Println("accounts", accounts)
	var user User
	result = db.First(&user, "id = ?", userID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{"Account": accounts,"User":user})
}


func GetAllAccounts(c *gin.Context) {

	var accounts []Account
	if err := db.Find(&accounts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve accounts"})
		return
	}

	c.JSON(http.StatusOK, accounts)
}
