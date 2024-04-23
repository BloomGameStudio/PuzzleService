package puzzle

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
)

type UseCase interface {
	CreatePuzzle(ctx context.Context, title string, solution []byte) error
	DeletePuzzle(ctx context.Context, id [32]byte) (bool, error)
	GetPuzzles(ctx context.Context) ([]*models.Puzzle, error)
}
