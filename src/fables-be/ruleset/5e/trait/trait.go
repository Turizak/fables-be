package trait

import (
	"net/http"

	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type Trait struct {
	Index              string         `gorm:"primaryKey;type:varchar" json:"index"`
	Name               string         `gorm:"type:varchar;not null" json:"name"`
	Description        datatypes.JSON `gorm:"type:jsonb;not null" json:"description"`
	Races              datatypes.JSON `gorm:"type:jsonb;not null" json:"races"`
	Subraces           datatypes.JSON `gorm:"type:jsonb;not null" json:"subraces"`
	Proficiencies      datatypes.JSON `gorm:"type:jsonb" json:"proficiencies"`
	ProficiencyChoices datatypes.JSON `gorm:"type:jsonb" json:"proficiency_choices"`
	TraitSpecific      datatypes.JSON `gorm:"type:jsonb" json:"trait_specific"`
	URL                string         `gorm:"type:varchar;not null" json:"url"`
}

// GetAllTraits5e retrieves all traits for 5e
func GetAllTraits5e(c *gin.Context) {
	traits, err := GetAllTraits5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve traits. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Traits retrieved successfully.", http.StatusOK, gin.H{"traits": mapTraitsToResponse(traits)})
}

// GetTraitByIndex5e retrieves a specific trait by index for 5e
func GetTraitByIndex5e(c *gin.Context) {
	traitIndex := c.Param("index")
	trait, err := GetTraitByIndex5eDB(traitIndex)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve trait. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Trait retrieved successfully.", http.StatusOK, gin.H{"trait": mapTraitToResponse(*trait)})
}

// Helper functions to map traits to their response formats
func mapTraitsToResponse(traits []Trait) []Trait {
	response := make([]Trait, len(traits))
	for i, trait := range traits {
		response[i] = mapTraitToResponse(trait)
	}
	return response
}

func mapTraitToResponse(trait Trait) Trait {
	return Trait{
		Index:              trait.Index,
		Name:               trait.Name,
		Description:        trait.Description,
		Races:              trait.Races,
		Subraces:           trait.Subraces,
		Proficiencies:      trait.Proficiencies,
		ProficiencyChoices: trait.ProficiencyChoices,
		TraitSpecific:      trait.TraitSpecific,
		URL:                trait.URL,
	}
}
