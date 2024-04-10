package main

import (
	"github.com/BloomGameStudio/PuzzleService/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {
    app := fiber.New()

    app.Get("/ping", func(c *fiber.Ctx) error {
        return c.SendString("Pong!")
    })

    puzzleHandlers := handlers.NewPuzzleHandler()
    app.Get("/puzzles", puzzleHandlers.GetPuzzles)
    app.Post("/puzzles", puzzleHandlers.CreatePuzzle)
    app.Put("/puzzles", puzzleHandlers.UpdatePuzzle)
    app.Delete("/puzzles", puzzleHandlers.DeletePuzzle)

    app.Listen(":3000")
}