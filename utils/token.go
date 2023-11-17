package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
)



func GenerateToken(userID uint) (string, error) {

	secretKey := []byte("kaleab") 
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func ValidateToken(tokenString string) (uint, error) {
	
	secretKey := []byte("kaleab") 
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return 0, errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return 0, errors.New("invalid token claims")
	}

	userID, ok := claims["user_id"].(float64)
	if !ok {
		return 0, errors.New("invalid user ID in token claims")
	}

	return uint(userID), nil
}


func GetUserIdFromToken(c *gin.Context) uint {
    token := c.GetHeader("Authorization")
	
    if token == "" {
        token = c.Query("token")
    }

    tokenParts := strings.Split(token, " ")
    if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
        return 0
    }

    token = tokenParts[1]
	userID, _ := ValidateToken(token)
    return userID
}