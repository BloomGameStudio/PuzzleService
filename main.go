package main

import (
	"net/http"

	"github.com/BloomGameStudio/PuzzleService/database"
	puzzlehttp "github.com/BloomGameStudio/PuzzleService/puzzle/http"
	puzzlegorm "github.com/BloomGameStudio/PuzzleService/puzzle/repository"
	puzzleuc "github.com/BloomGameStudio/PuzzleService/puzzle/usecase"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Initialize Database
	database.Init()

	// Create a new instance of the Fiber app.
	app := fiber.New()

	// Liveness
	app.Get("/liveness", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	// Puzzle
	puzzleRepository := puzzlegorm.NewPuzzleRepository(database.GetDB())
	puzzleUseCase := puzzleuc.NewPuzzleUseCase(puzzleRepository)
	puzzlehttp.RegisterHTTPEndpoints(app, puzzleUseCase)

	// Start the server and listen on port 3000
	app.Listen(":3000")
}
