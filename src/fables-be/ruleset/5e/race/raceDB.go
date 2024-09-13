package race

import "github.com/Turizak/fables-be/database"

func GetAllRaces5eDB() ([]Race, error) {
	races := []Race{}
	if result := database.DB.Find(&races); result.Error != nil {
		return nil, result.Error
	}
	return races, nil
}

func GetRaceByIndex5eDB(index string) (*Race, error) {
	race := new(Race)
	if result := database.DB.Where("index = ?", index).First(race); result.Error != nil {
		return nil, result.Error
	}
	return race, nil
}
