package ethereum

import (
	"context"
	"os"

	"github.com/BloomGameStudio/PuzzleService/contract"
	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

var Eth *ethclient.Client
var Keystore *keystore.KeyStore

func Init() {
	Eth = Connect()
	Keystore = keystore.NewKeyStore(
		os.Getenv("KEYSTORE_DIR"),
		keystore.StandardScryptN,
		keystore.StandardScryptP,
	)

	Keystore.Unlock(GetAccount(), os.Getenv("KEYSTORE_PASSWORD"))
}

func GetEth() *ethclient.Client {
	return Eth
}

func GetAccount() accounts.Account {
	return Keystore.Accounts()[0]
}

func GetAuth(ctx context.Context) (*bind.TransactOpts, error) {
	if Eth == nil || Keystore == nil {
		panic("Ethereum client not initialized")
	}

	chainId, err := Eth.ChainID(ctx)
	if err != nil {
		panic(err)
	}

	return bind.NewKeyStoreTransactorWithChainID(Keystore, GetAccount(), chainId)
}

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
