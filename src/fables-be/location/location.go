package location

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

func CreateLocation(c *gin.Context) {
	var location campaign.Location
	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	location.CreatorUUID = claims.UUID

	campaignUuid := c.Param("uuid")
	if err := c.BindJSON(&location); err != nil {
		utilities.ResponseMessage(c, "Could not create location. Please try again.", http.StatusBadRequest, nil)
		return
	}
	location.CampaignUUID = campaignUuid

	if err := CreateLocationDB(&location, location.CampaignUUID); err != nil {
		utilities.ResponseMessage(c, "Could not create location. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Location created successfully.", http.StatusCreated, gin.H{"location": CreateLocationResponse(location)})
}

func CreateLocationSession(c *gin.Context) {
	var location campaign.Location
	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	location.CreatorUUID = claims.UUID

	campaignUuid := c.Param("uuid")
	if err := c.BindJSON(&location); err != nil {
		utilities.ResponseMessage(c, "Could not create location. Please try again.", http.StatusBadRequest, nil)
		return
	}
	location.CampaignUUID = campaignUuid

	sessionUuid := c.Param("sessionUuid")

	if err := CreateLocationSessionDB(&location, location.CampaignUUID, sessionUuid); err != nil {
		utilities.ResponseMessage(c, "Could not create location. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Location created successfully.", http.StatusCreated, gin.H{"location": CreateLocationResponse(location)})
}

func GetLocationByUuid(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	campaignUuid := c.Param("uuid")
	locationUuid := c.Param("locationUuid")
	location, err := GetLocationByUuidDB(locationUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve location. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Location retrieved successfully.", http.StatusOK, gin.H{"location": CreateLocationResponse(*location)})
}

func GetLocationsByCampaignUuid(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	campaignUuid := c.Param("uuid")
	locations, err := GetLocationsByCampaignUuidDB(campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve locations. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Locations retrieved successfully.", http.StatusOK, gin.H{"locations": mapLocationsToResponse(locations)})
}

func GetLocationsByCreatorUuid(c *gin.Context) {
	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	locations, err := GetLocationsByCreatorUuidDB(claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve locations. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Locations retrieved successfully.", http.StatusOK, gin.H{"locations": mapLocationsToResponse(locations)})
}

func UpdateLocationByUuid(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	campaignUuid := c.Param("uuid")
	locationUuid := c.Param("locationUuid")
	location, err := GetLocationByUuidDB(locationUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve location. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	var updateLocation UpdateLocation
	if err := c.BindJSON(&updateLocation); err != nil {
		utilities.ResponseMessage(c, "Could not update location. Please try again.", http.StatusBadRequest, nil)
		return
	}

	if updateLocation.CampaignUUID != nil {
		location.CampaignUUID = *updateLocation.CampaignUUID
	}
	if updateLocation.Name != nil {
		location.Name = *updateLocation.Name
	}
	if updateLocation.Description != nil {
		location.Description = *updateLocation.Description
	}
	if updateLocation.Deleted != nil {
		location.Deleted = *updateLocation.Deleted
	}

	if err := UpdateLocationByUuidDB(location, campaignUuid); err != nil {
		utilities.ResponseMessage(c, "Could not update location. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Location updated successfully.", http.StatusOK, gin.H{"location": CreateLocationResponse(*location)})
}

// Helper functions to map locations to their response formats
func CreateLocationResponse(location campaign.Location) LocationResponse {
	return LocationResponse{
		UUID:         utilities.ToPointer(location.UUID),
		CampaignUUID: utilities.ToPointer(location.CampaignUUID),
		CreatorUUID:  utilities.ToPointer(location.CreatorUUID),
		Name:         utilities.ToPointer(location.Name),
		Description:  utilities.ToPointer(location.Description),
		Created:      location.Created,
		LastUpdated:  location.LastUpdated,
	}
}

func mapLocationsToResponse(locations []campaign.Location) []LocationResponse {
	response := make([]LocationResponse, len(locations))
	for i, location := range locations {
		response[i] = CreateLocationResponse(location)
	}
	return response
}
