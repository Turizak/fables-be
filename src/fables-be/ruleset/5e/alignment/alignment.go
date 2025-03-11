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

// GetAllAlignments5e retrieves all alignments for 5e
func GetAllAlignments5e(c *gin.Context) {
	alignments, err := GetAllAlignments5eDB()
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve alignments. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	responseAlignments := mapAlignmentsToResponse(alignments)
	utilities.ResponseMessage(c, "Alignments retrieved successfully.", http.StatusOK, gin.H{"alignments": responseAlignments})
}

// GetAlignmentByIndex5e retrieves a specific alignment by index for 5e
func GetAlignmentByIndex5e(c *gin.Context) {
	index := c.Param("index")
	alignment, err := GetAlignmentByIndex5eDB(index)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve alignment. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	responseAlignment := mapAlignmentToResponse(*alignment)
	utilities.ResponseMessage(c, "Alignment retrieved successfully.", http.StatusOK, gin.H{"alignment": responseAlignment})
}

// mapAlignmentsToResponse maps a slice of Alignment structs to their response form
func mapAlignmentsToResponse(alignments []Alignment) []Alignment {
	responseAlignments := make([]Alignment, len(alignments))
	for i, alignment := range alignments {
		responseAlignments[i] = mapAlignmentToResponse(alignment)
	}
	return responseAlignments
}

// mapAlignmentToResponse maps a single Alignment to its response form
func mapAlignmentToResponse(alignment Alignment) Alignment {
	return Alignment{
		Index:         alignment.Index,
		AlignmentName: alignment.AlignmentName,
		Abbreviation:  alignment.Abbreviation,
		Description:   alignment.Description,
		URL:           alignment.URL,
	}
}
