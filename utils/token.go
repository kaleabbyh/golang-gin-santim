package utils

import (
	"errors"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/google/uuid"
)


func GenerateToken(userID uuid.UUID,role string) (string, error) {

	secretKey := []byte("kaleab")
	claims := jwt.MapClaims{
		"user_id": userID,
		"role":    role,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}


func ValidateToken(tokenString string) (uuid.UUID,string, error) {
	secretKey := []byte("kaleab")
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})

	if err != nil || !token.Valid {
		return uuid.UUID{},"", errors.New("invalid token")
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return uuid.UUID{},"", errors.New("invalid token claims")
	}



	userIDStr, ok := claims["user_id"].(string)
	role,ok := claims["role"].(string)
	if !ok {
		return uuid.UUID{},"", errors.New("invalid user ID or role in token claims")
	}
	
	
	userID, err := uuid.Parse(userIDStr)
	if err != nil {
		return uuid.UUID{},"", errors.New("failed to parse user ID from token claims")
	}

	return userID,role, nil
}


func GetValuesFromToken(c *gin.Context) (uuid.UUID,string, error) {
	token := c.GetHeader("Authorization")

	if token == "" {
		token = c.Query("token")
	}

	tokenParts := strings.Split(token, " ")
	if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
		return uuid.UUID{},"", nil
	}

	token = tokenParts[1]
	userID,role, _ := ValidateToken(token)
	return userID,role, nil
}



