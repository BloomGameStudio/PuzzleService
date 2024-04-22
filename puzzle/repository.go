package puzzle

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
)

type Repository interface {
	CreatePuzzle(ctx context.Context, puzzle *models.Puzzle) error
	GetPuzzle(ctx context.Context) ([]*models.Puzzle, error)
}
