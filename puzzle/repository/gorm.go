package gorm

import (
	"context"

	"github.com/BloomGameStudio/PuzzleService/models"
	"gorm.io/gorm"
)

type Puzzle struct {
	ID       string `gorm:"primaryKey"`
	Solution string
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

	if result.Error != nil {
		panic(result.Error)
	}

	return nil
}

func (r PuzzleRepository) GetPuzzle(ctx context.Context) ([]*models.Puzzle, error) {
	var puzzles []Puzzle

	result := r.db.WithContext(ctx).Find(&puzzles)

	if result.Error != nil {
		panic(result.Error)
	}

	return toPuzzles(&puzzles), nil
}

func toModel(puzzle *models.Puzzle) *Puzzle {
	return &Puzzle{
		ID:       puzzle.ID,
		Solution: puzzle.Solution,
		Title:    puzzle.Title,
	}
}

func toPuzzle(puzzle *Puzzle) *models.Puzzle {
	return &models.Puzzle{
		ID:       puzzle.ID,
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
