package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/controllers"
	"github.com/kaleabbyh/golang-santim/middleware"
)

func SetupPaymentRoutes(router *gin.Engine) {
    router.POST("/payments/createpayment",middleware.IsLoggedIn(), controllers.CreatePayments)
    router.GET("/payments/getallpayments",middleware.RoleCheckMiddleware(), controllers.GetAllPayments)
	router.GET("/payments/getallpaymentsByUser",middleware.IsLoggedIn(), controllers.GetPaymentsByUser)
	router.GET("/payments/getpaymentById/:PaymentID",middleware.RoleCheckMiddleware(), controllers.GetPaymentsById)
}
