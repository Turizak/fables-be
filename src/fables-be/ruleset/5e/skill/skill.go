package skill

import (
	"net/http"

	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type Skill struct {
	Index        string         `gorm:"primaryKey;type:text" json:"index"`
	Name         string         `gorm:"type:text;not null" json:"name"`
	Description  datatypes.JSON `gorm:"type:jsonb;not null" json:"description"`
	AbilityScore datatypes.JSON `gorm:"type:jsonb;not null" json:"ability_score"`
	URL          string         `gorm:"type:text;not null" json:"url"`
}

func GetAllSkills5e(c *gin.Context) {
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
	skills, err := GetAllSkills5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve skills. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseSkills := []Skill{}
	for _, skill := range skills {
		responseSkill := Skill{
			Index:        skill.Index,
			Name:         skill.Name,
			Description:  skill.Description,
			AbilityScore: skill.AbilityScore,
			URL:          skill.URL,
		}
		responseSkills = append(responseSkills, responseSkill)
	}
	utilities.ResponseMessage(c, "Skills retrieved successfully.", http.StatusOK, gin.H{"skills": responseSkills})
}

func GetSkillByIndex5e(c *gin.Context) {
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
	skill, err := GetSkillByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve skill. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseSkill := Skill{
		Index:        skill.Index,
		Name:         skill.Name,
		Description:  skill.Description,
		AbilityScore: skill.AbilityScore,
		URL:          skill.URL,
	}
	utilities.ResponseMessage(c, "Skill retrieved successfully.", http.StatusOK, gin.H{"skill": responseSkill})
}
