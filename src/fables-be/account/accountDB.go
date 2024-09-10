package account

import (
	"strings"
	"time"

	"github.com/Turizak/fables-be/database"
	"github.com/Turizak/fables-be/utilities"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// CreateAccountDB inserts a new account into the database.
func CreateAccountDB(acc *Account) error {
	acc.Created = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	acc.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Time{}, Valid: false})
	acc.UUID = uuid.NewString()
	acc.Deleted = false
	if result := database.DB.Create(acc); result.Error != nil {
		return result.Error
	}
	return nil
}

// GetAccountByUuidDB retrieves an account from the database by its UUID.
func GetAccountByUuidDB(uuid string) (*Account, error) {
	var account Account
	result := database.DB.Where("uuid = ?", uuid).First(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

// GetAccountByEmailDB retrieves an account from the database by its email.
func GetAccountByEmailDB(email string) (*Account, error) {
	var account Account
	// Use LOWER(email) to make the search case-insensitive
	result := database.DB.Where("LOWER(email) = ?", strings.ToLower(email)).First(&account)
	if result.Error != nil {
		return nil, result.Error
	}
	return &account, nil
}

func UpdateAccountPasswordDB(account *Account) error {
	account.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	result := database.DB.Save(account)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
