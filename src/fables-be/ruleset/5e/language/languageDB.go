package language

import (
	"github.com/Turizak/fables-be/database"
)

func GetAllLanguages5eDB() ([]Language, error) {
	languages := []Language{}
	if result := database.DB.Find(&languages); result.Error != nil {
		return nil, result.Error
	}
	return languages, nil
}

func GetLanguageByIndex5eDB(index string) (*Language, error) {
	language := new(Language)
	if result := database.DB.Where("index = ?", index).First(language); result.Error != nil {
		return nil, result.Error
	}
	return language, nil
}
