package gorm

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
	"gorm.io/gorm"
)

type Puzzle struct {
	ID       []byte `gorm:"primaryKey;size:32"`
	Solution []byte
	Title    string
}

type PuzzleRepository struct {
	db *gorm.DB
}

func NewPuzzleRepository(db *gorm.DB) *PuzzleRepository {
	return &PuzzleRepository{
		db: db,
	}
}

func (r PuzzleRepository) CreatePuzzle(ctx context.Context, puzzle *models.Puzzle) error {
	result := r.db.WithContext(ctx).Create(toModel(puzzle))

	return result.Error
}

func (r PuzzleRepository) GetPuzzles(ctx context.Context) ([]*models.Puzzle, error) {
	var puzzles []Puzzle

	result := r.db.WithContext(ctx).Find(&puzzles)

	return toPuzzles(&puzzles), result.Error
}

func (r PuzzleRepository) GetPuzzle(ctx context.Context, id [32]byte) (*models.Puzzle, error) {
	var puzzle Puzzle

	result := r.db.WithContext(ctx).Where("id = ?", id[:]).First(&puzzle)

	return toPuzzle(&puzzle), result.Error
}

func (r PuzzleRepository) UpdatePuzzle(ctx context.Context, puzzle *models.Puzzle) (bool, error) {
	result := r.db.WithContext(ctx).Model(&Puzzle{}).Where("id = ?", puzzle.ID[:]).Updates(Puzzle{Title: puzzle.Title})

	return result.RowsAffected != 0, result.Error
}

func (r PuzzleRepository) DeletePuzzle(ctx context.Context, id [32]byte) (bool, error) {
	result := r.db.WithContext(ctx).Where("id = ?", id[:]).Delete(&Puzzle{})

	return result.RowsAffected != 0, result.Error
}

func toModel(puzzle *models.Puzzle) *Puzzle {
	return &Puzzle{
		ID:       puzzle.ID[:],
		Solution: puzzle.Solution,
		Title:    puzzle.Title,
	}
}

func toPuzzle(puzzle *Puzzle) *models.Puzzle {
	var id [32]byte
	copy(id[:], puzzle.ID)

	return &models.Puzzle{
		ID:       id,
		Solution: puzzle.Solution,
		Title:    puzzle.Title,
	}
}

func toPuzzles(puzzles *[]Puzzle) []*models.Puzzle {
	out := make([]*models.Puzzle, len(*puzzles))

	for i, puzzle := range *puzzles {
		out[i] = toPuzzle(&puzzle)
	}

	return out
}
