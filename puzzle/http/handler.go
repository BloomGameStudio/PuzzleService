package http

import (
	"github.com/gofiber/fiber/v2"
)

// VerifySolutionHandler handles the verification of puzzle solution.
func VerifySolutionHandler(c *fiber.Ctx) error {
    var request struct {
        PuzzleID string `json:"PuzzleID"`
        Solution string `json:"Solution"`
    }

    // Parse request body
    if err := c.BodyParser(&request); err != nil {
        return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
    }

    // verify the solution
    if request.Solution == "true" {
        return c.JSON(fiber.Map{"Solved": true})
    } else {
        return c.JSON(fiber.Map{"Solved": false})
    }
}
