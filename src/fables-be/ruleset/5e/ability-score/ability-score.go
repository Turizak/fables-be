package abilityscore

import (
	"net/http"

	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type AbilityScore struct {
	Index       string         `gorm:"primaryKey;type:varchar" json:"index"`
	AbilityName string         `gorm:"type:varchar" json:"ability_name"`
	FullName    string         `gorm:"type:varchar" json:"full_name"`
	Description datatypes.JSON `gorm:"type:jsonb" json:"description"`
	URL         string         `gorm:"type:varchar" json:"url"`
}

// GetAllAbilityScores5e retrieves all ability scores for 5e
func GetAllAbilityScores5e(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	abilityScores, err := GetAllAbilityScores5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve ability scores. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	responseAbilityScores := mapAbilityScoresToResponse(abilityScores)
	utilities.ResponseMessage(c, "Ability scores retrieved successfully.", http.StatusOK, gin.H{"abilityScores": responseAbilityScores})
}

// GetAbilityScoreByIndex5e retrieves a specific ability score by index for 5e
func GetAbilityScoreByIndex5e(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	index := c.Param("index")
	abilityScore, err := GetAbilityScoreByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve ability score. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	responseAbilityScore := mapAbilityScoreToResponse(*abilityScore)
	utilities.ResponseMessage(c, "Ability score retrieved successfully.", http.StatusOK, gin.H{"abilityScore": responseAbilityScore})
}

// mapAbilityScoresToResponse maps a slice of AbilityScore structs to their response form
func mapAbilityScoresToResponse(abilityScores []AbilityScore) []AbilityScore {
	responseAbilityScores := make([]AbilityScore, len(abilityScores))
	for i, abilityScore := range abilityScores {
		responseAbilityScores[i] = mapAbilityScoreToResponse(abilityScore)
	}
	return responseAbilityScores
}

// mapAbilityScoreToResponse maps a single AbilityScore to its response form
func mapAbilityScoreToResponse(abilityScore AbilityScore) AbilityScore {
	return AbilityScore{
		Index:       abilityScore.Index,
		AbilityName: abilityScore.AbilityName,
		FullName:    abilityScore.FullName,
		Description: abilityScore.Description,
		URL:         abilityScore.URL,
	}
}
