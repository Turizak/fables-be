package trait

import (
	"github.com/Turizak/fables-be/database"
)

func GetAllTraits5eDB() ([]Trait, error) {
	traits := []Trait{}
	if result := database.DB.Find(&traits); result.Error != nil {
		return nil, result.Error
	}
	return traits, nil
}

func GetTraitByIndex5eDB(index string) (*Trait, error) {
	trait := new(Trait)
	if result := database.DB.Where("index = ?", index).First(trait); result.Error != nil {
		return nil, result.Error
	}
	return trait, nil
}
