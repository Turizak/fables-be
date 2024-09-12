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

func GetAllTraits5e(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	_, validToken := utilities.ValidateAuthenticationToken(c, authHeader)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	traits, err := GetAllTraits5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve traits. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseTraits := []Trait{}
	for _, trait := range traits {
		responseTrait := Trait{
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
		responseTraits = append(responseTraits, responseTrait)
	}
	utilities.ResponseMessage(c, "Traits retrieved successfully.", http.StatusOK, gin.H{"traits": responseTraits})
}

func GetTraitByIndex5e(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	_, validToken := utilities.ValidateAuthenticationToken(c, authHeader)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	traitIndex := c.Param("index")
	trait, err := GetTraitByIndex5eDB(traitIndex)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve trait. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseTrait := Trait{
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
	utilities.ResponseMessage(c, "Trait retrieved successfully.", http.StatusOK, gin.H{"trait": responseTrait})
}
