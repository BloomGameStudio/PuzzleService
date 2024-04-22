package puzzle

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
)

type UseCase interface {
	CreatePuzzle(ctx context.Context, title string, solution []byte) error
	GetPuzzles(ctx context.Context) ([]*models.Puzzle, error)
}
