package quest

import "github.com/Turizak/fables-be/utilities"

type QuestResponse struct {
	UUID                *string            `json:"uuid"`
	CampaignUUID        *string            `json:"campaignUuid"`
	CreatorUUID         *string            `json:"creatorUuid"`
	Name                *string            `json:"name"`
	Description         *string            `json:"description"`
	QuestGiverUUID      *string            `json:"questGiver"`
	RewardUUIDs         []string           `json:"rewardUuids"`
	LocationUUIDs       []string           `json:"locationUuids"`
	NpcUUIDs            []string           `json:"npcUuids"`
	PartyUUIDs          []string           `json:"partyUuids"`
	BossUUIDs           []string           `json:"bossUuids"`
	StartingSessionUUID *string            `json:"startingSessionUuid"`
	EndingSessionUUID   *string            `json:"endingSessionUuid"`
	Status              *string            `json:"status"`
	Created             utilities.NullTime `json:"created"`
	LastUpdated         utilities.NullTime `json:"lastUpdated"`
}

type UpdateQuest struct {
	CampaignUUID        *string  `json:"campaignUuid"`
	Name                *string  `json:"name"`
	Description         *string  `json:"description"`
	PartyUUIDs          []string `json:"partyUuids"`
	NpcUUIDs            []string `json:"npcUuids"`
	BossUUIDs           []string `json:"bossUuids"`
	LocationUUIDs       []string `json:"locationUuids"`
	RewardUUIDs         []string `json:"rewardUuids"`
	StartingSessionUUID *string  `json:"startingSessionUuid"`
	EndingSessionUUID   *string  `json:"endingSessionUuid"`
	Status              *string  `json:"status"`
	Deleted             *bool    `json:"deleted"`
}
