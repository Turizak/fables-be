package quest

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

func CreateQuest(c *gin.Context) {
	var quest campaign.Quest
	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	quest.CreatorUUID = claims.UUID

	campaignUuid := c.Param("uuid")
	if err := c.BindJSON(&quest); err != nil {
		utilities.ResponseMessage(c, "Could not create Quest. Please try again.", http.StatusBadRequest, nil)
		return
	}
	quest.CampaignUUID = campaignUuid

	if err := CreateQuestDB(&quest, quest.CampaignUUID); err != nil {
		utilities.ResponseMessage(c, "Could not create Quest. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Quest created successfully.", http.StatusCreated, gin.H{"quest": CreateQuestResponse(quest)})
}

func GetQuestByUuid(c *gin.Context) {
	campaignUuid := c.Param("uuid")
	questUuid := c.Param("questUuid")
	quest, err := GetQuestByUuidDB(questUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve quest. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Quest retrieved successfully.", http.StatusOK, gin.H{"quest": CreateQuestResponse(quest)})
}

func GetQuestsByCampaignUuid(c *gin.Context) {
	campaignUuid := c.Param("uuid")
	quests, err := GetQuestsByCampaignUuidDB(campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve quests. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Quests retrieved successfully.", http.StatusOK, gin.H{"quests": mapQuestsToResponse(quests)})
}

func UpdateQuestByUuid(c *gin.Context) {
	campaignUuid := c.Param("uuid")
	questUuid := c.Param("questUuid")
	quest, err := GetQuestByUuidDB(questUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not update Quest. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	var updatedQuest UpdateQuest
	if err := c.BindJSON(&updatedQuest); err != nil {
		utilities.ResponseMessage(c, "Could not update Quest. Please try again.", http.StatusBadRequest, nil)
		return
	}

	if updatedQuest.CampaignUUID != nil {
		quest.CampaignUUID = *updatedQuest.CampaignUUID
	}
	if updatedQuest.Name != nil {
		quest.Name = *updatedQuest.Name
	}
	if updatedQuest.Description != nil {
		quest.Description = *updatedQuest.Description
	}
	if updatedQuest.PartyUUIDs != nil {
		quest.PartyUUIDs = updatedQuest.PartyUUIDs
	}
	if updatedQuest.NpcUUIDs != nil {
		quest.NpcUUIDs = updatedQuest.NpcUUIDs
	}
	if updatedQuest.BossUUIDs != nil {
		quest.BossUUIDs = updatedQuest.BossUUIDs
	}
	if updatedQuest.LocationUUIDs != nil {
		quest.LocationUUIDs = updatedQuest.LocationUUIDs
	}
	if updatedQuest.RewardUUIDs != nil {
		quest.RewardUUIDs = updatedQuest.RewardUUIDs
	}
	if updatedQuest.StartingSessionUUID != nil {
		quest.StartingSessionUUID = *updatedQuest.StartingSessionUUID
	}
	if updatedQuest.EndingSessionUUID != nil {
		quest.EndingSessionUUID = *updatedQuest.EndingSessionUUID
	}
	if updatedQuest.Status != nil {
		quest.Status = *updatedQuest.Status
	}
	if updatedQuest.Deleted != nil {
		quest.Deleted = *updatedQuest.Deleted
	}

	if err := UpdateQuestByUuidDB(&quest, campaignUuid); err != nil {
		utilities.ResponseMessage(c, "Could not update Quest. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Quest updated successfully.", http.StatusOK, gin.H{"quest": CreateQuestResponse(quest)})
}

func CreateQuestResponse(quest campaign.Quest) QuestResponse {
	return QuestResponse{
		UUID:                &quest.UUID,
		CampaignUUID:        &quest.CampaignUUID,
		CreatorUUID:         &quest.CreatorUUID,
		Name:                &quest.Name,
		Description:         &quest.Description,
		QuestGiverUUID:      &quest.QuestGiverUUID,
		RewardUUIDs:         quest.RewardUUIDs,
		LocationUUIDs:       quest.LocationUUIDs,
		NpcUUIDs:            quest.NpcUUIDs,
		PartyUUIDs:          quest.PartyUUIDs,
		BossUUIDs:           quest.BossUUIDs,
		StartingSessionUUID: &quest.StartingSessionUUID,
		EndingSessionUUID:   &quest.EndingSessionUUID,
		Status:              &quest.Status,
		Created:             quest.Created,
		LastUpdated:         quest.LastUpdated,
	}
}

func mapQuestsToResponse(npcs []campaign.Quest) []QuestResponse {
	response := make([]QuestResponse, len(npcs))
	for i, npc := range npcs {
		response[i] = CreateQuestResponse(npc)
	}
	return response
}
