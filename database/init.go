package database

import (
	"github.com/BloomGameStudio/PuzzleService/models"
	"gorm.io/gorm"
)

func AutoMigrate(db *gorm.DB) error {
    return db.AutoMigrate(&models.Puzzle{})
}