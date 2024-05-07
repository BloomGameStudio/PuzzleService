package puzzle

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
)

type Repository interface {
	CreatePuzzle(ctx context.Context, puzzle *models.Puzzle) error
	GetPuzzles(ctx context.Context) ([]*models.Puzzle, error)
	GetPuzzle(ctx context.Context, id [32]byte) (*models.Puzzle, error)
	UpdatePuzzle(ctx context.Context, puzzle *models.Puzzle) (bool, error)
	DeletePuzzle(ctx context.Context, id [32]byte) (bool, error)
}
