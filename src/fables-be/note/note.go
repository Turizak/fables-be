package note

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/token"
	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

func CreateNote(c *gin.Context) {
	var note campaign.Note
	claims, _ := c.Get("claims")

	note.CreatorUUID = claims.(*token.UserClaim).UUID

	campaignUuid := c.Param("uuid")
	sessionUuid := c.Param("sessionUuid")
	if err := c.BindJSON(&note); err != nil {
		utilities.ResponseMessage(c, "Could not create Note. Please try again.", http.StatusBadRequest, nil)
		return
	}
	note.CampaignUUID = campaignUuid
	note.SessionUuid = sessionUuid

	if err := CreateNoteDB(&note, note.CampaignUUID); err != nil {
		if err.Error() == "note already exists for this user and session" {
			utilities.ResponseMessage(c, "Note already exists for this user and session.", http.StatusBadRequest, nil)
			return
		} else {
			utilities.ResponseMessage(c, "Could not create Note. Please try again.", http.StatusInternalServerError, nil)
			return
		}
	}

	utilities.ResponseMessage(c, "Note created successfully.", http.StatusCreated, gin.H{"note": CreateNoteResponse(note)})
}

func GetNoteByUuid(c *gin.Context) {
	campaignUuid := c.Param("uuid")
	noteUuid := c.Param("noteUuid")
	note, err := GetNoteByUuidDB(noteUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve Note. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Note retrieved successfully.", http.StatusOK, gin.H{"note": CreateNoteResponse(*note)})
}

func GetNotesByCampaignUuid(c *gin.Context) {
	campaignUuid := c.Param("uuid")
	notes, err := GetNotesByCampaignUuidDB(campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve Notes. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Notes retrieved successfully.", http.StatusOK, gin.H{"notes": mapNotesToResponse(notes)})
}

func GetNotesByCreatorUuid(c *gin.Context) {
	claims, _ := c.Get("claims")

	notes, err := GetNotesByCreatorUuidDB(claims.(*token.UserClaim).UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve Notes. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Notes retrieved successfully.", http.StatusOK, gin.H{"notes": mapNotesToResponse(notes)})
}

func UpdateNoteByUuid(c *gin.Context) {
	campaignUuid := c.Param("uuid")
	noteUuid := c.Param("noteUuid")
	note, err := GetNoteByUuidDB(noteUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve note. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	var updateNote UpdateNote
	if err := c.BindJSON(&updateNote); err != nil {
		utilities.ResponseMessage(c, "Could not update note. Please try again.", http.StatusBadRequest, nil)
		return
	}

	if updateNote.Deleted != nil {
		note.Deleted = *updateNote.Deleted
	}
	if updateNote.Title != nil {
		note.Title = *updateNote.Title
	}
	if updateNote.Content != nil {
		note.Content = *updateNote.Content
	}

	if err := UpdateNoteByUuidDB(note, campaignUuid); err != nil {
		utilities.ResponseMessage(c, "Could not update note. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Note updated successfully.", http.StatusOK, gin.H{"note": CreateNoteResponse(*note)})
}

func CreateNoteResponse(note campaign.Note) NoteResponse {
	return NoteResponse{
		UUID:         utilities.ToPointer(note.UUID),
		CampaignUUID: utilities.ToPointer(note.CampaignUUID),
		CreatorUUID:  utilities.ToPointer(note.CreatorUUID),
		SessionUuid:  utilities.ToPointer(note.SessionUuid),
		Title:        utilities.ToPointer(note.Title),
		Content:      utilities.ToPointer(note.Content),
		Created:      note.Created,
		LastUpdated:  note.LastUpdated,
	}
}

func mapNotesToResponse(notes []campaign.Note) []NoteResponse {
	response := make([]NoteResponse, len(notes))
	for i, note := range notes {
		response[i] = CreateNoteResponse(note)
	}
	return response
}
