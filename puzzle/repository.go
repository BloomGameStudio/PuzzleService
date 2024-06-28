package puzzle

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
)

type Repository interface {
	Create(ctx context.Context, puzzle *models.Puzzle) error
	DeleteById(ctx context.Context, id [32]byte) (bool, error)
	ExistsById(ctx context.Context, id [32]byte) (bool, error)
	GetAll(ctx context.Context) ([]*models.Puzzle, error)
	GetById(ctx context.Context, id [32]byte) (*models.Puzzle, error)
	Update(ctx context.Context, puzzle *models.Puzzle) (bool, error)
}
