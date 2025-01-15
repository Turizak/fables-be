package note

import "github.com/Turizak/fables-be/utilities"

type NoteResponse struct {
	UUID         *string            `json:"uuid"`
	CampaignUUID *string            `json:"campaignUuid"`
	CreatorUUID  *string            `json:"creatorUuid"`
	SessionUuid  *string            `json:"sessionUuid"`
	Title        *string            `json:"title"`
	Content      *string            `json:"content"`
	Created      utilities.NullTime `json:"created"`
	LastUpdated  utilities.NullTime `json:"lastUpdated"`
}

type UpdateNote struct {
	UUID         *string            `json:"uuid"`
	CampaignUUID *string            `json:"campaignUuid"`
	CreatorUUID  *string            `json:"creatorUuid"`
	SessionUuid  *string            `json:"sessionUuid"`
	Title        *string            `json:"title"`
	Content      *string            `json:"content"`
	Created      utilities.NullTime `json:"created"`
	LastUpdated  utilities.NullTime `json:"lastUpdated"`
	Deleted      *bool              `json:"deleted"`
}
