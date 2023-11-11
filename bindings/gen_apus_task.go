// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// ApusDataClientConfig is an auto generated low-level Go binding around an user-defined struct.
type ApusDataClientConfig struct {
	Owner            common.Address
	Id               *big.Int
	Url              string
	MinFee           *big.Int
	MaxZkEvmInstance uint8
	CurInstance      uint8
}

// ApusDataTask is an auto generated low-level Go binding around an user-defined struct.
type ApusDataTask struct {
	Id       *big.Int
	ClientId *big.Int
	UniqID   *big.Int
	Assigner common.Address
	Input    []byte
	Tp       uint8
	Stat     uint8
	Result   []byte
	Reward   ApusDatarewardInfo
	Expiry   uint64
}

// ApusDatarewardInfo is an auto generated low-level Go binding around an user-defined struct.
type ApusDatarewardInfo struct {
	Token  common.Address
	Amount *big.Int
}

// ApusTaskMetaData contains all meta data concerning the ApusTask contract.
var ApusTaskMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_marketAddr\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_tokenAddr\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"fee\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"tokenAddress\",\"type\":\"address\"}],\"name\":\"eventPostTask\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tasks\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"clientId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"uniqID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assigner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"enumApusData.TaskType\",\"name\":\"_tp\",\"type\":\"uint8\"},{\"internalType\":\"enumApusData.TaskStatus\",\"name\":\"_stat\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structApusData.rewardInfo\",\"name\":\"reward\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"enumApusData.TaskType\",\"name\":\"_tp\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uniqID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"}],\"name\":\"submitTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"enumApusData.TaskType\",\"name\":\"_tp\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uniqID\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"clientId\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"uniqID\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"assigner\",\"type\":\"address\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"enumApusData.TaskType\",\"name\":\"_tp\",\"type\":\"uint8\"},{\"internalType\":\"enumApusData.TaskStatus\",\"name\":\"_stat\",\"type\":\"uint8\"},{\"internalType\":\"bytes\",\"name\":\"result\",\"type\":\"bytes\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structApusData.rewardInfo\",\"name\":\"reward\",\"type\":\"tuple\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"}],\"internalType\":\"structApusData.Task\",\"name\":\"\",\"type\":\"tuple\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"url\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"minFee\",\"type\":\"uint256\"},{\"internalType\":\"uint8\",\"name\":\"maxZkEvmInstance\",\"type\":\"uint8\"},{\"internalType\":\"uint8\",\"name\":\"curInstance\",\"type\":\"uint8\"}],\"internalType\":\"structApusData.ClientConfig\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\",\"constant\":true},{\"inputs\":[{\"internalType\":\"enumApusData.TaskType\",\"name\":\"_tp\",\"type\":\"uint8\"},{\"internalType\":\"uint256\",\"name\":\"uniqID\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"input\",\"type\":\"bytes\"},{\"internalType\":\"uint64\",\"name\":\"expiry\",\"type\":\"uint64\"},{\"components\":[{\"internalType\":\"address\",\"name\":\"token\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structApusData.rewardInfo\",\"name\":\"ri\",\"type\":\"tuple\"}],\"name\":\"postTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"taskID\",\"type\":\"uint256\"}],\"name\":\"dispatchTaskToClient\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ApusTaskABI is the input ABI used to generate the binding from.
// Deprecated: Use ApusTaskMetaData.ABI instead.
var ApusTaskABI = ApusTaskMetaData.ABI

// ApusTask is an auto generated Go binding around an Ethereum contract.
type ApusTask struct {
	ApusTaskCaller     // Read-only binding to the contract
	ApusTaskTransactor // Write-only binding to the contract
	ApusTaskFilterer   // Log filterer for contract events
}

// ApusTaskCaller is an auto generated read-only Go binding around an Ethereum contract.
type ApusTaskCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApusTaskTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ApusTaskTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApusTaskFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ApusTaskFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ApusTaskSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ApusTaskSession struct {
	Contract     *ApusTask         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ApusTaskCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ApusTaskCallerSession struct {
	Contract *ApusTaskCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// ApusTaskTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ApusTaskTransactorSession struct {
	Contract     *ApusTaskTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// ApusTaskRaw is an auto generated low-level Go binding around an Ethereum contract.
type ApusTaskRaw struct {
	Contract *ApusTask // Generic contract binding to access the raw methods on
}

// ApusTaskCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ApusTaskCallerRaw struct {
	Contract *ApusTaskCaller // Generic read-only contract binding to access the raw methods on
}

// ApusTaskTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ApusTaskTransactorRaw struct {
	Contract *ApusTaskTransactor // Generic write-only contract binding to access the raw methods on
}

