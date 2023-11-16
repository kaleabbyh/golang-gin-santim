package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/controllers"
	"gorm.io/gorm"
)

func SetupUserRoutes(router *gin.Engine, db *gorm.DB) {
    userController := &controllers.UserController{DB: db  }
    router.GET("/users", userController.GetUsers)
	router.POST("/users", userController.CreateUser)
}


