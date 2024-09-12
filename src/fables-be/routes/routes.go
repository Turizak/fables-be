package routes

import (
	"github.com/Turizak/fables-be/account"
	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/character"
	ruleset "github.com/Turizak/fables-be/ruleset/5e/race"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	// Accounts
	// POST
	router.POST("/api/account/create", account.CreateAccount)
	router.POST("/api/account/login", account.AccountLogin)
	router.POST("/api/account/token/refresh", account.RefreshAuthToken)
	// GET
	router.GET("/api/account/token/validate", account.ValidateAuthToken)
	router.GET("/api/account", account.GetAccount)
	router.GET("/api/account/created-characters", character.GetCharactersByCreatorUuid)
	router.GET("/api/account/owned-characters", character.GetCharactersByOwnerUuid)
	router.GET("/api/account/campaigns", campaign.GetCampaignsByCreatorUuid)
	// PATCH
	router.PATCH("/api/account/change-password", account.ChangePassword)

	//Campaigns
	// POST
	router.POST("/api/campaign/create", campaign.CreateCampaign)
	// GET
	router.GET("/api/campaign/:uuid", campaign.GetCampaignByUuid)
	router.GET("/api/campaign/:uuid/moniker", campaign.GetCampaignMonikerByUuid)

	// Campaign - Characters
	// POST
	router.POST("/api/campaign/:uuid/character/create", character.CreateCharacter)
	// GET
	router.GET("/api/campaign/:uuid/character/:characterUuid", character.GetCharacterByUuid)

	//Ruleset
	// GET
	router.GET("/api/ruleset/5e/races", ruleset.GetAllRaces5e)
	router.GET("/api/ruleset/5e/races/:index", ruleset.GetRaceByIndex5e)
}
