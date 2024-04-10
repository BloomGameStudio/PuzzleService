package main

import (
	"github.com/BloomGameStudio/PuzzleService/handlers"
	"github.com/gofiber/fiber/v2"
)

func main() {

    // Create a new instance of the Fiber app.
    app := fiber.New()

    // Create a new instance of PuzzleHandler
    puzzleHandler := handlers.NewPuzzleHandler()

    // Routes
    app.Get("/ping", func(c *fiber.Ctx) error {
        return c.SendString("Pong!")
    })

    // Routes for the puzzles
    app.Get("/puzzles", puzzleHandler.GetPuzzles)
    app.Get("/puzzles/:id", puzzleHandler.GetPuzzles)
    app.Post("/puzzles", puzzleHandler.CreatePuzzle)
    app.Patch("/puzzles/:id", puzzleHandler.UpdatePuzzle)
    app.Delete("/puzzles/:id", puzzleHandler.DeletePuzzle)

    // Verification routes
    app.Post("/verify/:puzzleID/:solution", func(c *fiber.Ctx) error {
        return c.SendString("Unimplemented!")
    })

    // Start the server and listen on port 3000
    app.Listen(":3000")
}