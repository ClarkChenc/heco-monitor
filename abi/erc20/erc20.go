// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package erc20

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

// Erc20MetaData contains all meta data concerning the Erc20 contract.
var Erc20MetaData = &bind.MetaData{
	ABI: "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"supply\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"class\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"modifyAdminAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"internalType\":\"uint8\",\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskHash\",\"type\":\"bytes32\"}],\"name\":\"dropTask\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOwnerRequiredNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"class\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"oneAddress\",\"type\":\"address\"}],\"name\":\"dropAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getLogicAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getStoreAddress\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isPaused\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"getOperatorRequiredNum\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"string\",\"name\":\"class\",\"type\":\"string\"}],\"name\":\"getAdminAddresses\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"class\",\"type\":\"string\"},{\"internalType\":\"address\",\"name\":\"oneAddress\",\"type\":\"address\"}],\"name\":\"addAddress\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"taskHash\",\"type\":\"bytes32\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"remaining\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"string\",\"name\":\"class\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"requiredNum\",\"type\":\"uint256\"}],\"name\":\"resetRequiredNum\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"taskHash\",\"type\":\"bytes32\"}],\"name\":\"cancelTask\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"internalType\":\"bytes32\",\"name\":\"taskHash\",\"type\":\"bytes32\"}],\"name\":\"burn\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner0\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner1\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"owner2\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"Burning\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"btcAddress\",\"type\":\"string\"}],\"name\":\"Burned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"minter\",\"type\":\"address\"}],\"name\":\"Minting\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"proof\",\"type\":\"string\"}],\"name\":\"Minted\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"TaskType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"class\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"oldAddress\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"newAddress\",\"type\":\"address\"}],\"name\":\"AdminChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"string\",\"name\":\"TaskType\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"class\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"previousNum\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"requiredNum\",\"type\":\"uint256\"}],\"name\":\"AdminRequiredNumChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bytes32\",\"name\":\"taskHash\",\"type\":\"bytes32\"}],\"name\":\"AdminTaskDropped\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"pauseState\",\"type\":\"bool\"}],\"name\":\"PauseChangedTo\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]",
}

// Erc20ABI is the input ABI used to generate the binding from.
// Deprecated: Use Erc20MetaData.ABI instead.
var Erc20ABI = Erc20MetaData.ABI

// Erc20 is an auto generated Go binding around an Ethereum contract.
type Erc20 struct {
	Erc20Caller     // Read-only binding to the contract
	Erc20Transactor // Write-only binding to the contract
	Erc20Filterer   // Log filterer for contract events
}

// Erc20Caller is an auto generated read-only Go binding around an Ethereum contract.
type Erc20Caller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Transactor is an auto generated write-only Go binding around an Ethereum contract.
type Erc20Transactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Filterer is an auto generated log filtering Go binding around an Ethereum contract events.
type Erc20Filterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// Erc20Session is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type Erc20Session struct {
	Contract     *Erc20            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20CallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type Erc20CallerSession struct {
	Contract *Erc20Caller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// Erc20TransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type Erc20TransactorSession struct {
	Contract     *Erc20Transactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// Erc20Raw is an auto generated low-level Go binding around an Ethereum contract.
type Erc20Raw struct {
	Contract *Erc20 // Generic contract binding to access the raw methods on
}

// Erc20CallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type Erc20CallerRaw struct {
	Contract *Erc20Caller // Generic read-only contract binding to access the raw methods on
}

// Erc20TransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type Erc20TransactorRaw struct {
	Contract *Erc20Transactor // Generic write-only contract binding to access the raw methods on
}

