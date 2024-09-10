package campaign

import (
	"fmt"
	"time"

	"github.com/Turizak/fables-be/database"
	"github.com/Turizak/fables-be/utilities"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// CreateCampaignDB creates a new campaign in the database and creates a new tables for the campaign
func CreateCampaignDB(campaign *Campaign) error {
	campaign.Created = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	campaign.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Time{}, Valid: false})
	campaign.UUID = uuid.NewString()
	if result := database.DB.Create(campaign); result.Error != nil {
		return result.Error
	}
	if err := CreateMonikerCharacterTable(campaign.Moniker); err != nil {
		return err
	}
	return nil
}

func GetCampaignByUuidDB(uuid string) (*Campaign, error) {
	campaign := new(Campaign)
	if result := database.DB.Where("uuid = ?", uuid).First(campaign); result.Error != nil {
		return nil, result.Error
	}
	return campaign, nil
}

func GetCampaignsByCreatorUuidDB(creatorUUID string) ([]Campaign, error) {
	campaigns := []Campaign{}
	if result := database.DB.Where("creator_uuid = ?", creatorUUID).Find(&campaigns); result.Error != nil {
		return nil, result.Error
	}
	return campaigns, nil
}

func CreateMonikerCharacterTable(moniker string) error {
	tableName := fmt.Sprintf("%s_characters", moniker)
	err := database.DB.Table(tableName).AutoMigrate(&Character{})
	return err
}