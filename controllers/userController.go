package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/models"
	"gorm.io/gorm"
)

type UserController struct {
    DB *gorm.DB
}

func (uc *UserController) GetUsers(c *gin.Context) {
    var users []models.User
    if err := uc.DB.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch users"})
        return
    }
    c.JSON(http.StatusOK, users)
}



func (uc *UserController) CreateUser(c *gin.Context) {
    // Create a new User instance
    user := models.User{}
    if err := c.ShouldBindJSON(&user); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    if err := uc.DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    c.JSON(http.StatusCreated, user)
}