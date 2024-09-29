package npc

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/utilities"

	"github.com/gin-gonic/gin"
)

func CreateNpc(c *gin.Context) {
	var npc campaign.Npc
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	campaignUuid := c.Param("uuid")
	claims, validToken := utilities.ValidateAuthenticationToken(c, authToken)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	if err := c.BindJSON(&npc); err != nil {
		utilities.ResponseMessage(c, "Could not create npc. Please try again.", http.StatusBadRequest, nil)
		return
	}
	npc.CreatorUUID = claims.UUID
	npc.CampaignUUID = campaignUuid
	if err := CreateNpcDB(&npc, npc.CampaignUUID); err != nil {
		utilities.ResponseMessage(c, "Could not create npc. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseNpc := CreateNpcResponse(npc)
	utilities.ResponseMessage(c, "Npc created successfully.", http.StatusCreated, gin.H{"npc": responseNpc})
}

func GetNpcByUuid(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	campaignUuid := c.Param("uuid")
	npcUuid := c.Param("npcUuid")
	_, validToken := utilities.ValidateAuthenticationToken(c, authToken)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	npc, err := GetNpcByUuidDB(npcUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve npc. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseNpc := CreateNpcResponse(*npc)
	utilities.ResponseMessage(c, "Npc retrieved successfully.", http.StatusOK, gin.H{"npc": responseNpc})
}

func GetNpcsByCampaignUuid(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	campaignUuid := c.Param("uuid")
	_, validToken := utilities.ValidateAuthenticationToken(c, authToken)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	npcs, err := GetNpcsByCampaignUuidDB(campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve npcs. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseNpcs := make([]NpcResponse, 0)
	for _, npc := range npcs {
		responseNpc := CreateNpcResponse(npc)
		responseNpcs = append(responseNpcs, responseNpc)
	}
	utilities.ResponseMessage(c, "Npcs retrieved successfully.", http.StatusOK, gin.H{"npcs": responseNpcs})
}

func GetNpcsByCreatorUuid(c *gin.Context) {
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
	npcs, err := GetNpcsByCreatorUuidDB(claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve npcs. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseNpcs := make([]NpcResponse, 0)
	for _, npc := range npcs {
		responseNpc := CreateNpcResponse(npc)
		responseNpcs = append(responseNpcs, responseNpc)
	}
	utilities.ResponseMessage(c, "Npcs retrieved successfully.", http.StatusOK, gin.H{"npcs": responseNpcs})
}

func CreateNpcResponse(npc campaign.Npc) NpcResponse {
	return NpcResponse{
		UUID:         utilities.ToPointer(npc.UUID),
		CampaignUUID: utilities.ToPointer(npc.CampaignUUID),
		CreatorUUID:  utilities.ToPointer(npc.CreatorUUID),
		FirstName:    utilities.ToPointer(npc.FirstName),
		LastName:     utilities.ToPointer(npc.LastName),
		Race:         utilities.ToPointer(npc.Race),
		Class:        utilities.ToPointer(npc.Class),
		Description:  utilities.ToPointer(npc.Description),
		IsQuestBoss:  npc.IsQuestBoss,
		Created:      npc.Created,
		LastUpdated:  npc.LastUpdated,
	}
}
