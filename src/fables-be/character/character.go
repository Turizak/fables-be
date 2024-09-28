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
	character.Public = true

	if err := CreateCharacterDB(&character, character.CampaignUUID); err != nil {
		utilities.ResponseMessage(c, "Could not create character. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseCharacter := CreateCharacterResponse(character)
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
	responseCharacter := CreateCharacterResponse(*character)
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
		responseCharacter := CreateCharacterResponse(character)
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
		responseCharacter := CreateCharacterResponse(character)
		responseCharacters = append(responseCharacters, responseCharacter)
	}
	utilities.ResponseMessage(c, "Characters retrieved successfully.", http.StatusOK, gin.H{"characters": responseCharacters})
}

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
