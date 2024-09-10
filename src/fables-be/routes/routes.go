package routes

import (
	"github.com/Turizak/fables-be/account"
	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/character"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Accounts
	// POST
	router.POST("/account/create", account.CreateAccount)
	router.POST("/account/login", account.AccountLogin)
	router.POST("/account/token/refresh", account.RefreshAuthToken)
	// GET
	router.GET("/account/token/validate", account.ValidateAuthToken)
	router.GET("/account", account.GetAccount)
	router.GET("/account/created-characters", character.GetCharactersByCreatorUuid)
	router.GET("account/owned-characters", character.GetCharactersByOwnerUuid)
	router.GET("/account/campaigns", campaign.GetCampaignsByCreatorUuid)
	// PATCH
	router.PATCH("/account/change-password", account.ChangePassword)

	//Campaigns
	// POST
	router.POST("/campaign/create", campaign.CreateCampaign)
	// GET
	router.GET("/campaign/:uuid", campaign.GetCampaignByUuid)
	router.GET("/campaign/:uuid/moniker", campaign.GetCampaignMonikerByUuid)

	// Campaign - Characters
	// POST
	router.POST("/campaign/:uuid/character/create", character.CreateCharacter)
	// GET
	router.GET("/campaign/:uuid/character/:characterUuid", character.GetCharacterByUuid)
}
