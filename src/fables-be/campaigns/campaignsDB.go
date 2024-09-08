package campaigns

import (
	"time"

	"github.com/Turizak/fables-be/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

type Campaign struct {
	ID          uint           `json:"id" gorm:"primary_key"`
	UUID        string         `json:"uuid" gorm:"column:uuid;type:varchar;not null;unique"`
	Name        string         `json:"name" gorm:"column:name;type:varchar;not null"`
	Moniker     string         `json:"moniker" gorm:"column:moniker;type:varchar(10);not null;unique"`
	CreatorUUID string         `json:"creatorUuid" gorm:"column:creator_uuid;type:varchar;not null"`
	DmUUID      string         `json:"dmUuid" gorm:"column:dm_uuid;type:varchar;not null"`
	PartyUUIDs  pq.StringArray `json:"partyUuids" gorm:"column:party_uuids;type:varchar[]"`
	Completed   bool           `json:"completed" gorm:"column:completed;not null;default:false"`
	Active      bool           `json:"active" gorm:"column:active;not null;default:true"`
	Ruleset     string         `json:"ruleset" gorm:"column:ruleset;type:varchar;not null"`
	MaxPlayers  int            `json:"maxPlayers" gorm:"column:max_players;not null"`
	Created     pq.NullTime    `json:"created" gorm:"column:created;type:timestamp with time zone;not null"`
	LastUpdated pq.NullTime    `json:"lastUpdated" gorm:"column:last_updated;type:timestamp with time zone"`
}

type CampaignResponse struct {
	UUID        string      `json:"uuid"`
	Name        string      `json:"name"`
	CreatorUUID string      `json:"creatorUuid"`
	DmUUID      string      `json:"dmUuid"`
	PartyUUIDs  []string    `json:"partyUuids"`
	Completed   bool        `json:"completed"`
	Active      bool        `json:"active"`
	Ruleset     string      `json:"ruleset"`
	MaxPlayers  int         `json:"maxPlayers"`
	Created     pq.NullTime `json:"created"`
	LastUpdated pq.NullTime `json:"lastUpdated"`
}

// CreateCampaignDB creates a new campaign in the database.
func CreateCampaignDB(campaign *Campaign) error {
	campaign.Created = pq.NullTime{Time: time.Now(), Valid: true}
	campaign.LastUpdated = pq.NullTime{Valid: false}
	campaign.UUID = uuid.NewString()
	if result := database.DB.Create(campaign); result.Error != nil {
		return result.Error
	}
	return nil
}
