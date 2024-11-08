package session

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/utilities"

	"github.com/gin-gonic/gin"
)

func CreateSession(c *gin.Context) {
	var session campaign.Session
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
	if err := c.BindJSON(&session); err != nil {
		utilities.ResponseMessage(c, "Could not create session. Please try again.", http.StatusBadRequest, nil)
		return
	}
	session.CreatorUUID = claims.UUID
	session.CampaignUUID = campaignUuid
	if err := CreateSessionDB(&session, session.CampaignUUID); err != nil {
		utilities.ResponseMessage(c, "Could not create session. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseSession := CreateSessionResponse(session)
	utilities.ResponseMessage(c, "Session created successfully.", http.StatusCreated, gin.H{"session": responseSession})
}

func GetSessionByUuid(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	campaignUuid := c.Param("uuid")
	sessionUuid := c.Param("sessionUuid")
	_, validToken := utilities.ValidateAuthenticationToken(c, authToken)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	session, err := GetSessionByUuidDB(sessionUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve session. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseSession := CreateSessionResponse(*session)
	utilities.ResponseMessage(c, "Session retrieved successfully.", http.StatusOK, gin.H{"session": responseSession})
}

func GetSessionsByCampaignUuid(c *gin.Context) {
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
	sessions, err := GetSessionsByCampaignUuidDB(campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve sessions. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseSessions := make([]SessionResponse, 0)
	for _, ses := range sessions {
		responseSession := CreateSessionResponse(ses)
		responseSessions = append(responseSessions, responseSession)
	}
	utilities.ResponseMessage(c, "Sessions retrieved successfully.", http.StatusOK, gin.H{"sessions": responseSessions})
}

func GetSessionsByCreatorUuid(c *gin.Context) {
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
	sessions, err := GetSessionsByCreatorUuidDB(claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve sessions. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseSessions := make([]SessionResponse, 0)
	for _, ses := range sessions {
		responseSession := CreateSessionResponse(ses)
		responseSessions = append(responseSessions, responseSession)
	}
	utilities.ResponseMessage(c, "Sessions retrieved successfully.", http.StatusOK, gin.H{"sessions": responseSessions})
}

func CreateSessionResponse(session campaign.Session) SessionResponse {
	return SessionResponse{
		ID:           session.ID,
		UUID:         utilities.ToPointer(session.UUID),
		CampaignUUID: utilities.ToPointer(session.CampaignUUID),
		CreatorUUID:  utilities.ToPointer(session.CreatorUUID),
		PartyUUIDs:   session.PartyUUIDs,
		DateOccured:  session.DateOccured,
		Created:      session.Created,
		LastUpdated:  session.LastUpdated,
	}
}
