package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/controllers"
)

func SetupTransactionRoutes(router *gin.Engine) {
    router.GET("/transactions/getalltransactions", controllers.GetAllTransactions)
    
   
}