// NewErc20 creates a new instance of Erc20, bound to a specific deployed contract.
func NewErc20(address common.Address, backend bind.ContractBackend) (*Erc20, error) {
	contract, err := bindErc20(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Erc20{Erc20Caller: Erc20Caller{contract: contract}, Erc20Transactor: Erc20Transactor{contract: contract}, Erc20Filterer: Erc20Filterer{contract: contract}}, nil
}

// NewErc20Caller creates a new read-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Caller(address common.Address, caller bind.ContractCaller) (*Erc20Caller, error) {
	contract, err := bindErc20(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Caller{contract: contract}, nil
}

// NewErc20Transactor creates a new write-only instance of Erc20, bound to a specific deployed contract.
func NewErc20Transactor(address common.Address, transactor bind.ContractTransactor) (*Erc20Transactor, error) {
	contract, err := bindErc20(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &Erc20Transactor{contract: contract}, nil
}

// NewErc20Filterer creates a new log filterer instance of Erc20, bound to a specific deployed contract.
func NewErc20Filterer(address common.Address, filterer bind.ContractFilterer) (*Erc20Filterer, error) {
	contract, err := bindErc20(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &Erc20Filterer{contract: contract}, nil
}

// bindErc20 binds a generic wrapper to an already deployed contract.
func bindErc20(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(Erc20ABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20Raw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.Erc20Caller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20Raw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20Raw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.Erc20Transactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Erc20 *Erc20CallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Erc20.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Erc20 *Erc20TransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Erc20 *Erc20TransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Erc20.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_Erc20 *Erc20Caller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "allowance", owner, spender)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_Erc20 *Erc20Session) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(address owner, address spender) view returns(uint256 remaining)
func (_Erc20 *Erc20CallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _Erc20.Contract.Allowance(&_Erc20.CallOpts, owner, spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_Erc20 *Erc20Caller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_Erc20 *Erc20Session) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256 balance)
func (_Erc20 *Erc20CallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _Erc20.Contract.BalanceOf(&_Erc20.CallOpts, owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20Caller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "decimals")

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20Session) Decimals() (uint8, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() view returns(uint8)
func (_Erc20 *Erc20CallerSession) Decimals() (uint8, error) {
	return _Erc20.Contract.Decimals(&_Erc20.CallOpts)
}

// GetAdminAddresses is a free data retrieval call binding the contract method 0xc8e369bf.
//
// Solidity: function getAdminAddresses(string class) view returns(address[])
func (_Erc20 *Erc20Caller) GetAdminAddresses(opts *bind.CallOpts, class string) ([]common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "getAdminAddresses", class)

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetAdminAddresses is a free data retrieval call binding the contract method 0xc8e369bf.
//
// Solidity: function getAdminAddresses(string class) view returns(address[])
func (_Erc20 *Erc20Session) GetAdminAddresses(class string) ([]common.Address, error) {
	return _Erc20.Contract.GetAdminAddresses(&_Erc20.CallOpts, class)
}

// GetAdminAddresses is a free data retrieval call binding the contract method 0xc8e369bf.
//
// Solidity: function getAdminAddresses(string class) view returns(address[])
func (_Erc20 *Erc20CallerSession) GetAdminAddresses(class string) ([]common.Address, error) {
	return _Erc20.Contract.GetAdminAddresses(&_Erc20.CallOpts, class)
}

// GetLogicAddress is a free data retrieval call binding the contract method 0xabd108ba.
//
// Solidity: function getLogicAddress() view returns(address)
func (_Erc20 *Erc20Caller) GetLogicAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "getLogicAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetLogicAddress is a free data retrieval call binding the contract method 0xabd108ba.
//
// Solidity: function getLogicAddress() view returns(address)
func (_Erc20 *Erc20Session) GetLogicAddress() (common.Address, error) {
	return _Erc20.Contract.GetLogicAddress(&_Erc20.CallOpts)
}

// GetLogicAddress is a free data retrieval call binding the contract method 0xabd108ba.
//
// Solidity: function getLogicAddress() view returns(address)
func (_Erc20 *Erc20CallerSession) GetLogicAddress() (common.Address, error) {
	return _Erc20.Contract.GetLogicAddress(&_Erc20.CallOpts)
}

// GetOperatorRequiredNum is a free data retrieval call binding the contract method 0xb1d3e05e.
//
// Solidity: function getOperatorRequiredNum() view returns(uint256)
func (_Erc20 *Erc20Caller) GetOperatorRequiredNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "getOperatorRequiredNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOperatorRequiredNum is a free data retrieval call binding the contract method 0xb1d3e05e.
//
// Solidity: function getOperatorRequiredNum() view returns(uint256)
func (_Erc20 *Erc20Session) GetOperatorRequiredNum() (*big.Int, error) {
	return _Erc20.Contract.GetOperatorRequiredNum(&_Erc20.CallOpts)
}

// GetOperatorRequiredNum is a free data retrieval call binding the contract method 0xb1d3e05e.
//
// Solidity: function getOperatorRequiredNum() view returns(uint256)
func (_Erc20 *Erc20CallerSession) GetOperatorRequiredNum() (*big.Int, error) {
	return _Erc20.Contract.GetOperatorRequiredNum(&_Erc20.CallOpts)
}

// GetOwnerRequiredNum is a free data retrieval call binding the contract method 0x5dd882dd.
//
// Solidity: function getOwnerRequiredNum() view returns(uint256)
func (_Erc20 *Erc20Caller) GetOwnerRequiredNum(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "getOwnerRequiredNum")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetOwnerRequiredNum is a free data retrieval call binding the contract method 0x5dd882dd.
//
// Solidity: function getOwnerRequiredNum() view returns(uint256)
func (_Erc20 *Erc20Session) GetOwnerRequiredNum() (*big.Int, error) {
	return _Erc20.Contract.GetOwnerRequiredNum(&_Erc20.CallOpts)
}

// GetOwnerRequiredNum is a free data retrieval call binding the contract method 0x5dd882dd.
//
// Solidity: function getOwnerRequiredNum() view returns(uint256)
func (_Erc20 *Erc20CallerSession) GetOwnerRequiredNum() (*big.Int, error) {
	return _Erc20.Contract.GetOwnerRequiredNum(&_Erc20.CallOpts)
}

// GetStoreAddress is a free data retrieval call binding the contract method 0xb07ed3a9.
//
// Solidity: function getStoreAddress() view returns(address)
func (_Erc20 *Erc20Caller) GetStoreAddress(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "getStoreAddress")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetStoreAddress is a free data retrieval call binding the contract method 0xb07ed3a9.
//
// Solidity: function getStoreAddress() view returns(address)
func (_Erc20 *Erc20Session) GetStoreAddress() (common.Address, error) {
	return _Erc20.Contract.GetStoreAddress(&_Erc20.CallOpts)
}

// GetStoreAddress is a free data retrieval call binding the contract method 0xb07ed3a9.
//
// Solidity: function getStoreAddress() view returns(address)
func (_Erc20 *Erc20CallerSession) GetStoreAddress() (common.Address, error) {
	return _Erc20.Contract.GetStoreAddress(&_Erc20.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Erc20 *Erc20Caller) IsPaused(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "isPaused")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Erc20 *Erc20Session) IsPaused() (bool, error) {
	return _Erc20.Contract.IsPaused(&_Erc20.CallOpts)
}

// IsPaused is a free data retrieval call binding the contract method 0xb187bd26.
//
// Solidity: function isPaused() view returns(bool)
func (_Erc20 *Erc20CallerSession) IsPaused() (bool, error) {
	return _Erc20.Contract.IsPaused(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Caller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20Session) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_Erc20 *Erc20CallerSession) Name() (string, error) {
	return _Erc20.Contract.Name(&_Erc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Caller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20Session) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_Erc20 *Erc20CallerSession) Symbol() (string, error) {
	return _Erc20.Contract.Symbol(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 supply)
func (_Erc20 *Erc20Caller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Erc20.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 supply)
func (_Erc20 *Erc20Session) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256 supply)
func (_Erc20 *Erc20CallerSession) TotalSupply() (*big.Int, error) {
	return _Erc20.Contract.TotalSupply(&_Erc20.CallOpts)
}

// AddAddress is a paid mutator transaction binding the contract method 0xceb35b0f.
//
// Solidity: function addAddress(string class, address oneAddress) returns(bool)
func (_Erc20 *Erc20Transactor) AddAddress(opts *bind.TransactOpts, class string, oneAddress common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "addAddress", class, oneAddress)
}

// AddAddress is a paid mutator transaction binding the contract method 0xceb35b0f.
//
// Solidity: function addAddress(string class, address oneAddress) returns(bool)
func (_Erc20 *Erc20Session) AddAddress(class string, oneAddress common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.AddAddress(&_Erc20.TransactOpts, class, oneAddress)
}

// AddAddress is a paid mutator transaction binding the contract method 0xceb35b0f.
//
// Solidity: function addAddress(string class, address oneAddress) returns(bool)
func (_Erc20 *Erc20TransactorSession) AddAddress(class string, oneAddress common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.AddAddress(&_Erc20.TransactOpts, class, oneAddress)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc20 *Erc20Transactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc20 *Erc20Session) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address spender, uint256 value) returns(bool)
func (_Erc20 *Erc20TransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Approve(&_Erc20.TransactOpts, spender, value)
}

// Burn is a paid mutator transaction binding the contract method 0xf691b71a.
//
// Solidity: function burn(address from, uint256 value, string btcAddress, string proof, bytes32 taskHash) returns(bool)
func (_Erc20 *Erc20Transactor) Burn(opts *bind.TransactOpts, from common.Address, value *big.Int, btcAddress string, proof string, taskHash [32]byte) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "burn", from, value, btcAddress, proof, taskHash)
}

