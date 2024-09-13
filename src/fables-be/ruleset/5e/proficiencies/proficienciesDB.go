package proficiencies

import (
	"github.com/Turizak/fables-be/database"
)

func GetAllProficiencies5eDB() ([]Proficiency, error) {
	proficiencys := []Proficiency{}
	if result := database.DB.Find(&proficiencys); result.Error != nil {
		return nil, result.Error
	}
	return proficiencys, nil
}

func GetProficiencyByIndex5eDB(index string) (*Proficiency, error) {
	proficiency := new(Proficiency)
	if result := database.DB.Where("index = ?", index).First(proficiency); result.Error != nil {
		return nil, result.Error
	}
	return proficiency, nil
}
