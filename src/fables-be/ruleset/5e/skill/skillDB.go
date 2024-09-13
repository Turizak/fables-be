package skill

import "github.com/Turizak/fables-be/database"

func GetAllSkills5eDB() ([]Skill, error) {
	skills := []Skill{}
	if result := database.DB.Find(&skills); result.Error != nil {
		return nil, result.Error
	}
	return skills, nil
}

func GetSkillByIndex5eDB(index string) (*Skill, error) {
	skill := new(Skill)
	if result := database.DB.Where("index = ?", index).First(skill); result.Error != nil {
		return nil, result.Error
	}
	return skill, nil
}