// Burn is a paid mutator transaction binding the contract method 0xf691b71a.
//
// Solidity: function burn(address from, uint256 value, string btcAddress, string proof, bytes32 taskHash) returns(bool)
func (_Erc20 *Erc20Session) Burn(from common.Address, value *big.Int, btcAddress string, proof string, taskHash [32]byte) (*types.Transaction, error) {
	return _Erc20.Contract.Burn(&_Erc20.TransactOpts, from, value, btcAddress, proof, taskHash)
}

// Burn is a paid mutator transaction binding the contract method 0xf691b71a.
//
// Solidity: function burn(address from, uint256 value, string btcAddress, string proof, bytes32 taskHash) returns(bool)
func (_Erc20 *Erc20TransactorSession) Burn(from common.Address, value *big.Int, btcAddress string, proof string, taskHash [32]byte) (*types.Transaction, error) {
	return _Erc20.Contract.Burn(&_Erc20.TransactOpts, from, value, btcAddress, proof, taskHash)
}

// CancelTask is a paid mutator transaction binding the contract method 0xee8ca3b5.
//
// Solidity: function cancelTask(bytes32 taskHash) returns(uint256)
func (_Erc20 *Erc20Transactor) CancelTask(opts *bind.TransactOpts, taskHash [32]byte) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "cancelTask", taskHash)
}

// CancelTask is a paid mutator transaction binding the contract method 0xee8ca3b5.
//
// Solidity: function cancelTask(bytes32 taskHash) returns(uint256)
func (_Erc20 *Erc20Session) CancelTask(taskHash [32]byte) (*types.Transaction, error) {
	return _Erc20.Contract.CancelTask(&_Erc20.TransactOpts, taskHash)
}

// CancelTask is a paid mutator transaction binding the contract method 0xee8ca3b5.
//
// Solidity: function cancelTask(bytes32 taskHash) returns(uint256)
func (_Erc20 *Erc20TransactorSession) CancelTask(taskHash [32]byte) (*types.Transaction, error) {
	return _Erc20.Contract.CancelTask(&_Erc20.TransactOpts, taskHash)
}

// DropAddress is a paid mutator transaction binding the contract method 0x91c40bf7.
//
// Solidity: function dropAddress(string class, address oneAddress) returns(bool)
func (_Erc20 *Erc20Transactor) DropAddress(opts *bind.TransactOpts, class string, oneAddress common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "dropAddress", class, oneAddress)
}

