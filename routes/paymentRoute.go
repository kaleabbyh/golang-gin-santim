package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/controllers"
)

func SetupPaymentRoutes(router *gin.Engine) {
    router.POST("/payments/createpayment", controllers.CreatePayments)
    router.GET("/payments/getallpayments", controllers.GetAllPayments)
	router.GET("/payments/getallpaymentsByUser", controllers.GetPaymentsByUser)
	router.GET("/payments/getpaymentById/:PaymentID", controllers.GetPaymentsById)
	
	
}
