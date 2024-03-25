package handlers

import (
	"github.com/BloomGameStudio/PuzzleService/models"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

type PuzzleHandler struct {
    db *gorm.DB
}

func NewPuzzleHandler(db *gorm.DB) *PuzzleHandler {
    return &PuzzleHandler{db: db}
}

func (h *PuzzleHandler) GetPuzzles(c *fiber.Ctx) error {
    id := c.Query("id")
    if id != "" {
        var puzzle models.Puzzle
        h.db.First(&puzzle, "id = ?", id)
        if puzzle.ID == 0 {
            return c.Status(fiber.StatusNotFound).SendString("No puzzle found with that ID")
        }
        return c.JSON(puzzle)
    } else {
        var puzzles []models.Puzzle
        h.db.Find(&puzzles)
        return c.JSON(puzzles)
    }
}

func (h *PuzzleHandler) CreatePuzzle(c *fiber.Ctx) error {
    puzzle := new(models.Puzzle)

    if err := c.BodyParser(puzzle); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    h.db.Create(puzzle)
    return c.JSON(puzzle)
}

func (h *PuzzleHandler) UpdatePuzzle(c *fiber.Ctx) error {
    id := c.Query("id")
    if id == "" {
        return c.Status(fiber.StatusBadRequest).SendString("No ID provided")
    }

    var puzzle models.Puzzle
    h.db.First(&puzzle, "id = ?", id)
    if puzzle.ID == 0 {
        return c.Status(fiber.StatusNotFound).SendString("No puzzle found with that ID")
    }

    var updatedPuzzle models.Puzzle
    if err := c.BodyParser(&updatedPuzzle); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString(err.Error())
    }

    puzzle.Name = updatedPuzzle.Name
    puzzle.PosX = updatedPuzzle.PosX
    puzzle.PosY = updatedPuzzle.PosY
    puzzle.PosZ = updatedPuzzle.PosZ
    h.db.Save(&puzzle)

    return c.JSON(puzzle)
}

func (h *PuzzleHandler) DeletePuzzle(c *fiber.Ctx) error {
    id := c.Query("id")
    if id == "" {
        return c.Status(fiber.StatusBadRequest).SendString("No ID provided")
    }

    var puzzle models.Puzzle
    h.db.First(&puzzle, "id = ?", id)
    if puzzle.ID == 0 {
        return c.Status(fiber.StatusNotFound).SendString("No puzzle found with that ID")
    }

    h.db.Delete(&puzzle)
    return c.SendString("Puzzle successfully deleted")
}