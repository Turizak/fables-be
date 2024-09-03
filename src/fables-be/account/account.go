package account

import (
	"net/http"
	"net/mail"
	"strings"

	"github.com/Turizak/fables-be/utilities"

	"github.com/gin-gonic/gin"
)

func CreateAccount(c *gin.Context) {
	var acc Account

	// Bind the request body to the Account struct
	if err := c.BindJSON(&acc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Could not create account. Please try again.",
			"status":     http.StatusBadRequest,
			"statusText": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	// Validate the email address
	addr, err := mail.ParseAddress(strings.ToLower(acc.Email))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Could not create account. Please try again.",
			"status":     http.StatusBadRequest,
			"statusText": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	// Validate the password
	validPass := utilities.ValidatePasswordComplexity(acc.Password)
	if !validPass {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Could not create account. Please try again.",
			"status":     http.StatusBadRequest,
			"statusText": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	// Hash the password
	hashedPassword, err := utilities.HashPassword(acc.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Could not create account. Please try again.",
			"status":     http.StatusBadRequest,
			"statusText": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	// Set the email address and password
	acc.Email = strings.ToLower(addr.Address)
	acc.Password = hashedPassword

	// Create the account
	if err := CreateAccountDB(&acc); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message":    "Could not create account. Please try again.",
			"status":     http.StatusBadRequest,
			"statusText": http.StatusText(http.StatusBadRequest),
		})
		return
	}

	// Return the account
	c.JSON(http.StatusCreated, gin.H{
		"message":    "Account created successfully.",
		"status":     http.StatusCreated,
		"statusText": http.StatusText(http.StatusCreated),
		"account": gin.H{
			"uuid":      acc.UUID,
			"email":     acc.Email,
			"firstName": acc.FirstName,
			"lastName":  acc.LastName,
		},
	})
}
