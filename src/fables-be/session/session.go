package session

import (
	"net/http"

	"github.com/Turizak/fables-be/campaign"
	"github.com/Turizak/fables-be/utilities"
	"github.com/gin-gonic/gin"
)

// CreateSession handles the creation of a new session
func CreateSession(c *gin.Context) {
	var session campaign.Session

	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	if err := c.BindJSON(&session); err != nil {
		utilities.ResponseMessage(c, "Could not create session. Please try again.", http.StatusBadRequest, nil)
		return
	}

	session.CampaignUUID = c.Param("uuid")
	session.CreatorUUID = claims.UUID

	if err := CreateSessionDB(&session, session.CampaignUUID); err != nil {
		utilities.ResponseMessage(c, "Could not create session. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Session created successfully.", http.StatusCreated, gin.H{
		"session": CreateSessionResponse(session),
	})
}

// GetSessionByUuid retrieves a session by its UUID
func GetSessionByUuid(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}
	session, err := GetSessionByUuidDB(c.Param("sessionUuid"), c.Param("uuid"))
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve session. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Session retrieved successfully.", http.StatusOK, gin.H{
		"session": CreateSessionResponse(*session),
	})
}

// GetSessionsByCampaignUuid retrieves sessions by campaign UUID
func GetSessionsByCampaignUuid(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	sessions, err := GetSessionsByCampaignUuidDB(c.Param("uuid"))
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve sessions. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Sessions retrieved successfully.", http.StatusOK, gin.H{
		"sessions": mapSessionsToResponses(sessions),
	})
}

// GetSessionsByCreatorUuid retrieves sessions by creator UUID
func GetSessionsByCreatorUuid(c *gin.Context) {
	claims, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	sessions, err := GetSessionsByCreatorUuidDB(claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve sessions. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Sessions retrieved successfully.", http.StatusOK, gin.H{
		"sessions": mapSessionsToResponses(sessions),
	})
}

func UpdateSessionByUuid(c *gin.Context) {
	_, authorized := utilities.AuthorizeRequest(c)
	if !authorized {
		return
	}

	campaignUuid := c.Param("uuid")
	sessionUuid := c.Param("sessionUuid")
	session, err := GetSessionByUuidDB(sessionUuid, campaignUuid)
	if err != nil {
		utilities.ResponseMessage(c, "Could not retrieve session. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	var updateSession UpdateSession
	if err := c.BindJSON(&updateSession); err != nil {
		utilities.ResponseMessage(c, "Could not update session. Please try again.", http.StatusBadRequest, nil)
		return
	}

	if updateSession.DateOccured.Valid {
		session.DateOccured = updateSession.DateOccured
	}
	if updateSession.PartyUUIDs != nil {
		session.PartyUUIDs = updateSession.PartyUUIDs
	}
	if updateSession.NpcUUIDs != nil {
		session.NpcUUIDs = updateSession.NpcUUIDs
	}
	if updateSession.LocationUUIDs != nil {
		session.LocationUUIDs = updateSession.LocationUUIDs
	}

	if err := UpdateSessionByUuidDB(session, campaignUuid); err != nil {
		utilities.ResponseMessage(c, "Could not update session. Please try again.", http.StatusInternalServerError, nil)
		return
	}

	utilities.ResponseMessage(c, "Session updated successfully.", http.StatusOK, gin.H{"session": CreateSessionResponse(*session)})
}

// CreateSessionResponse maps a session to a SessionResponse
func CreateSessionResponse(session campaign.Session) SessionResponse {
	return SessionResponse{
		SessionID:     session.SessionID,
		UUID:          utilities.ToPointer(session.UUID),
		CampaignUUID:  utilities.ToPointer(session.CampaignUUID),
		CreatorUUID:   utilities.ToPointer(session.CreatorUUID),
		PartyUUIDs:    session.PartyUUIDs,
		LocationUUIDs: session.LocationUUIDs,
		NpcUUIDs:      session.NpcUUIDs,
		DateOccured:   session.DateOccured,
		Created:       session.Created,
		LastUpdated:   session.LastUpdated,
	}
}

// mapSessionsToResponses maps a slice of sessions to session responses
func mapSessionsToResponses(sessions []campaign.Session) []SessionResponse {
	responseSessions := make([]SessionResponse, len(sessions))
	for i, session := range sessions {
		responseSessions[i] = CreateSessionResponse(session)
	}
	return responseSessions
}
