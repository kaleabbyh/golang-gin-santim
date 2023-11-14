package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/kaleabbyh/golang-santim/controllers"
)

func SetupRoutes(router *gin.Engine) {
    router.GET("/payments", controllers.Getpayments)
    router.GET("/payments/:id", controllers.GetPaymentByID)
    router.POST("/payments", controllers.PostPayments)
}
