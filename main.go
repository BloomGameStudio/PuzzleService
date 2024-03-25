package main

import (
	"log"

	"github.com/BloomGameStudio/PuzzleService/database"
	"github.com/BloomGameStudio/PuzzleService/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
    db, cleanup, err := database.InitDB()
    if err != nil {
        log.Fatalf("Failed to initialize database: %v", err)
    }
    defer cleanup()

    app := fiber.New()

    app.Get("/ping", func(c *fiber.Ctx) error {
        return c.SendString("Pong!")
    })

    puzzleHandlers := handlers.NewPuzzleHandler(db)
    app.Get("/puzzles", puzzleHandlers.GetPuzzles)
    app.Post("/puzzles", puzzleHandlers.CreatePuzzle)
    app.Put("/puzzles", puzzleHandlers.UpdatePuzzle)
    app.Delete("/puzzles", puzzleHandlers.DeletePuzzle)

    app.Listen(":3000")
}