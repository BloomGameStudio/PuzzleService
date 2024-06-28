package puzzle

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
)

type UseCase interface {
	Create(ctx context.Context, title string, solution []byte) error
	DeleteById(ctx context.Context, id [32]byte) (bool, error)
	ExistsBySolution(ctx context.Context, solution []byte) (bool, error)
	GetAll(ctx context.Context) ([]*models.Puzzle, error)
	GetById(ctx context.Context, id [32]byte) (*models.Puzzle, error)
	Update(ctx context.Context, id [32]byte, title string, revealed bool) (bool, error)
}
