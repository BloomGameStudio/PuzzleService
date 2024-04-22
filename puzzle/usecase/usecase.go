package usecase

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
	"github.com/BloomGameStudio/PuzzleService/puzzle"
)

type PuzzleUseCase struct {
	puzzleRepository puzzle.Repository
}

func NewPuzzleUseCase(repo puzzle.Repository) *PuzzleUseCase {
	return &PuzzleUseCase{
		puzzleRepository: repo,
	}
}

func (uc PuzzleUseCase) CreatePuzzle(ctx context.Context, title string, solution string, solutionTypes *[]string) error {
	puzzle := &models.Puzzle{
		ID:            title, // TODO Change to hash of the solution
		Solution:      solution,
		SolutionTypes: *solutionTypes,
		Title:         title,
	}

	uc.puzzleRepository.CreatePuzzle(ctx, puzzle)

	return nil
}

func (uc PuzzleUseCase) GetPuzzles(ctx context.Context) ([]*models.Puzzle, error) {
	return uc.puzzleRepository.GetPuzzle(ctx)
}
