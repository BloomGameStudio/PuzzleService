package puzzle

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
	"github.com/ethereum/go-ethereum/core/types"
)

type Onchain interface {
	GetById(ctx context.Context, id [32]byte) (*models.Puzzle, error)
	Commit(ctx context.Context, id [32]byte) (*types.Transaction, error)
	Reveal(ctx context.Context, solution []byte) (*types.Transaction, error)
}
