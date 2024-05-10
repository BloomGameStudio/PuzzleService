package database

import (
	publicModel "github.com/BloomGameStudio/PuzzleService/publicModels"
	puzzlegorm "github.com/BloomGameStudio/PuzzleService/puzzle/repository"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func Open() *gorm.DB {
	// memoryDB := "file::memory:?cache=shared"
	fileDB := "database/database.db"

	db, err := gorm.Open(sqlite.Open(fileDB), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
		// Logger:                 logger.Default.LogMode(logger.Info), // Show Gorm SQL Queries
	})

	if err != nil {
		panic("failed to connect database")
	}

	return db
}

func Init() {
	DB = Open()
	DB.AutoMigrate(&puzzlegorm.Puzzle{})
	DB.AutoMigrate(&publicModel.Puzzle{})
}

func GetDB() *gorm.DB {
	return DB
}
