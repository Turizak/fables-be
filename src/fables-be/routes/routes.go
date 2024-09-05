package routes

import (
	"github.com/Turizak/fables-be/account"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Account
	// POST
	router.POST("/account/create", account.CreateAccount)
	router.POST("/account/login", account.AccountLogin)
	router.POST("/account/token/refresh", account.RefreshAuthToken)
	// GET
	router.GET("/account/token/validate", account.ValidateAuthToken)
	router.GET("/account", account.GetAccount)
	// PATCH
	router.PATCH("/account/change-password", account.ChangePassword)
}
