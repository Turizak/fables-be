package account

import (
	"time"

	"github.com/Turizak/fables-be/database"
	"github.com/google/uuid"
	"github.com/lib/pq"
)

// Account represents the structure of the account table in the database.
type Account struct {
	ID        int32       `json:"id" gorm:"primaryKey:type:int32"`
	Email     string      `json:"email" gorm:"column:email"`
	Password  string      `json:"password" gorm:"column:password"`
	UUID      string      `json:"uuid" gorm:"column:uuid"`
	FirstName string      `json:"firstName" gorm:"column:first_name"`
	LastName  string      `json:"lastName" gorm:"column:last_name"`
	Created   pq.NullTime `json:"created" gorm:"column:created"`
}

// CreateAccountInDB inserts a new account into the database.
func CreateAccountDB(acc *Account) error {
	acc.Created = pq.NullTime{Time: time.Now(), Valid: true}
	acc.UUID = uuid.NewString()
	if result := database.DB.Create(acc); result.Error != nil {
		return result.Error
	}
	return nil
}
