package collection

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/character"
	"github.com/Turizak/fables-be/location"
	"github.com/Turizak/fables-be/note"
	"github.com/Turizak/fables-be/npc"
	"github.com/Turizak/fables-be/quest"
	"github.com/Turizak/fables-be/session"
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
	sessions, err := session.GetSessionsByCampaignUuidDB(uuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve sessions. Please try again.", http.StatusBadRequest, nil)
		return
	}
	quests, err := quest.GetQuestsByCampaignUuidDB(uuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve quests. Please try again.", http.StatusBadRequest, nil)
		return
	}
	notes, err := note.GetNotesByCampaignUuidDB(uuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve notes. Please try again.", http.StatusBadRequest, nil)
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
	responseSessions := []session.SessionResponse{}
	for _, sessionData := range sessions {
		responseSession := session.CreateSessionResponse(sessionData)
		responseSessions = append(responseSessions, responseSession)
	}
	responseQuests := []quest.QuestResponse{}
	for _, questData := range quests {
		responseQuest := quest.CreateQuestResponse(questData)
		responseQuests = append(responseQuests, responseQuest)
	}
	responseNotes := []note.NoteResponse{}
	for _, noteData := range notes {
		responseNote := note.CreateNoteResponse(noteData)
		responseNotes = append(responseNotes, responseNote)
	}
	utilities.ResponseMessage(c, "Campaign data retrieved successfully.", http.StatusOK, gin.H{"campaign": responseCampaign, "locations": responseLocations, "npcs": responseNpcs, "notes": responseNotes, "characters": responseCharacters, "sessions": responseSessions, "quests": responseQuests})
}

func GetAllSessionData(c *gin.Context) {
	uuid := c.Param("uuid")
	sessionUuid := c.Param("sessionUuid")

	sessionData, err := session.GetSessionByUuidDB(sessionUuid, uuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve session. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseCharacters := []character.CharacterResponse{}
	for _, characterUuid := range sessionData.PartyUUIDs {
		characterData, err := character.GetCharacterByUuidDB(characterUuid, uuid)
		if err != nil {
			utilities.ResponseMessage(c, "Could not retrieve character. Please try again.", http.StatusBadRequest, nil)
			return
		}
		responseCharacter := character.CreateCharacterResponse(*characterData)
		responseCharacters = append(responseCharacters, responseCharacter)
	}
	responseLocations := []location.LocationResponse{}
	for _, locationUuid := range sessionData.LocationUUIDs {
		locationData, err := location.GetLocationByUuidDB(locationUuid, uuid)
		if err != nil {
			utilities.ResponseMessage(c, "Could not retrieve location. Please try again.", http.StatusBadRequest, nil)
			return
		}
		responseLocation := location.CreateLocationResponse(*locationData)
		responseLocations = append(responseLocations, responseLocation)
	}
	responseNpcs := []npc.NpcResponse{}
	for _, npcUuid := range sessionData.NpcUUIDs {
		npcData, err := npc.GetNpcByUuidDB(npcUuid, uuid)
		if err != nil {
			utilities.ResponseMessage(c, "Could not retrieve npc. Please try again.", http.StatusBadRequest, nil)
			return
		}
		responseNpc := npc.CreateNpcResponse(*npcData)
		responseNpcs = append(responseNpcs, responseNpc)
	}
	utilities.ResponseMessage(c, "Session data retrieved successfully.", http.StatusOK, gin.H{"session": session.CreateSessionResponse(*sessionData), "characters": responseCharacters, "locations": responseLocations, "npcs": responseNpcs})
}
