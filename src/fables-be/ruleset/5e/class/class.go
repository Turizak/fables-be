package class

import (
	"net/http"

	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type Class struct {
	Index                    string                    `json:"index" gorm:"primaryKey"`
	Name                     string                    `json:"name"`
	HitDie                   int                       `json:"hit_die"`
	ClassLevels              string                    `json:"class_levels"`
	MultiClassing            datatypes.JSON            `json:"multi_classing" gorm:"type:jsonb"`
	Subclasses               datatypes.JSON            `json:"subclasses" gorm:"type:jsonb"`
	Spellcasting             datatypes.JSON            `json:"spellcasting" gorm:"type:jsonb"`
	Spells                   string                    `json:"spells"`
	URL                      string                    `json:"url"`
	ProficiencyChoices       []ProficiencyChoice       `json:"proficiency_choices" gorm:"foreignKey:ClassIndex;references:Index"`
	Proficiencies            []Proficiency             `json:"proficiencies" gorm:"foreignKey:ClassIndex;references:Index"`
	SavingThrows             []SavingThrow             `json:"saving_throws" gorm:"foreignKey:ClassIndex;references:Index"`
	StartingEquipment        []StartingEquipment       `json:"starting_equipment" gorm:"foreignKey:ClassIndex;references:Index"`
	StartingEquipmentOptions []StartingEquipmentOption `json:"starting_equipment_options" gorm:"foreignKey:ClassIndex;references:Index"`
}

type ClassDetailResponse struct {
	Index              string              `json:"index"`
	Name               string              `json:"name"`
	HitDie             int                 `json:"hit_die"`
	ClassLevels        string              `json:"class_levels"`
	MultiClassing      datatypes.JSON      `json:"multi_classing"`
	Subclasses         datatypes.JSON      `json:"subclasses"`
	Spellcasting       datatypes.JSON      `json:"spellcasting"`
	Spells             string              `json:"spells"`
	URL                string              `json:"url"`
	ProficiencyChoices []ProficiencyChoice `json:"proficiency_choices"`
	Proficiencies      []Proficiency       `json:"proficiencies"`
	SavingThrows       []SavingThrow       `json:"saving_throws"`
	// StartingEquipment        []StartingEquipment       `json:"starting_equipment"`
	// StartingEquipmentOptions []StartingEquipmentOption `json:"starting_equipment_options"`
}

type ClassResponse struct {
	Index         string         `json:"index"`
	Name          string         `json:"name"`
	HitDie        int            `json:"hit_die"`
	ClassLevels   string         `json:"class_levels"`
	MultiClassing datatypes.JSON `json:"multi_classing"`
	Subclasses    datatypes.JSON `json:"subclasses"`
	Spellcasting  datatypes.JSON `json:"spellcasting"`
	Spells        string         `json:"spells"`
	URL           string         `json:"url"`
}

type ProficiencyChoice struct {
	ClassIndex         string         `json:"class_index" gorm:"primaryKey"`
	Description        string         `json:"description"`
	ChooseNumber       int            `json:"choose_number"`
	ProficiencyType    string         `json:"proficiency_type"`
	ProficiencyOptions datatypes.JSON `json:"proficiency_options" gorm:"type:jsonb"`
}

type Proficiency struct {
	ClassIndex       string `json:"class_index" gorm:"primaryKey"`
	ProficiencyIndex string `json:"proficiency_index"`
	ProficiencyName  string `json:"proficiency_name"`
	ProficiencyURL   string `json:"proficiency_url"`
}

type SavingThrow struct {
	ClassIndex       string `json:"class_index" gorm:"primaryKey"`
	SavingThrowIndex string `json:"saving_throw_index"`
	SavingThrowName  string `json:"saving_throw_name"`
	SavingThrowURL   string `json:"saving_throw_url"`
}

type StartingEquipment struct {
	ClassIndex     string `json:"class_index" gorm:"primaryKey"`
	EquipmentIndex string `json:"equipment_index"`
	EquipmentName  string `json:"equipment_name"`
	EquipmentURL   string `json:"equipment_url"`
	Quantity       int    `json:"quantity"`
}

type StartingEquipmentOption struct {
	ClassIndex       string         `json:"class_index" gorm:"primaryKey"`
	Description      string         `json:"description"`
	Choose           int            `json:"choose"`
	EquipmentType    string         `json:"equipment_type"`
	EquipmentOptions datatypes.JSON `json:"equipment_options" gorm:"type:jsonb"`
}

func GetAllClasses5e(c *gin.Context) {
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

	// Get the "details" query parameter
	details := c.Query("details")

	var (
		classes []Class
		err     error
	)

	// If details=true, use GetAllClassesWithDetails function
	if details == "true" {
		classes, err = GetAllClassesWithDetails5eDB()
	} else {
		classes, err = GetAllClasses5eDB()
	}

	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve classes. Please try again.", http.StatusBadRequest, nil)
		return
	}

	if details == "true" {
		// Create slice of ClassDetailResponse
		responseClasses := []ClassDetailResponse{}
		for _, class := range classes {
			responseClass := ClassDetailResponse{
				Index:              class.Index,
				Name:               class.Name,
				HitDie:             class.HitDie,
				ClassLevels:        class.ClassLevels,
				MultiClassing:      class.MultiClassing,
				Subclasses:         class.Subclasses,
				Spellcasting:       class.Spellcasting,
				Spells:             class.Spells,
				URL:                class.URL,
				ProficiencyChoices: class.ProficiencyChoices,
				Proficiencies:      class.Proficiencies,
				SavingThrows:       class.SavingThrows,
			}
			responseClasses = append(responseClasses, responseClass)
		}
		utilities.ResponseMessage(c, "Classes retrieved successfully.", http.StatusOK, gin.H{"classes": responseClasses})
	} else {
		// Create slice of ClassResponse
		responseClasses := []ClassResponse{}
		for _, class := range classes {
			responseClass := ClassResponse{
				Index:         class.Index,
				Name:          class.Name,
				HitDie:        class.HitDie,
				ClassLevels:   class.ClassLevels,
				MultiClassing: class.MultiClassing,
				Subclasses:    class.Subclasses,
				Spellcasting:  class.Spellcasting,
				Spells:        class.Spells,
				URL:           class.URL,
			}
			responseClasses = append(responseClasses, responseClass)
		}
		utilities.ResponseMessage(c, "Classes retrieved successfully.", http.StatusOK, gin.H{"classes": responseClasses})
	}
}

