package puzzle

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
	"github.com/ethereum/go-ethereum/core/types"
)

type Onchain interface {
	Commit(ctx context.Context, id [32]byte) (*types.Transaction, error)
	GetById(ctx context.Context, id [32]byte) (*models.Puzzle, error)
}
