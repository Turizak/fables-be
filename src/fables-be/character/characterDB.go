package character

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

func CreateCharacterDB(character *campaign.Character, campaignUuid string) error {
	character.Created = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	character.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Time{}, Valid: false})
	character.UUID = uuid.NewString()
	character.Deleted = false
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}
	character.Ruleset = camp.Ruleset
	tableName := fmt.Sprintf("%s_characters", camp.Moniker)
	if result := database.DB.Table(tableName).Create(character); result.Error != nil {
		return result.Error
	}

	// Add character to Campaign PartyUUIDs
	// Check if the character.UUID is already in the Campaign PartyUUIDs array
	campaignPartyExists := false
	for _, uuid := range camp.PartyUUIDs {
		if uuid == character.UUID {
			campaignPartyExists = true
			break
		}
	}
	// Append the UUID only if it doesn't already exist
	if !campaignPartyExists {
		camp.PartyUUIDs = append(camp.PartyUUIDs, character.UUID)
	}
	// Update the campaign
	err = campaign.UpdateCampaignByUuidDB(campaignUuid, camp)
	if err != nil {
		return err
	}
	return nil
}

func CreateCharacterSessionDB(character *campaign.Character, campaignUuid string, sessionUuid string) error {
	character.Created = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	character.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Time{}, Valid: false})
	character.UUID = uuid.NewString()
	character.Deleted = false

	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}

	character.Ruleset = camp.Ruleset

	tableName := fmt.Sprintf("%s_characters", camp.Moniker)
	if result := database.DB.Table(tableName).Create(character); result.Error != nil {
		return result.Error
	}

	// Add character to Campaign PartyUUIDs
	// Check if the character.UUID is already in the Campaign PartyUUIDs array
	campaignPartyExists := false
	for _, uuid := range camp.PartyUUIDs {
		if uuid == character.UUID {
			campaignPartyExists = true
			break
		}
	}
	// Append the UUID only if it doesn't already exist
	if !campaignPartyExists {
		camp.PartyUUIDs = append(camp.PartyUUIDs, character.UUID)
	}
	// Update the campaign
	err = campaign.UpdateCampaignByUuidDB(campaignUuid, camp)
	if err != nil {
		return err
	}

	// add character to Session PartyUUIDs
	ses, err := session.GetSessionByUuidDB(sessionUuid, campaignUuid)
	if err != nil {
		return err
	}
	// Check if the character.UUID is already in the Session PartyUUIDs array
	sessionPartyExists := false
	for _, uuid := range ses.PartyUUIDs {
		if uuid == character.UUID {
			sessionPartyExists = true
			break
		}
	}
	// Append the UUID only if it doesn't already exist
	if !sessionPartyExists {
		ses.PartyUUIDs = append(ses.PartyUUIDs, character.UUID)
	}
	// Update the session
	err = session.UpdateSessionByUuidDB(ses, campaignUuid)
	if err != nil {
		return err
	}

	return nil
}

func GetCharacterByUuidDB(uuid string, campaignUuid string) (*campaign.Character, error) {
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return nil, err
	}
	character := new(campaign.Character)
	tableName := fmt.Sprintf("%s_characters", camp.Moniker)
	if result := database.DB.Table(tableName).Where("uuid = ?", uuid).First(character); result.Error != nil {
		return nil, result.Error
	}
	return character, nil
}

func GetCharactersByCampaignUuidDB(campaignUuid string) ([]campaign.Character, error) {
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return nil, err
	}
	var characters []campaign.Character
	tableName := fmt.Sprintf("%s_characters", camp.Moniker)
	if result := database.DB.Table(tableName).Find(&characters); result.Error != nil {
		return nil, result.Error
	}
	return characters, nil
}

func GetCharactersByCreatorUuidDB(creatorUuid string) ([]campaign.Character, error) {
	var allCharacters []campaign.Character

	// Fetch all monikers to loop through the respective tables
	monikers, err := campaign.GetAllMonikers()
	if err != nil {
		return nil, err
	}

	// Iterate over each moniker and query the respective table for characters
	for _, moniker := range monikers {
		var characters []campaign.Character
		tableName := fmt.Sprintf("%s_characters", moniker)

		// Query each table using the generated table name
		if result := database.DB.Table(tableName).Where("creator_uuid = ?", creatorUuid).Find(&characters); result.Error != nil {
			// If table doesn't exist, just skip and continue
			continue
		}

		// Append found characters to the result
		allCharacters = append(allCharacters, characters...)
	}

	return allCharacters, nil
}

func GetCharactersByOwnerUuidDB(ownerUuid string) ([]campaign.Character, error) {
	var allCharacters []campaign.Character

	// Fetch all monikers to loop through the respective tables
	monikers, err := campaign.GetAllMonikers()
	if err != nil {
		return nil, err
	}

	// Iterate over each moniker and query the respective table for characters
	for _, moniker := range monikers {
		var characters []campaign.Character
		tableName := fmt.Sprintf("%s_characters", moniker)

		// Query each table using the generated table name
		if result := database.DB.Table(tableName).Where("owner_uuid = ?", ownerUuid).Find(&characters); result.Error != nil {
			// If table doesn't exist, just skip and continue
			continue
		}

		// Append found characters to the result
		allCharacters = append(allCharacters, characters...)
	}

	return allCharacters, nil
}

func UpdateCharacterByUuidDB(character *campaign.Character, campaignUuid string) error {
	// Retrieve the campaign to determine the correct table
	camp, err := campaign.GetCampaignByUuidDB(campaignUuid)
	if err != nil {
		return err
	}

	// Set the table name based on the campaign's moniker
	tableName := fmt.Sprintf("%s_characters", camp.Moniker)

	// Update the LastUpdated field to the current time in UTC
	character.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})

	// Perform the update operation
	result := database.DB.Table(tableName).Save(character)

	if result.Error != nil {
		return result.Error
	}
	if result.RowsAffected == 0 {
		return fmt.Errorf("no character found with UUID %s in campaign %s", character.UUID, campaignUuid)
	}

	return nil
}
