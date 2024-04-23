package main

import (
	"github.com/BloomGameStudio/PuzzleService/database"
	publicmodels "github.com/BloomGameStudio/PuzzleService/publicModels"
)

func init() {

	// Initialize Database
	database.Init()

	// TODO: Cleanup Migration
	db := database.GetDB()

	db.AutoMigrate(&publicmodels.Puzzle{})
}
