package condition

import (
	"github.com/Turizak/fables-be/database"
)

func GetAllConditionsWithDescriptionsDB() ([]Condition, error) {
	var conditions []Condition

	// Preload the ConditionDescriptions to load related data
	err := database.DB.Preload("Descriptions").Find(&conditions).Error
	if err != nil {
		return nil, err
	}

	return conditions, nil
}

func GetConditionByIndexDB(index string) (Condition, error) {
	var condition Condition

	// Preload the ConditionDescriptions to load related data
	err := database.DB.Preload("Descriptions").First(&condition, "index = ?", index).Error
	if err != nil {
		return Condition{}, err
	}

	return condition, nil
}
