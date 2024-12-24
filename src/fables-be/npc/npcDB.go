package npc

import (
	"fmt"
	"time"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/database"
	"github.com/Turizak/fables-be/session"
	"github.com/Turizak/fables-be/utilities"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreateNpcDB(npc *campaign.Npc, campaignUuid string) error {
	npc.Created = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	npc.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Time{}, Valid: false})
	npc.UUID = uuid.NewString()
	npc.Deleted = false
	campaign, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}
	tableName := fmt.Sprintf("%s_npcs", campaign.Moniker)
	if result := database.DB.Table(tableName).Create(npc); result.Error != nil {
		return result.Error
	}
	return nil
}

func CreateNpcSessionDB(npc *campaign.Npc, campaignUuid string, sessionUuid string) error {
	npc.Created = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	npc.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Time{}, Valid: false})
	npc.UUID = uuid.NewString()
	npc.Deleted = false
	campaign, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}
	tableName := fmt.Sprintf("%s_npcs", campaign.Moniker)
	if result := database.DB.Table(tableName).Create(npc); result.Error != nil {
		return result.Error
	}
	ses, err := session.GetSessionByUuidDB(sessionUuid, campaignUuid)
	if err != nil {
		return err
	}
	sessionNpcUuids := false
	for _, uuid := range ses.NpcUUIDs {
		if uuid == npc.UUID {
			sessionNpcUuids = true
			break
		}
	}
	// Append the UUID only if it doesn't already exist
	if !sessionNpcUuids {
		ses.NpcUUIDs = append(ses.NpcUUIDs, npc.UUID)
	}
	// Update the session
	err = session.UpdateSessionByUuidDB(ses, campaignUuid)
	if err != nil {
		return err
	}
	return nil
}

func GetNpcByUuidDB(uuid string, campaignUuid string) (*campaign.Npc, error) {
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return nil, err
	}
	npc := new(campaign.Npc)
	tableName := fmt.Sprintf("%s_npcs", camp.Moniker)
	if result := database.DB.Table(tableName).Where("uuid = ?", uuid).First(npc); result.Error != nil {
		return nil, result.Error
	}
	return npc, nil
}

func GetNpcsByCampaignUuidDB(campaignUuid string) ([]campaign.Npc, error) {
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return nil, err
	}
	var npcs []campaign.Npc
	tableName := fmt.Sprintf("%s_npcs", camp.Moniker)
	if result := database.DB.Table(tableName).Find(&npcs); result.Error != nil {
		return nil, result.Error
	}
	return npcs, nil
}

func GetNpcsByCreatorUuidDB(creatorUuid string) ([]campaign.Npc, error) {
	var allNpcs []campaign.Npc

	// Fetch all monikers to loop through the respective tables
	monikers, err := campaign.GetAllMonikers()
	if err != nil {
		return nil, err
	}

	// Iterate over each moniker and query the respective table for npcs
	for _, moniker := range monikers {
		var npcs []campaign.Npc
		tableName := fmt.Sprintf("%s_npcs", moniker)

		// Query each table using the generated table name
		if result := database.DB.Table(tableName).Where("creator_uuid = ?", creatorUuid).Find(&npcs); result.Error != nil {
			// If table doesn't exist, just skip and continue
			continue
		}

		// Append found npcs to the result
		allNpcs = append(allNpcs, npcs...)
	}

	return allNpcs, nil
}

func UpdateNpcByUuidDB(npc *campaign.Npc, campaignUuid string) error {
	// Retrieve the campaign to determine the correct table
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}

	// Set the table name based on the campaign's moniker
	tableName := fmt.Sprintf("%s_npcs", camp.Moniker)

	// Update the LastUpdated field to the current time in UTC
	npc.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})

	// Perform the update operation
	result := database.DB.Table(tableName).Save(npc)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no npc found with UUID %s in campaign %s", npc.UUID, campaignUuid)
	}

	return nil
}

func MarkNpcAsBossDB(npcUuid string, campaignUuid string, questUuid string) error {
	npc, err := GetNpcByUuidDB(npcUuid, campaignUuid)
	if err != nil {
		return err
	}
	npc.IsQuestBoss = true
	npc.QuestBossUUID = questUuid
	return UpdateNpcByUuidDB(npc, campaignUuid)
}
