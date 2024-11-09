package character

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

func CreateCharacter(c *gin.Context) {
	var character campaign.Character
	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	character.CreatorUUID = claims.UUID
	character.OwnerUUID = claims.UUID

	campaignUuid := c.Param("uuid")
	if err := c.BindJSON(&character); err != nil {
		utilities.ResponseMessage(c, "Could not create character. Please try again.", http.StatusBadRequest, nil)
		return
	}
	character.CampaignUUID = campaignUuid
	character.Public = true

	if err := CreateCharacterDB(&character, character.CampaignUUID); err != nil {
		utilities.ResponseMessage(c, "Could not create character. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Character created successfully.", http.StatusCreated, gin.H{"character": CreateCharacterResponse(character)})
}

func GetCharacterByUuid(c *gin.Context) {
	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	campaignUuid := c.Param("uuid")
	characterUuid := c.Param("characterUuid")
	character, err := GetCharacterByUuidDB(characterUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve character. Please try again.", http.StatusInternalServerError, nil)
		return
	}
	if character.CreatorUUID != claims.UUID {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}

	utilities.ResponseMessage(c, "Character retrieved successfully.", http.StatusOK, gin.H{"character": CreateCharacterResponse(*character)})
}

func GetCharactersByCreatorUuid(c *gin.Context) {
	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	characters, err := GetCharactersByCreatorUuidDB(claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve characters. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Characters retrieved successfully.", http.StatusOK, gin.H{"characters": MapCharactersToResponse(characters)})
}

func GetCharactersByOwnerUuid(c *gin.Context) {
	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	characters, err := GetCharactersByOwnerUuidDB(claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve characters. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Characters retrieved successfully.", http.StatusOK, gin.H{"characters": MapCharactersToResponse(characters)})
}

func UpdateCharacterByUuid(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	campaignUuid := c.Param("uuid")
	characterUuid := c.Param("characterUuid")
	character, err := GetCharacterByUuidDB(characterUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve character. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	var updateCharacter UpdateCharacter
	if err := c.BindJSON(&updateCharacter); err != nil {
		utilities.ResponseMessage(c, "Could not update character. Please try again.", http.StatusBadRequest, nil)
		return
	}

	if updateCharacter.FirstName != nil {
		character.FirstName = *updateCharacter.FirstName
	}
	if updateCharacter.LastName != nil {
		character.LastName = *updateCharacter.LastName
	}
	if updateCharacter.OwnerUUID != nil {
		character.OwnerUUID = *updateCharacter.OwnerUUID
	}
	if updateCharacter.CampaignUUID != nil {
		character.CampaignUUID = *updateCharacter.CampaignUUID
	}
	if updateCharacter.Race != nil {
		character.Race = *updateCharacter.Race
	}
	if updateCharacter.Class != nil {
		character.Class = *updateCharacter.Class
	}
	if updateCharacter.Age != nil {
		character.Age = *updateCharacter.Age
	}
	if updateCharacter.Height != nil {
		character.Height = *updateCharacter.Height
	}
	if updateCharacter.Weight != nil {
		character.Weight = *updateCharacter.Weight
	}
	if updateCharacter.EyeColor != nil {
		character.EyeColor = *updateCharacter.EyeColor
	}
	if updateCharacter.SkinColor != nil {
		character.SkinColor = *updateCharacter.SkinColor
	}
	if updateCharacter.HairColor != nil {
		character.HairColor = *updateCharacter.HairColor
	}
	if updateCharacter.Gender != nil {
		character.Gender = *updateCharacter.Gender
	}
	if updateCharacter.Public != nil {
		character.Public = *updateCharacter.Public
	}
	if updateCharacter.Deleted != nil {
		character.Deleted = *updateCharacter.Deleted
	}

	if err := UpdateCharacterByUuidDB(character, campaignUuid); err != nil {
		utilities.ResponseMessage(c, "Could not update character. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Character updated successfully.", http.StatusOK, gin.H{"character": CreateCharacterResponse(*character)})
}

// Helper functions to map characters to their response formats
func CreateCharacterResponse(character campaign.Character) CharacterResponse {
	return CharacterResponse{
		UUID:         utilities.ToPointer(character.UUID),
		CampaignUUID: utilities.ToPointer(character.CampaignUUID),
		CreatorUUID:  utilities.ToPointer(character.CreatorUUID),
		OwnerUUID:    utilities.ToPointer(character.OwnerUUID),
		FirstName:    utilities.ToPointer(character.FirstName),
		LastName:     utilities.ToPointer(character.LastName),
		Race:         utilities.ToPointer(character.Race),
		Class:        utilities.ToPointer(character.Class),
		Age:          character.Age,
		Height:       character.Height,
		Weight:       character.Weight,
		EyeColor:     utilities.ToPointer(character.EyeColor),
		SkinColor:    utilities.ToPointer(character.SkinColor),
		HairColor:    utilities.ToPointer(character.HairColor),
		Ruleset:      utilities.ToPointer(character.Ruleset),
		Public:       character.Public,
		Gender:       utilities.ToPointer(character.Gender),
		Created:      character.Created,
		LastUpdated:  character.LastUpdated,
	}
}

func MapCharactersToResponse(characters []campaign.Character) []CharacterResponse {
	response := make([]CharacterResponse, len(characters))
	for i, character := range characters {
		response[i] = CreateCharacterResponse(character)
	}
	return response
}
