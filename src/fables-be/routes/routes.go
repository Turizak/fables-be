package routes

import (
	"github.com/Turizak/fables-be/account"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Account
	router.POST("/account/create", account.CreateAccount)
}
