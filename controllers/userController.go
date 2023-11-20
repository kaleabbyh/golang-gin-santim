package controllers

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/utils"
)


func RegisterUser(c *gin.Context) {
  
    var user User
    error := c.ShouldBindJSON(&user)
    if error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    encryptedPassword, err := utils.HashPassword(user.Password)
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to encrypt password"})
        return
    }

    user.Password = encryptedPassword

    if err := db.Create(&user).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create user"})
        return
    }

    token, _ := utils.GenerateToken(user.ID,user.Role)

    c.JSON(http.StatusOK, gin.H{
        "status": http.StatusOK,
        "user":   user,
        "token":  token,
    })
}


func LoginUser(c *gin.Context) {
    
    loginData := struct {
        Email    string `json:"email"`
        Password string `json:"password"`
    }{}

    err := c.ShouldBindJSON(&loginData)
    if err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
        return
    }

    var user User
    result := db.First(&user, "email = ?", strings.ToLower(loginData.Email))
    if result.Error != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
        return
    }

    token, _ := utils.GenerateToken(user.ID,user.Role)

    if err := utils.VerifyPassword(user.Password, loginData.Password); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"status": "fail", "message": "Invalid email or password"})
        return
    }

    c.JSON(http.StatusOK, gin.H{
        "status": "success",
        "user":   user,
        "token":  token,
    })
}


func GetAllUesrs(c *gin.Context){
    _, role, error := utils.GetUserIdFromToken(c)

	if error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
		return
	}
	if role != "admin" && role != "superadmin" {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized to search for account"})
		return
	}
    var users []User
    if err := db.Find(&users).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to retrieve users"})
        return
    }

    c.JSON(http.StatusOK, users)
}