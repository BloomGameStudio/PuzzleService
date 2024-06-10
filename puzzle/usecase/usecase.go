package usecase

import (
	"context"
	"fmt"

	"github.com/BloomGameStudio/PuzzleService/models"
	"github.com/BloomGameStudio/PuzzleService/puzzle"
	"github.com/ethereum/go-ethereum/crypto"
)

type PuzzleUseCase struct {
	puzzleRepository puzzle.Repository
	puzzleOnchain    puzzle.Onchain
}

func NewPuzzleUseCase(gorm puzzle.Repository, ethereum puzzle.Onchain) *PuzzleUseCase {
	return &PuzzleUseCase{
		puzzleRepository: gorm,
		puzzleOnchain:    ethereum,
	}
}

func (uc PuzzleUseCase) CreatePuzzle(ctx context.Context, title string, solution []byte) error {
	puzzle := &models.Puzzle{
		ID:       crypto.Keccak256Hash(solution),
		Solution: solution,
		Title:    title,
	}

	return uc.puzzleRepository.CreatePuzzle(ctx, puzzle)
}

func (uc PuzzleUseCase) GetPuzzles(ctx context.Context) ([]*models.Puzzle, error) {
	return uc.puzzleRepository.GetPuzzles(ctx)
}

func (uc PuzzleUseCase) GetPuzzle(ctx context.Context, id [32]byte) (*models.Puzzle, error) {
	puzzle, err := uc.puzzleOnchain.GetPuzzle(ctx, id)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(puzzle.Committed)
	fmt.Println(puzzle.Revealed)

	return uc.puzzleRepository.GetPuzzle(ctx, id)
}

func (uc PuzzleUseCase) UpdatePuzzle(ctx context.Context, id [32]byte, title string) (bool, error) {
	puzzle := &models.Puzzle{
		ID:       id,
		Solution: nil,
		Title:    title,
	}

	return uc.puzzleRepository.UpdatePuzzle(ctx, puzzle)
}

func (uc PuzzleUseCase) DeletePuzzle(ctx context.Context, id [32]byte) (bool, error) {
	return uc.puzzleRepository.DeletePuzzle(ctx, id)
}
