package database

import (
	puzzlegorm "github.com/BloomGameStudio/PuzzleService/puzzle/repository/gorm"
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
}

func GetDB() *gorm.DB {
	return DB
}