// DropAddress is a paid mutator transaction binding the contract method 0x91c40bf7.
//
// Solidity: function dropAddress(string class, address oneAddress) returns(bool)
func (_Erc20 *Erc20Session) DropAddress(class string, oneAddress common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DropAddress(&_Erc20.TransactOpts, class, oneAddress)
}

// DropAddress is a paid mutator transaction binding the contract method 0x91c40bf7.
//
// Solidity: function dropAddress(string class, address oneAddress) returns(bool)
func (_Erc20 *Erc20TransactorSession) DropAddress(class string, oneAddress common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.DropAddress(&_Erc20.TransactOpts, class, oneAddress)
}

// DropTask is a paid mutator transaction binding the contract method 0x521cb590.
//
// Solidity: function dropTask(bytes32 taskHash) returns(bool)
func (_Erc20 *Erc20Transactor) DropTask(opts *bind.TransactOpts, taskHash [32]byte) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "dropTask", taskHash)
}

// DropTask is a paid mutator transaction binding the contract method 0x521cb590.
//
// Solidity: function dropTask(bytes32 taskHash) returns(bool)
func (_Erc20 *Erc20Session) DropTask(taskHash [32]byte) (*types.Transaction, error) {
	return _Erc20.Contract.DropTask(&_Erc20.TransactOpts, taskHash)
}

// DropTask is a paid mutator transaction binding the contract method 0x521cb590.
//
// Solidity: function dropTask(bytes32 taskHash) returns(bool)
func (_Erc20 *Erc20TransactorSession) DropTask(taskHash [32]byte) (*types.Transaction, error) {
	return _Erc20.Contract.DropTask(&_Erc20.TransactOpts, taskHash)
}

// Mint is a paid mutator transaction binding the contract method 0xd4bf51a7.
//
// Solidity: function mint(address to, uint256 value, string proof, bytes32 taskHash) returns(bool)
func (_Erc20 *Erc20Transactor) Mint(opts *bind.TransactOpts, to common.Address, value *big.Int, proof string, taskHash [32]byte) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "mint", to, value, proof, taskHash)
}

// Mint is a paid mutator transaction binding the contract method 0xd4bf51a7.
//
// Solidity: function mint(address to, uint256 value, string proof, bytes32 taskHash) returns(bool)
func (_Erc20 *Erc20Session) Mint(to common.Address, value *big.Int, proof string, taskHash [32]byte) (*types.Transaction, error) {
	return _Erc20.Contract.Mint(&_Erc20.TransactOpts, to, value, proof, taskHash)
}

// Mint is a paid mutator transaction binding the contract method 0xd4bf51a7.
//
// Solidity: function mint(address to, uint256 value, string proof, bytes32 taskHash) returns(bool)
func (_Erc20 *Erc20TransactorSession) Mint(to common.Address, value *big.Int, proof string, taskHash [32]byte) (*types.Transaction, error) {
	return _Erc20.Contract.Mint(&_Erc20.TransactOpts, to, value, proof, taskHash)
}

// ModifyAdminAddress is a paid mutator transaction binding the contract method 0x2c2df742.
//
// Solidity: function modifyAdminAddress(string class, address oldAddress, address newAddress) returns()
func (_Erc20 *Erc20Transactor) ModifyAdminAddress(opts *bind.TransactOpts, class string, oldAddress common.Address, newAddress common.Address) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "modifyAdminAddress", class, oldAddress, newAddress)
}

// ModifyAdminAddress is a paid mutator transaction binding the contract method 0x2c2df742.
//
// Solidity: function modifyAdminAddress(string class, address oldAddress, address newAddress) returns()
func (_Erc20 *Erc20Session) ModifyAdminAddress(class string, oldAddress common.Address, newAddress common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.ModifyAdminAddress(&_Erc20.TransactOpts, class, oldAddress, newAddress)
}

// ModifyAdminAddress is a paid mutator transaction binding the contract method 0x2c2df742.
//
// Solidity: function modifyAdminAddress(string class, address oldAddress, address newAddress) returns()
func (_Erc20 *Erc20TransactorSession) ModifyAdminAddress(class string, oldAddress common.Address, newAddress common.Address) (*types.Transaction, error) {
	return _Erc20.Contract.ModifyAdminAddress(&_Erc20.TransactOpts, class, oldAddress, newAddress)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Erc20 *Erc20Transactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Erc20 *Erc20Session) Pause() (*types.Transaction, error) {
	return _Erc20.Contract.Pause(&_Erc20.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_Erc20 *Erc20TransactorSession) Pause() (*types.Transaction, error) {
	return _Erc20.Contract.Pause(&_Erc20.TransactOpts)
}

// ResetRequiredNum is a paid mutator transaction binding the contract method 0xe9e211bd.
//
// Solidity: function resetRequiredNum(string class, uint256 requiredNum) returns(bool)
func (_Erc20 *Erc20Transactor) ResetRequiredNum(opts *bind.TransactOpts, class string, requiredNum *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "resetRequiredNum", class, requiredNum)
}

// ResetRequiredNum is a paid mutator transaction binding the contract method 0xe9e211bd.
//
// Solidity: function resetRequiredNum(string class, uint256 requiredNum) returns(bool)
func (_Erc20 *Erc20Session) ResetRequiredNum(class string, requiredNum *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.ResetRequiredNum(&_Erc20.TransactOpts, class, requiredNum)
}

