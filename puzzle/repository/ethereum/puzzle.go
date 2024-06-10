package ethereum

import (
	"context"
	"fmt"

	"github.com/BloomGameStudio/PuzzleService/contract"
	"github.com/BloomGameStudio/PuzzleService/ethereum"
	"github.com/BloomGameStudio/PuzzleService/models"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

type PuzzleOnchain struct {
	eth              *ethclient.Client
	puzzleRegistry   *contract.BloomPuzzleRegistry
	solutionRegistry *contract.BloomSolutionRegistry
}

func NewPuzzleOnchain(eth *ethclient.Client) *PuzzleOnchain {
	return &PuzzleOnchain{
		eth:              eth,
		puzzleRegistry:   ethereum.CreateBloomPuzzleRegistry(eth),
		solutionRegistry: ethereum.CreateBloomSolutionRegistry(eth),
	}
}

type Puzzle struct {
	ID        [32]byte
	Committed bool
	Revealed  bool
}

func (o PuzzleOnchain) GetPuzzle(ctx context.Context, id [32]byte) (*models.Puzzle, error) {
	fmt.Println(id)
	fmt.Println(o.eth.ChainID(context.TODO()))
	owner, _ := o.puzzleRegistry.Owner(&bind.CallOpts{})
	fmt.Println(owner)
	committed, err := o.puzzleRegistry.Committed(&bind.CallOpts{}, id)
	if err != nil {
		fmt.Println(err.Error())
		panic("could not retrieve \"Committed\" data")
	}
	fmt.Println(committed)

	revealed, err := o.puzzleRegistry.SolutionRevealed(&bind.CallOpts{}, id)
	if err != nil {
		fmt.Println(err.Error())
		panic("could not retrieve \"Solution\" data")
	}
	fmt.Println(revealed)

	return toPuzzle(&Puzzle{
		ID:        id,
		Committed: committed,
		Revealed:  revealed,
	}), nil
}

func toPuzzle(puzzle *Puzzle) *models.Puzzle {
	return &models.Puzzle{
		ID:        puzzle.ID,
		Solution:  nil,
		Title:     "",
		Committed: puzzle.Committed,
		Revealed:  puzzle.Revealed,
	}
}
