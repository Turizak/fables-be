package condition

import (
	"net/http"

	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

type Condition struct {
	Index        string                 `gorm:"primaryKey;column:index" json:"index"`
	Name         string                 `gorm:"column:name" json:"name"`
	Description  string                 `gorm:"column:description" json:"description"`
	URL          string                 `gorm:"column:url" json:"url"`
	Descriptions []ConditionDescription `gorm:"foreignKey:ConditionIndex;references:Index" json:"descriptions"`
}

type ConditionDescription struct {
	ConditionIndex string `gorm:"column:index" json:"index"`
	Description    string `gorm:"column:description" json:"description"`
}

// GetAllConditionsWithDescriptions retrieves all conditions with descriptions
func GetAllConditionsWithDescriptions(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	conditions, err := GetAllConditionsWithDescriptionsDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve conditions. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Conditions retrieved successfully.", http.StatusOK, gin.H{"conditions": mapConditionsToResponse(conditions)})
}

// GetConditionByIndex retrieves a specific condition by index
func GetConditionByIndex(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	index := c.Param("index")
	condition, err := GetConditionByIndexDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve condition. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Condition retrieved successfully.", http.StatusOK, gin.H{"condition": mapConditionToResponse(condition)})
}

// Helper functions to map conditions to their response formats
func mapConditionsToResponse(conditions []Condition) []Condition {
	response := make([]Condition, len(conditions))
	for i, condition := range conditions {
		response[i] = mapConditionToResponse(condition)
	}
	return response
}

func mapConditionToResponse(condition Condition) Condition {
	return Condition{
		Index:        condition.Index,
		Name:         condition.Name,
		Description:  condition.Description,
		URL:          condition.URL,
		Descriptions: condition.Descriptions,
	}
}
