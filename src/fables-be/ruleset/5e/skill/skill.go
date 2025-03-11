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

// GetAllSkills5e retrieves all skills for 5e
func GetAllSkills5e(c *gin.Context) {
	skills, err := GetAllSkills5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve skills. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Skills retrieved successfully.", http.StatusOK, gin.H{"skills": mapSkillsToResponse(skills)})
}

// GetSkillByIndex5e retrieves a specific skill by index for 5e
func GetSkillByIndex5e(c *gin.Context) {
	index := c.Param("index")
	skill, err := GetSkillByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve skill. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Skill retrieved successfully.", http.StatusOK, gin.H{"skill": mapSkillToResponse(*skill)})
}

// Helper functions to map skills to their response formats
func mapSkillsToResponse(skills []Skill) []Skill {
	response := make([]Skill, len(skills))
	for i, skill := range skills {
		response[i] = mapSkillToResponse(skill)
	}
	return response
}

func mapSkillToResponse(skill Skill) Skill {
	return Skill{
		Index:        skill.Index,
		Name:         skill.Name,
		Description:  skill.Description,
		AbilityScore: skill.AbilityScore,
		URL:          skill.URL,
	}
}
