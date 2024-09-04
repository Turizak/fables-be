package token

import (
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

var secretKey = []byte(os.Getenv("TOKEN_SECRET"))
var refreshSecretKey = []byte(os.Getenv("REFRESH_TOKEN_SECRET"))

type UserClaim struct {
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	ExpiresAt int64  `json:"exp"`
	IssuedAt  int64  `json:"iat"`
	jwt.StandardClaims
}

type RefreshUserClaim struct {
	UUID      string `json:"uuid"`
	Email     string `json:"email"`
	ExpiresAt int64  `json:"exp"`
	IssuedAt  int64  `json:"iat"`
	jwt.StandardClaims
}

// CreateToken generates a JWT token with the provided email and UUID.
func CreateToken(email string, uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, UserClaim{
		UUID:      uuid,
		Email:     email,
		ExpiresAt: time.Now().Add(time.Minute * 2).Unix(),
		IssuedAt:  time.Now().Unix(),
	})
	tokenString, err := token.SignedString(secretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func CreateRefreshToken(email string, uuid string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, RefreshUserClaim{
		UUID:      uuid,
		Email:     email,
		ExpiresAt: time.Now().Add(time.Minute * 120).Unix(),
		IssuedAt:  time.Now().Unix(),
	})
	tokenString, err := token.SignedString(refreshSecretKey)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// VerifyToken verifies the validity of a JWT token.
func VerifyToken(tokenString string) error {
	authToken := strings.Split(tokenString, " ")[1]
	token, err := jwt.Parse(authToken, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

// VerifyRefreshToken verifies the validity of a Refresh JWT token.
func VerifyRefreshToken(tokenString string) error {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return refreshSecretKey, nil
	})
	if err != nil {
		return err
	}
	if !token.Valid {
		return fmt.Errorf("invalid token")
	}
	return nil
}

func ParseToken(tokenString string) (*UserClaim, error) {
	authToken := strings.Split(tokenString, " ")[1]
	token, err := jwt.ParseWithClaims(authToken, &UserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return secretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*UserClaim)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}

func ParseRefreshToken(tokenString string) (*RefreshUserClaim, error) {
	token, err := jwt.ParseWithClaims(tokenString, &RefreshUserClaim{}, func(token *jwt.Token) (interface{}, error) {
		return refreshSecretKey, nil
	})
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(*RefreshUserClaim)
	if !ok {
		return nil, fmt.Errorf("invalid token claims")
	}
	return claims, nil
}
