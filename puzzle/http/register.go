package http

import (
	"github.com/BloomGameStudio/PuzzleService/puzzle"
	"github.com/gofiber/fiber/v2"
)

func RegisterHTTPEndpoints(app *fiber.App, uc puzzle.UseCase) {
	h := NewHandler(uc)

	puzzles := app.Group("/puzzles")
	{
		puzzles.Post("", h.Create)
		puzzles.Get("", h.Get)
	}
}
