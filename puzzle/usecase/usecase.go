package usecase

import (
	"context"

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
	puzzles, err := uc.puzzleRepository.GetPuzzles(ctx)
	if err != nil {
		panic(err.Error())
	}

	for i := 0; i < len(puzzles); i++ {
		puzzles[i], err = uc.SyncPuzzleState(ctx, puzzles[i])
		if err != nil {
			panic(err.Error())
		}
	}

	return puzzles, nil
}

func (uc PuzzleUseCase) GetPuzzle(ctx context.Context, id [32]byte) (*models.Puzzle, error) {
	puzzle, err := uc.puzzleRepository.GetPuzzle(ctx, id)
	if err != nil {
		panic(err.Error())
	}

	return uc.SyncPuzzleState(ctx, puzzle)
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

func (uc PuzzleUseCase) SyncPuzzleState(ctx context.Context, puzzle *models.Puzzle) (*models.Puzzle, error) {
	// Short circuit if monotonic onchain variables are in their final (permanent) state
	if puzzle.Committed && puzzle.Revealed {
		return puzzle, nil
	}

	ocPuzzle, err := uc.puzzleOnchain.GetPuzzle(ctx, puzzle.ID)
	if err != nil {
		panic(err.Error())
	}

	_, err = uc.puzzleRepository.UpdatePuzzle(ctx, ocPuzzle)
	if err != nil {
		panic(err.Error())
	}

	return uc.puzzleRepository.GetPuzzle(ctx, puzzle.ID)
}
