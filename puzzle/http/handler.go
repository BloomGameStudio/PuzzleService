package http

import (
	"encoding/hex"
	"fmt"
	"net/http"

	"github.com/BloomGameStudio/PuzzleService/models"
	"github.com/BloomGameStudio/PuzzleService/puzzle"
	"github.com/gofiber/fiber/v2"
)

type Puzzle struct {
	ID       string `json:"id"`
	Solution string `json:"solution"`
	Title    string `json:"title"`
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
	Solution string `json:"solution"`
	Title    string `json:"title"`
}

func (h *Handler) Create(c *fiber.Ctx) error {
	body := new(CreateBody)

	if err := c.BodyParser(body); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}

	if body.Solution[:2] != "0x" {
		c.Status(http.StatusBadRequest)
		return nil
	}

	solutionByteSlice, err := hex.DecodeString(body.Solution[2:])
	if err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}

	if err := h.UseCase.CreatePuzzle(c.Context(), body.Title, solutionByteSlice); err != nil {
		c.Status(http.StatusInternalServerError)
		return nil
	}

	c.Status(http.StatusCreated)
	return nil
}

type GetResponse struct {
	Puzzles []*Puzzle `json:"puzzles"`
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
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

func (h *Handler) GetById(c *fiber.Ctx) error {
	c.Status(http.StatusNotImplemented)
	return nil
}

func (h *Handler) PatchById(c *fiber.Ctx) error {
	c.Status(http.StatusNotImplemented)
	return nil
}

type DeleteParams struct {
	ID string `json:"id"`
}

func (h *Handler) DeleteById(c *fiber.Ctx) error {
	params := new(DeleteParams)

	if err := c.ParamsParser(params); err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}

	if params.ID[:2] != "0x" {
		c.Status(http.StatusBadRequest)
		return nil
	}

	idByteSlice, err := hex.DecodeString(params.ID[2:])
	if err != nil {
		c.Status(http.StatusBadRequest)
		return nil
	}

	if len(idByteSlice) != 32 {
		c.Status(http.StatusBadRequest)
		return nil
	}

	var id [32]byte
	copy(id[:], idByteSlice)

	rowsAffected, err := h.UseCase.DeletePuzzle(c.Context(), id)
	if err != nil {
		c.Status(http.StatusInternalServerError)
		return nil
	}

	if rowsAffected {
		c.Status(http.StatusNoContent)
	} else {
		c.Status(http.StatusNotFound)
	}

	return nil
}

func toPuzzle(puzzle *models.Puzzle) *Puzzle {
	return &Puzzle{
		ID:       fmt.Sprintf("0x%x", puzzle.ID),
		Solution: fmt.Sprintf("0x%x", puzzle.Solution),
		Title:    puzzle.Title,
	}
}

func toPuzzles(puzzles []*models.Puzzle) []*Puzzle {
	out := make([]*Puzzle, len(puzzles))

	for i, puzzle := range puzzles {
		out[i] = toPuzzle(puzzle)
	}

	return out
}
