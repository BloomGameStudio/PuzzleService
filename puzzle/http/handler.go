package http

import (
	"net/http"

	"github.com/BloomGameStudio/PuzzleService/models"
	"github.com/BloomGameStudio/PuzzleService/puzzle"
	"github.com/gofiber/fiber/v2"
)

type Puzzle struct {
	ID            string   `json:"id"`
	Solution      string   `json:"solution"`
	SolutionTypes []string `json:"solutionTypes"`
	Title         string   `json:"title"`
}

type Handler struct {
	UseCase puzzle.UseCase
}

func NewHandler(uc puzzle.UseCase) *Handler {
	return &Handler{
		UseCase: uc,
	}
}

type CreateBody struct {
	Solution      string   `json:"solution"`
	SolutionTypes []string `json:"solutionTypes"`
	Title         string   `json:"title"`
}

func (h *Handler) Create(c *fiber.Ctx) error {
	body := new(CreateBody)
	if err := c.BodyParser(body); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}

	if err := h.UseCase.CreatePuzzle(c.Context(), body.Title, body.Solution, &body.SolutionTypes); err != nil {
		c.Status(http.StatusInternalServerError)
		return nil
	}

	c.Status(http.StatusCreated)
	return nil
}

type GetResponse struct {
	Puzzles []*Puzzle `json:"puzzles"`
}

func (h *Handler) Get(c *fiber.Ctx) error {
	puzzles, err := h.UseCase.GetPuzzles(c.Context())

	if err != nil {
		c.Status(http.StatusInternalServerError)
		return nil
	}

	c.Status(http.StatusOK).JSON(&GetResponse{
		Puzzles: toPuzzles(puzzles),
	})

	return nil
}

func toPuzzle(puzzle *models.Puzzle) *Puzzle {
	return &Puzzle{
		ID:            puzzle.ID,
		Solution:      puzzle.Solution,
		SolutionTypes: puzzle.SolutionTypes,
		Title:         puzzle.Title,
	}
}

func toPuzzles(puzzles []*models.Puzzle) []*Puzzle {
	out := make([]*Puzzle, len(puzzles))

	for i, puzzle := range puzzles {
		out[i] = toPuzzle(puzzle)
	}

	return out
}
