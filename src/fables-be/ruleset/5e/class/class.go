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

// GetAllClasses5e retrieves all classes for 5e
func GetAllClasses5e(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	details := c.Query("details") == "true"

	var (
		classes []Class
		err     error
	)

	if details {
		classes, err = GetAllClassesWithDetails5eDB()
	} else {
		classes, err = GetAllClasses5eDB()
	}

	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve classes. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	if details {
		utilities.ResponseMessage(c, "Classes retrieved successfully.", http.StatusOK, gin.H{"classes": mapClassesToDetailResponse(classes)})
	} else {
		utilities.ResponseMessage(c, "Classes retrieved successfully.", http.StatusOK, gin.H{"classes": mapClassesToResponse(classes)})
	}
}

// GetClassByIndex5e retrieves a specific class by index for 5e
func GetClassByIndex5e(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	index := c.Param("index")
	details := c.Query("details") == "true"

	var (
		class *Class
		err   error
	)

	if details {
		class, err = GetClassByIndexWithDetails5eDB(index)
	} else {
		class, err = GetClassByIndex5eDB(index)
	}

	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve class. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	if details {
		utilities.ResponseMessage(c, "Class retrieved successfully.", http.StatusOK, gin.H{"class": mapClassToDetailResponse(*class)})
	} else {
		utilities.ResponseMessage(c, "Class retrieved successfully.", http.StatusOK, gin.H{"class": mapClassToResponse(*class)})
	}
}

// Helper functions to map classes to their respective response formats
func mapClassesToDetailResponse(classes []Class) []ClassDetailResponse {
	response := make([]ClassDetailResponse, len(classes))
	for i, class := range classes {
		response[i] = mapClassToDetailResponse(class)
	}
	return response
}

func mapClassesToResponse(classes []Class) []ClassResponse {
	response := make([]ClassResponse, len(classes))
	for i, class := range classes {
		response[i] = mapClassToResponse(class)
	}
	return response
}

func mapClassToDetailResponse(class Class) ClassDetailResponse {
	return ClassDetailResponse{
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
}

func mapClassToResponse(class Class) ClassResponse {
	return ClassResponse{
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
}
