package location

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/utilities"

	"github.com/gin-gonic/gin"
)

func CreateLocation(c *gin.Context) {
	var location campaign.Location
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
	if err := c.BindJSON(&location); err != nil {
		utilities.ResponseMessage(c, "Could not create location. Please try again.", http.StatusBadRequest, nil)
		return
	}
	location.CreatorUUID = claims.UUID
	location.CampaignUUID = campaignUuid
	if err := CreateLocationDB(&location, location.CampaignUUID); err != nil {
		utilities.ResponseMessage(c, "Could not create location. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseLocation := LocationResponse{
		UUID:         location.UUID,
		CampaignUUID: location.CampaignUUID,
		CreatorUUID:  location.CreatorUUID,
		Name:         location.Name,
		Description:  location.Description,
		Created:      location.Created,
		LastUpdated:  location.LastUpdated,
	}
	utilities.ResponseMessage(c, "Location created successfully.", http.StatusCreated, gin.H{"location": responseLocation})
}

func GetLocationByUuid(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	campaignUuid := c.Param("uuid")
	locationUuid := c.Param("locationUuid")
	_, validToken := utilities.ValidateAuthenticationToken(c, authToken)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	location, err := GetLocationByUuidDB(locationUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve location. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseLocation := LocationResponse{
		UUID:         location.UUID,
		CampaignUUID: location.CampaignUUID,
		CreatorUUID:  location.CreatorUUID,
		Name:         location.Name,
		Description:  location.Description,
		Created:      location.Created,
		LastUpdated:  location.LastUpdated,
	}
	utilities.ResponseMessage(c, "Location retrieved successfully.", http.StatusOK, gin.H{"location": responseLocation})
}

func GetLocationsByCampaignUuid(c *gin.Context) {
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
	locations, err := GetLocationsByCampaignUuidDB(campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve locations. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseLocations := make([]LocationResponse, 0)
	for _, location := range locations {
		responseLocation := LocationResponse{
			UUID:         location.UUID,
			CampaignUUID: location.CampaignUUID,
			CreatorUUID:  location.CreatorUUID,
			Name:         location.Name,
			Description:  location.Description,
			Created:      location.Created,
			LastUpdated:  location.LastUpdated,
		}
		responseLocations = append(responseLocations, responseLocation)
	}
	utilities.ResponseMessage(c, "Locations retrieved successfully.", http.StatusOK, gin.H{"locations": responseLocations})
}

func GetLocationsByCreatorUuid(c *gin.Context) {
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
	locations, err := GetLocationsByCreatorUuidDB(claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve locations. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseLocations := make([]LocationResponse, 0)
	for _, location := range locations {
		responseLocation := LocationResponse{
			UUID:         location.UUID,
			CampaignUUID: location.CampaignUUID,
			CreatorUUID:  location.CreatorUUID,
			Name:         location.Name,
			Description:  location.Description,
			Created:      location.Created,
			LastUpdated:  location.LastUpdated,
		}
		responseLocations = append(responseLocations, responseLocation)
	}
	utilities.ResponseMessage(c, "Locations retrieved successfully.", http.StatusOK, gin.H{"locations": responseLocations})
}
