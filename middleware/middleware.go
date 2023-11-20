package middleware

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/kaleabbyh/golang-santim/utils"
)
func RoleCheckMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, role, err := utils.GetValuesFromToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if userID == uuid.Nil  {
			c.JSON(http.StatusBadRequest, gin.H{"error": "no user is logged in"})
			c.Abort()
			return
		}

		if role != "admin" && role != "superadmin" {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "not authorized"})
			c.Abort()
			return
		}
		
		c.Set("userID", userID)
		c.Set("role", role)

		c.Next()
	}
}

func IsLoggedIn() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID, role, err := utils.GetValuesFromToken(c)

		if err != nil {
			c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid token"})
			c.Abort()
			return
		}

		if userID == uuid.Nil || role == "" {
			c.JSON(http.StatusBadRequest, gin.H{"error": "UserID or role missing"})
			c.Abort()
			return
		}

		c.Set("userID", userID)
		c.Set("role", role)

		c.Next()
	}
}