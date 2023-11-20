package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/controllers"
	"github.com/kaleabbyh/golang-santim/middleware"
)

func SetupUserRoutes(router *gin.Engine) {
	router.POST("/register", controllers.RegisterUser)
	router.POST("/login", controllers.LoginUser)
	router.GET("/user/getallusers",middleware.RoleCheckMiddleware(),controllers.GetAllUesrs)
}