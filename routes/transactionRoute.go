package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/controllers"
	"github.com/kaleabbyh/golang-santim/middleware"
)

func SetupTransactionRoutes(router *gin.Engine) {
    router.GET("/transactions/getalltransactions",middleware.RoleCheckMiddleware(), controllers.GetAllTransactions)
}