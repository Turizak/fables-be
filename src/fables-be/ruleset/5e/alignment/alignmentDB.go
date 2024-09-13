package alignment

import (
	"github.com/Turizak/fables-be/database"
)

func GetAllAlignments5eDB() ([]Alignment, error) {
	alignments := []Alignment{}
	if result := database.DB.Find(&alignments); result.Error != nil {
		return nil, result.Error
	}
	return alignments, nil
}

func GetAlignmentByIndex5eDB(index string) (*Alignment, error) {
	alignment := new(Alignment)
	if result := database.DB.Where("index = ?", index).First(alignment); result.Error != nil {
		return nil, result.Error
	}
	return alignment, nil
}
