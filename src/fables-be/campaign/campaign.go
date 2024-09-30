package campaign

import (
	"net/http"

	"github.com/Turizak/fables-be/account"
	"github.com/Turizak/fables-be/database"
	"github.com/Turizak/fables-be/utilities"

	"github.com/gin-gonic/gin"
)

func GetAllMonikers() ([]string, error) {
	var monikers []string
	if result := database.DB.Model(&Campaign{}).Select("moniker").Find(&monikers); result.Error != nil {
		return nil, result.Error
	}
	return monikers, nil
}

func CreateCampaign(c *gin.Context) {
	var campaign Campaign
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
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
	responseCampaign := CreateCampaignResponse(campaign)

	utilities.ResponseMessage(c, "Campaign created successfully.", http.StatusCreated, gin.H{"campaign": responseCampaign})
}

func GetCampaignByUuid(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	_, validToken := utilities.ValidateAuthenticationToken(c, authToken)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	uuid := c.Param("uuid")
	campaign, err := GetCampaignByUuidDB(uuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve campaign. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseCampaign := CreateCampaignResponse(*campaign)
	utilities.ResponseMessage(c, "Campaign retrieved successfully.", http.StatusOK, gin.H{"campaign": responseCampaign})
}

func GetCampaignsByCreatorUuid(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	claims, validToken := utilities.ValidateAuthenticationToken(c, authToken)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	account, err := account.GetAccountByEmailDB(claims.Email)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve account. Please try again.", http.StatusBadRequest, nil)
		return
	}
	campaigns, err := GetCampaignsByCreatorUuidDB(account.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve campaigns. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseCampaigns := []CampaignResponse{}
	for _, campaign := range campaigns {
		responseCampaign := CreateCampaignResponse(campaign)
		responseCampaigns = append(responseCampaigns, responseCampaign)
	}
	if len(responseCampaigns) == 0 {
		utilities.ResponseMessage(c, "No campaigns found.", http.StatusOK, nil)
		return
	}
	utilities.ResponseMessage(c, "Campaigns retrieved successfully.", http.StatusOK, gin.H{"campaigns": responseCampaigns})
}

func GetCampaignMonikerByUuid(c *gin.Context) {
	uuid := c.Param("uuid")
	campaign, err := GetCampaignByUuidDB(uuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve campaign. Please try again.", http.StatusBadRequest, nil)
		return
	}
	utilities.ResponseMessage(c, "Campaign moniker retrieved successfully.", http.StatusOK, gin.H{"moniker": campaign.Moniker})
}

func CreateCampaignResponse(campaign Campaign) CampaignResponse {
	return CampaignResponse{
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
		LastUpdated: campaign.LastUpdated,
	}
}
