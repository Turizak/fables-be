package damagetype

import (
	"github.com/Turizak/fables-be/database"
)

func GetAllDamageTypes5eDB() ([]DamageType, error) {
	damageTypes := []DamageType{}
	if result := database.DB.Find(&damageTypes); result.Error != nil {
		return nil, result.Error
	}
	return damageTypes, nil
}

func GetDamageTypeByIndex5eDB(index string) (*DamageType, error) {
	damageType := new(DamageType)
	if result := database.DB.Where("index = ?", index).First(damageType); result.Error != nil {
		return nil, result.Error
	}
	return damageType, nil
}
