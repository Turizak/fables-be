package note

import (
	"fmt"
	"time"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/database"
	"github.com/Turizak/fables-be/utilities"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreateNoteDB(note *campaign.Note) error {
	note.Created = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	note.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Time{}, Valid: false})
	note.UUID = uuid.NewString()
	note.Deleted = false
	note.Published = false

	campaign, err := campaign.GetCampaignByUuidDB(note.CampaignUUID)
	if err != nil {
		return err
	}
	tableName := fmt.Sprintf("%s_notes", campaign.Moniker)
	var rowCount int64
	if err := database.DB.Table(tableName).
		Where("session_uuid = ? AND creator_uuid = ?", note.SessionUUID, note.CreatorUUID).
		Count(&rowCount).Error; err != nil {
		return err
	}
	note.Version = uint(rowCount + 1)
	if result := database.DB.Table(tableName).Create(note); result.Error != nil {
		return result.Error
	}
	// Handle the case where a new note (draft is created, we need to unpublish the previous note)
	//    This might actually happen in a different function
	return nil
}
