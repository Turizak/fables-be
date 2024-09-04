package routes

import (
	"github.com/Turizak/fables-be/account"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Account
	router.POST("/account/create", account.CreateAccount)
	router.POST("/account/login", account.AccountLogin)
	router.POST("/account/token/refresh", account.RefreshAuthToken)
	router.GET("/account/token/validate", account.ValidateAuthToken)
}
