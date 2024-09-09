package utilities

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ResponseMessage creates a standardized JSON response
func ResponseMessage(c *gin.Context, message string, status int, data gin.H) {
	response := gin.H{
		"message":    message,
		"status":     status,
		"statusText": http.StatusText(status),
		"timestamp":  GetTimestamp(),
	}

	// Conditionally add the "data" field if it's not nil
	if data != nil {
		response["data"] = data
	}

	c.JSON(status, response)
}

// GetTimestamp returns the current timestamp
func GetTimestamp() string {
	return time.Now().Format(time.RFC3339Nano)
}
