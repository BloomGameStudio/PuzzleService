package main

import (
	"net/http"

	"github.com/BloomGameStudio/PuzzleService/database"
	"github.com/BloomGameStudio/PuzzleService/ethereum"
	puzzlehttp "github.com/BloomGameStudio/PuzzleService/puzzle/http"
	puzzleonchain "github.com/BloomGameStudio/PuzzleService/puzzle/repository/ethereum"
	puzzlegorm "github.com/BloomGameStudio/PuzzleService/puzzle/repository/gorm"
	puzzleuc "github.com/BloomGameStudio/PuzzleService/puzzle/usecase"
	"github.com/gofiber/fiber/v2"
)

func main() {
	// Create a new instance of the Fiber app.
	app := fiber.New()

	// Liveness
	app.Get("/liveness", func(c *fiber.Ctx) error {
		return c.SendStatus(http.StatusOK)
	})

	// Puzzle
	puzzleGormRepository := puzzlegorm.NewPuzzleRepository(database.GetDB())
	puzzleEthereumRepository := puzzleonchain.NewPuzzleOnchain(ethereum.GetEth())
	puzzleUseCase := puzzleuc.NewPuzzleUseCase(puzzleGormRepository, puzzleEthereumRepository)
	puzzlehttp.RegisterHTTPEndpoints(app, puzzleUseCase)

	// Start the server and listen on port 3000
	app.Listen(":3000")
}
