package account

import (
	"net/http"
	"net/mail"
	"strings"

	"github.com/Turizak/fables-be/database"
	"github.com/Turizak/fables-be/token"
	"github.com/Turizak/fables-be/utilities"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var acc Account
	if err := c.BindJSON(&acc); err != nil {
		utilities.ResponseMessage(c, "Could not create account. Please try again.", http.StatusBadRequest, nil)
		return
	}
	addr, err := mail.ParseAddress(strings.ToLower(acc.Email))
	if err != nil {
		utilities.ResponseMessage(c, "Could not create account. Please try again.", http.StatusBadRequest, nil)
		return
	}
	validPass := utilities.ValidatePasswordComplexity(acc.Password)
	if !validPass {
		utilities.ResponseMessage(c, "Could not create account. Please try again.", http.StatusBadRequest, nil)
		return
	}
	hashedPassword, err := utilities.HashPassword(acc.Password)
	if err != nil {
		utilities.ResponseMessage(c, "Could not create account. Please try again.", http.StatusBadRequest, nil)
		return
	}
	acc.Email = strings.ToLower(addr.Address)
	acc.Password = hashedPassword
	if err := CreateAccountDB(&acc); err != nil {
		utilities.ResponseMessage(c, "Could not create account. Please try again.", http.StatusBadRequest, nil)
		return
	}
	utilities.ResponseMessage(c, "Account created successfully.", http.StatusCreated, gin.H{
		"account": gin.H{
			"uuid":      acc.UUID,
			"email":     acc.Email,
			"firstName": acc.FirstName,
			"lastName":  acc.LastName,
		},
	})
}

func AccountLogin(c *gin.Context) {
	var acc Account
	var account Account
	if err := c.BindJSON(&acc); err != nil {
		return
	}
	if result := database.DB.Where("email = ?", strings.ToLower(acc.Email)).First(&account); result.Error != nil {
		utilities.ResponseMessage(c, "Login Failed. Please try again.", http.StatusBadRequest, nil)
		return
	}
	if !utilities.DoPasswordsMatch(account.Password, acc.Password) {
		utilities.ResponseMessage(c, "Login Failed. Please try again.", http.StatusBadRequest, nil)
		return
	}
	authToken, err := token.CreateToken(account.Email, account.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Login Failed. Please try again.", http.StatusBadRequest, nil)
		return
	}
	refreshToken, err := token.CreateRefreshToken(account.Email, account.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "Login Failed. Please try again.", http.StatusBadRequest, nil)
		return
	}
	utilities.ResponseMessage(c, "Login Successful.", http.StatusOK, gin.H{
		"token":        authToken,
		"refreshToken": refreshToken,
	})
}

func ValidateAuthToken(c *gin.Context) {
	authToken := c.GetHeader("Authorization")
	_, validToken := utilities.ValidateAuthToken(c, authToken)
	if !validToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	utilities.ResponseMessage(c, "Token is valid.", http.StatusOK, nil)
}

func RefreshAuthToken(c *gin.Context) {
	var refreshToken RefreshToken
	if err := c.BindJSON(&refreshToken); err != nil {
		return
	}
	_, validRefToken := utilities.ValidateRefreshAuthentication(c, refreshToken.RefreshToken)
	if !validRefToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	claims, err := token.ParseRefreshToken(refreshToken.RefreshToken)
	if err != nil {
		utilities.ResponseMessage(c, "An error occurred parsing the refresh token. Please try again.", http.StatusBadRequest, nil)
		return
	}
	authToken, err := token.CreateToken(claims.Email, claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "An error occurred creating the token. Please try again.", http.StatusBadRequest, nil)
		return
	}
	refToken, err := token.CreateRefreshToken(claims.Email, claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "An error occurred creating the refresh token. Please try again.", http.StatusBadRequest, nil)
		return
	}
	utilities.ResponseMessage(c, "Token refreshed successfully.", http.StatusOK, gin.H{
		"token":        authToken,
		"refreshToken": refToken,
	})
}
