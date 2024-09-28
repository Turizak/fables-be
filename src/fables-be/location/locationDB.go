package location

import (
	"fmt"
	"time"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/database"
	"github.com/Turizak/fables-be/utilities"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

func CreateLocationDB(location *campaign.Location, campaignUuid string) error {
	location.Created = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	location.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Time{}, Valid: false})
	location.UUID = uuid.NewString()
	campaign, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}
	tableName := fmt.Sprintf("%s_locations", campaign.Moniker)
	if result := database.DB.Table(tableName).Create(location); result.Error != nil {
		return result.Error
	}
	return nil
}

func GetLocationByUuidDB(uuid string, campaignUuid string) (*campaign.Location, error) {
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return nil, err
	}
	location := new(campaign.Location)
	tableName := fmt.Sprintf("%s_locations", camp.Moniker)
	if result := database.DB.Table(tableName).Where("uuid = ?", uuid).First(location); result.Error != nil {
		return nil, result.Error
	}
	return location, nil
}

func GetLocationsByCampaignUuidDB(campaignUuid string) ([]campaign.Location, error) {
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return nil, err
	}
	var locations []campaign.Location
	tableName := fmt.Sprintf("%s_locations", camp.Moniker)
	if result := database.DB.Table(tableName).Find(&locations); result.Error != nil {
		return nil, result.Error
	}
	return locations, nil
}

func GetLocationsByCreatorUuidDB(creatorUuid string) ([]campaign.Location, error) {
	var allLocations []campaign.Location

	// Fetch all monikers to loop through the respective tables
	monikers, err := campaign.GetAllMonikers()
	if err != nil {
		return nil, err
	}

	// Iterate over each moniker and query the respective table for locations
	for _, moniker := range monikers {
		var locations []campaign.Location
		tableName := fmt.Sprintf("%s_locations", moniker)

		// Query each table using the generated table name
		if result := database.DB.Table(tableName).Where("creator_uuid = ?", creatorUuid).Find(&locations); result.Error != nil {
			// If table doesn't exist, just skip and continue
			continue
		}

		// Append found locations to the result
		allLocations = append(allLocations, locations...)
	}

	return allLocations, nil
}
