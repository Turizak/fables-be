package proficiencies

import (
	"github.com/Turizak/fables-be/database"
)

func GetAllProficiencies5eDB() ([]Proficiencie, error) {
	proficiencie := []Proficiencie{}
	if result := database.DB.Find(&proficiencie); result.Error != nil {
		return nil, result.Error
	}
	return proficiencie, nil
}

func GetProficiencyByIndex5eDB(index string) (*Proficiencie, error) {
	proficiencie := new(Proficiencie)
	if result := database.DB.Where("class_index = ?", index).First(proficiencie); result.Error != nil {
		return nil, result.Error
	}
	return proficiencie, nil
}
