package alignment

import (
	"net/http"

	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

type Alignment struct {
	Index         string `gorm:"primaryKey;column:index" json:"index"`
	AlignmentName string `gorm:"column:alignment_name" json:"alignment_name"`
	Abbreviation  string `gorm:"column:abbreviation" json:"abbreviation"`
	Description   string `gorm:"column:description" json:"description"`
	URL           string `gorm:"column:url" json:"url"`
}

func GetAllAlignments5e(c *gin.Context) {
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
	alignments, err := GetAllAlignments5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve alignments. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseAlignments := []Alignment{}
	for _, alignment := range alignments {
		responseAlignment := Alignment{
			Index:         alignment.Index,
			AlignmentName: alignment.AlignmentName,
			Abbreviation:  alignment.Abbreviation,
			Description:   alignment.Description,
			URL:           alignment.URL,
		}
		responseAlignments = append(responseAlignments, responseAlignment)
	}
	utilities.ResponseMessage(c, "Alignments retrieved successfully.", http.StatusOK, gin.H{"alignments": responseAlignments})
}

func GetAlignmentByIndex5e(c *gin.Context) {
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
	alignment, err := GetAlignmentByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve alignment. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseAlignment := Alignment{
		Index:         alignment.Index,
		AlignmentName: alignment.AlignmentName,
		Abbreviation:  alignment.Abbreviation,
		Description:   alignment.Description,
		URL:           alignment.URL,
	}
	utilities.ResponseMessage(c, "Alignment retrieved successfully.", http.StatusOK, gin.H{"alignment": responseAlignment})
}
