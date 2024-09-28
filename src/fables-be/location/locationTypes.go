package location

import "github.com/Turizak/fables-be/utilities"

type LocationResponse struct {
	UUID         string             `json:"uuid"`
	CampaignUUID string             `json:"campaignUuid"`
	CreatorUUID  string             `json:"creatorUuid"`
	Name         string             `json:"name"`
	Description  string             `json:"description"`
	Created      utilities.NullTime `json:"created"`
	LastUpdated  utilities.NullTime `json:"lastUpdated"`
}
