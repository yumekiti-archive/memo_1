package utils

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"backend/domain"
)

func NewDB() *gorm.DB {
	db, err := gorm.Open(sqlite.Open("./data.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	db.AutoMigrate(&domain.Memo{})

	return db
}
