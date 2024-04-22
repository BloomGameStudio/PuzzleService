package main

import (
	"github.com/gofiber/fiber/v2"

	"github.com/BloomGameStudio/PuzzleService/database"
	puzzlehttp "github.com/BloomGameStudio/PuzzleService/puzzle/http"
	puzzlegorm "github.com/BloomGameStudio/PuzzleService/puzzle/repository"
	puzzleuc "github.com/BloomGameStudio/PuzzleService/puzzle/usecase"
)

func main() {
	// Initialize Database
	database.Init()

	// Create a new instance of the Fiber app.
	app := fiber.New()

	puzzleRepository := puzzlegorm.NewPuzzleRepository(database.GetDB())
	puzzleUseCase := puzzleuc.NewPuzzleUseCase(puzzleRepository)
	puzzlehttp.RegisterHTTPEndpoints(app, puzzleUseCase)

	// Routes
	app.Get("/ping", func(c *fiber.Ctx) error {
		return c.SendString("Pong!")
	})

	// Routes for the puzzles
	app.Get("/puzzles/:id", func(c *fiber.Ctx) error {
		return c.SendString("Unimplemented!")
	})

	app.Patch("/puzzles", func(c *fiber.Ctx) error {
		return c.SendString("Unimplemented!")
	})

	app.Delete("/puzzles", func(c *fiber.Ctx) error {
		return c.SendString("Unimplemented!")
	})

	// Verification routes
	app.Post("/verify/:puzzleID/:solution", func(c *fiber.Ctx) error {
		return c.SendString("Unimplemented!")
	})

	// Start the server and listen on port 3000
	app.Listen(":3000")
}
