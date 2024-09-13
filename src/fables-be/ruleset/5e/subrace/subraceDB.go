package subrace

import "github.com/Turizak/fables-be/database"

func GetAllSubraces5eDB() ([]Subrace, error) {
	subraces := []Subrace{}
	if result := database.DB.Find(&subraces); result.Error != nil {
		return nil, result.Error
	}
	return subraces, nil
}

func GetSubraceByIndex5eDB(index string) (*Subrace, error) {
	subrace := new(Subrace)
	if result := database.DB.Where("index = ?", index).First(subrace); result.Error != nil {
		return nil, result.Error
	}
	return subrace, nil
}
