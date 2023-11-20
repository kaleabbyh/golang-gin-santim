package controllers

import (
	//"fmt"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/utils"
)

// Create Account by if User is LoggedIn
func CreateAccount(c *gin.Context) {
	
	userID, role, error := utils.GetUserIdFromToken(c)

	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	if role != "admin" && role != "superadmin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized to create account"})
		return
	}
	var accountData Account
	if err := c.ShouldBindJSON(&accountData); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	accountData.CreatedBy = userID

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
	_, role, error := utils.GetUserIdFromToken(c)

	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	if role != "admin" && role != "superadmin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized to search for account"})
		return
	}
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
	_, role, error := utils.GetUserIdFromToken(c)

	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	if role != "admin" && role != "superadmin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized to search for account"})
		return
	}

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
	UserID := c.Param("UserID")
	var accounts []Account
	result := db.Where("user_id = ?", UserID).Find(&accounts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accounts"})
		return
	}

	var user User
	result = db.First(&user, "id = ?", UserID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{ "Account": accounts, "User":    user })
}

// Get Accounts by Acount LoggedIn User
func GetAccountByLoggedInUser(c *gin.Context) {
	UserID, _, error := utils.GetUserIdFromToken(c)
	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	var accounts []Account
	result := db.Where("user_id = ?", UserID).Find(&accounts)

	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch accounts"})
		return
	}
	fmt.Println("accounts", accounts)
	var user User
	result = db.First(&user, "id = ?", UserID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}


	c.JSON(http.StatusOK, gin.H{"Account": accounts,"User":user})
}

func GetAllAccounts(c *gin.Context) {
	_, role, error := utils.GetUserIdFromToken(c)

	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	if role != "admin" && role != "superadmin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized to see user accounts"})
		return
	}
	var accounts []Account
	if err := db.Find(&accounts).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve accounts"})
		return
	}

	c.JSON(http.StatusOK, accounts)
}
