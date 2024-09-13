package abilityscore

import (
	"github.com/Turizak/fables-be/database"
)

func GetAllAbilityScores5eDB() ([]AbilityScore, error) {
	abilityScores := []AbilityScore{}
	if result := database.DB.Find(&abilityScores); result.Error != nil {
		return nil, result.Error
	}
	return abilityScores, nil
}

func GetAbilityScoreByIndex5eDB(index string) (*AbilityScore, error) {
	abilityScore := new(AbilityScore)
	if result := database.DB.Where("index = ?", index).First(abilityScore); result.Error != nil {
		return nil, result.Error
	}
	return abilityScore, nil
}
