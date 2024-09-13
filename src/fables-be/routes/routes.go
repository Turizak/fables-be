package routes

import (
	"github.com/Turizak/fables-be/account"
	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/character"
	abilityscore "github.com/Turizak/fables-be/ruleset/5e/ability-score"
	damagetype "github.com/Turizak/fables-be/ruleset/5e/damage-type"
	"github.com/Turizak/fables-be/ruleset/5e/language"
	"github.com/Turizak/fables-be/ruleset/5e/proficiencies"
	"github.com/Turizak/fables-be/ruleset/5e/race"
	"github.com/Turizak/fables-be/ruleset/5e/skill"
	"github.com/Turizak/fables-be/ruleset/5e/trait"
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
	// 5e
	// Races
	router.GET("/api/ruleset/5e/races", race.GetAllRaces5e)
	router.GET("/api/ruleset/5e/races/:index", race.GetRaceByIndex5e)
	// Languages
	router.GET("/api/ruleset/5e/languages", language.GetAllLanguages5e)
	router.GET("/api/ruleset/5e/languages/:index", language.GetLanguageByIndex5e)
	// Traits
	router.GET("/api/ruleset/5e/traits", trait.GetAllTraits5e)
	router.GET("/api/ruleset/5e/traits/:index", trait.GetTraitByIndex5e)
	// Skills
	router.GET("/api/ruleset/5e/skills", skill.GetAllSkills5e)
	router.GET("/api/ruleset/5e/skills/:index", skill.GetSkillByIndex5e)
	// Proficiencies
	router.GET("/api/ruleset/5e/proficiencies", proficiencies.GetAllProficiencies5e)
	router.GET("/api/ruleset/5e/proficiencies/:index", proficiencies.GetProficiencyByIndex5e)
	// Damage Types
	router.GET("/api/ruleset/5e/damage-types", damagetype.GetAllDamageTypes5e)
	router.GET("/api/ruleset/5e/damage-types/:index", damagetype.GetDamageTypeByIndex5e)
	// Ability Scores
	router.GET("/api/ruleset/5e/ability-scores", abilityscore.GetAllAbilityScores5e)
	router.GET("/api/ruleset/5e/ability-scores/:index", abilityscore.GetAbilityScoreByIndex5e)
}
