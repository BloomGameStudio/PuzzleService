package puzzle

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
)

type UseCase interface {
	CreatePuzzle(ctx context.Context, title string, solution []byte) error
	GetPuzzles(ctx context.Context) ([]*models.Puzzle, error)
	GetPuzzle(ctx context.Context, id [32]byte) (*models.Puzzle, error)
	UpdatePuzzle(ctx context.Context, id [32]byte, title string) (bool, error)
	DeletePuzzle(ctx context.Context, id [32]byte) (bool, error)
}
