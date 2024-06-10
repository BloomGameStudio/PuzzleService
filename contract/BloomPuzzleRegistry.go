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

// BloomPuzzleRegistryMetaData contains all meta data concerning the BloomPuzzleRegistry contract.
var BloomPuzzleRegistryMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"commit\",\"inputs\":[{\"name\":\"_puzzle\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"committed\",\"inputs\":[{\"name\":\"_puzzle\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setSolution\",\"inputs\":[{\"name\":\"_solution\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"solution\",\"inputs\":[{\"name\":\"_puzzle\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"solutionRevealed\",\"inputs\":[{\"name\":\"_puzzle\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"BloomPuzzleRegistry_PuzzleAlreadyCommitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BloomPuzzleRegistry_PuzzleNotCommitted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"BloomPuzzleRegistry_SolutionNotRevealed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// BloomPuzzleRegistryABI is the input ABI used to generate the binding from.
// Deprecated: Use BloomPuzzleRegistryMetaData.ABI instead.
var BloomPuzzleRegistryABI = BloomPuzzleRegistryMetaData.ABI

// BloomPuzzleRegistry is an auto generated Go binding around an Ethereum contract.
type BloomPuzzleRegistry struct {
	BloomPuzzleRegistryCaller     // Read-only binding to the contract
	BloomPuzzleRegistryTransactor // Write-only binding to the contract
	BloomPuzzleRegistryFilterer   // Log filterer for contract events
}

// BloomPuzzleRegistryCaller is an auto generated read-only Go binding around an Ethereum contract.
type BloomPuzzleRegistryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BloomPuzzleRegistryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BloomPuzzleRegistryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BloomPuzzleRegistryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BloomPuzzleRegistryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BloomPuzzleRegistrySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BloomPuzzleRegistrySession struct {
	Contract     *BloomPuzzleRegistry // Generic contract binding to set the session for
	CallOpts     bind.CallOpts        // Call options to use throughout this session
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// BloomPuzzleRegistryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BloomPuzzleRegistryCallerSession struct {
	Contract *BloomPuzzleRegistryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts              // Call options to use throughout this session
}

// BloomPuzzleRegistryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BloomPuzzleRegistryTransactorSession struct {
	Contract     *BloomPuzzleRegistryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts              // Transaction auth options to use throughout this session
}

// BloomPuzzleRegistryRaw is an auto generated low-level Go binding around an Ethereum contract.
type BloomPuzzleRegistryRaw struct {
	Contract *BloomPuzzleRegistry // Generic contract binding to access the raw methods on
}

// BloomPuzzleRegistryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BloomPuzzleRegistryCallerRaw struct {
	Contract *BloomPuzzleRegistryCaller // Generic read-only contract binding to access the raw methods on
}

// BloomPuzzleRegistryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BloomPuzzleRegistryTransactorRaw struct {
	Contract *BloomPuzzleRegistryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBloomPuzzleRegistry creates a new instance of BloomPuzzleRegistry, bound to a specific deployed contract.