// ResetRequiredNum is a paid mutator transaction binding the contract method 0xe9e211bd.
//
// Solidity: function resetRequiredNum(string class, uint256 requiredNum) returns(bool)
func (_Erc20 *Erc20TransactorSession) ResetRequiredNum(class string, requiredNum *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.ResetRequiredNum(&_Erc20.TransactOpts, class, requiredNum)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc20 *Erc20Transactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc20 *Erc20Session) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(address to, uint256 value) returns(bool)
func (_Erc20 *Erc20TransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.Transfer(&_Erc20.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc20 *Erc20Transactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc20 *Erc20Session) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 value) returns(bool)
func (_Erc20 *Erc20TransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _Erc20.Contract.TransferFrom(&_Erc20.TransactOpts, from, to, value)
}

// Erc20AdminChangedIterator is returned from FilterAdminChanged and is used to iterate over the raw logs and unpacked data for AdminChanged events raised by the Erc20 contract.
type Erc20AdminChangedIterator struct {
	Event *Erc20AdminChanged // Event containing the contract specifics and raw log

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
func (it *Erc20AdminChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20AdminChanged)
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
		it.Event = new(Erc20AdminChanged)
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
func (it *Erc20AdminChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20AdminChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20AdminChanged represents a AdminChanged event raised by the Erc20 contract.
type Erc20AdminChanged struct {
	TaskType   string
	Class      string
	OldAddress common.Address
	NewAddress common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterAdminChanged is a free log retrieval operation binding the contract event 0xb02ddf2800cd4468f7eb14268a7c9f1b6e816aa425ed542d13d8d3a96fa44566.
//
// Solidity: event AdminChanged(string TaskType, string class, address oldAddress, address newAddress)
func (_Erc20 *Erc20Filterer) FilterAdminChanged(opts *bind.FilterOpts) (*Erc20AdminChangedIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return &Erc20AdminChangedIterator{contract: _Erc20.contract, event: "AdminChanged", logs: logs, sub: sub}, nil
}

// WatchAdminChanged is a free log subscription operation binding the contract event 0xb02ddf2800cd4468f7eb14268a7c9f1b6e816aa425ed542d13d8d3a96fa44566.
//
// Solidity: event AdminChanged(string TaskType, string class, address oldAddress, address newAddress)
func (_Erc20 *Erc20Filterer) WatchAdminChanged(opts *bind.WatchOpts, sink chan<- *Erc20AdminChanged) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "AdminChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20AdminChanged)
				if err := _Erc20.contract.UnpackLog(event, "AdminChanged", log); err != nil {
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

// ParseAdminChanged is a log parse operation binding the contract event 0xb02ddf2800cd4468f7eb14268a7c9f1b6e816aa425ed542d13d8d3a96fa44566.
//
// Solidity: event AdminChanged(string TaskType, string class, address oldAddress, address newAddress)
func (_Erc20 *Erc20Filterer) ParseAdminChanged(log types.Log) (*Erc20AdminChanged, error) {
	event := new(Erc20AdminChanged)
	if err := _Erc20.contract.UnpackLog(event, "AdminChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20AdminRequiredNumChangedIterator is returned from FilterAdminRequiredNumChanged and is used to iterate over the raw logs and unpacked data for AdminRequiredNumChanged events raised by the Erc20 contract.
type Erc20AdminRequiredNumChangedIterator struct {
	Event *Erc20AdminRequiredNumChanged // Event containing the contract specifics and raw log

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
func (it *Erc20AdminRequiredNumChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20AdminRequiredNumChanged)
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
		it.Event = new(Erc20AdminRequiredNumChanged)
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
func (it *Erc20AdminRequiredNumChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20AdminRequiredNumChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20AdminRequiredNumChanged represents a AdminRequiredNumChanged event raised by the Erc20 contract.
type Erc20AdminRequiredNumChanged struct {
	TaskType    string
	Class       string
	PreviousNum *big.Int
	RequiredNum *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterAdminRequiredNumChanged is a free log retrieval operation binding the contract event 0xce49fc04234925e87b95750e0e50cac6d4ffcf8a696e3316b3e13bedc84ee7a8.
//
// Solidity: event AdminRequiredNumChanged(string TaskType, string class, uint256 previousNum, uint256 requiredNum)
func (_Erc20 *Erc20Filterer) FilterAdminRequiredNumChanged(opts *bind.FilterOpts) (*Erc20AdminRequiredNumChangedIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "AdminRequiredNumChanged")
	if err != nil {
		return nil, err
	}
	return &Erc20AdminRequiredNumChangedIterator{contract: _Erc20.contract, event: "AdminRequiredNumChanged", logs: logs, sub: sub}, nil
}

// WatchAdminRequiredNumChanged is a free log subscription operation binding the contract event 0xce49fc04234925e87b95750e0e50cac6d4ffcf8a696e3316b3e13bedc84ee7a8.
//
// Solidity: event AdminRequiredNumChanged(string TaskType, string class, uint256 previousNum, uint256 requiredNum)
func (_Erc20 *Erc20Filterer) WatchAdminRequiredNumChanged(opts *bind.WatchOpts, sink chan<- *Erc20AdminRequiredNumChanged) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "AdminRequiredNumChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20AdminRequiredNumChanged)
				if err := _Erc20.contract.UnpackLog(event, "AdminRequiredNumChanged", log); err != nil {
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

// ParseAdminRequiredNumChanged is a log parse operation binding the contract event 0xce49fc04234925e87b95750e0e50cac6d4ffcf8a696e3316b3e13bedc84ee7a8.
//
// Solidity: event AdminRequiredNumChanged(string TaskType, string class, uint256 previousNum, uint256 requiredNum)
func (_Erc20 *Erc20Filterer) ParseAdminRequiredNumChanged(log types.Log) (*Erc20AdminRequiredNumChanged, error) {
	event := new(Erc20AdminRequiredNumChanged)
	if err := _Erc20.contract.UnpackLog(event, "AdminRequiredNumChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20AdminTaskDroppedIterator is returned from FilterAdminTaskDropped and is used to iterate over the raw logs and unpacked data for AdminTaskDropped events raised by the Erc20 contract.
type Erc20AdminTaskDroppedIterator struct {
	Event *Erc20AdminTaskDropped // Event containing the contract specifics and raw log

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
func (it *Erc20AdminTaskDroppedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20AdminTaskDropped)
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
		it.Event = new(Erc20AdminTaskDropped)
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
func (it *Erc20AdminTaskDroppedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20AdminTaskDroppedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20AdminTaskDropped represents a AdminTaskDropped event raised by the Erc20 contract.
type Erc20AdminTaskDropped struct {
	TaskHash [32]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterAdminTaskDropped is a free log retrieval operation binding the contract event 0x468b360fa155a4c5fdf0cb38ce238407e41508a56023e7536088c21b2cd64139.
//
// Solidity: event AdminTaskDropped(bytes32 taskHash)
func (_Erc20 *Erc20Filterer) FilterAdminTaskDropped(opts *bind.FilterOpts) (*Erc20AdminTaskDroppedIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "AdminTaskDropped")
	if err != nil {
		return nil, err
	}
	return &Erc20AdminTaskDroppedIterator{contract: _Erc20.contract, event: "AdminTaskDropped", logs: logs, sub: sub}, nil
}

// WatchAdminTaskDropped is a free log subscription operation binding the contract event 0x468b360fa155a4c5fdf0cb38ce238407e41508a56023e7536088c21b2cd64139.
//
// Solidity: event AdminTaskDropped(bytes32 taskHash)
func (_Erc20 *Erc20Filterer) WatchAdminTaskDropped(opts *bind.WatchOpts, sink chan<- *Erc20AdminTaskDropped) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "AdminTaskDropped")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20AdminTaskDropped)
				if err := _Erc20.contract.UnpackLog(event, "AdminTaskDropped", log); err != nil {
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

// ParseAdminTaskDropped is a log parse operation binding the contract event 0x468b360fa155a4c5fdf0cb38ce238407e41508a56023e7536088c21b2cd64139.
//
// Solidity: event AdminTaskDropped(bytes32 taskHash)
func (_Erc20 *Erc20Filterer) ParseAdminTaskDropped(log types.Log) (*Erc20AdminTaskDropped, error) {
	event := new(Erc20AdminTaskDropped)
	if err := _Erc20.contract.UnpackLog(event, "AdminTaskDropped", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20ApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the Erc20 contract.
type Erc20ApprovalIterator struct {
	Event *Erc20Approval // Event containing the contract specifics and raw log

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
func (it *Erc20ApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Approval)
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
		it.Event = new(Erc20Approval)
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
func (it *Erc20ApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20ApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Approval represents a Approval event raised by the Erc20 contract.
type Erc20Approval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*Erc20ApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &Erc20ApprovalIterator{contract: _Erc20.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *Erc20Approval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Approval)
				if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
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

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed spender, uint256 value)
func (_Erc20 *Erc20Filterer) ParseApproval(log types.Log) (*Erc20Approval, error) {
	event := new(Erc20Approval)
	if err := _Erc20.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20BurnedIterator is returned from FilterBurned and is used to iterate over the raw logs and unpacked data for Burned events raised by the Erc20 contract.
type Erc20BurnedIterator struct {
	Event *Erc20Burned // Event containing the contract specifics and raw log

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
func (it *Erc20BurnedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Burned)
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
		it.Event = new(Erc20Burned)
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
func (it *Erc20BurnedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20BurnedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Burned represents a Burned event raised by the Erc20 contract.
type Erc20Burned struct {
	From       common.Address
	Value      *big.Int
	Proof      string
	BtcAddress string
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurned is a free log retrieval operation binding the contract event 0xe7fe72e51b458dcd29475a3be9675669af7aa5c3d7e9161fdb6cbba71803dd50.
//
// Solidity: event Burned(address indexed from, uint256 value, string proof, string btcAddress)
func (_Erc20 *Erc20Filterer) FilterBurned(opts *bind.FilterOpts, from []common.Address) (*Erc20BurnedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Burned", fromRule)
	if err != nil {
		return nil, err
	}
	return &Erc20BurnedIterator{contract: _Erc20.contract, event: "Burned", logs: logs, sub: sub}, nil
}

// WatchBurned is a free log subscription operation binding the contract event 0xe7fe72e51b458dcd29475a3be9675669af7aa5c3d7e9161fdb6cbba71803dd50.
//
// Solidity: event Burned(address indexed from, uint256 value, string proof, string btcAddress)
func (_Erc20 *Erc20Filterer) WatchBurned(opts *bind.WatchOpts, sink chan<- *Erc20Burned, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Burned", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Burned)
				if err := _Erc20.contract.UnpackLog(event, "Burned", log); err != nil {
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

// ParseBurned is a log parse operation binding the contract event 0xe7fe72e51b458dcd29475a3be9675669af7aa5c3d7e9161fdb6cbba71803dd50.
//
// Solidity: event Burned(address indexed from, uint256 value, string proof, string btcAddress)
func (_Erc20 *Erc20Filterer) ParseBurned(log types.Log) (*Erc20Burned, error) {
	event := new(Erc20Burned)
	if err := _Erc20.contract.UnpackLog(event, "Burned", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20BurningIterator is returned from FilterBurning and is used to iterate over the raw logs and unpacked data for Burning events raised by the Erc20 contract.
type Erc20BurningIterator struct {
	Event *Erc20Burning // Event containing the contract specifics and raw log

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
func (it *Erc20BurningIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Burning)
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
		it.Event = new(Erc20Burning)
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
func (it *Erc20BurningIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20BurningIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Burning represents a Burning event raised by the Erc20 contract.
type Erc20Burning struct {
	From       common.Address
	Value      *big.Int
	Proof      string
	BtcAddress string
	Burner     common.Address
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterBurning is a free log retrieval operation binding the contract event 0xf5577aec900779fdc4060207f95a90077d57e4ac39af49c805a073b9b1b852ea.
//
// Solidity: event Burning(address indexed from, uint256 value, string proof, string btcAddress, address burner)
func (_Erc20 *Erc20Filterer) FilterBurning(opts *bind.FilterOpts, from []common.Address) (*Erc20BurningIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Burning", fromRule)
	if err != nil {
		return nil, err
	}
	return &Erc20BurningIterator{contract: _Erc20.contract, event: "Burning", logs: logs, sub: sub}, nil
}

// WatchBurning is a free log subscription operation binding the contract event 0xf5577aec900779fdc4060207f95a90077d57e4ac39af49c805a073b9b1b852ea.
//
// Solidity: event Burning(address indexed from, uint256 value, string proof, string btcAddress, address burner)
func (_Erc20 *Erc20Filterer) WatchBurning(opts *bind.WatchOpts, sink chan<- *Erc20Burning, from []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Burning", fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Burning)
				if err := _Erc20.contract.UnpackLog(event, "Burning", log); err != nil {
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

// ParseBurning is a log parse operation binding the contract event 0xf5577aec900779fdc4060207f95a90077d57e4ac39af49c805a073b9b1b852ea.
//
// Solidity: event Burning(address indexed from, uint256 value, string proof, string btcAddress, address burner)
func (_Erc20 *Erc20Filterer) ParseBurning(log types.Log) (*Erc20Burning, error) {
	event := new(Erc20Burning)
	if err := _Erc20.contract.UnpackLog(event, "Burning", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20MintedIterator is returned from FilterMinted and is used to iterate over the raw logs and unpacked data for Minted events raised by the Erc20 contract.
type Erc20MintedIterator struct {
	Event *Erc20Minted // Event containing the contract specifics and raw log

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
func (it *Erc20MintedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Minted)
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
		it.Event = new(Erc20Minted)
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
func (it *Erc20MintedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20MintedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Minted represents a Minted event raised by the Erc20 contract.
type Erc20Minted struct {
	To    common.Address
	Value *big.Int
	Proof string
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterMinted is a free log retrieval operation binding the contract event 0xe7cd4ce7f2a465edc730269a1305e8a48bad821e8fb7e152ec413829c01a53c4.
//
// Solidity: event Minted(address indexed to, uint256 value, string proof)
func (_Erc20 *Erc20Filterer) FilterMinted(opts *bind.FilterOpts, to []common.Address) (*Erc20MintedIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20MintedIterator{contract: _Erc20.contract, event: "Minted", logs: logs, sub: sub}, nil
}

// WatchMinted is a free log subscription operation binding the contract event 0xe7cd4ce7f2a465edc730269a1305e8a48bad821e8fb7e152ec413829c01a53c4.
//
// Solidity: event Minted(address indexed to, uint256 value, string proof)
func (_Erc20 *Erc20Filterer) WatchMinted(opts *bind.WatchOpts, sink chan<- *Erc20Minted, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Minted", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Minted)
				if err := _Erc20.contract.UnpackLog(event, "Minted", log); err != nil {
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

// ParseMinted is a log parse operation binding the contract event 0xe7cd4ce7f2a465edc730269a1305e8a48bad821e8fb7e152ec413829c01a53c4.
//
// Solidity: event Minted(address indexed to, uint256 value, string proof)
func (_Erc20 *Erc20Filterer) ParseMinted(log types.Log) (*Erc20Minted, error) {
	event := new(Erc20Minted)
	if err := _Erc20.contract.UnpackLog(event, "Minted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20MintingIterator is returned from FilterMinting and is used to iterate over the raw logs and unpacked data for Minting events raised by the Erc20 contract.
type Erc20MintingIterator struct {
	Event *Erc20Minting // Event containing the contract specifics and raw log

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
func (it *Erc20MintingIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Minting)
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
		it.Event = new(Erc20Minting)
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
func (it *Erc20MintingIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20MintingIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Minting represents a Minting event raised by the Erc20 contract.
type Erc20Minting struct {
	To     common.Address
	Value  *big.Int
	Proof  string
	Minter common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMinting is a free log retrieval operation binding the contract event 0x0416081e24ea042d7f2160192c1d6bef74c76f5f4fffeb36c5dcc9e4a5dcf9b4.
//
// Solidity: event Minting(address indexed to, uint256 value, string proof, address minter)
func (_Erc20 *Erc20Filterer) FilterMinting(opts *bind.FilterOpts, to []common.Address) (*Erc20MintingIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Minting", toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20MintingIterator{contract: _Erc20.contract, event: "Minting", logs: logs, sub: sub}, nil
}

// WatchMinting is a free log subscription operation binding the contract event 0x0416081e24ea042d7f2160192c1d6bef74c76f5f4fffeb36c5dcc9e4a5dcf9b4.
//
// Solidity: event Minting(address indexed to, uint256 value, string proof, address minter)
func (_Erc20 *Erc20Filterer) WatchMinting(opts *bind.WatchOpts, sink chan<- *Erc20Minting, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Minting", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Minting)
				if err := _Erc20.contract.UnpackLog(event, "Minting", log); err != nil {
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

// ParseMinting is a log parse operation binding the contract event 0x0416081e24ea042d7f2160192c1d6bef74c76f5f4fffeb36c5dcc9e4a5dcf9b4.
//
// Solidity: event Minting(address indexed to, uint256 value, string proof, address minter)
func (_Erc20 *Erc20Filterer) ParseMinting(log types.Log) (*Erc20Minting, error) {
	event := new(Erc20Minting)
	if err := _Erc20.contract.UnpackLog(event, "Minting", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20PauseChangedToIterator is returned from FilterPauseChangedTo and is used to iterate over the raw logs and unpacked data for PauseChangedTo events raised by the Erc20 contract.
type Erc20PauseChangedToIterator struct {
	Event *Erc20PauseChangedTo // Event containing the contract specifics and raw log

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
func (it *Erc20PauseChangedToIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20PauseChangedTo)
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
		it.Event = new(Erc20PauseChangedTo)
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
func (it *Erc20PauseChangedToIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20PauseChangedToIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20PauseChangedTo represents a PauseChangedTo event raised by the Erc20 contract.
type Erc20PauseChangedTo struct {
	PauseState bool
	Raw        types.Log // Blockchain specific contextual infos
}

// FilterPauseChangedTo is a free log retrieval operation binding the contract event 0x0619930b74b56b9cdbdbd8709cef274ec3e9ed0616d7e996599558acccac8e6e.
//
// Solidity: event PauseChangedTo(bool pauseState)
func (_Erc20 *Erc20Filterer) FilterPauseChangedTo(opts *bind.FilterOpts) (*Erc20PauseChangedToIterator, error) {

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "PauseChangedTo")
	if err != nil {
		return nil, err
	}
	return &Erc20PauseChangedToIterator{contract: _Erc20.contract, event: "PauseChangedTo", logs: logs, sub: sub}, nil
}

// WatchPauseChangedTo is a free log subscription operation binding the contract event 0x0619930b74b56b9cdbdbd8709cef274ec3e9ed0616d7e996599558acccac8e6e.
//
// Solidity: event PauseChangedTo(bool pauseState)
func (_Erc20 *Erc20Filterer) WatchPauseChangedTo(opts *bind.WatchOpts, sink chan<- *Erc20PauseChangedTo) (event.Subscription, error) {

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "PauseChangedTo")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20PauseChangedTo)
				if err := _Erc20.contract.UnpackLog(event, "PauseChangedTo", log); err != nil {
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

// ParsePauseChangedTo is a log parse operation binding the contract event 0x0619930b74b56b9cdbdbd8709cef274ec3e9ed0616d7e996599558acccac8e6e.
//
// Solidity: event PauseChangedTo(bool pauseState)
func (_Erc20 *Erc20Filterer) ParsePauseChangedTo(log types.Log) (*Erc20PauseChangedTo, error) {
	event := new(Erc20PauseChangedTo)
	if err := _Erc20.contract.UnpackLog(event, "PauseChangedTo", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// Erc20TransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the Erc20 contract.
type Erc20TransferIterator struct {
	Event *Erc20Transfer // Event containing the contract specifics and raw log

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
func (it *Erc20TransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(Erc20Transfer)
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
		it.Event = new(Erc20Transfer)
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
func (it *Erc20TransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *Erc20TransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// Erc20Transfer represents a Transfer event raised by the Erc20 contract.
type Erc20Transfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*Erc20TransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &Erc20TransferIterator{contract: _Erc20.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *Erc20Transfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _Erc20.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(Erc20Transfer)
				if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 value)
func (_Erc20 *Erc20Filterer) ParseTransfer(log types.Log) (*Erc20Transfer, error) {
	event := new(Erc20Transfer)
	if err := _Erc20.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
