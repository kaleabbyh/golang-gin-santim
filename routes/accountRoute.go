package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/kaleabbyh/golang-santim/controllers"
	"github.com/kaleabbyh/golang-santim/middleware"
)

func SetupAccountRoutes(router *gin.Engine) {

	router.POST("/account/createaccount", middleware.RoleCheckMiddleware(), controllers.CreateAccount)
	router.GET("/account/accounts/:id", middleware.RoleCheckMiddleware(), controllers.GetAccountByID)
	router.GET("/account/accounts", middleware.RoleCheckMiddleware(), controllers.GetAccountByAccountNumber)
	router.GET("/account/useraccounts/:UserID",middleware.IsLoggedIn(), controllers.GetAccountByUser)
	router.GET("/account/loggedinuseraccounts",middleware.IsLoggedIn(), controllers.GetAccountByLoggedInUser)
	router.GET("/account/getallaccounts", middleware.RoleCheckMiddleware(), controllers.GetAllAccounts)
	
}