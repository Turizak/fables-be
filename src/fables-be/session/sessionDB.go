package session

import (
	"fmt"
	"time"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/database"
	"github.com/Turizak/fables-be/utilities"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreateSessionDB(session *campaign.Session, campaignUuid string) error {
	session.Created = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	session.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Time{}, Valid: false})
	session.Deleted = false
	session.UUID = uuid.NewString()
	campaign, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}
	tableName := fmt.Sprintf("%s_sessions", campaign.Moniker)

	var rowCount int64
	if err := database.DB.Table(tableName).Count(&rowCount).Error; err != nil {
		return err
	}
	session.SessionID = uint(rowCount + 1)
	if result := database.DB.Table(tableName).Create(session); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetSessionByUuidDB(uuid string, campaignUuid string) (*campaign.Session, error) {
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return nil, err
	}
	session := new(campaign.Session)
	tableName := fmt.Sprintf("%s_sessions", camp.Moniker)
	if result := database.DB.Table(tableName).Where("uuid = ?", uuid).First(session); result.Error != nil {
		return nil, result.Error
	}
	return session, nil
}

func GetSessionsByCampaignUuidDB(campaignUuid string) ([]campaign.Session, error) {
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return nil, err
	}
	var sessions []campaign.Session
	tableName := fmt.Sprintf("%s_sessions", camp.Moniker)
	if result := database.DB.Table(tableName).Find(&sessions); result.Error != nil {
		return nil, result.Error
	}
	return sessions, nil
}

func GetSessionsByCreatorUuidDB(creatorUuid string) ([]campaign.Session, error) {
	var allSessions []campaign.Session

	// Fetch all monikers to loop through the respective tables
	monikers, err := campaign.GetAllMonikers()
	if err != nil {
		return nil, err
	}

	// Iterate over each moniker and query the respective table for sessions
	for _, moniker := range monikers {
		var sessions []campaign.Session
		tableName := fmt.Sprintf("%s_sessions", moniker)

		// Query each table using the generated table name
		if result := database.DB.Table(tableName).Where("creator_uuid = ?", creatorUuid).Find(&sessions); result.Error != nil {
			// If table doesn't exist, just skip and continue
			continue
		}

		// Append found sessions to the result
		allSessions = append(allSessions, sessions...)
	}

	return allSessions, nil
}

func UpdateSessionByUuidDB(session *campaign.Session, campaignUuid string) error {
	// Retrieve the campaign to determine the correct table
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}

	// Set the table name based on the campaign's moniker
	tableName := fmt.Sprintf("%s_sessions", camp.Moniker)

	// Update the LastUpdated field to the current time in UTC
	session.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})

	// Perform the update operation
	result := database.DB.Table(tableName).Save(session)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no session found with UUID %s in campaign %s", session.UUID, campaignUuid)
	}

	return nil
}
