package session

import "github.com/Turizak/fables-be/utilities"

type SessionResponse struct {
	SessionID    uint               `json:"sessionId"`
	UUID         *string            `json:"uuid"`
	CampaignUUID *string            `json:"campaignUuid"`
	CreatorUUID  *string            `json:"creatorUuid"`
	PartyUUIDs   []string           `json:"partyUuids"`
	DateOccured  utilities.NullTime `json:"dateOccured"`
	Created      utilities.NullTime `json:"created"`
	LastUpdated  utilities.NullTime `json:"lastUpdated"`
}
