package database

import (
	"github.com/BloomGameStudio/PuzzleService/models"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func ConnectDB() (*gorm.DB, error) {
    db, err := gorm.Open(sqlite.Open("database/database.db"), &gorm.Config{})
    if err != nil {
        return nil, err
    }

    // Migrate the schema
    db.AutoMigrate(&models.Puzzle{})

    return db, nil
}

func InitDB() (*gorm.DB, func(), error) {
    db, err := ConnectDB()
    if err != nil {
        return nil, nil, err
    }

    return db, func() {
        sqlDB, err := db.DB()
        if err != nil {
			return
        }
        sqlDB.Close()
    }, nil
}