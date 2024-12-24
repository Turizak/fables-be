package quest

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

func CreateQuest(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

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
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}
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
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}
	campaignUuid := c.Param("uuid")
	quests, err := GetQuestsByCampaignUuidDB(campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve quests. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Quests retrieved successfully.", http.StatusOK, gin.H{"quests": mapQuestsToResponse(quests)})
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
