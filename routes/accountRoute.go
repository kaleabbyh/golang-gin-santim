package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/controllers"
	// "gorm.io/gorm"
)

func SetupAccountRoutes(router *gin.Engine) {

	router.POST("/acount/createacount", controllers.CreateAccount)
	router.GET("/acount/accounts/:id", controllers.GetAccountByID)
	router.GET("/acount/useraccounts/:UserID", controllers.GetAccountByUser)
	router.GET("/acount/loggedinuseraccounts", controllers.GetAccountByLoggedInUser)
	router.GET("/acount/getallaccounts", controllers.GetAllAccounts)
	
	
}