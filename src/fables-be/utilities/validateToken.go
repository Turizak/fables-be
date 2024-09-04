package utilities

import (
	"net/http"
	"time"

	"github.com/Turizak/fables-be/token"
	"github.com/gin-gonic/gin"
)

func ValidateAuthToken(c *gin.Context, authToken string) (*token.UserClaim, bool) {
	validToken := CheckToken(c, authToken)
	expire := CheckTokenNotExpired(c, authToken)

	if !validToken {
		return nil, false
	}
	if !expire {
		return nil, false
	}
	claims, err := token.ParseToken(authToken)
	if err != nil {
		ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return nil, false
	}
	return claims, true
}

func ValidateRefreshAuthentication(c *gin.Context, refreshToken string) (*token.RefreshUserClaim, bool) {
	validToken := CheckRefreshToken(c, refreshToken)
	expire := CheckRefreshTokenNotExpired(c, refreshToken)

	if !validToken {
		ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return nil, false
	}
	if !expire {
		ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return nil, false
	}
	claims, err := token.ParseRefreshToken(refreshToken)
	if err != nil {
		ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return nil, false
	}
	return claims, true
}

func CheckToken(c *gin.Context, authToken string) bool {
	if authToken == "" {
		ResponseMessage(c, "No token found.", http.StatusBadRequest, nil)
		return false
	}
	if err := token.VerifyToken(authToken); err != nil {
		ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return false
	}
	return true
}

func CheckRefreshToken(c *gin.Context, refreshToken string) bool {
	if refreshToken == "" {
		ResponseMessage(c, "No token found.", http.StatusBadRequest, nil)
		return false
	}
	if err := token.VerifyRefreshToken(refreshToken); err != nil {
		ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return false
	}
	return true
}

func CheckTokenNotExpired(c *gin.Context, authToken string) bool {
	claims, err := token.ParseToken(authToken)
	if err != nil {
		ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return false
	}
	if time.Now().Unix() > claims.ExpiresAt {
		ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return false
	}
	return true
}

func CheckRefreshTokenNotExpired(c *gin.Context, refreshToken string) bool {
	claims, err := token.ParseRefreshToken(refreshToken)
	if err != nil {
		ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return false
	}
	if time.Now().Unix() > claims.ExpiresAt {
		ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return false
	}
	return true
}
