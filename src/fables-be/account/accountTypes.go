package account

import "github.com/Turizak/fables-be/utilities"

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
	Deleted     bool               `json:"deleted" gorm:"column:deleted"`
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
