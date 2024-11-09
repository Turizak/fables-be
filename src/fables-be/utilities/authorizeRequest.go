package utilities

import (
	"net/http"

	"github.com/Turizak/fables-be/token"
	"github.com/gin-gonic/gin"
)

func AuthorizeRequest(c *gin.Context) (*token.UserClaim, bool) {
	authToken := c.GetHeader("Authorization")
	if authToken == "" {
		ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return nil, false
	}
	claims, validToken := ValidateAuthenticationToken(c, authToken)
	if !validToken {
		ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return nil, false
	}
	return claims, true
}
