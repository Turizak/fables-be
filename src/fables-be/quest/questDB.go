package quest

import (
	"fmt"
	"time"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/database"
	"github.com/Turizak/fables-be/npc"
	"github.com/Turizak/fables-be/utilities"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreateQuestDB(quest *campaign.Quest, campaignUuid string) error {
	quest.Created = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	quest.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Time{}, Valid: false})
	quest.UUID = uuid.NewString()
	quest.Deleted = false
	quest.Status = "active"
	quest.EndingSessionUUID = ""
	campaign, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}
	tableName := fmt.Sprintf("%s_quests", campaign.Moniker)
	if result := database.DB.Table(tableName).Create(quest); result.Error != nil {
		return result.Error
	}
	if err := markBossNpcs(quest, campaignUuid); err != nil {
		return err
	}
	return nil
}

func GetQuestByUuidDB(questUuid string, campaignUuid string) (campaign.Quest, error) {
	var quest campaign.Quest
	campaign, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return quest, err
	}
	tableName := fmt.Sprintf("%s_quests", campaign.Moniker)
	if result := database.DB.Table(tableName).Where("uuid = ?", questUuid).First(&quest); result.Error != nil {
		return quest, result.Error
	}
	return quest, nil
}

func GetQuestsByCampaignUuidDB(campaignUuid string) ([]campaign.Quest, error) {
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return nil, err
	}
	var quests []campaign.Quest
	tableName := fmt.Sprintf("%s_quests", camp.Moniker)
	if result := database.DB.Table(tableName).Find(&quests); result.Error != nil {
		return quests, result.Error
	}
	return quests, nil
}

func UpdateQuestByUuidDB(quest *campaign.Quest, campaignUuid string) error {
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}
	tableName := fmt.Sprintf("%s_quests", camp.Moniker)
	quest.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})

	result := database.DB.Table(tableName).Save(quest)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no quest found with UUID %s in campaign %s", quest.UUID, campaignUuid)
	}

	if err := markBossNpcs(quest, campaignUuid); err != nil {
		return err
	}
	return nil
}

func markBossNpcs(quest *campaign.Quest, campaignUuid string) error {
	if len(quest.BossUUIDs) > 0 {
		for _, bossUUID := range quest.BossUUIDs {
			if err := npc.MarkNpcAsBossDB(bossUUID, campaignUuid, quest.UUID); err != nil {
				return err
			}
		}
	}
	return nil
}
