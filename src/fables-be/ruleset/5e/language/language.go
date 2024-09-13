package language

import (
	"fmt"
	"net/http"

	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
	"gorm.io/datatypes"
)

type Language struct {
	Index           string         `gorm:"primaryKey;type:varchar" json:"index"`
	Name            string         `gorm:"type:varchar" json:"name"`
	Description     string         `gorm:"type:text" json:"description"`
	Type            string         `gorm:"type:varchar" json:"type"`
	TypicalSpeakers datatypes.JSON `gorm:"type:jsonb" json:"typical_speakers"`
	Script          string         `gorm:"type:varchar" json:"script"`
	URL             string         `gorm:"type:varchar" json:"url"`
}

func GetAllLanguages5e(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	_, validToken := utilities.ValidateAuthenticationToken(c, authHeader)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	languages, err := GetAllLanguages5eDB()
	fmt.Println(err)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve languages. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseLanguages := []Language{}
	for _, language := range languages {
		responseLanguage := Language{
			Index:           language.Index,
			Name:            language.Name,
			Description:     language.Description,
			Type:            language.Type,
			TypicalSpeakers: language.TypicalSpeakers,
			Script:          language.Script,
			URL:             language.URL,
		}
		responseLanguages = append(responseLanguages, responseLanguage)
	}
	utilities.ResponseMessage(c, "Languages retrieved successfully.", http.StatusOK, gin.H{"languages": responseLanguages})
}

func GetLanguageByIndex5e(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")
	if authHeader == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	_, validToken := utilities.ValidateAuthenticationToken(c, authHeader)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	index := c.Param("index")
	language, err := GetLanguageByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve language. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseLanguage := Language{
		Index:           language.Index,
		Name:            language.Name,
		Description:     language.Description,
		Type:            language.Type,
		TypicalSpeakers: language.TypicalSpeakers,
		Script:          language.Script,
		URL:             language.URL,
	}
	utilities.ResponseMessage(c, "Language retrieved successfully.", http.StatusOK, gin.H{"language": responseLanguage})
}
