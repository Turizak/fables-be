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

func CreateNoteDB(notes *campaign.Note, campaignUuid string) error {
	notes.Created = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	notes.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Time{}, Valid: false})
	notes.UUID = uuid.NewString()
	notes.Deleted = false
	campaign, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}

	// Check note for this user and session does not already exist
	existingNotes, err := GetNotesByCampaignUuidDB(campaignUuid)
	if err != nil {
		return err
	}
	for _, note := range existingNotes {
		if note.CreatorUUID == notes.CreatorUUID && note.SessionUuid == notes.SessionUuid {
			return fmt.Errorf("note already exists for this user and session")
		}
	}

	tableName := fmt.Sprintf("%s_notes", campaign.Moniker)
	if result := database.DB.Table(tableName).Create(notes); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetNoteByUuidDB(uuid string, campaignUuid string) (*campaign.Note, error) {
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return nil, err
	}
	note := new(campaign.Note)
	tableName := fmt.Sprintf("%s_notes", camp.Moniker)
	if result := database.DB.Table(tableName).Where("uuid = ?", uuid).First(note); result.Error != nil {
		return nil, result.Error
	}
	return note, nil
}

func GetNotesByCampaignUuidDB(campaignUuid string) ([]campaign.Note, error) {
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return nil, err
	}
	var notes []campaign.Note
	tableName := fmt.Sprintf("%s_notes", camp.Moniker)
	if result := database.DB.Table(tableName).Find(&notes); result.Error != nil {
		return nil, result.Error
	}
	return notes, nil
}

func GetNotesByCreatorUuidDB(creatorUuid string) ([]campaign.Note, error) {
	var allNotes []campaign.Note

	// Fetch all monikers to loop through the respective tables
	monikers, err := campaign.GetAllMonikers()
	if err != nil {
		return nil, err
	}

	// Iterate over each moniker and query the respective table for notes
	for _, moniker := range monikers {
		var notes []campaign.Note
		tableName := fmt.Sprintf("%s_notes", moniker)

		// Query each table using the generated table name
		if result := database.DB.Table(tableName).Where("creator_uuid = ?", creatorUuid).Find(&notes); result.Error != nil {
			// If table doesn't exist, just skip and continue
			continue
		}

		// Append found notes to the result
		allNotes = append(allNotes, notes...)
	}

	return allNotes, nil
}

func UpdateNoteByUuidDB(note *campaign.Note, campaignUuid string) error {
	// Retrieve the campaign to determine the correct table
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}

	// Set the table name based on the campaign's moniker
	tableName := fmt.Sprintf("%s_notes", camp.Moniker)

	// Update the LastUpdated field to the current time in UTC
	note.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})

	// Perform the update operation
	result := database.DB.Table(tableName).Save(note)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no note found with UUID %s in campaign %s", note.UUID, campaignUuid)
	}

	return nil
}
