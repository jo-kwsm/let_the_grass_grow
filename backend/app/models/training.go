package models

import (
	"grass_backend/db"
	"log"

	"gorm.io/gorm"
)

type training struct {
	gorm.Model
	Date  string `json:"date"`
	Count int16  `json:"count"`
}

func (t training) GetByDate(date string) training {
	db := db.GetDB()

	var record training
	if result := db.Where("date = ?", date).First(&record); result != nil {
		log.Fatal("Could not find record")
	}

	return record
}
