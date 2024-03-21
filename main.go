package main

import (
	"fmt"

	"github.com/BloomGameStudio/PuzzleService/controllers/puzzles"
	"github.com/labstack/echo/v4"
)

func main() {
	// Create an instance of Echo
	e := echo.New()

	// Puzzle Routes
	e.POST("/solve", puzzles.SolveHandler)
	// End of Puzzle Routes

	// Indicate that the server is starting
	fmt.Println("Starting service on port 8080")

	// Start the HTTP server
	e.Logger.Fatal(e.Start(":8080"))
}
