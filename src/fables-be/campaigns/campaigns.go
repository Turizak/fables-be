package campaigns

import (
	"net/http"

	"github.com/Turizak/fables-be/utilities"

	"github.com/gin-gonic/gin"
)

func CreateCampaign(c *gin.Context) {
	var campaign Campaign
	authToken := c.GetHeader("Authorization")
	claims, validToken := utilities.ValidateAuthenticationToken(c, authToken)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	if err := c.BindJSON(&campaign); err != nil {
		utilities.ResponseMessage(c, "Could not create campaign. Please try again.", http.StatusBadRequest, nil)
		return
	}
	campaign.CreatorUUID = claims.UUID
	campaign.DmUUID = claims.UUID
	campaign.Completed = false
	campaign.Active = true
	moniker, err := utilities.GenerateMoniker()
	if err != nil {
		utilities.ResponseMessage(c, "Could not create campaign. Please try again.", http.StatusBadRequest, nil)
		return
	}
	campaign.Moniker = moniker
	if err := CreateCampaignDB(&campaign); err != nil {
		utilities.ResponseMessage(c, "Could not create campaign. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseCampaign := CampaignResponse{
		UUID:        campaign.UUID,
		Name:        campaign.Name,
		CreatorUUID: campaign.CreatorUUID,
		DmUUID:      campaign.DmUUID,
		PartyUUIDs:  campaign.PartyUUIDs,
		Completed:   campaign.Completed,
		Active:      campaign.Active,
		Ruleset:     campaign.Ruleset,
		MaxPlayers:  campaign.MaxPlayers,
		Created:     campaign.Created,
	}
	utilities.ResponseMessage(c, "Campaign created successfully.", http.StatusCreated, gin.H{"campaign": responseCampaign})
}
