package http

import (
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"

	"github.com/BloomGameStudio/PuzzleService/models"
	"github.com/BloomGameStudio/PuzzleService/puzzle"
	"github.com/gofiber/fiber/v2"
)

type Puzzle struct {
	Title     string `json:"title"`
	ID        string `json:"id"`
	Solution  string `json:"solution"`
	Committed bool   `json:"committed"`
	Revealed  bool   `json:"revealed"`
}

type Handler struct {
	UseCase puzzle.UseCase
}

func NewHandler(uc puzzle.UseCase) *Handler {
	return &Handler{
		UseCase: uc,
	}
}

type CreateRequest struct {
	Solution string `json:"solution"`
	Title    string `json:"title"`
}

func (h *Handler) Create(c *fiber.Ctx) error {
	body := new(CreateRequest)

	if err := c.BodyParser(body); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	if len(body.Solution) < 2 || body.Solution[:2] != "0x" {
		return c.SendStatus(http.StatusBadRequest)
	}

	solutionByteSlice, err := hex.DecodeString(body.Solution[2:])
	if err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	if err := h.UseCase.CreatePuzzle(c.Context(), body.Title, solutionByteSlice); err != nil {
		c.SendStatus(http.StatusInternalServerError)
	}

	return c.SendStatus(http.StatusCreated)
}

type GetAllResponse struct {
	Puzzles []*Puzzle `json:"puzzles"`
}

func (h *Handler) GetAll(c *fiber.Ctx) error {
	puzzles, err := h.UseCase.GetPuzzles(c.Context())

	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	return c.Status(http.StatusOK).JSON(&GetAllResponse{
		Puzzles: toPuzzles(puzzles),
	})
}

func (h *Handler) GetById(c *fiber.Ctx) error {
	id, err := toID(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	puzzle, err := h.UseCase.GetPuzzle(c.Context(), id)

	if err != nil {
		return c.SendStatus(http.StatusNotFound)
	}

	return c.Status(http.StatusOK).JSON(toPuzzle(puzzle))
}

type PatchByIdRequest struct {
	Title string `json:"title"`
}

func (h *Handler) PatchById(c *fiber.Ctx) error {
	id, err := toID(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	body := new(PatchByIdRequest)

	if err := c.BodyParser(body); err != nil {
		return c.SendStatus(http.StatusBadRequest)
	}

	rowsAffected, err := h.UseCase.UpdatePuzzle(c.Context(), id, body.Title)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	if rowsAffected {
		return c.SendStatus(http.StatusNoContent)
	} else {
		return c.SendStatus(http.StatusNotFound)
	}
}

func (h *Handler) DeleteById(c *fiber.Ctx) error {
	id, err := toID(c.Params("id"))

	if err != nil {
		return c.Status(http.StatusBadRequest).JSON(fiber.Map{"error": err.Error()})
	}

	rowsAffected, err := h.UseCase.DeletePuzzle(c.Context(), id)
	if err != nil {
		return c.SendStatus(http.StatusInternalServerError)
	}

	if rowsAffected {
		return c.SendStatus(http.StatusNoContent)
	} else {
		return c.SendStatus(http.StatusNotFound)
	}
}

func toID(id string) ([32]byte, error) {
	var out [32]byte

	if len(id) < 2 || id[:2] != "0x" {
		return out, errors.New("'id' must start with '0x'")
	}

	idByteSlice, err := hex.DecodeString(id[2:])
	if err != nil {
		return out, errors.New("'id' is not a valid hex string")
	}

	if len(idByteSlice) != 32 {
		return out, errors.New("'id' must be a 32 byte hex string")
	}

	copy(out[:], idByteSlice)

	return out, nil
}

func toPuzzle(puzzle *models.Puzzle) *Puzzle {
	return &Puzzle{
		Title:     puzzle.Title,
		ID:        fmt.Sprintf("0x%x", puzzle.ID),
		Solution:  fmt.Sprintf("0x%x", puzzle.Solution),
		Committed: puzzle.Committed,
		Revealed:  puzzle.Revealed,
	}
}

func toPuzzles(puzzles []*models.Puzzle) []*Puzzle {
	out := make([]*Puzzle, len(puzzles))

	for i, puzzle := range puzzles {
		out[i] = toPuzzle(puzzle)
	}

	return out
}
