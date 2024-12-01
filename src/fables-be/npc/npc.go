package npc

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

func CreateNpc(c *gin.Context) {
	var npc campaign.Npc
	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	npc.CreatorUUID = claims.UUID

	campaignUuid := c.Param("uuid")
	if err := c.BindJSON(&npc); err != nil {
		utilities.ResponseMessage(c, "Could not create NPC. Please try again.", http.StatusBadRequest, nil)
		return
	}
	npc.CampaignUUID = campaignUuid

	if err := CreateNpcDB(&npc, npc.CampaignUUID); err != nil {
		utilities.ResponseMessage(c, "Could not create NPC. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "NPC created successfully.", http.StatusCreated, gin.H{"npc": CreateNpcResponse(npc)})
}

func GetNpcByUuid(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	campaignUuid := c.Param("uuid")
	npcUuid := c.Param("npcUuid")
	npc, err := GetNpcByUuidDB(npcUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve NPC. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "NPC retrieved successfully.", http.StatusOK, gin.H{"npc": CreateNpcResponse(*npc)})
}

func GetNpcsByCampaignUuid(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	campaignUuid := c.Param("uuid")
	npcs, err := GetNpcsByCampaignUuidDB(campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve NPCs. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "NPCs retrieved successfully.", http.StatusOK, gin.H{"npcs": mapNpcsToResponse(npcs)})
}

func GetNpcsByCreatorUuid(c *gin.Context) {
	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	npcs, err := GetNpcsByCreatorUuidDB(claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve NPCs. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "NPCs retrieved successfully.", http.StatusOK, gin.H{"npcs": mapNpcsToResponse(npcs)})
}

func UpdateNpcByUuid(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	campaignUuid := c.Param("uuid")
	npcUuid := c.Param("npcUuid")
	npc, err := GetNpcByUuidDB(npcUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve npc. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	var updateNpc UpdateNpc
	if err := c.BindJSON(&updateNpc); err != nil {
		utilities.ResponseMessage(c, "Could not update npc. Please try again.", http.StatusBadRequest, nil)
		return
	}

	if updateNpc.CampaignUUID != nil {
		npc.CampaignUUID = *updateNpc.CampaignUUID
	}
	if updateNpc.FirstName != nil {
		npc.FirstName = *updateNpc.FirstName
	}
	if updateNpc.LastName != nil {
		npc.LastName = *updateNpc.LastName
	}
	if updateNpc.Description != nil {
		npc.Description = *updateNpc.Description
	}
	if updateNpc.Class != nil {
		npc.Class = *updateNpc.Class
	}
	if updateNpc.Race != nil {
		npc.Race = *updateNpc.Race
	}
	if updateNpc.IsQuestBoss != nil {
		npc.IsQuestBoss = *updateNpc.IsQuestBoss
	}
	if updateNpc.Deleted != nil {
		npc.Deleted = *updateNpc.Deleted
	}

	if err := UpdateNpcByUuidDB(npc, campaignUuid); err != nil {
		utilities.ResponseMessage(c, "Could not update npc. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Npc updated successfully.", http.StatusOK, gin.H{"npc": CreateNpcResponse(*npc)})
}

// Helper functions to map NPCs to their response formats
func CreateNpcResponse(npc campaign.Npc) NpcResponse {
	return NpcResponse{
		UUID:          utilities.ToPointer(npc.UUID),
		CampaignUUID:  utilities.ToPointer(npc.CampaignUUID),
		CreatorUUID:   utilities.ToPointer(npc.CreatorUUID),
		FirstName:     utilities.ToPointer(npc.FirstName),
		LastName:      utilities.ToPointer(npc.LastName),
		Race:          utilities.ToPointer(npc.Race),
		Class:         utilities.ToPointer(npc.Class),
		Description:   utilities.ToPointer(npc.Description),
		IsQuestBoss:   npc.IsQuestBoss,
		QuestBossUUID: utilities.ToPointer(npc.QuestBossUUID),
		Created:       npc.Created,
		LastUpdated:   npc.LastUpdated,
	}
}

func mapNpcsToResponse(npcs []campaign.Npc) []NpcResponse {
	response := make([]NpcResponse, len(npcs))
	for i, npc := range npcs {
		response[i] = CreateNpcResponse(npc)
	}
	return response
}
