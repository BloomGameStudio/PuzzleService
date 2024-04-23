package http

import (
	"github.com/BloomGameStudio/PuzzleService/puzzle"
	"github.com/gofiber/fiber/v2"
)

func RegisterHTTPEndpoints(app *fiber.App, uc puzzle.UseCase) {
	h := NewHandler(uc)

	puzzles := app.Group("/puzzles")
	{
		// Create
		puzzles.Post("", h.Create)

		// Read
		puzzles.Get("", h.GetAll)
		puzzles.Get("/:id", h.GetById)

		// Update
		puzzles.Patch("/:id", h.PatchById)

		// Delete
		puzzles.Delete("/:id", h.DeleteById)
	}
}
