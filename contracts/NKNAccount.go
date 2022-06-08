// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contracts

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
)

// NKNAccountNKNAddress is an auto generated low-level Go binding around an user-defined struct.
type NKNAccountNKNAddress struct {
	Identifier string
	PublicKey  [32]byte
}

// NKNAccountMetaData contains all meta data concerning the NKNAccount contract.
var NKNAccountMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[],\"name\":\"del\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"publicKey\",\"type\":\"address\"}],\"name\":\"getNKNAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"publicKey\",\"type\":\"bytes32\"}],\"internalType\":\"structNKNAccount.NKNAddress\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNKNAddr\",\"outputs\":[{\"components\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"publicKey\",\"type\":\"bytes32\"}],\"internalType\":\"structNKNAccount.NKNAddress\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"identifier\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"publicKey\",\"type\":\"bytes32\"}],\"name\":\"set\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// NKNAccountABI is the input ABI used to generate the binding from.
// Deprecated: Use NKNAccountMetaData.ABI instead.
var NKNAccountABI = NKNAccountMetaData.ABI

// NKNAccount is an auto generated Go binding around an Ethereum contract.
type NKNAccount struct {
	NKNAccountCaller     // Read-only binding to the contract
	NKNAccountTransactor // Write-only binding to the contract
	NKNAccountFilterer   // Log filterer for contract events
}

// NKNAccountCaller is an auto generated read-only Go binding around an Ethereum contract.
type NKNAccountCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NKNAccountTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NKNAccountTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NKNAccountFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NKNAccountFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NKNAccountSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NKNAccountSession struct {
	Contract     *NKNAccount       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NKNAccountCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NKNAccountCallerSession struct {
	Contract *NKNAccountCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// NKNAccountTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NKNAccountTransactorSession struct {
	Contract     *NKNAccountTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// NKNAccountRaw is an auto generated low-level Go binding around an Ethereum contract.
type NKNAccountRaw struct {
	Contract *NKNAccount // Generic contract binding to access the raw methods on
}

// NKNAccountCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NKNAccountCallerRaw struct {
	Contract *NKNAccountCaller // Generic read-only contract binding to access the raw methods on
}

// NKNAccountTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NKNAccountTransactorRaw struct {
	Contract *NKNAccountTransactor // Generic write-only contract binding to access the raw methods on
}

