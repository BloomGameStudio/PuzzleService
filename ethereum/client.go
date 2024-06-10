package ethereum

import (
	"os"

	"github.com/BloomGameStudio/PuzzleService/contract"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Eth *ethclient.Client

func Connect() *ethclient.Client {
	providerUrl := os.Getenv("ETHEREUM_PROVIDER_URL")

	if providerUrl == "" {
		panic("missing Ethereum provider")
	}

	client, err := ethclient.Dial(providerUrl)
	if err != nil {
		panic("failed to connect Ethereum client")
	}

	return client
}

func Init() {
	Eth = Connect()
}

func GetEth() *ethclient.Client {
	return Eth
}

func CreateBloomPuzzleRegistry(eth *ethclient.Client) *contract.BloomPuzzleRegistry {
	instance, err := contract.NewBloomPuzzleRegistry(
		common.HexToAddress(os.Getenv("BLOOM_PUZZLE_REGISTRY_ADDRESS")),
		eth,
	)
	if err != nil {
		panic("failed to initialize \"BloomPuzzleRegistry\"")
	}

	return instance
}

func CreateBloomSolutionRegistry(eth *ethclient.Client) *contract.BloomSolutionRegistry {
	instance, err := contract.NewBloomSolutionRegistry(
		common.HexToAddress(os.Getenv("BLOOM_SOLUTION_REGISTRY_ADDRESS")),
		eth,
	)
	if err != nil {
		panic("failed to initialize \"BloomSolutionRegistry\"")
	}

	return instance
}
