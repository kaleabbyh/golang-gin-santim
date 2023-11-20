package controllers

import (
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/models"
	"gorm.io/gorm"
)


func CreatePayments(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve userID"})
		return
	}

	var newPayment Payment
	if err := c.ShouldBindJSON(&newPayment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if newPayment.ReceiverAccount == newPayment.PayerAccount {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "self transfer is not allowed"})
		return
	}
 
	var loggedInUser User
	result := db.First(&loggedInUser, userID)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	var ReceiverAccount models.Account
	result = db.Where("account_number = ?", newPayment.ReceiverAccount).Find(&ReceiverAccount)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error(), "message": "reciever account doesnt ext"})
		return
	}

	var PayerAccount models.Account
	result = db.Where("account_number = ?", newPayment.PayerAccount).Find(&PayerAccount)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error(), "message": "this account doesnt ext"})
		return
	}

	if PayerAccount.Balance-25 < newPayment.Amount {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "insufficient funds"})
		return
	}

	if PayerAccount.UserID != userID {
		c.JSON(http.StatusInternalServerError, gin.H{"message": "this account does not belongs to you"})
		return
	}

	newPayment.UserID = PayerAccount.UserID
	if err := db.Create(&newPayment).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create newPayment"})
		return
	}

	// Update the payer account balance
	PayerAccount.Balance -= newPayment.Amount
	result = db.Save(&PayerAccount)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update account balance"})
		return
	}

	payerTransaction := Transaction{
		PaymentID:   newPayment.ID,
		UserID:      newPayment.UserID,
		Type:        "payed",
		Amount:      newPayment.Amount,
		TranferedTo: newPayment.ReceiverAccount,
	}

	if err := db.Create(&payerTransaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create newPayment"})
		return
	}

	// Update the reciever account balance
	ReceiverAccount.Balance += newPayment.Amount
	result = db.Save(&ReceiverAccount)
	if result.Error != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update account balance"})
		return
	}

	recieverTransaction := Transaction{
		PaymentID:     newPayment.ID,
		UserID:        ReceiverAccount.UserID,
		Type:          "recieved",
		TranferedFrom: newPayment.PayerAccount,
		Amount:        newPayment.Amount,
	}

	if err := db.Create(&recieverTransaction).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create newPayment"})
		return
	}

	response := PaymentResponse{
		Payment: newPayment,
		User:    loggedInUser,
		Message: "Payed successfully",
	}
	c.JSON(http.StatusCreated, response)
}


func GetPaymentsByUser(c *gin.Context) {
	userID, exists := c.Get("userID")
	if !exists {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve userID"})
		return
	}

	var payments []Payment
	err := db.Where("user_id = ?", userID).Find(&payments)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
	}

	c.JSON(http.StatusOK, payments)
}

func GetPaymentsById(c *gin.Context) {
	paymentID := c.Param("PaymentID")
	var payment Payment
	result := db.First(&payment,"id=?", paymentID)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			c.JSON(http.StatusNotFound, gin.H{"error": "Payment not found"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, payment)
}

func GetAllPayments(c *gin.Context) {
	var payments []Payment
	if err := db.Find(&payments).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve payments"})
		return
	}

	c.JSON(http.StatusOK, payments)
}
