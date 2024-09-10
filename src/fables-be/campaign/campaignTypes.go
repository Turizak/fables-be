package campaign

import (
	"github.com/Turizak/fables-be/utilities"
	"github.com/lib/pq"
)

type Campaign struct {
	ID          uint               `json:"id" gorm:"primary_key"`
	UUID        string             `json:"uuid" gorm:"column:uuid;type:varchar;not null;unique"`
	Name        string             `json:"name" gorm:"column:name;type:varchar;not null"`
	Moniker     string             `json:"moniker" gorm:"column:moniker;type:varchar(10);not null;unique"`
	CreatorUUID string             `json:"creatorUuid" gorm:"column:creator_uuid;type:varchar;not null"`
	DmUUID      string             `json:"dmUuid" gorm:"column:dm_uuid;type:varchar;not null"`
	PartyUUIDs  pq.StringArray     `json:"partyUuids" gorm:"column:party_uuids;type:varchar[]"`
	Completed   bool               `json:"completed" gorm:"column:completed;not null;default:false"`
	Active      bool               `json:"active" gorm:"column:active;not null;default:true"`
	Ruleset     string             `json:"ruleset" gorm:"column:ruleset;type:varchar;not null"`
	MaxPlayers  int                `json:"maxPlayers" gorm:"column:max_players;not null"`
	Created     utilities.NullTime `json:"created" gorm:"column:created;type:timestamp with time zone;not null"`
	LastUpdated utilities.NullTime `json:"lastUpdated" gorm:"column:last_updated;type:timestamp with time zone"`
}

type CampaignResponse struct {
	UUID        string             `json:"uuid"`
	Name        string             `json:"name"`
	CreatorUUID string             `json:"creatorUuid"`
	DmUUID      string             `json:"dmUuid"`
	PartyUUIDs  []string           `json:"partyUuids"`
	Completed   bool               `json:"completed"`
	Active      bool               `json:"active"`
	Ruleset     string             `json:"ruleset"`
	MaxPlayers  int                `json:"maxPlayers"`
	Created     utilities.NullTime `json:"created"`
	LastUpdated utilities.NullTime `json:"lastUpdated"`
}

type Character struct {
	ID           uint               `json:"id" gorm:"primary_key"`
	CampaignUUID string             `json:"campaignUuid" gorm:"column:campaign_uuid;type:varchar;not null"`
	UUID         string             `json:"uuid" gorm:"column:uuid;type:varchar;not null;unique"`
	CreatorUUID  string             `json:"creatorUuid" gorm:"column:creator_uuid;type:varchar;not null"`
	OwnerUUID    string             `json:"ownerUuid" gorm:"column:owner_uuid;type:varchar;not null"`
	FirstName    string             `json:"firstName" gorm:"column:first_name;type:varchar;not null"`
	LastName     string             `json:"lastName" gorm:"column:last_name;type:varchar;not null"`
	Race         string             `json:"race" gorm:"column:race;type:varchar;not null"`
	Class        string             `json:"class" gorm:"column:class;type:varchar;not null"`
	Age          int                `json:"age" gorm:"column:age;not null"`
	Height       int                `json:"height" gorm:"column:height;not null"`
	Weight       int                `json:"weight" gorm:"column:weight;not null"`
	EyeColor     string             `json:"eyeColor" gorm:"column:eye_color;type:varchar;not null"`
	SkinColor    string             `json:"skinColor" gorm:"column:skin_color;type:varchar;not null"`
	HairColor    string             `json:"hairColor" gorm:"column:hair_color;type:varchar;not null"`
	Ruleset      string             `json:"ruleset" gorm:"column:ruleset;type:varchar;not null"`
	Created      utilities.NullTime `json:"created" gorm:"column:created;type:timestamp with time zone;not null"`
	LastUpdated  utilities.NullTime `json:"lastUpdated" gorm:"column:last_updated;type:timestamp with time zone"`
}

type CharacterResponse struct {
	UUID         string             `json:"uuid"`
	CampaignUUID string             `json:"campaignUuid"`
	CreatorUUID  string             `json:"creatorUuid"`
	OwnerUUID    string             `json:"ownerUuid"`
	FirstName    string             `json:"firstName"`
	LastName     string             `json:"lastName"`
	Race         string             `json:"race"`
	Class        string             `json:"class"`
	Age          int                `json:"age"`
	Height       int                `json:"height"`
	Weight       int                `json:"weight"`
	EyeColor     string             `json:"eyeColor"`
	SkinColor    string             `json:"skinColor"`
	HairColor    string             `json:"hairColor"`
	Ruleset      string             `json:"ruleset"`
	Created      utilities.NullTime `json:"created"`
	LastUpdated  utilities.NullTime `json:"lastUpdated"`
}