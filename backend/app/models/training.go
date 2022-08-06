package models

import (
	"grass_backend/db"
	"log"
)

type Training struct {
	Date  string `json:"date"`
	Count int16  `json:"count"`
}

func (t Training) GetByDate(date string) (*Training, error) {
	db := db.GetDB()

	var record Training
	if result := db.Where("date = ?", date).First(&record); result.Error != nil {
		log.Fatal("Could not find record")
		return nil, result.Error
	}

	return &record, nil
}

func (t Training) FindAll() ([]Training, error) {
	db := db.GetDB()

	var records []Training

	if result := db.Find(&records); result.Error != nil {
		log.Fatal("Could not find record")
		return nil, result.Error
	}

	return records, nil
}
