package campaign

import (
	"net/http"

	"github.com/Turizak/fables-be/account"
	"github.com/Turizak/fables-be/database"
	"github.com/Turizak/fables-be/token"
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
	claims, _ := c.Get("claims")

	campaign.CreatorUUID = claims.(*token.UserClaim).UUID
	campaign.DmUUID = claims.(*token.UserClaim).UUID

	if err := c.BindJSON(&campaign); err != nil {
		utilities.ResponseMessage(c, "Could not create campaign. Please try again.", http.StatusBadRequest, nil)
		return
	}

	campaign.Completed = false
	campaign.Active = true
	moniker, err := utilities.GenerateMoniker()
	if err != nil {
		utilities.ResponseMessage(c, "Could not create campaign. Please try again.", http.StatusInternalServerError, nil)
		return
	}
	campaign.Moniker = moniker

	if err := CreateCampaignDB(&campaign); err != nil {
		utilities.ResponseMessage(c, "Could not create campaign. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Campaign created successfully.", http.StatusCreated, gin.H{"campaign": CreateCampaignResponse(campaign)})
}

func GetCampaignByUuid(c *gin.Context) {
	uuid := c.Param("uuid")
	campaign, err := GetCampaignByUuidDB(uuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve campaign. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Campaign retrieved successfully.", http.StatusOK, gin.H{"campaign": CreateCampaignResponse(*campaign)})
}

func GetCampaignsByCreatorUuid(c *gin.Context) {
	claims, _ := c.Get("claims")

	account, err := account.GetAccountByEmailDB(claims.(*token.UserClaim).Email)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve account. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	campaigns, err := GetCampaignsByCreatorUuidDB(account.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve campaigns. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	responseCampaigns := mapCampaignsToResponse(campaigns)
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
		utilities.ResponseMessage(c, "Could not retrieve campaign. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Campaign moniker retrieved successfully.", http.StatusOK, gin.H{"moniker": campaign.Moniker})
}

func UpdateCampaignByUuid(c *gin.Context) {
	claims, _ := c.Get("claims")

	uuid := c.Param("uuid")
	campaign, err := GetCampaignByUuidDB(uuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve campaign. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	if campaign.CreatorUUID != claims.(*token.UserClaim).UUID {
		utilities.ResponseMessage(c, "Unauthorized to update campaign.", http.StatusUnauthorized, nil)
		return
	}

	var updateData UpdateCampaign
	if err := c.BindJSON(&updateData); err != nil {
		utilities.ResponseMessage(c, "Could not update campaign. Please try again.", http.StatusBadRequest, nil)
		return
	}

	// Update only the allowed fields
	if updateData.Name != nil {
		campaign.Name = *updateData.Name
	}
	if updateData.PartyUUIDs != nil {
		campaign.PartyUUIDs = *updateData.PartyUUIDs
	}
	if updateData.Completed != nil {
		campaign.Completed = *updateData.Completed
	}
	if updateData.Active != nil {
		campaign.Active = *updateData.Active
	}
	if updateData.MaxPlayers != nil {
		campaign.MaxPlayers = *updateData.MaxPlayers
	}
	if updateData.Deleted != nil {
		campaign.Deleted = *updateData.Deleted
	}

	if err := UpdateCampaignByUuidDB(uuid, campaign); err != nil {
		utilities.ResponseMessage(c, "Could not update campaign. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Campaign updated successfully.", http.StatusOK, gin.H{"campaign": CreateCampaignResponse(*campaign)})
}

// Helper function to map multiple campaigns to response format
func mapCampaignsToResponse(campaigns []Campaign) []CampaignResponse {
	response := make([]CampaignResponse, len(campaigns))
	for i, campaign := range campaigns {
		response[i] = CreateCampaignResponse(campaign)
	}
	return response
}

// Creates a single campaign response
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
