package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/controllers"
)

func SetupDemoRoutes(router *gin.Engine) {
	router.POST("/createdemodata", controllers.CreateDemo)
}