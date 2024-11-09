package character

import "github.com/Turizak/fables-be/utilities"

type CharacterResponse struct {
	UUID         *string            `json:"uuid"`
	CampaignUUID *string            `json:"campaignUuid"`
	CreatorUUID  *string            `json:"creatorUuid"`
	OwnerUUID    *string            `json:"ownerUuid"`
	FirstName    *string            `json:"firstName"`
	LastName     *string            `json:"lastName"`
	Race         *string            `json:"race"`
	Class        *string            `json:"class"`
	Age          int                `json:"age"`
	Height       int                `json:"height"`
	Weight       int                `json:"weight"`
	EyeColor     *string            `json:"eyeColor"`
	SkinColor    *string            `json:"skinColor"`
	HairColor    *string            `json:"hairColor"`
	Ruleset      *string            `json:"ruleset"`
	Public       bool               `json:"public"`
	Gender       *string            `json:"gender"`
	Created      utilities.NullTime `json:"created"`
	LastUpdated  utilities.NullTime `json:"lastUpdated"`
}

type UpdateCharacter struct {
	FirstName    *string `json:"firstName"`
	LastName     *string `json:"lastName"`
	OwnerUUID    *string `json:"ownerUuid"`
	CampaignUUID *string `json:"campaignUuid"`
	Race         *string `json:"race"`
	Class        *string `json:"class"`
	Age          *int    `json:"age"`
	Height       *int    `json:"height"`
	Weight       *int    `json:"weight"`
	EyeColor     *string `json:"eyeColor"`
	SkinColor    *string `json:"skinColor"`
	HairColor    *string `json:"hairColor"`
	Gender       *string `json:"gender"`
	Public       *bool   `json:"public"`
	Deleted      *bool   `json:"deleted"`
}
