package handlers

import (
	"strconv"

	"github.com/gofiber/fiber/v2"
)

type PuzzleHandler struct {
    puzzles map[string]map[string]interface{}
}

func NewPuzzleHandler() *PuzzleHandler {
    return &PuzzleHandler{puzzles: make(map[string]map[string]interface{})}
}

func (h *PuzzleHandler) GetPuzzles(c *fiber.Ctx) error {
    id := c.Query("id")
    if id != "" {
        puzzle, ok := h.puzzles[id]
        if !ok {
            return c.Status(fiber.StatusNotFound).SendString("No puzzle found with that ID")
        }
        return c.JSON(puzzle)
    } else {
        return c.JSON(h.puzzles)
    }
}

func (h *PuzzleHandler) CreatePuzzle(c *fiber.Ctx) error {
    puzzle := make(map[string]interface{})

    if err := c.BodyParser(&puzzle); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    id := strconv.Itoa(len(h.puzzles) + 1)
    h.puzzles[id] = puzzle
    return c.JSON(puzzle)
}

func (h *PuzzleHandler) UpdatePuzzle(c *fiber.Ctx) error {
    id := c.Query("id")
    if id == "" {
        return c.Status(fiber.StatusBadRequest).SendString("No ID provided")
    }

    puzzle, ok := h.puzzles[id]
    if !ok {
        return c.Status(fiber.StatusNotFound).SendString("No puzzle found with that ID")
    }

    var updatedPuzzle map[string]interface{}
    if err := c.BodyParser(&updatedPuzzle); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    for key, value := range updatedPuzzle {
        puzzle[key] = value
    }
    h.puzzles[id] = puzzle

    return c.JSON(puzzle)
}

func (h *PuzzleHandler) DeletePuzzle(c *fiber.Ctx) error {
    id := c.Query("id")
    if id == "" {
        return c.Status(fiber.StatusBadRequest).SendString("No ID provided")
    }

    _, ok := h.puzzles[id]
    if !ok {
        return c.Status(fiber.StatusNotFound).SendString("No puzzle found with that ID")
    }

    delete(h.puzzles, id)
    return c.SendString("Puzzle successfully deleted")
}