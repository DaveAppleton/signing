// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"constant\":true,\"inputs\":[{\"name\":\"message\",\"type\":\"string\"},{\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"recoverSig\",\"outputs\":[{\"name\":\"addr\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"message\",\"type\":\"string\"},{\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"helperVisit\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"helper\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"whoWasVisited\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"message\",\"type\":\"string\"},{\"indexed\":false,\"name\":\"date\",\"type\":\"uint256\"}],\"name\":\"LogHelperVisit\",\"type\":\"event\"}]"

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// RecoverSig is a free data retrieval call binding the contract method 0x0fc156d6.
//
// Solidity: function recoverSig(message string, sig bytes) constant returns(addr address)
func (_Main *MainCaller) RecoverSig(opts *bind.CallOpts, message string, sig []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "recoverSig", message, sig)
	return *ret0, err
}

// RecoverSig is a free data retrieval call binding the contract method 0x0fc156d6.
//
// Solidity: function recoverSig(message string, sig bytes) constant returns(addr address)
func (_Main *MainSession) RecoverSig(message string, sig []byte) (common.Address, error) {
	return _Main.Contract.RecoverSig(&_Main.CallOpts, message, sig)
}

// RecoverSig is a free data retrieval call binding the contract method 0x0fc156d6.
//
// Solidity: function recoverSig(message string, sig bytes) constant returns(addr address)
func (_Main *MainCallerSession) RecoverSig(message string, sig []byte) (common.Address, error) {
	return _Main.Contract.RecoverSig(&_Main.CallOpts, message, sig)
}

// HelperVisit is a paid mutator transaction binding the contract method 0xfd37a336.
//
// Solidity: function helperVisit(message string, sig bytes) returns()
func (_Main *MainTransactor) HelperVisit(opts *bind.TransactOpts, message string, sig []byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "helperVisit", message, sig)
}

// HelperVisit is a paid mutator transaction binding the contract method 0xfd37a336.
//
// Solidity: function helperVisit(message string, sig bytes) returns()
func (_Main *MainSession) HelperVisit(message string, sig []byte) (*types.Transaction, error) {
	return _Main.Contract.HelperVisit(&_Main.TransactOpts, message, sig)
}

// HelperVisit is a paid mutator transaction binding the contract method 0xfd37a336.
//
// Solidity: function helperVisit(message string, sig bytes) returns()
func (_Main *MainTransactorSession) HelperVisit(message string, sig []byte) (*types.Transaction, error) {
	return _Main.Contract.HelperVisit(&_Main.TransactOpts, message, sig)
}

// MainLogHelperVisitIterator is returned from FilterLogHelperVisit and is used to iterate over the raw logs and unpacked data for LogHelperVisit events raised by the Main contract.
type MainLogHelperVisitIterator struct {
	Event *MainLogHelperVisit // Event containing the contract specifics and raw log

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
func (it *MainLogHelperVisitIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MainLogHelperVisit)
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
		it.Event = new(MainLogHelperVisit)
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
func (it *MainLogHelperVisitIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MainLogHelperVisitIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MainLogHelperVisit represents a LogHelperVisit event raised by the Main contract.
type MainLogHelperVisit struct {
	Helper        common.Address
	WhoWasVisited common.Address
	Message       string
	Date          *big.Int
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterLogHelperVisit is a free log retrieval operation binding the contract event 0x770359b02a985a1dc07df6d37388c7283945a9710ae5e3064c37be323439d97a.
//
// Solidity: e LogHelperVisit(helper address, whoWasVisited address, message string, date uint256)
func (_Main *MainFilterer) FilterLogHelperVisit(opts *bind.FilterOpts) (*MainLogHelperVisitIterator, error) {

	logs, sub, err := _Main.contract.FilterLogs(opts, "LogHelperVisit")
	if err != nil {
		return nil, err
	}
	return &MainLogHelperVisitIterator{contract: _Main.contract, event: "LogHelperVisit", logs: logs, sub: sub}, nil
}

// WatchLogHelperVisit is a free log subscription operation binding the contract event 0x770359b02a985a1dc07df6d37388c7283945a9710ae5e3064c37be323439d97a.
//
// Solidity: e LogHelperVisit(helper address, whoWasVisited address, message string, date uint256)
func (_Main *MainFilterer) WatchLogHelperVisit(opts *bind.WatchOpts, sink chan<- *MainLogHelperVisit) (event.Subscription, error) {

	logs, sub, err := _Main.contract.WatchLogs(opts, "LogHelperVisit")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MainLogHelperVisit)
				if err := _Main.contract.UnpackLog(event, "LogHelperVisit", log); err != nil {
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