func NewBloomPuzzleRegistry(address common.Address, backend bind.ContractBackend) (*BloomPuzzleRegistry, error) {
	contract, err := bindBloomPuzzleRegistry(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BloomPuzzleRegistry{BloomPuzzleRegistryCaller: BloomPuzzleRegistryCaller{contract: contract}, BloomPuzzleRegistryTransactor: BloomPuzzleRegistryTransactor{contract: contract}, BloomPuzzleRegistryFilterer: BloomPuzzleRegistryFilterer{contract: contract}}, nil
}

// NewBloomPuzzleRegistryCaller creates a new read-only instance of BloomPuzzleRegistry, bound to a specific deployed contract.
func NewBloomPuzzleRegistryCaller(address common.Address, caller bind.ContractCaller) (*BloomPuzzleRegistryCaller, error) {
	contract, err := bindBloomPuzzleRegistry(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BloomPuzzleRegistryCaller{contract: contract}, nil
}

// NewBloomPuzzleRegistryTransactor creates a new write-only instance of BloomPuzzleRegistry, bound to a specific deployed contract.
func NewBloomPuzzleRegistryTransactor(address common.Address, transactor bind.ContractTransactor) (*BloomPuzzleRegistryTransactor, error) {
	contract, err := bindBloomPuzzleRegistry(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BloomPuzzleRegistryTransactor{contract: contract}, nil
}

// NewBloomPuzzleRegistryFilterer creates a new log filterer instance of BloomPuzzleRegistry, bound to a specific deployed contract.
func NewBloomPuzzleRegistryFilterer(address common.Address, filterer bind.ContractFilterer) (*BloomPuzzleRegistryFilterer, error) {
	contract, err := bindBloomPuzzleRegistry(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BloomPuzzleRegistryFilterer{contract: contract}, nil
}

// bindBloomPuzzleRegistry binds a generic wrapper to an already deployed contract.
func bindBloomPuzzleRegistry(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := BloomPuzzleRegistryMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BloomPuzzleRegistry *BloomPuzzleRegistryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BloomPuzzleRegistry.Contract.BloomPuzzleRegistryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BloomPuzzleRegistry *BloomPuzzleRegistryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.Contract.BloomPuzzleRegistryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BloomPuzzleRegistry *BloomPuzzleRegistryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.Contract.BloomPuzzleRegistryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BloomPuzzleRegistry *BloomPuzzleRegistryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BloomPuzzleRegistry.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BloomPuzzleRegistry *BloomPuzzleRegistryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BloomPuzzleRegistry *BloomPuzzleRegistryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.Contract.contract.Transact(opts, method, params...)
}

// Committed is a free data retrieval call binding the contract method 0xb5ee95d3.
//
// Solidity: function committed(bytes32 _puzzle) view returns(bool)
func (_BloomPuzzleRegistry *BloomPuzzleRegistryCaller) Committed(opts *bind.CallOpts, _puzzle [32]byte) (bool, error) {
	var out []interface{}
	err := _BloomPuzzleRegistry.contract.Call(opts, &out, "committed", _puzzle)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Committed is a free data retrieval call binding the contract method 0xb5ee95d3.
//
// Solidity: function committed(bytes32 _puzzle) view returns(bool)
func (_BloomPuzzleRegistry *BloomPuzzleRegistrySession) Committed(_puzzle [32]byte) (bool, error) {
	return _BloomPuzzleRegistry.Contract.Committed(&_BloomPuzzleRegistry.CallOpts, _puzzle)
}

// Committed is a free data retrieval call binding the contract method 0xb5ee95d3.
//
// Solidity: function committed(bytes32 _puzzle) view returns(bool)
func (_BloomPuzzleRegistry *BloomPuzzleRegistryCallerSession) Committed(_puzzle [32]byte) (bool, error) {
	return _BloomPuzzleRegistry.Contract.Committed(&_BloomPuzzleRegistry.CallOpts, _puzzle)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BloomPuzzleRegistry *BloomPuzzleRegistryCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BloomPuzzleRegistry.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BloomPuzzleRegistry *BloomPuzzleRegistrySession) Owner() (common.Address, error) {
	return _BloomPuzzleRegistry.Contract.Owner(&_BloomPuzzleRegistry.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_BloomPuzzleRegistry *BloomPuzzleRegistryCallerSession) Owner() (common.Address, error) {
	return _BloomPuzzleRegistry.Contract.Owner(&_BloomPuzzleRegistry.CallOpts)
}

// Solution is a free data retrieval call binding the contract method 0xb7f2e2cf.
//
// Solidity: function solution(bytes32 _puzzle) view returns(bytes)
func (_BloomPuzzleRegistry *BloomPuzzleRegistryCaller) Solution(opts *bind.CallOpts, _puzzle [32]byte) ([]byte, error) {
	var out []interface{}
	err := _BloomPuzzleRegistry.contract.Call(opts, &out, "solution", _puzzle)

	if err != nil {
		return *new([]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([]byte)).(*[]byte)

	return out0, err

}

// Solution is a free data retrieval call binding the contract method 0xb7f2e2cf.
//
// Solidity: function solution(bytes32 _puzzle) view returns(bytes)
func (_BloomPuzzleRegistry *BloomPuzzleRegistrySession) Solution(_puzzle [32]byte) ([]byte, error) {
	return _BloomPuzzleRegistry.Contract.Solution(&_BloomPuzzleRegistry.CallOpts, _puzzle)
}

// Solution is a free data retrieval call binding the contract method 0xb7f2e2cf.
//
// Solidity: function solution(bytes32 _puzzle) view returns(bytes)
func (_BloomPuzzleRegistry *BloomPuzzleRegistryCallerSession) Solution(_puzzle [32]byte) ([]byte, error) {
	return _BloomPuzzleRegistry.Contract.Solution(&_BloomPuzzleRegistry.CallOpts, _puzzle)
}

// SolutionRevealed is a free data retrieval call binding the contract method 0x1287f8f7.
//
// Solidity: function solutionRevealed(bytes32 _puzzle) view returns(bool)
func (_BloomPuzzleRegistry *BloomPuzzleRegistryCaller) SolutionRevealed(opts *bind.CallOpts, _puzzle [32]byte) (bool, error) {
	var out []interface{}
	err := _BloomPuzzleRegistry.contract.Call(opts, &out, "solutionRevealed", _puzzle)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SolutionRevealed is a free data retrieval call binding the contract method 0x1287f8f7.
//
// Solidity: function solutionRevealed(bytes32 _puzzle) view returns(bool)
func (_BloomPuzzleRegistry *BloomPuzzleRegistrySession) SolutionRevealed(_puzzle [32]byte) (bool, error) {
	return _BloomPuzzleRegistry.Contract.SolutionRevealed(&_BloomPuzzleRegistry.CallOpts, _puzzle)
}

// SolutionRevealed is a free data retrieval call binding the contract method 0x1287f8f7.
//
// Solidity: function solutionRevealed(bytes32 _puzzle) view returns(bool)
func (_BloomPuzzleRegistry *BloomPuzzleRegistryCallerSession) SolutionRevealed(_puzzle [32]byte) (bool, error) {
	return _BloomPuzzleRegistry.Contract.SolutionRevealed(&_BloomPuzzleRegistry.CallOpts, _puzzle)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 _puzzle) returns()
func (_BloomPuzzleRegistry *BloomPuzzleRegistryTransactor) Commit(opts *bind.TransactOpts, _puzzle [32]byte) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.contract.Transact(opts, "commit", _puzzle)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 _puzzle) returns()
func (_BloomPuzzleRegistry *BloomPuzzleRegistrySession) Commit(_puzzle [32]byte) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.Contract.Commit(&_BloomPuzzleRegistry.TransactOpts, _puzzle)
}

// Commit is a paid mutator transaction binding the contract method 0xf14fcbc8.
//
// Solidity: function commit(bytes32 _puzzle) returns()
func (_BloomPuzzleRegistry *BloomPuzzleRegistryTransactorSession) Commit(_puzzle [32]byte) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.Contract.Commit(&_BloomPuzzleRegistry.TransactOpts, _puzzle)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BloomPuzzleRegistry *BloomPuzzleRegistryTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BloomPuzzleRegistry *BloomPuzzleRegistrySession) RenounceOwnership() (*types.Transaction, error) {
	return _BloomPuzzleRegistry.Contract.RenounceOwnership(&_BloomPuzzleRegistry.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_BloomPuzzleRegistry *BloomPuzzleRegistryTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _BloomPuzzleRegistry.Contract.RenounceOwnership(&_BloomPuzzleRegistry.TransactOpts)
}

// SetSolution is a paid mutator transaction binding the contract method 0x68e317e2.
//
// Solidity: function setSolution(bytes _solution) returns()
func (_BloomPuzzleRegistry *BloomPuzzleRegistryTransactor) SetSolution(opts *bind.TransactOpts, _solution []byte) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.contract.Transact(opts, "setSolution", _solution)
}

// SetSolution is a paid mutator transaction binding the contract method 0x68e317e2.
//
// Solidity: function setSolution(bytes _solution) returns()
func (_BloomPuzzleRegistry *BloomPuzzleRegistrySession) SetSolution(_solution []byte) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.Contract.SetSolution(&_BloomPuzzleRegistry.TransactOpts, _solution)
}

// SetSolution is a paid mutator transaction binding the contract method 0x68e317e2.
//
// Solidity: function setSolution(bytes _solution) returns()
func (_BloomPuzzleRegistry *BloomPuzzleRegistryTransactorSession) SetSolution(_solution []byte) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.Contract.SetSolution(&_BloomPuzzleRegistry.TransactOpts, _solution)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BloomPuzzleRegistry *BloomPuzzleRegistryTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BloomPuzzleRegistry *BloomPuzzleRegistrySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.Contract.TransferOwnership(&_BloomPuzzleRegistry.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_BloomPuzzleRegistry *BloomPuzzleRegistryTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _BloomPuzzleRegistry.Contract.TransferOwnership(&_BloomPuzzleRegistry.TransactOpts, newOwner)
}

// BloomPuzzleRegistryOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the BloomPuzzleRegistry contract.
type BloomPuzzleRegistryOwnershipTransferredIterator struct {
	Event *BloomPuzzleRegistryOwnershipTransferred // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *BloomPuzzleRegistryOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BloomPuzzleRegistryOwnershipTransferred)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(BloomPuzzleRegistryOwnershipTransferred)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *BloomPuzzleRegistryOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BloomPuzzleRegistryOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BloomPuzzleRegistryOwnershipTransferred represents a OwnershipTransferred event raised by the BloomPuzzleRegistry contract.
type BloomPuzzleRegistryOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BloomPuzzleRegistry *BloomPuzzleRegistryFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*BloomPuzzleRegistryOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BloomPuzzleRegistry.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &BloomPuzzleRegistryOwnershipTransferredIterator{contract: _BloomPuzzleRegistry.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BloomPuzzleRegistry *BloomPuzzleRegistryFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *BloomPuzzleRegistryOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _BloomPuzzleRegistry.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BloomPuzzleRegistryOwnershipTransferred)
				if err := _BloomPuzzleRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_BloomPuzzleRegistry *BloomPuzzleRegistryFilterer) ParseOwnershipTransferred(log types.Log) (*BloomPuzzleRegistryOwnershipTransferred, error) {
	event := new(BloomPuzzleRegistryOwnershipTransferred)
	if err := _BloomPuzzleRegistry.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
