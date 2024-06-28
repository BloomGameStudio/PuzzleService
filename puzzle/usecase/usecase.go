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

func (uc PuzzleUseCase) Create(ctx context.Context, title string, solution []byte) error {
	puzzle := &models.Puzzle{
		ID:       crypto.Keccak256Hash(solution),
		Solution: solution,
		Title:    title,
	}

	_, err := uc.puzzleOnchain.Commit(ctx, puzzle.ID)
	if err != nil {
		return err
	}

	return uc.puzzleRepository.Create(ctx, puzzle)
}

func (uc PuzzleUseCase) DeleteById(ctx context.Context, id [32]byte) (bool, error) {
	return uc.puzzleRepository.DeleteById(ctx, id)
}

func (uc PuzzleUseCase) ExistsBySolution(ctx context.Context, solution []byte) (bool, error) {
	exists, err := uc.puzzleRepository.ExistsById(ctx, crypto.Keccak256Hash(solution))

	if err != nil {
		return false, err
	}

	return exists, nil
}

func (uc PuzzleUseCase) GetAll(ctx context.Context) ([]*models.Puzzle, error) {
	puzzles, err := uc.puzzleRepository.GetAll(ctx)
	if err != nil {
		return nil, err
	}

	for i := 0; i < len(puzzles); i++ {
		puzzles[i], err = uc.SyncState(ctx, puzzles[i])
		if err != nil {
			return nil, err
		}
	}

	return puzzles, nil
}

func (uc PuzzleUseCase) GetById(ctx context.Context, id [32]byte) (*models.Puzzle, error) {
	puzzle, err := uc.puzzleRepository.GetById(ctx, id)
	if err != nil {
		return nil, err
	}

	return uc.SyncState(ctx, puzzle)
}

func (uc PuzzleUseCase) Update(ctx context.Context, id [32]byte, title string, revealed bool) (bool, error) {
	puzzle, err := uc.puzzleRepository.GetById(ctx, id)
	if err != nil {
		return false, err
	}

	if revealed && !puzzle.Revealed {
		_, err := uc.puzzleOnchain.Reveal(ctx, puzzle.Solution)
		if err != nil {
			return false, nil
		}
	}

	return uc.puzzleRepository.Update(ctx, &models.Puzzle{
		ID:       id,
		Solution: nil,
		Title:    title,
		Revealed: revealed || puzzle.Revealed,
	})
}

func (uc PuzzleUseCase) SyncState(ctx context.Context, puzzle *models.Puzzle) (*models.Puzzle, error) {
	// Short circuit if monotonic onchain variables are in their final (permanent) state
	if puzzle.Committed && puzzle.Revealed {
		return puzzle, nil
	}

	ocPuzzle, err := uc.puzzleOnchain.GetById(ctx, puzzle.ID)
	if err != nil {
		return nil, err
	}

	_, err = uc.puzzleRepository.Update(ctx, ocPuzzle)
	if err != nil {
		return nil, err
	}

	return uc.puzzleRepository.GetById(ctx, puzzle.ID)
}