// NewApusTask creates a new instance of ApusTask, bound to a specific deployed contract.
func NewApusTask(address common.Address, backend bind.ContractBackend) (*ApusTask, error) {
	contract, err := bindApusTask(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ApusTask{ApusTaskCaller: ApusTaskCaller{contract: contract}, ApusTaskTransactor: ApusTaskTransactor{contract: contract}, ApusTaskFilterer: ApusTaskFilterer{contract: contract}}, nil
}

// NewApusTaskCaller creates a new read-only instance of ApusTask, bound to a specific deployed contract.
func NewApusTaskCaller(address common.Address, caller bind.ContractCaller) (*ApusTaskCaller, error) {
	contract, err := bindApusTask(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ApusTaskCaller{contract: contract}, nil
}

// NewApusTaskTransactor creates a new write-only instance of ApusTask, bound to a specific deployed contract.
func NewApusTaskTransactor(address common.Address, transactor bind.ContractTransactor) (*ApusTaskTransactor, error) {
	contract, err := bindApusTask(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ApusTaskTransactor{contract: contract}, nil
}

// NewApusTaskFilterer creates a new log filterer instance of ApusTask, bound to a specific deployed contract.
func NewApusTaskFilterer(address common.Address, filterer bind.ContractFilterer) (*ApusTaskFilterer, error) {
	contract, err := bindApusTask(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ApusTaskFilterer{contract: contract}, nil
}

// bindApusTask binds a generic wrapper to an already deployed contract.
func bindApusTask(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ApusTaskMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApusTask *ApusTaskRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApusTask.Contract.ApusTaskCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApusTask *ApusTaskRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApusTask.Contract.ApusTaskTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApusTask *ApusTaskRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApusTask.Contract.ApusTaskTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ApusTask *ApusTaskCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ApusTask.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ApusTask *ApusTaskTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ApusTask.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ApusTask *ApusTaskTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ApusTask.Contract.contract.Transact(opts, method, params...)
}

// GetTask is a free data retrieval call binding the contract method 0x3253da0f.
//
// Solidity: function getTask(uint8 _tp, uint256 uniqID) view returns((uint256,uint256,uint256,address,bytes,uint8,uint8,bytes,(address,uint256),uint64), (address,uint256,string,uint256,uint8,uint8))
func (_ApusTask *ApusTaskCaller) GetTask(opts *bind.CallOpts, _tp uint8, uniqID *big.Int) (ApusDataTask, ApusDataClientConfig, error) {
	var out []interface{}
	err := _ApusTask.contract.Call(opts, &out, "getTask", _tp, uniqID)

	if err != nil {
		return *new(ApusDataTask), *new(ApusDataClientConfig), err
	}

	out0 := *abi.ConvertType(out[0], new(ApusDataTask)).(*ApusDataTask)
	out1 := *abi.ConvertType(out[1], new(ApusDataClientConfig)).(*ApusDataClientConfig)

	return out0, out1, err

}

// GetTask is a free data retrieval call binding the contract method 0x3253da0f.
//
// Solidity: function getTask(uint8 _tp, uint256 uniqID) view returns((uint256,uint256,uint256,address,bytes,uint8,uint8,bytes,(address,uint256),uint64), (address,uint256,string,uint256,uint8,uint8))
func (_ApusTask *ApusTaskSession) GetTask(_tp uint8, uniqID *big.Int) (ApusDataTask, ApusDataClientConfig, error) {
	return _ApusTask.Contract.GetTask(&_ApusTask.CallOpts, _tp, uniqID)
}

// GetTask is a free data retrieval call binding the contract method 0x3253da0f.
//
// Solidity: function getTask(uint8 _tp, uint256 uniqID) view returns((uint256,uint256,uint256,address,bytes,uint8,uint8,bytes,(address,uint256),uint64), (address,uint256,string,uint256,uint8,uint8))
func (_ApusTask *ApusTaskCallerSession) GetTask(_tp uint8, uniqID *big.Int) (ApusDataTask, ApusDataClientConfig, error) {
	return _ApusTask.Contract.GetTask(&_ApusTask.CallOpts, _tp, uniqID)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, uint256 clientId, uint256 uniqID, address assigner, bytes input, uint8 _tp, uint8 _stat, bytes result, (address,uint256) reward, uint64 expiry)
func (_ApusTask *ApusTaskCaller) Tasks(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Id       *big.Int
	ClientId *big.Int
	UniqID   *big.Int
	Assigner common.Address
	Input    []byte
	Tp       uint8
	Stat     uint8
	Result   []byte
	Reward   ApusDatarewardInfo
	Expiry   uint64
}, error) {
	var out []interface{}
	err := _ApusTask.contract.Call(opts, &out, "tasks", arg0)

	outstruct := new(struct {
		Id       *big.Int
		ClientId *big.Int
		UniqID   *big.Int
		Assigner common.Address
		Input    []byte
		Tp       uint8
		Stat     uint8
		Result   []byte
		Reward   ApusDatarewardInfo
		Expiry   uint64
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Id = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.ClientId = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.UniqID = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Assigner = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.Input = *abi.ConvertType(out[4], new([]byte)).(*[]byte)
	outstruct.Tp = *abi.ConvertType(out[5], new(uint8)).(*uint8)
	outstruct.Stat = *abi.ConvertType(out[6], new(uint8)).(*uint8)
	outstruct.Result = *abi.ConvertType(out[7], new([]byte)).(*[]byte)
	outstruct.Reward = *abi.ConvertType(out[8], new(ApusDatarewardInfo)).(*ApusDatarewardInfo)
	outstruct.Expiry = *abi.ConvertType(out[9], new(uint64)).(*uint64)

	return *outstruct, err

}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, uint256 clientId, uint256 uniqID, address assigner, bytes input, uint8 _tp, uint8 _stat, bytes result, (address,uint256) reward, uint64 expiry)
func (_ApusTask *ApusTaskSession) Tasks(arg0 *big.Int) (struct {
	Id       *big.Int
	ClientId *big.Int
	UniqID   *big.Int
	Assigner common.Address
	Input    []byte
	Tp       uint8
	Stat     uint8
	Result   []byte
	Reward   ApusDatarewardInfo
	Expiry   uint64
}, error) {
	return _ApusTask.Contract.Tasks(&_ApusTask.CallOpts, arg0)
}

// Tasks is a free data retrieval call binding the contract method 0x8d977672.
//
// Solidity: function tasks(uint256 ) view returns(uint256 id, uint256 clientId, uint256 uniqID, address assigner, bytes input, uint8 _tp, uint8 _stat, bytes result, (address,uint256) reward, uint64 expiry)
func (_ApusTask *ApusTaskCallerSession) Tasks(arg0 *big.Int) (struct {
	Id       *big.Int
	ClientId *big.Int
	UniqID   *big.Int
	Assigner common.Address
	Input    []byte
	Tp       uint8
	Stat     uint8
	Result   []byte
	Reward   ApusDatarewardInfo
	Expiry   uint64
}, error) {
	return _ApusTask.Contract.Tasks(&_ApusTask.CallOpts, arg0)
}

// DispatchTaskToClient is a paid mutator transaction binding the contract method 0xe5ab75eb.
//
// Solidity: function dispatchTaskToClient(uint256 taskID) returns()
func (_ApusTask *ApusTaskTransactor) DispatchTaskToClient(opts *bind.TransactOpts, taskID *big.Int) (*types.Transaction, error) {
	return _ApusTask.contract.Transact(opts, "dispatchTaskToClient", taskID)
}

// DispatchTaskToClient is a paid mutator transaction binding the contract method 0xe5ab75eb.
//
// Solidity: function dispatchTaskToClient(uint256 taskID) returns()
func (_ApusTask *ApusTaskSession) DispatchTaskToClient(taskID *big.Int) (*types.Transaction, error) {
	return _ApusTask.Contract.DispatchTaskToClient(&_ApusTask.TransactOpts, taskID)
}

// DispatchTaskToClient is a paid mutator transaction binding the contract method 0xe5ab75eb.
//
// Solidity: function dispatchTaskToClient(uint256 taskID) returns()
func (_ApusTask *ApusTaskTransactorSession) DispatchTaskToClient(taskID *big.Int) (*types.Transaction, error) {
	return _ApusTask.Contract.DispatchTaskToClient(&_ApusTask.TransactOpts, taskID)
}

// PostTask is a paid mutator transaction binding the contract method 0x02feed60.
//
// Solidity: function postTask(uint8 _tp, uint256 uniqID, bytes input, uint64 expiry, (address,uint256) ri) returns()
func (_ApusTask *ApusTaskTransactor) PostTask(opts *bind.TransactOpts, _tp uint8, uniqID *big.Int, input []byte, expiry uint64, ri ApusDatarewardInfo) (*types.Transaction, error) {
	return _ApusTask.contract.Transact(opts, "postTask", _tp, uniqID, input, expiry, ri)
}

// PostTask is a paid mutator transaction binding the contract method 0x02feed60.
//
// Solidity: function postTask(uint8 _tp, uint256 uniqID, bytes input, uint64 expiry, (address,uint256) ri) returns()
func (_ApusTask *ApusTaskSession) PostTask(_tp uint8, uniqID *big.Int, input []byte, expiry uint64, ri ApusDatarewardInfo) (*types.Transaction, error) {
	return _ApusTask.Contract.PostTask(&_ApusTask.TransactOpts, _tp, uniqID, input, expiry, ri)
}

// PostTask is a paid mutator transaction binding the contract method 0x02feed60.
//
// Solidity: function postTask(uint8 _tp, uint256 uniqID, bytes input, uint64 expiry, (address,uint256) ri) returns()
func (_ApusTask *ApusTaskTransactorSession) PostTask(_tp uint8, uniqID *big.Int, input []byte, expiry uint64, ri ApusDatarewardInfo) (*types.Transaction, error) {
	return _ApusTask.Contract.PostTask(&_ApusTask.TransactOpts, _tp, uniqID, input, expiry, ri)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x150645d8.
//
// Solidity: function submitTask(uint8 _tp, uint256 uniqID, bytes result) returns()
func (_ApusTask *ApusTaskTransactor) SubmitTask(opts *bind.TransactOpts, _tp uint8, uniqID *big.Int, result []byte) (*types.Transaction, error) {
	return _ApusTask.contract.Transact(opts, "submitTask", _tp, uniqID, result)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x150645d8.
//
// Solidity: function submitTask(uint8 _tp, uint256 uniqID, bytes result) returns()
func (_ApusTask *ApusTaskSession) SubmitTask(_tp uint8, uniqID *big.Int, result []byte) (*types.Transaction, error) {
	return _ApusTask.Contract.SubmitTask(&_ApusTask.TransactOpts, _tp, uniqID, result)
}

// SubmitTask is a paid mutator transaction binding the contract method 0x150645d8.
//
// Solidity: function submitTask(uint8 _tp, uint256 uniqID, bytes result) returns()
func (_ApusTask *ApusTaskTransactorSession) SubmitTask(_tp uint8, uniqID *big.Int, result []byte) (*types.Transaction, error) {
	return _ApusTask.Contract.SubmitTask(&_ApusTask.TransactOpts, _tp, uniqID, result)
}

// ApusTaskEventPostTaskIterator is returned from FilterEventPostTask and is used to iterate over the raw logs and unpacked data for EventPostTask events raised by the ApusTask contract.
type ApusTaskEventPostTaskIterator struct {
	Event *ApusTaskEventPostTask // Event containing the contract specifics and raw log

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
func (it *ApusTaskEventPostTaskIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ApusTaskEventPostTask)
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
		it.Event = new(ApusTaskEventPostTask)
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
func (it *ApusTaskEventPostTaskIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ApusTaskEventPostTaskIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ApusTaskEventPostTask represents a EventPostTask event raised by the ApusTask contract.
type ApusTaskEventPostTask struct {
	TaskId       *big.Int
	Fee          *big.Int
	TokenAddress common.Address
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterEventPostTask is a free log retrieval operation binding the contract event 0x03d907ed51dcdee925c010362ff54d1e02ff47a76efaae5b2ef10b6b9a9aeceb.
//
// Solidity: event eventPostTask(uint256 taskId, uint256 fee, address tokenAddress)
func (_ApusTask *ApusTaskFilterer) FilterEventPostTask(opts *bind.FilterOpts) (*ApusTaskEventPostTaskIterator, error) {

	logs, sub, err := _ApusTask.contract.FilterLogs(opts, "eventPostTask")
	if err != nil {
		return nil, err
	}
	return &ApusTaskEventPostTaskIterator{contract: _ApusTask.contract, event: "eventPostTask", logs: logs, sub: sub}, nil
}

// WatchEventPostTask is a free log subscription operation binding the contract event 0x03d907ed51dcdee925c010362ff54d1e02ff47a76efaae5b2ef10b6b9a9aeceb.
//
// Solidity: event eventPostTask(uint256 taskId, uint256 fee, address tokenAddress)
func (_ApusTask *ApusTaskFilterer) WatchEventPostTask(opts *bind.WatchOpts, sink chan<- *ApusTaskEventPostTask) (event.Subscription, error) {

	logs, sub, err := _ApusTask.contract.WatchLogs(opts, "eventPostTask")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ApusTaskEventPostTask)
				if err := _ApusTask.contract.UnpackLog(event, "eventPostTask", log); err != nil {
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

// ParseEventPostTask is a log parse operation binding the contract event 0x03d907ed51dcdee925c010362ff54d1e02ff47a76efaae5b2ef10b6b9a9aeceb.
//
// Solidity: event eventPostTask(uint256 taskId, uint256 fee, address tokenAddress)
func (_ApusTask *ApusTaskFilterer) ParseEventPostTask(log types.Log) (*ApusTaskEventPostTask, error) {
	event := new(ApusTaskEventPostTask)
	if err := _ApusTask.contract.UnpackLog(event, "eventPostTask", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
