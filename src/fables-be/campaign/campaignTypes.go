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
	Deleted     bool               `json:"deleted" gorm:"column:deleted"`
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
	Age          int                `json:"age" gorm:"column:age"`
	Height       int                `json:"height" gorm:"column:height"`
	Weight       int                `json:"weight" gorm:"column:weight"`
	EyeColor     string             `json:"eyeColor" gorm:"column:eye_color;type:varchar"`
	SkinColor    string             `json:"skinColor" gorm:"column:skin_color;type:varchar"`
	HairColor    string             `json:"hairColor" gorm:"column:hair_color;type:varchar"`
	Gender       string             `json:"gender" gorm:"column:gender;type:varchar"`
	Ruleset      string             `json:"ruleset" gorm:"column:ruleset;type:varchar;not null"`
	Created      utilities.NullTime `json:"created" gorm:"column:created;type:timestamp with time zone;not null"`
	LastUpdated  utilities.NullTime `json:"lastUpdated" gorm:"column:last_updated;type:timestamp with time zone"`
	Deleted      bool               `json:"deleted" gorm:"column:deleted"`
	Public       bool               `json:"public" gorm:"column:is_public"`
}

type Location struct {
	ID           uint               `json:"id" gorm:"primary_key"`
	CampaignUUID string             `json:"campaignUuid" gorm:"column:campaign_uuid;type:varchar;not null"`
	UUID         string             `json:"uuid" gorm:"column:uuid;type:varchar;not null;unique"`
	CreatorUUID  string             `json:"creatorUuid" gorm:"column:creator_uuid;type:varchar;not null"`
	Name         string             `json:"name" gorm:"column:name;type:varchar;not null"`
	Description  string             `json:"description" gorm:"column:description;type:varchar;not null"`
	Created      utilities.NullTime `json:"created" gorm:"column:created;type:timestamp with time zone;not null"`
	LastUpdated  utilities.NullTime `json:"lastUpdated" gorm:"column:last_updated;type:timestamp with time zone"`
}

type Npc struct {
	ID           uint               `json:"id" gorm:"primary_key"`
	CampaignUUID string             `json:"campaignUuid" gorm:"column:campaign_uuid;type:varchar;not null"`
	UUID         string             `json:"uuid" gorm:"column:uuid;type:varchar;not null;unique"`
	CreatorUUID  string             `json:"creatorUuid" gorm:"column:creator_uuid;type:varchar;not null"`
	FirstName    string             `json:"firstName" gorm:"column:first_name;type:varchar;not null"`
	LastName     string             `json:"lastName" gorm:"column:last_name;type:varchar;not null"`
	Race         string             `json:"race" gorm:"column:race;type:varchar;not null"`
	Class        string             `json:"class" gorm:"column:class;type:varchar;not null"`
	Description  string             `json:"description" gorm:"column:description;type:varchar;not null"`
	IsQuestBoss  bool               `json:"isQuestBoss" gorm:"column:is_quest_boss;not null"`
	Created      utilities.NullTime `json:"created" gorm:"column:created;type:timestamp with time zone;not null"`
	LastUpdated  utilities.NullTime `json:"lastUpdated" gorm:"column:last_updated;type:timestamp with time zone"`
}

type Session struct {
	ID           uint               `json:"id" gorm:"primary_key"`
	SessionID    uint               `json:"sessionId" gorm:"column:session_id;type:integer"`
	CampaignUUID string             `json:"campaignUuid" gorm:"column:campaign_uuid;type:varchar;not null"`
	CreatorUUID  string             `json:"creatorUuid" gorm:"column:creator_uuid;type:varchar;not null"`
	UUID         string             `json:"uuid" gorm:"column:uuid;type:varchar;not null;unique"`
	PartyUUIDs   pq.StringArray     `json:"partyUuids" gorm:"column:party_uuids;type:varchar[]"`
	DateOccured  utilities.NullTime `json:"dateOccured" gorm:"column:date_occured;type:timestamp with time zone;not null"`
	Created      utilities.NullTime `json:"created" gorm:"column:created;type:timestamp with time zone;not null"`
	LastUpdated  utilities.NullTime `json:"lastUpdated" gorm:"column:last_updated;type:timestamp with time zone"`
	Deleted      bool               `json:"deleted" gorm:"column:deleted"`
}
