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

// GetAllDamageTypes5e retrieves all damage types for 5e
func GetAllDamageTypes5e(c *gin.Context) {
	damageTypes, err := GetAllDamageTypes5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve damage types. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Damage types retrieved successfully.", http.StatusOK, gin.H{"damageTypes": mapDamageTypesToResponse(damageTypes)})
}

// GetDamageTypeByIndex5e retrieves a specific damage type by index for 5e
func GetDamageTypeByIndex5e(c *gin.Context) {
	index := c.Param("index")
	damageType, err := GetDamageTypeByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve damage type. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Damage type retrieved successfully.", http.StatusOK, gin.H{"damageType": mapDamageTypeToResponse(*damageType)})
}

// Helper functions to map damage types to their response formats
func mapDamageTypesToResponse(damageTypes []DamageType) []DamageType {
	response := make([]DamageType, len(damageTypes))
	for i, damageType := range damageTypes {
		response[i] = mapDamageTypeToResponse(damageType)
	}
	return response
}

func mapDamageTypeToResponse(damageType DamageType) DamageType {
	return DamageType{
		Index:       damageType.Index,
		Name:        damageType.Name,
		Description: damageType.Description,
		URL:         damageType.URL,
	}
}
