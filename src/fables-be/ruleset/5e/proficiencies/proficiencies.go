package proficiencies

import (
	"net/http"

	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type Proficiencie struct {
	Index     string         `gorm:"primaryKey;type:varchar" json:"index"`
	Type      string         `gorm:"type:varchar" json:"type"`
	Name      string         `gorm:"type:varchar" json:"name"`
	Classes   datatypes.JSON `gorm:"type:jsonb" json:"classes"`
	Races     datatypes.JSON `gorm:"type:jsonb" json:"races"`
	URL       string         `gorm:"type:varchar" json:"url"`
	Reference datatypes.JSON `gorm:"type:jsonb" json:"reference"`
}

// GetAllProficiencies5e retrieves all proficiencies for 5e
func GetAllProficiencies5e(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	proficiencies, err := GetAllProficiencies5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve proficiencies. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Proficiencies retrieved successfully.", http.StatusOK, gin.H{"proficiencies": mapProficienciesToResponse(proficiencies)})
}

// GetProficiencyByIndex5e retrieves a specific proficiency by index for 5e
func GetProficiencyByIndex5e(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	index := c.Param("index")
	proficiency, err := GetProficiencyByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve proficiency. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Proficiency retrieved successfully.", http.StatusOK, gin.H{"proficiency": mapProficiencyToResponse(*proficiency)})
}

// Helper functions to map proficiencies to their response formats
func mapProficienciesToResponse(proficiencies []Proficiencie) []Proficiencie {
	response := make([]Proficiencie, len(proficiencies))
	for i, proficiency := range proficiencies {
		response[i] = mapProficiencyToResponse(proficiency)
	}
	return response
}

func mapProficiencyToResponse(proficiency Proficiencie) Proficiencie {
	return Proficiencie{
		Index:     proficiency.Index,
		Type:      proficiency.Type,
		Name:      proficiency.Name,
		Classes:   proficiency.Classes,
		Races:     proficiency.Races,
		URL:       proficiency.URL,
		Reference: proficiency.Reference,
	}
}
