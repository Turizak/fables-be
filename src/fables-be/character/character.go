package character

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/utilities"

	"github.com/gin-gonic/gin"
)

func CreateCharacter(c *gin.Context) {
	var character campaign.Character
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
	if err := c.BindJSON(&character); err != nil {
		utilities.ResponseMessage(c, "Could not create character. Please try again.", http.StatusBadRequest, nil)
		return
	}
	character.CreatorUUID = claims.UUID
	character.OwnerUUID = claims.UUID
	character.CampaignUUID = campaignUuid
	if err := CreateCharacterDB(&character, character.CampaignUUID); err != nil {
		utilities.ResponseMessage(c, "Could not create character. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseCharacter := CharacterResponse{
		UUID:         character.UUID,
		CampaignUUID: character.CampaignUUID,
		CreatorUUID:  character.CreatorUUID,
		OwnerUUID:    character.OwnerUUID,
		FirstName:    character.FirstName,
		LastName:     character.LastName,
		Race:         character.Race,
		Class:        character.Class,
		Age:          character.Age,
		Height:       character.Height,
		Weight:       character.Weight,
		EyeColor:     character.EyeColor,
		SkinColor:    character.SkinColor,
		HairColor:    character.HairColor,
		Ruleset:      character.Ruleset,
		Created:      character.Created,
		LastUpdated:  character.LastUpdated,
	}
	utilities.ResponseMessage(c, "Character created successfully.", http.StatusCreated, gin.H{"character": responseCharacter})
}

func GetCharacterByUuid(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	campaignUuid := c.Param("uuid")
	characterUuid := c.Param("characterUuid")
	claims, validToken := utilities.ValidateAuthenticationToken(c, authToken)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	character, err := GetCharacterByUuidDB(characterUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve character. Please try again.", http.StatusBadRequest, nil)
		return
	}
	if character.CreatorUUID != claims.UUID {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	responseCharacter := CharacterResponse{
		UUID:         character.UUID,
		CampaignUUID: character.CampaignUUID,
		CreatorUUID:  character.CreatorUUID,
		OwnerUUID:    character.OwnerUUID,
		FirstName:    character.FirstName,
		LastName:     character.LastName,
		Race:         character.Race,
		Class:        character.Class,
		Age:          character.Age,
		Height:       character.Height,
		Weight:       character.Weight,
		EyeColor:     character.EyeColor,
		SkinColor:    character.SkinColor,
		HairColor:    character.HairColor,
		Ruleset:      character.Ruleset,
		Created:      character.Created,
		LastUpdated:  character.LastUpdated,
	}
	utilities.ResponseMessage(c, "Character retrieved successfully.", http.StatusOK, gin.H{"character": responseCharacter})
}

func GetCharactersByCreatorUuid(c *gin.Context) {
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
	characters, err := GetCharactersByCreatorUuidDB(claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve characters. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseCharacters := make([]CharacterResponse, 0)
	for _, character := range characters {
		responseCharacter := CharacterResponse{
			UUID:         character.UUID,
			CampaignUUID: character.CampaignUUID,
			CreatorUUID:  character.CreatorUUID,
			OwnerUUID:    character.OwnerUUID,
			FirstName:    character.FirstName,
			LastName:     character.LastName,
			Race:         character.Race,
			Class:        character.Class,
			Age:          character.Age,
			Height:       character.Height,
			Weight:       character.Weight,
			EyeColor:     character.EyeColor,
			SkinColor:    character.SkinColor,
			HairColor:    character.HairColor,
			Ruleset:      character.Ruleset,
			Created:      character.Created,
			LastUpdated:  character.LastUpdated,
		}
		responseCharacters = append(responseCharacters, responseCharacter)
	}
	utilities.ResponseMessage(c, "Characters retrieved successfully.", http.StatusOK, gin.H{"characters": responseCharacters})
}

func GetCharactersByOwnerUuid(c *gin.Context) {
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
	characters, err := GetCharactersByOwnerUuidDB(claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve characters. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseCharacters := make([]CharacterResponse, 0)
	for _, character := range characters {
		responseCharacter := CharacterResponse{
			UUID:         character.UUID,
			CampaignUUID: character.CampaignUUID,
			CreatorUUID:  character.CreatorUUID,
			OwnerUUID:    character.OwnerUUID,
			FirstName:    character.FirstName,
			LastName:     character.LastName,
			Race:         character.Race,
			Class:        character.Class,
			Age:          character.Age,
			Height:       character.Height,
			Weight:       character.Weight,
			EyeColor:     character.EyeColor,
			SkinColor:    character.SkinColor,
			HairColor:    character.HairColor,
			Ruleset:      character.Ruleset,
			Created:      character.Created,
			LastUpdated:  character.LastUpdated,
		}
		responseCharacters = append(responseCharacters, responseCharacter)
	}
	utilities.ResponseMessage(c, "Characters retrieved successfully.", http.StatusOK, gin.H{"characters": responseCharacters})
}