func GetClassByIndex5e(c *gin.Context) {
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

	// Get the "details" query parameter
	details := c.Query("details")

	var (
		class *Class
		err   error
	)

	// If details=true, use GetClassWithDetailsByIndex function
	if details == "true" {
		class, err = GetClassByIndexWithDetails5eDB(index) // Assuming the function exists
	} else {
		class, err = GetClassByIndex5eDB(index)
	}

	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve class. Please try again.", http.StatusBadRequest, nil)
		return
	}

	if details == "true" {
		responseClass := ClassDetailResponse{
			Index:              class.Index,
			Name:               class.Name,
			HitDie:             class.HitDie,
			ClassLevels:        class.ClassLevels,
			MultiClassing:      class.MultiClassing,
			Subclasses:         class.Subclasses,
			Spellcasting:       class.Spellcasting,
			Spells:             class.Spells,
			URL:                class.URL,
			ProficiencyChoices: class.ProficiencyChoices,
			Proficiencies:      class.Proficiencies,
			SavingThrows:       class.SavingThrows,
		}

		utilities.ResponseMessage(c, "Class retrieved successfully.", http.StatusOK, gin.H{"class": responseClass})

	} else {
		responseClass := ClassResponse{
			Index:         class.Index,
			Name:          class.Name,
			HitDie:        class.HitDie,
			ClassLevels:   class.ClassLevels,
			MultiClassing: class.MultiClassing,
			Subclasses:    class.Subclasses,
			Spellcasting:  class.Spellcasting,
			Spells:        class.Spells,
			URL:           class.URL,
		}

		utilities.ResponseMessage(c, "Class retrieved successfully.", http.StatusOK, gin.H{"class": responseClass})
	}
}
