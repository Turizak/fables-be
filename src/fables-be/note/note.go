package note

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	var note campaign.Note

	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	if err := c.BindJSON(&note); err != nil {
		utilities.ResponseMessage(c, "Could not create Note. Please try again.", http.StatusBadRequest, nil)
		return
	}

	note.CreatorUUID = claims.UUID
	note.CampaignUUID = c.Param("uuid")
	note.SessionUUID = c.Param("sessionUuid")

	if err := CreateNoteDB(&note); err != nil {
		utilities.ResponseMessage(c, "Could not create Note. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Note created successfully.", http.StatusCreated, gin.H{"note": CreateNoteResponse(note)})
}

func GetNoteByUuid(c *gin.Context) {
	// GetNote retrieves a note
}

func GetNotesByCampaignUuid(c *gin.Context) {
	// GetNotes retrieves all notes
}

func GetNotesBySessionUuid(c *gin.Context) {
	// GetNotes retrieves all notes
}

func GetNotesByCreatorUuid(c *gin.Context) {
	// GetNotes retrieves all notes
}

func UpdateNotePublishStatusByUuid(c *gin.Context) {
	// UpdateNotePublishStatus updates the publish status of a note
}

func UpdateNoteByUuid(c *gin.Context) {
	// UpdateNote updates a note
}

func CreateNoteResponse(note campaign.Note) NoteResponse {
	return NoteResponse{
		UUID:          &note.UUID,
		CampaignUUID:  &note.CampaignUUID,
		CreatorUUID:   &note.CreatorUUID,
		SessionUUID:   &note.SessionUUID,
		Published:     note.Published,
		Title:         &note.Title,
		Content:       &note.Content,
		Created:       note.Created,
		LastUpdated:   note.LastUpdated,
		LastPublished: note.LastPublished,
		Version:       note.Version,
	}
}
