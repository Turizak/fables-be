package routes

import (
	"github.com/Turizak/fables-be/account"
	"github.com/Turizak/fables-be/auth"
	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/character"
	"github.com/Turizak/fables-be/collection"
	"github.com/Turizak/fables-be/location"
	"github.com/Turizak/fables-be/note"
	"github.com/Turizak/fables-be/npc"
	"github.com/Turizak/fables-be/quest"
	abilityscore "github.com/Turizak/fables-be/ruleset/5e/ability-score"
	"github.com/Turizak/fables-be/ruleset/5e/alignment"
	"github.com/Turizak/fables-be/ruleset/5e/class"
	"github.com/Turizak/fables-be/ruleset/5e/condition"
	damagetype "github.com/Turizak/fables-be/ruleset/5e/damage-type"
	"github.com/Turizak/fables-be/ruleset/5e/language"
	"github.com/Turizak/fables-be/ruleset/5e/proficiencies"
	"github.com/Turizak/fables-be/ruleset/5e/race"
	"github.com/Turizak/fables-be/ruleset/5e/skill"
	"github.com/Turizak/fables-be/ruleset/5e/subrace"
	"github.com/Turizak/fables-be/ruleset/5e/trait"
	"github.com/Turizak/fables-be/session"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	unprotectedAccountRoutes := router.Group("/api/account")
	{
		unprotectedAccountRoutes.POST("/create", account.CreateAccount)
		unprotectedAccountRoutes.POST("/login", account.AccountLogin)
		unprotectedAccountRoutes.POST("/token/refresh", account.RefreshAuthToken)
	}

	protectedAccountRoutes := router.Group("/api/account")
	protectedAccountRoutes.Use(auth.AuthorizeMiddleware())
	{
		protectedAccountRoutes.GET("/token/validate", account.ValidateAuthToken)
		protectedAccountRoutes.GET("/detail", account.GetAccount)
		protectedAccountRoutes.GET("/created-characters", character.GetCharactersByCreatorUuid)
		protectedAccountRoutes.GET("/owned-characters", character.GetCharactersByOwnerUuid)
		protectedAccountRoutes.GET("/campaigns", campaign.GetCampaignsByCreatorUuid)
		protectedAccountRoutes.GET("/created-locations", location.GetLocationsByCreatorUuid)
		protectedAccountRoutes.GET("/created-npcs", npc.GetNpcsByCreatorUuid)
		protectedAccountRoutes.GET("/created-notes", note.GetNotesByCreatorUuid)
		protectedAccountRoutes.PATCH("/change-password", account.ChangePassword)
	}

	protectedCampaignRoutes := router.Group("/api/campaign")
	protectedCampaignRoutes.Use(auth.AuthorizeMiddleware())
	{
		protectedCampaignRoutes.PATCH("/:uuid/update", campaign.UpdateCampaignByUuid)
		protectedCampaignRoutes.POST("/create", campaign.CreateCampaign)
		protectedCampaignRoutes.GET("/:uuid", campaign.GetCampaignByUuid)
		protectedCampaignRoutes.GET("/:uuid/moniker", campaign.GetCampaignMonikerByUuid)
		protectedCampaignRoutes.GET("/:uuid/all", collection.GetCampaignAllData)

		protectedCampaignRoutes.PATCH("/:uuid/character/:characterUuid/update", character.UpdateCharacterByUuid)
		protectedCampaignRoutes.POST("/:uuid/character/create", character.CreateCharacter)
		protectedCampaignRoutes.POST("/:uuid/session/:sessionUuid/character/create", character.CreateCharacterSession)
		protectedCampaignRoutes.GET("/:uuid/character/:characterUuid", character.GetCharacterByUuid)

		protectedCampaignRoutes.PATCH("/:uuid/location/:locationUuid/update", location.UpdateLocationByUuid)
		protectedCampaignRoutes.POST("/:uuid/location/create", location.CreateLocation)
		protectedCampaignRoutes.POST("/:uuid/session/:sessionUuid/location/create", location.CreateLocationSession)
		protectedCampaignRoutes.GET("/:uuid/location/:locationUuid", location.GetLocationByUuid)
		protectedCampaignRoutes.GET("/:uuid/locations", location.GetLocationsByCampaignUuid)

		protectedCampaignRoutes.PATCH("/:uuid/npc/:npcUuid/update", npc.UpdateNpcByUuid)
		protectedCampaignRoutes.POST("/:uuid/npc/create", npc.CreateNpc)
		protectedCampaignRoutes.POST("/:uuid/session/:sessionUuid/npc/create", npc.CreateNpcSession)
		protectedCampaignRoutes.GET("/:uuid/npc/:npcUuid", npc.GetNpcByUuid)
		protectedCampaignRoutes.GET("/:uuid/npcs", npc.GetNpcsByCampaignUuid)

		protectedCampaignRoutes.PATCH("/:uuid/session/:sessionUuid/update", session.UpdateSessionByUuid)
		protectedCampaignRoutes.POST("/:uuid/session/create", session.CreateSession)
		protectedCampaignRoutes.GET("/:uuid/session/:sessionUuid", session.GetSessionByUuid)
		protectedCampaignRoutes.GET("/:uuid/sessions", session.GetSessionsByCampaignUuid)

		protectedCampaignRoutes.GET("/:uuid/session/:sessionUuid/all", collection.GetAllSessionData)
		//Quest
		protectedCampaignRoutes.PATCH("/:uuid/quest/:questUuid/update", quest.UpdateQuestByUuid)
		protectedCampaignRoutes.POST("/:uuid/quest/create", quest.CreateQuest)
		protectedCampaignRoutes.GET("/:uuid/quest/:questUuid", quest.GetQuestByUuid)
		protectedCampaignRoutes.GET("/:uuid/quests", quest.GetQuestsByCampaignUuid)
		// Note
		protectedCampaignRoutes.PATCH("/:uuid/note/:noteUuid/update", note.UpdateNoteByUuid)
		protectedCampaignRoutes.POST("/:uuid/session/:sessionUuid/note/create", note.CreateNote)
		protectedCampaignRoutes.GET("/:uuid/note/:noteUuid", note.GetNoteByUuid)
		protectedCampaignRoutes.GET("/:uuid/notes", note.GetNotesByCampaignUuid)
	}

	protectedRulesetRoutes := router.Group("/api/ruleset")
	protectedRulesetRoutes.Use(auth.AuthorizeMiddleware())
	{
		// Races
		protectedRulesetRoutes.GET("/5e/races", race.GetAllRaces5e)
		protectedRulesetRoutes.GET("/5e/races/:index", race.GetRaceByIndex5e)
		// Languages
		protectedRulesetRoutes.GET("/5e/languages", language.GetAllLanguages5e)
		protectedRulesetRoutes.GET("/5e/languages/:index", language.GetLanguageByIndex5e)
		// Traits
		protectedRulesetRoutes.GET("/5e/traits", trait.GetAllTraits5e)
		protectedRulesetRoutes.GET("/5e/traits/:index", trait.GetTraitByIndex5e)
		// Skills
		protectedRulesetRoutes.GET("/5e/skills", skill.GetAllSkills5e)
		protectedRulesetRoutes.GET("/5e/skills/:index", skill.GetSkillByIndex5e)
		// Proficiencies
		protectedRulesetRoutes.GET("/5e/proficiencies", proficiencies.GetAllProficiencies5e)
		protectedRulesetRoutes.GET("/5e/proficiencies/:index", proficiencies.GetProficiencyByIndex5e)
		// Damage Types
		protectedRulesetRoutes.GET("/5e/damage-types", damagetype.GetAllDamageTypes5e)
		protectedRulesetRoutes.GET("/5e/damage-types/:index", damagetype.GetDamageTypeByIndex5e)
		// Ability Scores
		protectedRulesetRoutes.GET("/5e/ability-scores", abilityscore.GetAllAbilityScores5e)
		protectedRulesetRoutes.GET("/5e/ability-scores/:index", abilityscore.GetAbilityScoreByIndex5e)
		// Alignment
		protectedRulesetRoutes.GET("/5e/alignments", alignment.GetAllAlignments5e)
		protectedRulesetRoutes.GET("/5e/alignments/:index", alignment.GetAlignmentByIndex5e)
		// Conditions
		protectedRulesetRoutes.GET("/5e/conditions", condition.GetAllConditionsWithDescriptions)
		protectedRulesetRoutes.GET("/5e/conditions/:index", condition.GetConditionByIndex)
		// Classes
		protectedRulesetRoutes.GET("/5e/classes", class.GetAllClasses5e)
		protectedRulesetRoutes.GET("/5e/classes/:index", class.GetClassByIndex5e)
		// Subraces
		protectedRulesetRoutes.GET("/5e/subraces", subrace.GetAllSubraces5e)
		protectedRulesetRoutes.GET("/5e/subraces/:index", subrace.GetSubraceByIndex5e)
	}
}
