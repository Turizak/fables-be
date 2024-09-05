package utilities

import (
	"time"

	"github.com/Turizak/fables-be/token"
	"github.com/gin-gonic/gin"
)

func ValidateAuthenticationToken(c *gin.Context, authToken string) (*token.UserClaim, bool) {
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
		return nil, false
	}
	return claims, true
}

func ValidateRefreshToken(c *gin.Context, refreshToken string) (*token.RefreshUserClaim, bool) {
	validToken := CheckRefreshToken(c, refreshToken)
	expire := CheckRefreshTokenNotExpired(c, refreshToken)

	if !validToken {
		return nil, false
	}
	if !expire {
		return nil, false
	}
	claims, err := token.ParseRefreshToken(refreshToken)
	if err != nil {
		return nil, false
	}
	return claims, true
}

func CheckToken(c *gin.Context, authToken string) bool {
	if authToken == "" {
		return false
	}
	if err := token.VerifyToken(authToken); err != nil {
		return false
	}
	return true
}

func CheckRefreshToken(c *gin.Context, refreshToken string) bool {
	if refreshToken == "" {
		return false
	}
	if err := token.VerifyRefreshToken(refreshToken); err != nil {
		return false
	}
	return true
}

func CheckTokenNotExpired(c *gin.Context, authToken string) bool {
	claims, err := token.ParseToken(authToken)
	if err != nil {
		return false
	}
	if time.Now().Unix() > claims.ExpiresAt {
		return false
	}
	return true
}

func CheckRefreshTokenNotExpired(c *gin.Context, refreshToken string) bool {
	claims, err := token.ParseRefreshToken(refreshToken)
	if err != nil {
		return false
	}
	if time.Now().Unix() > claims.ExpiresAt {
		return false
	}
	return true
}
