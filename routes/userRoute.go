package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/controllers"
)

func SetupUserRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
}


