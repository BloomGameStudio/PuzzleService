package puzzles

import (
	"net/http"

	"github.com/BloomGameStudio/PuzzleService/models"
	"github.com/labstack/echo/v4"
)

func SolutionCheck(puzzleSolution models.PuzzleSolution, correctPuzzleID int, correctSolution string) bool {
	return puzzleSolution.Solution == correctSolution && puzzleSolution.PuzzleID == correctPuzzleID
}

func SolveHandler(c echo.Context) error {
	var puzzleSolution models.PuzzleSolution
	if err := c.Bind(&puzzleSolution); err != nil {
		return c.String(http.StatusBadRequest, "Invalid request body")
	}

	// These values should not be hardcoded
	correctPuzzleID := 1
	correctSolution := "correct_solution"

	// Check if the puzzle has already been solved
	if puzzleSolution.Solved {
		return c.String(http.StatusBadRequest, "Puzzle has already been solved")
	}

	// Check if the solution is valid
	isSolutionValid := SolutionCheck(puzzleSolution, correctPuzzleID, correctSolution)

	if isSolutionValid {
		// Mark the puzzle as solved
		puzzleSolution.Solved = true

		// Update the state of Solved somewhere for the player

		// Return ERC-20 token
		erc20 := models.ERC20{
			TokenName:   "Test ERC-20 Token",
			TokenSymbol: "ERC20",
			Decimals:    18,
			TotalSupply: 100000,
		}
		return c.JSON(http.StatusOK, erc20)
	}

	// If invalid return a 400 error
	return c.String(http.StatusBadRequest, "Invalid solution")
}
