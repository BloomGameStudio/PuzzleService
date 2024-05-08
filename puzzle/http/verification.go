package http

import (
	"fmt"

	publicmodels "github.com/BloomGameStudio/PuzzleService/publicModels"
	"github.com/gofiber/fiber/v2"
)

// Mock function to simulate fetching a puzzle solution
// For example, return a hardcoded correct solution for ID 1

func getCorrectSolutionForPuzzle(id int) (string, error) {
    if id == 1 {
        return "Correct!", nil
    }
    return "", fmt.Errorf("no solution found for id %d", id)
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
