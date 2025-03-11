package race

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
	AbilityBonuses             datatypes.JSON `gorm:"column:ability_bonuses" json:"ability_bonuses"`
	AbilityBonusOptions        datatypes.JSON `gorm:"column:ability_bonus_options" json:"ability_bonus_options"`
	Alignment                  string         `gorm:"column:alignment" json:"alignment"`
	Age                        string         `gorm:"column:age" json:"age"`
	Size                       string         `gorm:"column:size" json:"size"`
	SizeDescription            string         `gorm:"column:size_description" json:"size_description"`
	StartingProficiencies      datatypes.JSON `gorm:"column:starting_proficiencies" json:"starting_proficiencies"`
	StartingProficiencyOptions datatypes.JSON `gorm:"column:starting_proficiency_options" json:"starting_proficiency_options"`
	Languages                  datatypes.JSON `gorm:"column:languages" json:"languages"`
	LanguageDesc               string         `gorm:"column:language_desc" json:"language_desc"`
	LanguageOptions            datatypes.JSON `gorm:"column:language_options" json:"language_options"`
	Traits                     datatypes.JSON `gorm:"column:traits" json:"traits"`
	Subraces                   datatypes.JSON `gorm:"column:subraces" json:"subraces"`
	URL                        string         `gorm:"column:url" json:"url"`
}

// GetAllRaces5e retrieves all races for 5e
func GetAllRaces5e(c *gin.Context) {
	races, err := GetAllRaces5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve races. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	if len(races) == 0 {
		utilities.ResponseMessage(c, "No races found.", http.StatusOK, nil)
		return
	}

	utilities.ResponseMessage(c, "Races retrieved successfully.", http.StatusOK, gin.H{"races": mapRacesToResponse(races)})
}

// GetRaceByIndex5e retrieves a specific race by index for 5e
func GetRaceByIndex5e(c *gin.Context) {
	index := c.Param("index")
	race, err := GetRaceByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve race. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Race retrieved successfully.", http.StatusOK, gin.H{"race": mapRaceToResponse(*race)})
}

// Helper functions to map races to their response formats
func mapRacesToResponse(races []Race) []Race {
	response := make([]Race, len(races))
	for i, race := range races {
		response[i] = mapRaceToResponse(race)
	}
	return response
}

func mapRaceToResponse(race Race) Race {
	return Race{
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
}
