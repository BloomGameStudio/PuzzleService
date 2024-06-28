// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// BloomSolutionRegistryMetaData contains all meta data concerning the BloomSolutionRegistry contract.
var BloomSolutionRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"function\",\"name\":\"commit\",\"inputs\":[{\"name\":\"_puzzle\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_solution\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_registry\",\"type\":\"address\",\"internalType\":\"contractIBloomPuzzleRegistry\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"commits\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"registry\",\"type\":\"address\",\"internalType\":\"contractIBloomPuzzleRegistry\"},{\"name\":\"solution\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"verifyCommit\",\"inputs\":[{\"name\":\"_puzzle\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"_solver\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"error\",\"name\":\"BloomSolutionRegistry_PuzzleNotCommitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BloomSolutionRegistry_SolutionNotRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BloomSolutionRegistry_SolutionRevealed\",\"inputs\":[]}]",
}

// BloomSolutionRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use BloomSolutionRegistryMetaData.ABI instead.
var BloomSolutionRegistryABI = BloomSolutionRegistryMetaData.ABI

// BloomSolutionRegistry is an auto generated Go binding around an Ethereum contract.
type BloomSolutionRegistry struct {
	BloomSolutionRegistryCaller     // Read-only binding to the contract
	BloomSolutionRegistryTransactor // Write-only binding to the contract
	BloomSolutionRegistryFilterer   // Log filterer for contract events
}

// BloomSolutionRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BloomSolutionRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BloomSolutionRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BloomSolutionRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BloomSolutionRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BloomSolutionRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BloomSolutionRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BloomSolutionRegistrySession struct {
	Contract     *BloomSolutionRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts          // Call options to use throughout this session
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BloomSolutionRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BloomSolutionRegistryCallerSession struct {
	Contract *BloomSolutionRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts                // Call options to use throughout this session
}

// BloomSolutionRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BloomSolutionRegistryTransactorSession struct {
	Contract     *BloomSolutionRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts                // Transaction auth options to use throughout this session
}

// BloomSolutionRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BloomSolutionRegistryRaw struct {
	Contract *BloomSolutionRegistry // Generic contract binding to access the raw methods on
}

// BloomSolutionRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BloomSolutionRegistryCallerRaw struct {
	Contract *BloomSolutionRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BloomSolutionRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BloomSolutionRegistryTransactorRaw struct {
	Contract *BloomSolutionRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBloomSolutionRegistry creates a new instance of BloomSolutionRegistry, bound to a specific deployed contract.
func NewBloomSolutionRegistry(address common.Address, backend bind.ContractBackend) (*BloomSolutionRegistry, error) {
	contract, err := bindBloomSolutionRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BloomSolutionRegistry{BloomSolutionRegistryCaller: BloomSolutionRegistryCaller{contract: contract}, BloomSolutionRegistryTransactor: BloomSolutionRegistryTransactor{contract: contract}, BloomSolutionRegistryFilterer: BloomSolutionRegistryFilterer{contract: contract}}, nil
}

// NewBloomSolutionRegistryCaller creates a new read-only instance of BloomSolutionRegistry, bound to a specific deployed contract.
func NewBloomSolutionRegistryCaller(address common.Address, caller bind.ContractCaller) (*BloomSolutionRegistryCaller, error) {
	contract, err := bindBloomSolutionRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BloomSolutionRegistryCaller{contract: contract}, nil
}

// NewBloomSolutionRegistryTransactor creates a new write-only instance of BloomSolutionRegistry, bound to a specific deployed contract.
func NewBloomSolutionRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BloomSolutionRegistryTransactor, error) {
	contract, err := bindBloomSolutionRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BloomSolutionRegistryTransactor{contract: contract}, nil
}

// NewBloomSolutionRegistryFilterer creates a new log filterer instance of BloomSolutionRegistry, bound to a specific deployed contract.
func NewBloomSolutionRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BloomSolutionRegistryFilterer, error) {
	contract, err := bindBloomSolutionRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BloomSolutionRegistryFilterer{contract: contract}, nil
}

// bindBloomSolutionRegistry binds a generic wrapper to an already deployed contract.
func bindBloomSolutionRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BloomSolutionRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BloomSolutionRegistry *BloomSolutionRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BloomSolutionRegistry.Contract.BloomSolutionRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BloomSolutionRegistry *BloomSolutionRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BloomSolutionRegistry.Contract.BloomSolutionRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BloomSolutionRegistry *BloomSolutionRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BloomSolutionRegistry.Contract.BloomSolutionRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BloomSolutionRegistry *BloomSolutionRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BloomSolutionRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BloomSolutionRegistry *BloomSolutionRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BloomSolutionRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BloomSolutionRegistry *BloomSolutionRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BloomSolutionRegistry.Contract.contract.Transact(opts, method, params...)
}

