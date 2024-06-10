package main

import (
	"github.com/BloomGameStudio/PuzzleService/database"
	"github.com/BloomGameStudio/PuzzleService/ethereum"
	_ "github.com/joho/godotenv/autoload"
)

func init() {
	// Initialize Database
	database.Init()

	// Initialize Ethereum client
	ethereum.Init()
}
