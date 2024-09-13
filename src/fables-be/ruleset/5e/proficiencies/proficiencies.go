package proficiencies

import (
	"net/http"

	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type Proficiency struct {
	Index     string         `gorm:"primaryKey;type:varchar" json:"index"`
	Type      string         `gorm:"type:varchar" json:"type"`
	Name      string         `gorm:"type:varchar" json:"name"`
	Classes   datatypes.JSON `gorm:"type:jsonb" json:"classes"`
	Races     datatypes.JSON `gorm:"type:jsonb" json:"races"`
	URL       string         `gorm:"type:varchar" json:"url"`
	Reference datatypes.JSON `gorm:"type:jsonb" json:"reference"`
}

func GetAllProficiencies5e(c *gin.Context) {
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
	proficiencies, err := GetAllProficiencies5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve proficiencies. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseProficiencies := []Proficiency{}
	for _, proficiency := range proficiencies {
		responseProficiency := Proficiency{
			Index:     proficiency.Index,
			Type:      proficiency.Type,
			Name:      proficiency.Name,
			Classes:   proficiency.Classes,
			Races:     proficiency.Races,
			URL:       proficiency.URL,
			Reference: proficiency.Reference,
		}
		responseProficiencies = append(responseProficiencies, responseProficiency)
	}
	utilities.ResponseMessage(c, "Proficiencies retrieved successfully.", http.StatusOK, gin.H{"proficiencies": responseProficiencies})
}

func GetProficiencyByIndex5e(c *gin.Context) {
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
	proficiency, err := GetProficiencyByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve proficiency. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseProficiency := Proficiency{
		Index:     proficiency.Index,
		Type:      proficiency.Type,
		Name:      proficiency.Name,
		Classes:   proficiency.Classes,
		Races:     proficiency.Races,
		URL:       proficiency.URL,
		Reference: proficiency.Reference,
	}
	utilities.ResponseMessage(c, "Proficiency retrieved successfully.", http.StatusOK, gin.H{"proficiency": responseProficiency})
}
