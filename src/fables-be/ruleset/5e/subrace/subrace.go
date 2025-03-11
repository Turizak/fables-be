package subrace

import (
	"net/http"

	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type Subrace struct {
	Index                 string         `gorm:"primaryKey;column:index" json:"index"`
	Name                  string         `gorm:"column:name" json:"name"`
	Race                  datatypes.JSON `gorm:"column:race" json:"race"`
	Description           string         `gorm:"column:description" json:"description"`
	AbilityBonuses        datatypes.JSON `gorm:"column:ability_bonuses" json:"ability_bonuses"`
	StartingProficiencies datatypes.JSON `gorm:"column:starting_proficiencies" json:"starting_proficiencies"`
	Languages             datatypes.JSON `gorm:"column:languages" json:"languages"`
	LanguageOptions       datatypes.JSON `gorm:"column:language_options" json:"language_options"`
	RacialTraits          datatypes.JSON `gorm:"column:racial_traits" json:"racial_traits"`
	URL                   string         `gorm:"column:url" json:"url"`
}

// GetAllSubraces5e retrieves all subraces for 5e
func GetAllSubraces5e(c *gin.Context) {
	subraces, err := GetAllSubraces5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve subraces. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Subraces retrieved successfully.", http.StatusOK, gin.H{"subraces": mapSubracesToResponse(subraces)})
}

// GetSubraceByIndex5e retrieves a specific subrace by index for 5e
func GetSubraceByIndex5e(c *gin.Context) {
	index := c.Param("index")
	subrace, err := GetSubraceByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve subrace. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Subrace retrieved successfully.", http.StatusOK, gin.H{"subrace": mapSubraceToResponse(*subrace)})
}

// Helper functions to map subraces to their response formats
func mapSubracesToResponse(subraces []Subrace) []Subrace {
	response := make([]Subrace, len(subraces))
	for i, subrace := range subraces {
		response[i] = mapSubraceToResponse(subrace)
	}
	return response
}

func mapSubraceToResponse(subrace Subrace) Subrace {
	return Subrace{
		Index:                 subrace.Index,
		Name:                  subrace.Name,
		Race:                  subrace.Race,
		Description:           subrace.Description,
		AbilityBonuses:        subrace.AbilityBonuses,
		StartingProficiencies: subrace.StartingProficiencies,
		Languages:             subrace.Languages,
		LanguageOptions:       subrace.LanguageOptions,
		RacialTraits:          subrace.RacialTraits,
		URL:                   subrace.URL,
	}
}
