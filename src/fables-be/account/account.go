package account

import (
	"net/http"
	"net/mail"
	"strings"

	"github.com/Turizak/fables-be/token"
	"github.com/Turizak/fables-be/utilities"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var account Account
	if err := c.BindJSON(&account); err != nil {
		utilities.ResponseMessage(c, "Could not create account. Please try again.", http.StatusBadRequest, nil)
		return
	}
	addr, err := mail.ParseAddress(strings.ToLower(account.Email))
	if err != nil {
		utilities.ResponseMessage(c, "Could not create account. Please try again.", http.StatusBadRequest, nil)
		return
	}
	validPass := utilities.ValidatePasswordComplexity(account.Password)
	if !validPass {
		utilities.ResponseMessage(c, "Could not create account. Please try again.", http.StatusBadRequest, nil)
		return
	}
	hashedPassword, err := utilities.HashPassword(account.Password)
	if err != nil {
		utilities.ResponseMessage(c, "Could not create account. Please try again.", http.StatusBadRequest, nil)
		return
	}
	account.Email = strings.ToLower(addr.Address)
	account.Password = hashedPassword
	account.Username = strings.ToLower(account.Username)
	if err := CreateAccountDB(&account); err != nil {
		utilities.ResponseMessage(c, "Could not create account. Please try again.", http.StatusBadRequest, nil)
		return
	}
	responseAccount := AccountResponse{
		Username:  account.Username,
		Email:     account.Email,
		UUID:      account.UUID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Created:   account.Created,
	}
	utilities.ResponseMessage(c, "Account created successfully.", http.StatusCreated, gin.H{
		"account": responseAccount,
	})
}

func AccountLogin(c *gin.Context) {
	var requestAccount Account
	if err := c.BindJSON(&requestAccount); err != nil {
		utilities.ResponseMessage(c, "Login Failed. Please try again.", http.StatusBadRequest, nil)
		return
	}
	account, err := GetAccountByEmailDB(requestAccount.Email)
	if err != nil {
		utilities.ResponseMessage(c, "Login Failed. Please try again.", http.StatusBadRequest, nil)
		return
	}
	if !utilities.DoPasswordsMatch(account.Password, requestAccount.Password) {
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
	var tokenResponse = LoginResponse{
		AccessToken:  authToken,
		RefreshToken: refreshToken,
	}
	utilities.ResponseMessage(c, "Login Successful.", http.StatusOK, gin.H{"tokens": tokenResponse})
}

func ValidateAuthToken(c *gin.Context) {
	utilities.ResponseMessage(c, "Token is valid.", http.StatusOK, nil)
}

func RefreshAuthToken(c *gin.Context) {
	var requestRefreshToken RefreshToken
	if err := c.BindJSON(&requestRefreshToken); err != nil {
		return
	}
	_, validRefToken := utilities.ValidateRefreshToken(c, requestRefreshToken.RefreshToken)
	if !validRefToken {
		utilities.ResponseMessage(c, "Unauthorized.", http.StatusUnauthorized, nil)
		return
	}
	claims, err := token.ParseRefreshToken(requestRefreshToken.RefreshToken)
	if err != nil {
		utilities.ResponseMessage(c, "An error occurred parsing the refresh token. Please try again.", http.StatusBadRequest, nil)
		return
	}
	authToken, err := token.CreateToken(claims.Email, claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "An error occurred creating the token. Please try again.", http.StatusBadRequest, nil)
		return
	}
	refreshToken, err := token.CreateRefreshToken(claims.Email, claims.UUID)
	if err != nil {
		utilities.ResponseMessage(c, "An error occurred creating the refresh token. Please try again.", http.StatusBadRequest, nil)
		return
	}
	tokenResponse := LoginResponse{
		AccessToken:  authToken,
		RefreshToken: refreshToken,
	}
	utilities.ResponseMessage(c, "Token refreshed successfully.", http.StatusOK, gin.H{"tokens": tokenResponse})
}

func GetAccount(c *gin.Context) {
	claims, _ := c.Get("claims")
	account, err := GetAccountByEmailDB(claims.(*token.UserClaim).Email)
	if err != nil {
		utilities.ResponseMessage(c, "Account not found.", http.StatusNotFound, nil)
		return
	}
	var responseAccount = AccountResponse{
		Username:  account.Username,
		Email:     account.Email,
		UUID:      account.UUID,
		FirstName: account.FirstName,
		LastName:  account.LastName,
		Created:   account.Created,
	}
	utilities.ResponseMessage(c, "Account found.", http.StatusOK, gin.H{"account": responseAccount})
}

func ChangePassword(c *gin.Context) {
	claims, _ := c.Get("claims")
	var requestChangePassword UpdatePassword
	if err := c.BindJSON(&requestChangePassword); err != nil {
		utilities.ResponseMessage(c, "Error: A error occurred. Please try again.", http.StatusBadRequest, nil)
		return
	}
	validPass := utilities.ValidatePasswordComplexity(requestChangePassword.NewPassword)
	if !validPass {
		utilities.ResponseMessage(c, "New Password does not meet complexity requirements.", http.StatusBadRequest, nil)
		return
	}
	account, err := GetAccountByEmailDB(claims.(*token.UserClaim).Email)
	if err != nil {
		utilities.ResponseMessage(c, "Account not found.", http.StatusNotFound, nil)
		return
	}
	if !utilities.DoPasswordsMatch(account.Password, requestChangePassword.OldPassword) {
		utilities.ResponseMessage(c, "Could not verify passwords. Please try again.", http.StatusBadRequest, nil)
		return
	}
	if utilities.DoPasswordsMatch(account.Password, requestChangePassword.NewPassword) {
		utilities.ResponseMessage(c, "New password cannot match old password. Please try again.", http.StatusBadRequest, nil)
		return
	}
	hashedPassword, err := utilities.HashPassword(requestChangePassword.NewPassword)
	if err != nil {
		utilities.ResponseMessage(c, "An error occured. Please try again.", http.StatusBadRequest, nil)
		return
	}
	account.Password = hashedPassword
	if err := UpdateAccountPasswordDB(account); err != nil {
		utilities.ResponseMessage(c, "Could not update account. Please try again.", http.StatusBadRequest, nil)
		return
	}
	utilities.ResponseMessage(c, "Password successfully updated.", http.StatusOK, nil)
}
