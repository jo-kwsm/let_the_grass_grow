package models

import (
	"grass_backend/db"
	"log"
)

type Training struct {
	User     string `json:"user"`
	Date     string `json:"date"`
	Walking  int16  `json:"walking"`
	Running  int16  `json:"running"`
	Cycling  int16  `json:"cycling"`
	Swimming int16  `json:swimming`
}

func (t Training) GetByDate(user, date string) (*Training, error) {
	db := db.GetDB()

	var record Training
	if result := db.Where("user = ? AND date = ?", user, date).First(&record); result.Error != nil {
		log.Fatal("Could not find record")
		return nil, result.Error
	}

	return &record, nil
}

func (t Training) Find(user, start, end string) ([]Training, error) {
	db := db.GetDB()

	var records []Training

	if result := db.Where("user = ? AND date BETWEEN ? AND ?", user, start, end).Find(&records); result.Error != nil {
		log.Fatal("Could not find record")
		return nil, result.Error
	}

	return records, nil
}

func (t Training) FindAll(user string) ([]Training, error) {
	db := db.GetDB()

	var records []Training

	if result := db.Where("user = ?", user).Find(&records); result.Error != nil {
		log.Fatal("Could not find record")
		return nil, result.Error
	}

	return records, nil
}
