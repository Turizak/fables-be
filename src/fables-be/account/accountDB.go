package account

import (
	"strings"
	"time"

	"github.com/Turizak/fables-be/database"
	"github.com/Turizak/fables-be/utilities"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Account represents the structure of the account table in the database.
type Account struct {
	ID          int32              `json:"id" gorm:"primaryKey:type:int32"`
	UUID        string             `json:"uuid" gorm:"column:uuid"`
	Email       string             `json:"email" gorm:"column:email"`
	Username    string             `json:"username" gorm:"column:username"`
	Password    string             `json:"password" gorm:"column:password"`
	FirstName   string             `json:"firstName" gorm:"column:first_name"`
	LastName    string             `json:"lastName" gorm:"column:last_name"`
	Created     utilities.NullTime `json:"created" gorm:"column:created"`
	LastUpdated utilities.NullTime `json:"lastUpdated" gorm:"column:last_updated"`
}

type AccountResponse struct {
	UUID        string             `json:"uuid"`
	Email       string             `json:"email"`
	Username    string             `json:"username"`
	FirstName   string             `json:"firstName"`
	LastName    string             `json:"lastName"`
	Created     utilities.NullTime `json:"created"`
	LastUpdated utilities.NullTime `json:"lastUpdated" gorm:"column:last_updated"`
}

type RefreshToken struct {
	RefreshToken string `json:"refreshToken"`
}

type LoginResponse struct {
	AccessToken  string `json:"accessToken"`
	RefreshToken string `json:"refreshToken"`
}

type UpdatePassword struct {
	OldPassword string `json:"oldPassword"`
	NewPassword string `json:"newPassword"`
}

// CreateAccountDB inserts a new account into the database.
func CreateAccountDB(acc *Account) error {
	acc.Created = utilities.ToNullTime(pq.NullTime{Time: time.Now(), Valid: true})
	acc.LastUpdated = utilities.ToNullTime(pq.NullTime{Time: time.Time{}, Valid: false})
	acc.UUID = uuid.NewString()
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
