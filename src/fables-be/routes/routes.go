package routes

import (
	"github.com/Turizak/fables-be/account"
	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/character"
	"github.com/Turizak/fables-be/collection"
	"github.com/Turizak/fables-be/location"
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
	router.GET("/api/account/created-locations", location.GetLocationsByCreatorUuid)
	router.GET("/api/account/created-npcs", npc.GetNpcsByCreatorUuid)
	// PATCH
	router.PATCH("/api/account/change-password", account.ChangePassword)

	//Campaigns
	//PATCH
	router.PATCH("/api/campaign/:uuid/update", campaign.UpdateCampaignByUuid)
	// POST
	router.POST("/api/campaign/create", campaign.CreateCampaign)
	// GET
	router.GET("/api/campaign/:uuid", campaign.GetCampaignByUuid)
	router.GET("/api/campaign/:uuid/moniker", campaign.GetCampaignMonikerByUuid)
	router.GET("/api/campaign/:uuid/all", collection.GetCampaignAllData)

	// Campaign - Characters
	//PATCH
	router.PATCH("/api/campaign/:uuid/character/:characterUuid/update", character.UpdateCharacterByUuid)
	// POST
	router.POST("/api/campaign/:uuid/character/create", character.CreateCharacter)
	router.POST("/api/campaign/:uuid/session/:sessionUuid/character/create", character.CreateCharacterSession)
	// GET
	router.GET("/api/campaign/:uuid/character/:characterUuid", character.GetCharacterByUuid)

	// Campaign - Locations
	//PATCH
	router.PATCH("/api/campaign/:uuid/location/:locationUuid/update", location.UpdateLocationByUuid)
	// POST
	router.POST("/api/campaign/:uuid/location/create", location.CreateLocation)
	router.POST("/api/campaign/:uuid/session/:sessionUuid/location/create", location.CreateLocationSession)
	// GET
	router.GET("/api/campaign/:uuid/location/:locationUuid", location.GetLocationByUuid)
	router.GET("/api/campaign/:uuid/locations", location.GetLocationsByCampaignUuid)

	// Campaign - Npcs
	//PATCH
	router.PATCH("/api/campaign/:uuid/npc/:npcUuid/update", npc.UpdateNpcByUuid)
	// POST
	router.POST("/api/campaign/:uuid/npc/create", npc.CreateNpc)
	router.POST("/api/campaign/:uuid/session/:sessionUuid/npc/create", npc.CreateNpcSession)
	// GET
	router.GET("/api/campaign/:uuid/npc/:npcUuid", npc.GetNpcByUuid)
	router.GET("/api/campaign/:uuid/npcs", npc.GetNpcsByCampaignUuid)

	// Campaign - Sessions
	//PATCH
	router.PATCH("/api/campaign/:uuid/session/:sessionUuid/update", session.UpdateSessionByUuid)
	// POST
	router.POST("/api/campaign/:uuid/session/create", session.CreateSession)
	// GET
	router.GET("/api/campaign/:uuid/session/:sessionUuid", session.GetSessionByUuid)
	router.GET("/api/campaign/:uuid/sessions", session.GetSessionsByCampaignUuid)
	router.GET("/api/campaign/:uuid/session/:sessionUuid/all", collection.GetAllSessionData)

	// Campaign - Quests
	// POST
	router.POST("/api/campaign/:uuid/quest/create", quest.CreateQuest)
	// GET
	router.GET("/api/campaign/:uuid/quest/:questUuid", quest.GetQuestByUuid)
	router.GET("/api/campaign/:uuid/quests", quest.GetQuestsByCampaignUuid)

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
	// Alignment
	router.GET("/api/ruleset/5e/alignments", alignment.GetAllAlignments5e)
	router.GET("/api/ruleset/5e/alignments/:index", alignment.GetAlignmentByIndex5e)
	// Conditions
	router.GET("/api/ruleset/5e/conditions", condition.GetAllConditionsWithDescriptions)
	router.GET("/api/ruleset/5e/conditions/:index", condition.GetConditionByIndex)
	// Classes
	router.GET("/api/ruleset/5e/classes", class.GetAllClasses5e)
	router.GET("/api/ruleset/5e/classes/:index", class.GetClassByIndex5e)
	// Subraces
	router.GET("/api/ruleset/5e/subraces", subrace.GetAllSubraces5e)
	router.GET("/api/ruleset/5e/subraces/:index", subrace.GetSubraceByIndex5e)
}
