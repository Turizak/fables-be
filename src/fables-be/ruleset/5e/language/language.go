package language

import (
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

// GetAllLanguages5e retrieves all languages for 5e
func GetAllLanguages5e(c *gin.Context) {
	languages, err := GetAllLanguages5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve languages. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Languages retrieved successfully.", http.StatusOK, gin.H{"languages": mapLanguagesToResponse(languages)})
}

// GetLanguageByIndex5e retrieves a specific language by index for 5e
func GetLanguageByIndex5e(c *gin.Context) {
	index := c.Param("index")
	language, err := GetLanguageByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve language. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Language retrieved successfully.", http.StatusOK, gin.H{"language": mapLanguageToResponse(*language)})
}

// Helper functions to map languages to their response formats
func mapLanguagesToResponse(languages []Language) []Language {
	response := make([]Language, len(languages))
	for i, language := range languages {
		response[i] = mapLanguageToResponse(language)
	}
	return response
}

func mapLanguageToResponse(language Language) Language {
	return Language{
		Index:           language.Index,
		Name:            language.Name,
		Description:     language.Description,
		Type:            language.Type,
		TypicalSpeakers: language.TypicalSpeakers,
		Script:          language.Script,
		URL:             language.URL,
	}
}
