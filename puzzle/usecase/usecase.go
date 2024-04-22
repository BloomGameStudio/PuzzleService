package usecase

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
	"github.com/BloomGameStudio/PuzzleService/puzzle"
	"github.com/ethereum/go-ethereum/crypto"
)

type PuzzleUseCase struct {
	puzzleRepository puzzle.Repository
}

func NewPuzzleUseCase(repo puzzle.Repository) *PuzzleUseCase {
	return &PuzzleUseCase{
		puzzleRepository: repo,
	}
}

func (uc PuzzleUseCase) CreatePuzzle(ctx context.Context, title string, solution []byte) error {
	puzzle := &models.Puzzle{
		ID:       crypto.Keccak256Hash(solution),
		Solution: solution,
		Title:    title,
	}

	uc.puzzleRepository.CreatePuzzle(ctx, puzzle)

	return nil
}

func (uc PuzzleUseCase) GetPuzzles(ctx context.Context) ([]*models.Puzzle, error) {
	return uc.puzzleRepository.GetPuzzle(ctx)
}
