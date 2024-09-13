package damagetype

import (
	"net/http"

	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

type DamageType struct {
	Index       string `gorm:"primaryKey;type:varchar" json:"index"`
	Name        string `gorm:"type:varchar" json:"name"`
	Description string `gorm:"type:text" json:"description"`
	URL         string `gorm:"type:varchar" json:"url"`
}

func GetAllDamageTypes5e(c *gin.Context) {
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
	damageTypes, err := GetAllDamageTypes5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve damage types. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseDamageTypes := []DamageType{}
	for _, damageType := range damageTypes {
		responseDamageType := DamageType{
			Index:       damageType.Index,
			Name:        damageType.Name,
			Description: damageType.Description,
			URL:         damageType.URL,
		}
		responseDamageTypes = append(responseDamageTypes, responseDamageType)
	}
	utilities.ResponseMessage(c, "Damage types retrieved successfully.", http.StatusOK, gin.H{"damageTypes": responseDamageTypes})
}

func GetDamageTypeByIndex5e(c *gin.Context) {
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
	damageType, err := GetDamageTypeByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve damage type. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseDamageType := DamageType{
		Index:       damageType.Index,
		Name:        damageType.Name,
		Description: damageType.Description,
		URL:         damageType.URL,
	}
	utilities.ResponseMessage(c, "Damage type retrieved successfully.", http.StatusOK, gin.H{"damageType": responseDamageType})
}
