package note

import "github.com/Turizak/fables-be/utilities"

type NoteResponse struct {
	UUID          *string            `json:"uuid"`
	CampaignUUID  *string            `json:"campaignUuid"`
	CreatorUUID   *string            `json:"creatorUuid"`
	SessionUUID   *string            `json:"sessionUuid"`
	Published     bool               `json:"published"`
	Title         *string            `json:"title"`
	Content       *string            `json:"content"`
	Created       utilities.NullTime `json:"created"`
	LastUpdated   utilities.NullTime `json:"lastUpdated"`
	LastPublished utilities.NullTime `json:"lastPublished"`
	Version       uint               `json:"version"`
}
