package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/controllers"
	// "gorm.io/gorm"
)

func SetupAccountRoutes(router *gin.Engine) {

	router.POST("/account/createaccount", controllers.CreateAccount)
	router.GET("/account/accounts/:id", controllers.GetAccountByID)
	router.GET("/account/useraccounts/:UserID", controllers.GetAccountByUser)
	router.GET("/account/loggedinuseraccounts", controllers.GetAccountByLoggedInUser)
	router.GET("/account/getallaccounts", controllers.GetAllAccounts)
	
	
}