package model

import (
	"github.com/infotitanz/dancerapy-api/internal/database"
	"gorm.io/gorm"
)

type Training struct {
	gorm.Model
	Title       string `json:"title"`
	Description string `json:"description"`
	Tutor       string `json:"tutor"`
	UserID      uint
}

func GetTraining(ID uint) (Training, error) {
	var training Training
	if result := database.Database.First(&training, ID); result.Error != nil {
		return Training{}, result.Error
	}
	return training, nil
}

func (training *Training) CreateTraining() (*Training, error) {
	err := database.Database.Create(&training).Error
	if err != nil {
		return training, nil
	}

	return training, nil
}

func UpdateTraining(ID uint, newTraining Training) (Training, error) {
	training, err := GetTraining(ID)
	if err != nil {
		return Training{}, err
	}

	if result := database.Database.Model(&training).Updates(newTraining); result.Error != nil {
		return Training{}, result.Error
	}
	return training, nil
}

func DeleteTraining(ID uint) error {
	training, err := GetTraining(ID)
	if err != nil {
		return err
	}

	if result := database.Database.Delete(&training); result.Error != nil {
		return result.Error
	}
	return nil
}
