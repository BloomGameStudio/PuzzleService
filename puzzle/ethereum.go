package puzzle

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
)

type Onchain interface {
	GetPuzzle(ctx context.Context, id [32]byte) (*models.Puzzle, error)
}
