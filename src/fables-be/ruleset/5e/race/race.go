package ruleset

import (
	"net/http"

	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type Race struct {
	Index                      string         `gorm:"primaryKey;column:index" json:"index"`
	Name                       string         `gorm:"column:name" json:"name"`
	Speed                      int            `gorm:"column:speed" json:"speed"`
	AbilityBonuses             datatypes.JSON `gorm:"column:ability_bonuses" json:"ability_bonuses"`             // JSONB field
	AbilityBonusOptions        datatypes.JSON `gorm:"column:ability_bonus_options" json:"ability_bonus_options"` // JSONB field
	Alignment                  string         `gorm:"column:alignment" json:"alignment"`
	Age                        string         `gorm:"column:age" json:"age"`
	Size                       string         `gorm:"column:size" json:"size"`
	SizeDescription            string         `gorm:"column:size_description" json:"size_description"`
	StartingProficiencies      datatypes.JSON `gorm:"column:starting_proficiencies" json:"starting_proficiencies"`             // JSONB field
	StartingProficiencyOptions datatypes.JSON `gorm:"column:starting_proficiency_options" json:"starting_proficiency_options"` // JSONB field
	Languages                  datatypes.JSON `gorm:"column:languages" json:"languages"`                                       // JSONB field
	LanguageDesc               string         `gorm:"column:language_desc" json:"language_desc"`
	LanguageOptions            datatypes.JSON `gorm:"column:language_options" json:"language_options"` // JSONB field
	Traits                     datatypes.JSON `gorm:"column:traits" json:"traits"`                     // JSONB field
	Subraces                   datatypes.JSON `gorm:"column:subraces" json:"subraces"`                 // JSONB field
	URL                        string         `gorm:"column:url" json:"url"`
}

func GetAllRaces5e(c *gin.Context) {
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
	races, err := GetAllRaces5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve campaigns. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseRaces := []Race{}
	for _, race := range races {
		responseRace := Race{
			Index:                      race.Index,
			Name:                       race.Name,
			Speed:                      race.Speed,
			AbilityBonuses:             race.AbilityBonuses,
			AbilityBonusOptions:        race.AbilityBonusOptions,
			Alignment:                  race.Alignment,
			Age:                        race.Age,
			Size:                       race.Size,
			SizeDescription:            race.SizeDescription,
			StartingProficiencies:      race.StartingProficiencies,
			StartingProficiencyOptions: race.StartingProficiencyOptions,
			Languages:                  race.Languages,
			LanguageDesc:               race.LanguageDesc,
			LanguageOptions:            race.LanguageOptions,
			Traits:                     race.Traits,
			Subraces:                   race.Subraces,
			URL:                        race.URL,
		}
		responseRaces = append(responseRaces, responseRace)
	}
	if len(responseRaces) == 0 {
		utilities.ResponseMessage(c, "No races found.", http.StatusOK, nil)
		return
	}
	utilities.ResponseMessage(c, "Races retrieved successfully.", http.StatusOK, gin.H{"races": responseRaces})
}

func GetRaceByIndex5e(c *gin.Context) {
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
	race, err := GetRaceByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve race. Please try again.", http.StatusBadRequest, nil)
		return
	}
	utilities.ResponseMessage(c, "Race retrieved successfully.", http.StatusOK, gin.H{"race": race})
}
