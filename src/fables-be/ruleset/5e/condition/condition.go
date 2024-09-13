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

func GetAllConditionsWithDescriptions(c *gin.Context) {
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
	conditions, err := GetAllConditionsWithDescriptionsDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve conditions. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseConditions := []Condition{}
	for _, condition := range conditions {
		responseCondition := Condition{
			Index:        condition.Index,
			Name:         condition.Name,
			Description:  condition.Description,
			URL:          condition.URL,
			Descriptions: condition.Descriptions,
		}
		responseConditions = append(responseConditions, responseCondition)
	}
	utilities.ResponseMessage(c, "Conditions retrieved successfully.", http.StatusOK, gin.H{"conditions": responseConditions})
}

func GetConditionByIndex(c *gin.Context) {
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
	condition, err := GetConditionByIndexDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve condition. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseCondition := Condition{
		Index:        condition.Index,
		Name:         condition.Name,
		Description:  condition.Description,
		URL:          condition.URL,
		Descriptions: condition.Descriptions,
	}
	utilities.ResponseMessage(c, "Condition retrieved successfully.", http.StatusOK, gin.H{"condition": responseCondition})
}
