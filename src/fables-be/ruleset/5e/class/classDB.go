package class

import (
	"github.com/Turizak/fables-be/database"
)

func GetAllClasses5eDB() ([]Class, error) {
	classes := []Class{}
	if result := database.DB.Find(&classes); result.Error != nil {
		return nil, result.Error
	}
	return classes, nil
}

func GetAllClassesWithDetails5eDB() ([]Class, error) {
	var classes []Class

	// Use Preload to load all related tables
	if err := database.DB.Preload("ProficiencyChoices").
		Preload("Proficiencies").
		Preload("SavingThrows").
		// Preload("StartingEquipment").
		// Preload("StartingEquipmentOptions").
		Find(&classes).Error; err != nil {
		return nil, err
	}

	return classes, nil
}

func GetClassByIndex5eDB(index string) (*Class, error) {
	class := new(Class)
	if result := database.DB.Where("index = ?", index).First(class); result.Error != nil {
		return nil, result.Error
	}
	return class, nil
}

func GetClassByIndexWithDetails5eDB(index string) (*Class, error) {
	class := new(Class)

	// Use Preload to load all related tables
	if err := database.DB.Preload("ProficiencyChoices").
		Preload("Proficiencies").
		Preload("SavingThrows").
		// Preload("StartingEquipment").
		// Preload("StartingEquipmentOptions").
		Where("index = ?", index).First(class).Error; err != nil {
		return nil, err
	}

	return class, nil
}
