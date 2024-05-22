package http

import (
	"fmt"
	"log"

	"github.com/BloomGameStudio/PuzzleService/database"
	publicmodels "github.com/BloomGameStudio/PuzzleService/publicModels"
	"github.com/gofiber/fiber/v2"
	"gorm.io/gorm"
)

// getCorrectSolutionForPuzzle retrieves the correct solution for the puzzle with the given ID.
func getCorrectSolutionForPuzzle(id int) (string, error) {
    var puzzle publicmodels.Puzzle

    // Use the databse to find the puzzle with the given ID
    result := database.GetDB().First(&puzzle, id)

    // If the puzzle was not found, return an error
    if result.Error != nil {
        if result.Error == gorm.ErrRecordNotFound {
			return "", fmt.Errorf("no puzzle found for id %d", id)
		}
    	return "", fmt.Errorf("database error: %v", result.Error)
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
	if err := c.BodyParser(&request); err != nil {
		log.Printf("Error parsing request body: %v", err)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	if request.ID == 0 || request.Data == "" {
		log.Printf("Invalid request data: %v", request)
		return c.Status(fiber.StatusBadRequest).SendString("Invalid request")
	}

	// Retrieve the correct solution
	correctSolution, err := getCorrectSolutionForPuzzle(request.ID)
	if err != nil {
		log.Printf("Error retrieving correct solution: %v", err)
		return c.Status(fiber.StatusNotFound).SendString(err.Error())
	}

	// Verify the solution by comparing the request data with the retrieved correct solution
	if request.Data != correctSolution {
		log.Printf("Incorrect solution for ID %d", request.ID)
		return c.Status(fiber.StatusBadRequest).SendString("Incorrect solution")
	}

	// Return if successful
	log.Printf("Correct solution for ID %d", request.ID)
	return c.Status(fiber.StatusOK).SendString("Correct!")
}
