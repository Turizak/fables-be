package collection

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/character"
	"github.com/Turizak/fables-be/location"
	"github.com/Turizak/fables-be/npc"
	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

func GetCampaignAllData(c *gin.Context) {
	uuid := c.Param("uuid")
	campaignData, err := campaign.GetCampaignByUuidDB(uuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve campaign. Please try again.", http.StatusBadRequest, nil)
		return
	}
	locations, err := location.GetLocationsByCampaignUuidDB(uuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve locations. Please try again.", http.StatusBadRequest, nil)
		return
	}
	npcs, err := npc.GetNpcsByCampaignUuidDB(uuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve npcs. Please try again.", http.StatusBadRequest, nil)
		return
	}
	characters, err := character.GetCharactersByCampaignUuidDB(uuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve characters. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseCampaign := campaign.CreateCampaignResponse(*campaignData)
	responseLocations := []location.LocationResponse{}
	for _, locationData := range locations {
		responseLocation := location.CreateLocationResponse(locationData)
		responseLocations = append(responseLocations, responseLocation)
	}
	responseNpcs := []npc.NpcResponse{}
	for _, npcData := range npcs {
		responseNpc := npc.CreateNpcResponse(npcData)
		responseNpcs = append(responseNpcs, responseNpc)
	}
	responseCharacters := []character.CharacterResponse{}
	for _, characterData := range characters {
		responseCharacter := character.CreateCharacterResponse(characterData)
		responseCharacters = append(responseCharacters, responseCharacter)
	}
	utilities.ResponseMessage(c, "Campaign data retrieved successfully.", http.StatusOK, gin.H{"campaign": responseCampaign, "locations": responseLocations, "npcs": responseNpcs, "characters": responseCharacters})
}
