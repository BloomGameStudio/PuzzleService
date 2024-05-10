package http

import (
	"fmt"

	"github.com/BloomGameStudio/PuzzleService/database"
	publicmodels "github.com/BloomGameStudio/PuzzleService/publicModels"
	"github.com/gofiber/fiber/v2"
)

// getCorrectSolutionForPuzzle retrieves the correct solution for the puzzle with the given ID.
func getCorrectSolutionForPuzzle(id int) (string, error) {
    var puzzle publicmodels.Puzzle

    // Use the databse to find the puzzle with the given ID
    result := database.GetDB().First(&puzzle, id)

    // If the puzzle was not found, return an error
    if result.Error != nil {
        return "", fmt.Errorf("no solution found for id %d", id)
    }

    // Return the solution of the puzzle
    return puzzle.Data, nil
}

// VerifySolutionHandler handles the verification of puzzle solution.
func VerifySolutionHandler(c *fiber.Ctx) error {
	// There is differences between publicModels/puzzle.go struct
	// and what is in the README.md
	// For now I am working based off the publicModel
	var request publicmodels.Puzzle

	// Parse request body and validate fields
	if err := c.BodyParser(&request); err != nil || request.ID == 0 || request.Data == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	// Retrieve the correct solution
	correctSolution, err := getCorrectSolutionForPuzzle(request.ID)
	if err != nil {
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	// Verify the solution by comparing the request data with the retrieved correct solution
	if request.Data != correctSolution {
		return c.Status(fiber.StatusBadRequest).SendString("Incorrect solution")
	}

	return c.Status(fiber.StatusOK).SendString("Correct!")
}
