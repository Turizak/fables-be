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

func GetAllAbilityScores5e(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	_, validToken := utilities.ValidateAuthenticationToken(c, authToken)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	abilityScores, err := GetAllAbilityScores5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve ability scores. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseAbilityScores := []AbilityScore{}
	for _, abilityScore := range abilityScores {
		responseAbilityScore := AbilityScore{
			Index:       abilityScore.Index,
			AbilityName: abilityScore.AbilityName,
			FullName:    abilityScore.FullName,
			Description: abilityScore.Description,
			URL:         abilityScore.URL,
		}
		responseAbilityScores = append(responseAbilityScores, responseAbilityScore)
	}
	utilities.ResponseMessage(c, "Ability scores retrieved successfully.", http.StatusOK, gin.H{"abilityScores": responseAbilityScores})
}

func GetAbilityScoreByIndex5e(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	_, validToken := utilities.ValidateAuthenticationToken(c, authToken)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	index := c.Param("index")
	abilityScore, err := GetAbilityScoreByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve ability score. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseAbilityScore := AbilityScore{
		Index:       abilityScore.Index,
		AbilityName: abilityScore.AbilityName,
		FullName:    abilityScore.FullName,
		Description: abilityScore.Description,
		URL:         abilityScore.URL,
	}
	utilities.ResponseMessage(c, "Ability score retrieved successfully.", http.StatusOK, gin.H{"abilityScore": responseAbilityScore})
}
