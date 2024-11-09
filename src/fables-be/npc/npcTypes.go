package npc

import "github.com/Turizak/fables-be/utilities"

type NpcResponse struct {
	UUID         *string            `json:"uuid"`
	CampaignUUID *string            `json:"campaignUuid"`
	CreatorUUID  *string            `json:"creatorUuid"`
	FirstName    *string            `json:"firstName"`
	LastName     *string            `json:"lastName"`
	Race         *string            `json:"race"`
	Class        *string            `json:"class"`
	Description  *string            `json:"description"`
	IsQuestBoss  bool               `json:"isQuestBoss"`
	Created      utilities.NullTime `json:"created"`
	LastUpdated  utilities.NullTime `json:"lastUpdated"`
}

type UpdateNpc struct {
	CampaignUUID *string `json:"campaignUuid"`
	FirstName    *string `json:"firstName"`
	LastName     *string `json:"lastName"`
	Race         *string `json:"race"`
	Class        *string `json:"class"`
	Description  *string `json:"description"`
	IsQuestBoss  *bool   `json:"isQuestBoss"`
}