// Commits is a free data retrieval call binding the contract method 0x75165280.
//
// Solidity: function commits(address , bytes32 ) view returns(address registry, bytes32 solution)
func (_BloomSolutionRegistry *BloomSolutionRegistryCaller) Commits(opts *bind.CallOpts, arg0 common.Address, arg1 [32]byte) (struct {
	Registry common.Address
	Solution [32]byte
}, error) {
	var out []interface{}
	err := _BloomSolutionRegistry.contract.Call(opts, &out, "commits", arg0, arg1)

	outstruct := new(struct {
		Registry common.Address
		Solution [32]byte
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Registry = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Solution = *abi.ConvertType(out[1], new([32]byte)).(*[32]byte)

	return *outstruct, err

}

// Commits is a free data retrieval call binding the contract method 0x75165280.
//
// Solidity: function commits(address , bytes32 ) view returns(address registry, bytes32 solution)
func (_BloomSolutionRegistry *BloomSolutionRegistrySession) Commits(arg0 common.Address, arg1 [32]byte) (struct {
	Registry common.Address
	Solution [32]byte
}, error) {
	return _BloomSolutionRegistry.Contract.Commits(&_BloomSolutionRegistry.CallOpts, arg0, arg1)
}

// Commits is a free data retrieval call binding the contract method 0x75165280.
//
// Solidity: function commits(address , bytes32 ) view returns(address registry, bytes32 solution)
func (_BloomSolutionRegistry *BloomSolutionRegistryCallerSession) Commits(arg0 common.Address, arg1 [32]byte) (struct {
	Registry common.Address
	Solution [32]byte
}, error) {
	return _BloomSolutionRegistry.Contract.Commits(&_BloomSolutionRegistry.CallOpts, arg0, arg1)
}

// VerifyCommit is a free data retrieval call binding the contract method 0x1d634ec0.
//
// Solidity: function verifyCommit(bytes32 _puzzle, address _solver) view returns(bool)
func (_BloomSolutionRegistry *BloomSolutionRegistryCaller) VerifyCommit(opts *bind.CallOpts, _puzzle [32]byte, _solver common.Address) (bool, error) {
	var out []interface{}
	err := _BloomSolutionRegistry.contract.Call(opts, &out, "verifyCommit", _puzzle, _solver)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// VerifyCommit is a free data retrieval call binding the contract method 0x1d634ec0.
//
// Solidity: function verifyCommit(bytes32 _puzzle, address _solver) view returns(bool)
func (_BloomSolutionRegistry *BloomSolutionRegistrySession) VerifyCommit(_puzzle [32]byte, _solver common.Address) (bool, error) {
	return _BloomSolutionRegistry.Contract.VerifyCommit(&_BloomSolutionRegistry.CallOpts, _puzzle, _solver)
}

// VerifyCommit is a free data retrieval call binding the contract method 0x1d634ec0.
//
// Solidity: function verifyCommit(bytes32 _puzzle, address _solver) view returns(bool)
func (_BloomSolutionRegistry *BloomSolutionRegistryCallerSession) VerifyCommit(_puzzle [32]byte, _solver common.Address) (bool, error) {
	return _BloomSolutionRegistry.Contract.VerifyCommit(&_BloomSolutionRegistry.CallOpts, _puzzle, _solver)
}

// Commit is a paid mutator transaction binding the contract method 0xd9cd8d1c.
//
// Solidity: function commit(bytes32 _puzzle, bytes32 _solution, address _registry) returns()
func (_BloomSolutionRegistry *BloomSolutionRegistryTransactor) Commit(opts *bind.TransactOpts, _puzzle [32]byte, _solution [32]byte, _registry common.Address) (*types.Transaction, error) {
	return _BloomSolutionRegistry.contract.Transact(opts, "commit", _puzzle, _solution, _registry)
}

// Commit is a paid mutator transaction binding the contract method 0xd9cd8d1c.
//
// Solidity: function commit(bytes32 _puzzle, bytes32 _solution, address _registry) returns()
func (_BloomSolutionRegistry *BloomSolutionRegistrySession) Commit(_puzzle [32]byte, _solution [32]byte, _registry common.Address) (*types.Transaction, error) {
	return _BloomSolutionRegistry.Contract.Commit(&_BloomSolutionRegistry.TransactOpts, _puzzle, _solution, _registry)
}

// Commit is a paid mutator transaction binding the contract method 0xd9cd8d1c.
//
// Solidity: function commit(bytes32 _puzzle, bytes32 _solution, address _registry) returns()
func (_BloomSolutionRegistry *BloomSolutionRegistryTransactorSession) Commit(_puzzle [32]byte, _solution [32]byte, _registry common.Address) (*types.Transaction, error) {
	return _BloomSolutionRegistry.Contract.Commit(&_BloomSolutionRegistry.TransactOpts, _puzzle, _solution, _registry)
}
