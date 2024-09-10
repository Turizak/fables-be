package character

import "github.com/Turizak/fables-be/utilities"

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