// NewNKNAccount creates a new instance of NKNAccount, bound to a specific deployed contract.
func NewNKNAccount(address common.Address, backend bind.ContractBackend) (*NKNAccount, error) {
	contract, err := bindNKNAccount(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &NKNAccount{NKNAccountCaller: NKNAccountCaller{contract: contract}, NKNAccountTransactor: NKNAccountTransactor{contract: contract}, NKNAccountFilterer: NKNAccountFilterer{contract: contract}}, nil
}

// NewNKNAccountCaller creates a new read-only instance of NKNAccount, bound to a specific deployed contract.
func NewNKNAccountCaller(address common.Address, caller bind.ContractCaller) (*NKNAccountCaller, error) {
	contract, err := bindNKNAccount(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NKNAccountCaller{contract: contract}, nil
}

// NewNKNAccountTransactor creates a new write-only instance of NKNAccount, bound to a specific deployed contract.
func NewNKNAccountTransactor(address common.Address, transactor bind.ContractTransactor) (*NKNAccountTransactor, error) {
	contract, err := bindNKNAccount(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NKNAccountTransactor{contract: contract}, nil
}

// NewNKNAccountFilterer creates a new log filterer instance of NKNAccount, bound to a specific deployed contract.
func NewNKNAccountFilterer(address common.Address, filterer bind.ContractFilterer) (*NKNAccountFilterer, error) {
	contract, err := bindNKNAccount(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NKNAccountFilterer{contract: contract}, nil
}

// bindNKNAccount binds a generic wrapper to an already deployed contract.
func bindNKNAccount(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NKNAccountABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NKNAccount *NKNAccountRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NKNAccount.Contract.NKNAccountCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NKNAccount *NKNAccountRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NKNAccount.Contract.NKNAccountTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NKNAccount *NKNAccountRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NKNAccount.Contract.NKNAccountTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_NKNAccount *NKNAccountCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _NKNAccount.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_NKNAccount *NKNAccountTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NKNAccount.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_NKNAccount *NKNAccountTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _NKNAccount.Contract.contract.Transact(opts, method, params...)
}

// GetNKNAddr is a free data retrieval call binding the contract method 0xadfddd7e.
//
// Solidity: function getNKNAddr(address publicKey) view returns((string,bytes32))
func (_NKNAccount *NKNAccountCaller) GetNKNAddr(opts *bind.CallOpts, publicKey common.Address) (NKNAccountNKNAddress, error) {
	var out []interface{}
	err := _NKNAccount.contract.Call(opts, &out, "getNKNAddr", publicKey)

	if err != nil {
		return *new(NKNAccountNKNAddress), err
	}

	out0 := *abi.ConvertType(out[0], new(NKNAccountNKNAddress)).(*NKNAccountNKNAddress)

	return out0, err

}

// GetNKNAddr is a free data retrieval call binding the contract method 0xadfddd7e.
//
// Solidity: function getNKNAddr(address publicKey) view returns((string,bytes32))
func (_NKNAccount *NKNAccountSession) GetNKNAddr(publicKey common.Address) (NKNAccountNKNAddress, error) {
	return _NKNAccount.Contract.GetNKNAddr(&_NKNAccount.CallOpts, publicKey)
}

// GetNKNAddr is a free data retrieval call binding the contract method 0xadfddd7e.
//
// Solidity: function getNKNAddr(address publicKey) view returns((string,bytes32))
func (_NKNAccount *NKNAccountCallerSession) GetNKNAddr(publicKey common.Address) (NKNAccountNKNAddress, error) {
	return _NKNAccount.Contract.GetNKNAddr(&_NKNAccount.CallOpts, publicKey)
}

// GetNKNAddr0 is a free data retrieval call binding the contract method 0xba1bf4b7.
//
// Solidity: function getNKNAddr() view returns((string,bytes32))
func (_NKNAccount *NKNAccountCaller) GetNKNAddr0(opts *bind.CallOpts) (NKNAccountNKNAddress, error) {
	var out []interface{}
	err := _NKNAccount.contract.Call(opts, &out, "getNKNAddr0")

	if err != nil {
		return *new(NKNAccountNKNAddress), err
	}

	out0 := *abi.ConvertType(out[0], new(NKNAccountNKNAddress)).(*NKNAccountNKNAddress)

	return out0, err

}

// GetNKNAddr0 is a free data retrieval call binding the contract method 0xba1bf4b7.
//
// Solidity: function getNKNAddr() view returns((string,bytes32))
func (_NKNAccount *NKNAccountSession) GetNKNAddr0() (NKNAccountNKNAddress, error) {
	return _NKNAccount.Contract.GetNKNAddr0(&_NKNAccount.CallOpts)
}

// GetNKNAddr0 is a free data retrieval call binding the contract method 0xba1bf4b7.
//
// Solidity: function getNKNAddr() view returns((string,bytes32))
func (_NKNAccount *NKNAccountCallerSession) GetNKNAddr0() (NKNAccountNKNAddress, error) {
	return _NKNAccount.Contract.GetNKNAddr0(&_NKNAccount.CallOpts)
}

// Del is a paid mutator transaction binding the contract method 0xb6588ffd.
//
// Solidity: function del() returns()
func (_NKNAccount *NKNAccountTransactor) Del(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _NKNAccount.contract.Transact(opts, "del")
}

// Del is a paid mutator transaction binding the contract method 0xb6588ffd.
//
// Solidity: function del() returns()
func (_NKNAccount *NKNAccountSession) Del() (*types.Transaction, error) {
	return _NKNAccount.Contract.Del(&_NKNAccount.TransactOpts)
}

// Del is a paid mutator transaction binding the contract method 0xb6588ffd.
//
// Solidity: function del() returns()
func (_NKNAccount *NKNAccountTransactorSession) Del() (*types.Transaction, error) {
	return _NKNAccount.Contract.Del(&_NKNAccount.TransactOpts)
}

// Set is a paid mutator transaction binding the contract method 0x2e3196a5.
//
// Solidity: function set(string identifier, bytes32 publicKey) returns()
func (_NKNAccount *NKNAccountTransactor) Set(opts *bind.TransactOpts, identifier string, publicKey [32]byte) (*types.Transaction, error) {
	return _NKNAccount.contract.Transact(opts, "set", identifier, publicKey)
}

// Set is a paid mutator transaction binding the contract method 0x2e3196a5.
//
// Solidity: function set(string identifier, bytes32 publicKey) returns()
func (_NKNAccount *NKNAccountSession) Set(identifier string, publicKey [32]byte) (*types.Transaction, error) {
	return _NKNAccount.Contract.Set(&_NKNAccount.TransactOpts, identifier, publicKey)
}

// Set is a paid mutator transaction binding the contract method 0x2e3196a5.
//
// Solidity: function set(string identifier, bytes32 publicKey) returns()
func (_NKNAccount *NKNAccountTransactorSession) Set(identifier string, publicKey [32]byte) (*types.Transaction, error) {
	return _NKNAccount.Contract.Set(&_NKNAccount.TransactOpts, identifier, publicKey)
}
