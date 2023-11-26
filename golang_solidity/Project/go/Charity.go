// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package GO

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// CharityFunder is an auto generated low-level Go binding around an user-defined struct.
type CharityFunder struct {
	Addr   common.Address
	Amount *big.Int
}

// CharityFunding is an auto generated low-level Go binding around an user-defined struct.
type CharityFunding struct {
	Initiator  common.Address
	Title      string
	Info       string
	Goal       *big.Int
	Exist      bool
	Id         *big.Int
	Success    bool
	Amount     *big.Int
	NumFunders *big.Int
	NumUses    *big.Int
}

// CalldemoABI is the input ABI used to generate the binding from.
const CalldemoABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"GetAllfundings\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numFunders\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numUses\",\"type\":\"uint256\"}],\"internalType\":\"structCharity.Funding[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"GetFunder\",\"outputs\":[{\"components\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"internalType\":\"structCharity.Funder\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"InsCollectFunding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allfundings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numFunders\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numUses\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"delOneCollectFunding\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"deleteFunding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"funders\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"fundings\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numFunders\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numUses\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"key\",\"type\":\"uint256\"}],\"name\":\"getFunding\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numFunders\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numUses\",\"type\":\"uint256\"}],\"internalType\":\"structCharity.Funding\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMsgBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"getMyFundings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTime\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getallCollectFunding\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numFunders\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numUses\",\"type\":\"uint256\"}],\"internalType\":\"structCharity.Funding[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"addr\",\"type\":\"address\"}],\"name\":\"getallContribute\",\"outputs\":[{\"components\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numFunders\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numUses\",\"type\":\"uint256\"}],\"internalType\":\"structCharity.Funding[]\",\"name\":\"\",\"type\":\"tuple[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"host\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"msgFunder\",\"outputs\":[{\"internalType\":\"addresspayable\",\"name\":\"addr\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"}],\"name\":\"newFunding\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numCollect\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numContributeFundings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"numFundings\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ID\",\"type\":\"uint256\"}],\"name\":\"returnMoney\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tfCollectFunding\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"initiator\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"title\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"info\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"goal\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"exist\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"id\",\"type\":\"uint256\"},{\"internalType\":\"bool\",\"name\":\"success\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"amount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numFunders\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"numUses\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"tfContribute\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// Calldemo is an auto generated Go binding around an Ethereum contract.
type Calldemo struct {
	CalldemoCaller     // Read-only binding to the contract
	CalldemoTransactor // Write-only binding to the contract
	CalldemoFilterer   // Log filterer for contract events
}

// CalldemoCaller is an auto generated read-only Go binding around an Ethereum contract.
type CalldemoCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalldemoTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CalldemoTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalldemoFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CalldemoFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CalldemoSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CalldemoSession struct {
	Contract     *Calldemo         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CalldemoCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CalldemoCallerSession struct {
	Contract *CalldemoCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CalldemoTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CalldemoTransactorSession struct {
	Contract     *CalldemoTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CalldemoRaw is an auto generated low-level Go binding around an Ethereum contract.
type CalldemoRaw struct {
	Contract *Calldemo // Generic contract binding to access the raw methods on
}

// CalldemoCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CalldemoCallerRaw struct {
	Contract *CalldemoCaller // Generic read-only contract binding to access the raw methods on
}

// CalldemoTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CalldemoTransactorRaw struct {
	Contract *CalldemoTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCalldemo creates a new instance of Calldemo, bound to a specific deployed contract.
func NewCalldemo(address common.Address, backend bind.ContractBackend) (*Calldemo, error) {
	contract, err := bindCalldemo(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Calldemo{CalldemoCaller: CalldemoCaller{contract: contract}, CalldemoTransactor: CalldemoTransactor{contract: contract}, CalldemoFilterer: CalldemoFilterer{contract: contract}}, nil
}

// NewCalldemoCaller creates a new read-only instance of Calldemo, bound to a specific deployed contract.
func NewCalldemoCaller(address common.Address, caller bind.ContractCaller) (*CalldemoCaller, error) {
	contract, err := bindCalldemo(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CalldemoCaller{contract: contract}, nil
}

// NewCalldemoTransactor creates a new write-only instance of Calldemo, bound to a specific deployed contract.
func NewCalldemoTransactor(address common.Address, transactor bind.ContractTransactor) (*CalldemoTransactor, error) {
	contract, err := bindCalldemo(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CalldemoTransactor{contract: contract}, nil
}

// NewCalldemoFilterer creates a new log filterer instance of Calldemo, bound to a specific deployed contract.
func NewCalldemoFilterer(address common.Address, filterer bind.ContractFilterer) (*CalldemoFilterer, error) {
	contract, err := bindCalldemo(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CalldemoFilterer{contract: contract}, nil
}

// bindCalldemo binds a generic wrapper to an already deployed contract.
func bindCalldemo(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CalldemoABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calldemo *CalldemoRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Calldemo.Contract.CalldemoCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calldemo *CalldemoRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calldemo.Contract.CalldemoTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calldemo *CalldemoRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Calldemo.Contract.CalldemoTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Calldemo *CalldemoCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Calldemo.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Calldemo *CalldemoTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Calldemo.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Calldemo *CalldemoTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Calldemo.Contract.contract.Transact(opts, method, params...)
}

// GetAllfundings is a free data retrieval call binding the contract method 0xbf5bd742.
//
// Solidity: function GetAllfundings() view returns((address,string,string,uint256,bool,uint256,bool,uint256,uint256,uint256)[])
func (_Calldemo *CalldemoCaller) GetAllfundings(opts *bind.CallOpts) ([]CharityFunding, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "GetAllfundings")

	if err != nil {
		return *new([]CharityFunding), err
	}

	out0 := *abi.ConvertType(out[0], new([]CharityFunding)).(*[]CharityFunding)

	return out0, err

}

// GetAllfundings is a free data retrieval call binding the contract method 0xbf5bd742.
//
// Solidity: function GetAllfundings() view returns((address,string,string,uint256,bool,uint256,bool,uint256,uint256,uint256)[])
func (_Calldemo *CalldemoSession) GetAllfundings() ([]CharityFunding, error) {
	return _Calldemo.Contract.GetAllfundings(&_Calldemo.CallOpts)
}

// GetAllfundings is a free data retrieval call binding the contract method 0xbf5bd742.
//
// Solidity: function GetAllfundings() view returns((address,string,string,uint256,bool,uint256,bool,uint256,uint256,uint256)[])
func (_Calldemo *CalldemoCallerSession) GetAllfundings() ([]CharityFunding, error) {
	return _Calldemo.Contract.GetAllfundings(&_Calldemo.CallOpts)
}

// GetFunder is a free data retrieval call binding the contract method 0x5358e660.
//
// Solidity: function GetFunder(uint256 key) view returns((address,uint256))
func (_Calldemo *CalldemoCaller) GetFunder(opts *bind.CallOpts, key *big.Int) (CharityFunder, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "GetFunder", key)

	if err != nil {
		return *new(CharityFunder), err
	}

	out0 := *abi.ConvertType(out[0], new(CharityFunder)).(*CharityFunder)

	return out0, err

}

// GetFunder is a free data retrieval call binding the contract method 0x5358e660.
//
// Solidity: function GetFunder(uint256 key) view returns((address,uint256))
func (_Calldemo *CalldemoSession) GetFunder(key *big.Int) (CharityFunder, error) {
	return _Calldemo.Contract.GetFunder(&_Calldemo.CallOpts, key)
}

// GetFunder is a free data retrieval call binding the contract method 0x5358e660.
//
// Solidity: function GetFunder(uint256 key) view returns((address,uint256))
func (_Calldemo *CalldemoCallerSession) GetFunder(key *big.Int) (CharityFunder, error) {
	return _Calldemo.Contract.GetFunder(&_Calldemo.CallOpts, key)
}

// Allfundings is a free data retrieval call binding the contract method 0x67dfc8f2.
//
// Solidity: function allfundings(uint256 , uint256 ) view returns(address initiator, string title, string info, uint256 goal, bool exist, uint256 id, bool success, uint256 amount, uint256 numFunders, uint256 numUses)
func (_Calldemo *CalldemoCaller) Allfundings(opts *bind.CallOpts, arg0 *big.Int, arg1 *big.Int) (struct {
	Initiator  common.Address
	Title      string
	Info       string
	Goal       *big.Int
	Exist      bool
	Id         *big.Int
	Success    bool
	Amount     *big.Int
	NumFunders *big.Int
	NumUses    *big.Int
}, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "allfundings", arg0, arg1)

	outstruct := new(struct {
		Initiator  common.Address
		Title      string
		Info       string
		Goal       *big.Int
		Exist      bool
		Id         *big.Int
		Success    bool
		Amount     *big.Int
		NumFunders *big.Int
		NumUses    *big.Int
	})

	outstruct.Initiator = out[0].(common.Address)
	outstruct.Title = out[1].(string)
	outstruct.Info = out[2].(string)
	outstruct.Goal = out[3].(*big.Int)
	outstruct.Exist = out[4].(bool)
	outstruct.Id = out[5].(*big.Int)
	outstruct.Success = out[6].(bool)
	outstruct.Amount = out[7].(*big.Int)
	outstruct.NumFunders = out[8].(*big.Int)
	outstruct.NumUses = out[9].(*big.Int)

	return *outstruct, err

}

// Allfundings is a free data retrieval call binding the contract method 0x67dfc8f2.
//
// Solidity: function allfundings(uint256 , uint256 ) view returns(address initiator, string title, string info, uint256 goal, bool exist, uint256 id, bool success, uint256 amount, uint256 numFunders, uint256 numUses)
func (_Calldemo *CalldemoSession) Allfundings(arg0 *big.Int, arg1 *big.Int) (struct {
	Initiator  common.Address
	Title      string
	Info       string
	Goal       *big.Int
	Exist      bool
	Id         *big.Int
	Success    bool
	Amount     *big.Int
	NumFunders *big.Int
	NumUses    *big.Int
}, error) {
	return _Calldemo.Contract.Allfundings(&_Calldemo.CallOpts, arg0, arg1)
}

// Allfundings is a free data retrieval call binding the contract method 0x67dfc8f2.
//
// Solidity: function allfundings(uint256 , uint256 ) view returns(address initiator, string title, string info, uint256 goal, bool exist, uint256 id, bool success, uint256 amount, uint256 numFunders, uint256 numUses)
func (_Calldemo *CalldemoCallerSession) Allfundings(arg0 *big.Int, arg1 *big.Int) (struct {
	Initiator  common.Address
	Title      string
	Info       string
	Goal       *big.Int
	Exist      bool
	Id         *big.Int
	Success    bool
	Amount     *big.Int
	NumFunders *big.Int
	NumUses    *big.Int
}, error) {
	return _Calldemo.Contract.Allfundings(&_Calldemo.CallOpts, arg0, arg1)
}

// Funders is a free data retrieval call binding the contract method 0xdc0d3dff.
//
// Solidity: function funders(uint256 ) view returns(address addr, uint256 amount)
func (_Calldemo *CalldemoCaller) Funders(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Addr   common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "funders", arg0)

	outstruct := new(struct {
		Addr   common.Address
		Amount *big.Int
	})

	outstruct.Addr = out[0].(common.Address)
	outstruct.Amount = out[1].(*big.Int)

	return *outstruct, err

}

// Funders is a free data retrieval call binding the contract method 0xdc0d3dff.
//
// Solidity: function funders(uint256 ) view returns(address addr, uint256 amount)
func (_Calldemo *CalldemoSession) Funders(arg0 *big.Int) (struct {
	Addr   common.Address
	Amount *big.Int
}, error) {
	return _Calldemo.Contract.Funders(&_Calldemo.CallOpts, arg0)
}

// Funders is a free data retrieval call binding the contract method 0xdc0d3dff.
//
// Solidity: function funders(uint256 ) view returns(address addr, uint256 amount)
func (_Calldemo *CalldemoCallerSession) Funders(arg0 *big.Int) (struct {
	Addr   common.Address
	Amount *big.Int
}, error) {
	return _Calldemo.Contract.Funders(&_Calldemo.CallOpts, arg0)
}

// Fundings is a free data retrieval call binding the contract method 0xd13cb1fb.
//
// Solidity: function fundings(uint256 ) view returns(address initiator, string title, string info, uint256 goal, bool exist, uint256 id, bool success, uint256 amount, uint256 numFunders, uint256 numUses)
func (_Calldemo *CalldemoCaller) Fundings(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Initiator  common.Address
	Title      string
	Info       string
	Goal       *big.Int
	Exist      bool
	Id         *big.Int
	Success    bool
	Amount     *big.Int
	NumFunders *big.Int
	NumUses    *big.Int
}, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "fundings", arg0)

	outstruct := new(struct {
		Initiator  common.Address
		Title      string
		Info       string
		Goal       *big.Int
		Exist      bool
		Id         *big.Int
		Success    bool
		Amount     *big.Int
		NumFunders *big.Int
		NumUses    *big.Int
	})

	outstruct.Initiator = out[0].(common.Address)
	outstruct.Title = out[1].(string)
	outstruct.Info = out[2].(string)
	outstruct.Goal = out[3].(*big.Int)
	outstruct.Exist = out[4].(bool)
	outstruct.Id = out[5].(*big.Int)
	outstruct.Success = out[6].(bool)
	outstruct.Amount = out[7].(*big.Int)
	outstruct.NumFunders = out[8].(*big.Int)
	outstruct.NumUses = out[9].(*big.Int)

	return *outstruct, err

}

// Fundings is a free data retrieval call binding the contract method 0xd13cb1fb.
//
// Solidity: function fundings(uint256 ) view returns(address initiator, string title, string info, uint256 goal, bool exist, uint256 id, bool success, uint256 amount, uint256 numFunders, uint256 numUses)
func (_Calldemo *CalldemoSession) Fundings(arg0 *big.Int) (struct {
	Initiator  common.Address
	Title      string
	Info       string
	Goal       *big.Int
	Exist      bool
	Id         *big.Int
	Success    bool
	Amount     *big.Int
	NumFunders *big.Int
	NumUses    *big.Int
}, error) {
	return _Calldemo.Contract.Fundings(&_Calldemo.CallOpts, arg0)
}

// Fundings is a free data retrieval call binding the contract method 0xd13cb1fb.
//
// Solidity: function fundings(uint256 ) view returns(address initiator, string title, string info, uint256 goal, bool exist, uint256 id, bool success, uint256 amount, uint256 numFunders, uint256 numUses)
func (_Calldemo *CalldemoCallerSession) Fundings(arg0 *big.Int) (struct {
	Initiator  common.Address
	Title      string
	Info       string
	Goal       *big.Int
	Exist      bool
	Id         *big.Int
	Success    bool
	Amount     *big.Int
	NumFunders *big.Int
	NumUses    *big.Int
}, error) {
	return _Calldemo.Contract.Fundings(&_Calldemo.CallOpts, arg0)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Calldemo *CalldemoCaller) GetBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "getBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Calldemo *CalldemoSession) GetBalance() (*big.Int, error) {
	return _Calldemo.Contract.GetBalance(&_Calldemo.CallOpts)
}

// GetBalance is a free data retrieval call binding the contract method 0x12065fe0.
//
// Solidity: function getBalance() view returns(uint256)
func (_Calldemo *CalldemoCallerSession) GetBalance() (*big.Int, error) {
	return _Calldemo.Contract.GetBalance(&_Calldemo.CallOpts)
}

// GetFunding is a free data retrieval call binding the contract method 0xebed4bd4.
//
// Solidity: function getFunding(uint256 key) view returns((address,string,string,uint256,bool,uint256,bool,uint256,uint256,uint256))
func (_Calldemo *CalldemoCaller) GetFunding(opts *bind.CallOpts, key *big.Int) (CharityFunding, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "getFunding", key)

	if err != nil {
		return *new(CharityFunding), err
	}

	out0 := *abi.ConvertType(out[0], new(CharityFunding)).(*CharityFunding)

	return out0, err

}

// GetFunding is a free data retrieval call binding the contract method 0xebed4bd4.
//
// Solidity: function getFunding(uint256 key) view returns((address,string,string,uint256,bool,uint256,bool,uint256,uint256,uint256))
func (_Calldemo *CalldemoSession) GetFunding(key *big.Int) (CharityFunding, error) {
	return _Calldemo.Contract.GetFunding(&_Calldemo.CallOpts, key)
}

// GetFunding is a free data retrieval call binding the contract method 0xebed4bd4.
//
// Solidity: function getFunding(uint256 key) view returns((address,string,string,uint256,bool,uint256,bool,uint256,uint256,uint256))
func (_Calldemo *CalldemoCallerSession) GetFunding(key *big.Int) (CharityFunding, error) {
	return _Calldemo.Contract.GetFunding(&_Calldemo.CallOpts, key)
}

// GetMsgBalance is a free data retrieval call binding the contract method 0xa2ef5410.
//
// Solidity: function getMsgBalance() view returns(uint256)
func (_Calldemo *CalldemoCaller) GetMsgBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "getMsgBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMsgBalance is a free data retrieval call binding the contract method 0xa2ef5410.
//
// Solidity: function getMsgBalance() view returns(uint256)
func (_Calldemo *CalldemoSession) GetMsgBalance() (*big.Int, error) {
	return _Calldemo.Contract.GetMsgBalance(&_Calldemo.CallOpts)
}

// GetMsgBalance is a free data retrieval call binding the contract method 0xa2ef5410.
//
// Solidity: function getMsgBalance() view returns(uint256)
func (_Calldemo *CalldemoCallerSession) GetMsgBalance() (*big.Int, error) {
	return _Calldemo.Contract.GetMsgBalance(&_Calldemo.CallOpts)
}

// GetMyFundings is a free data retrieval call binding the contract method 0x716f7147.
//
// Solidity: function getMyFundings(address addr, uint256 ID) view returns(uint256)
func (_Calldemo *CalldemoCaller) GetMyFundings(opts *bind.CallOpts, addr common.Address, ID *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "getMyFundings", addr, ID)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMyFundings is a free data retrieval call binding the contract method 0x716f7147.
//
// Solidity: function getMyFundings(address addr, uint256 ID) view returns(uint256)
func (_Calldemo *CalldemoSession) GetMyFundings(addr common.Address, ID *big.Int) (*big.Int, error) {
	return _Calldemo.Contract.GetMyFundings(&_Calldemo.CallOpts, addr, ID)
}

// GetMyFundings is a free data retrieval call binding the contract method 0x716f7147.
//
// Solidity: function getMyFundings(address addr, uint256 ID) view returns(uint256)
func (_Calldemo *CalldemoCallerSession) GetMyFundings(addr common.Address, ID *big.Int) (*big.Int, error) {
	return _Calldemo.Contract.GetMyFundings(&_Calldemo.CallOpts, addr, ID)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint256)
func (_Calldemo *CalldemoCaller) GetTime(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "getTime")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint256)
func (_Calldemo *CalldemoSession) GetTime() (*big.Int, error) {
	return _Calldemo.Contract.GetTime(&_Calldemo.CallOpts)
}

// GetTime is a free data retrieval call binding the contract method 0x557ed1ba.
//
// Solidity: function getTime() view returns(uint256)
func (_Calldemo *CalldemoCallerSession) GetTime() (*big.Int, error) {
	return _Calldemo.Contract.GetTime(&_Calldemo.CallOpts)
}

// GetallCollectFunding is a free data retrieval call binding the contract method 0xf6e8c234.
//
// Solidity: function getallCollectFunding(address addr) view returns((address,string,string,uint256,bool,uint256,bool,uint256,uint256,uint256)[])
func (_Calldemo *CalldemoCaller) GetallCollectFunding(opts *bind.CallOpts, addr common.Address) ([]CharityFunding, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "getallCollectFunding", addr)

	if err != nil {
		return *new([]CharityFunding), err
	}

	out0 := *abi.ConvertType(out[0], new([]CharityFunding)).(*[]CharityFunding)

	return out0, err

}

// GetallCollectFunding is a free data retrieval call binding the contract method 0xf6e8c234.
//
// Solidity: function getallCollectFunding(address addr) view returns((address,string,string,uint256,bool,uint256,bool,uint256,uint256,uint256)[])
func (_Calldemo *CalldemoSession) GetallCollectFunding(addr common.Address) ([]CharityFunding, error) {
	return _Calldemo.Contract.GetallCollectFunding(&_Calldemo.CallOpts, addr)
}

// GetallCollectFunding is a free data retrieval call binding the contract method 0xf6e8c234.
//
// Solidity: function getallCollectFunding(address addr) view returns((address,string,string,uint256,bool,uint256,bool,uint256,uint256,uint256)[])
func (_Calldemo *CalldemoCallerSession) GetallCollectFunding(addr common.Address) ([]CharityFunding, error) {
	return _Calldemo.Contract.GetallCollectFunding(&_Calldemo.CallOpts, addr)
}

// GetallContribute is a free data retrieval call binding the contract method 0x566d4516.
//
// Solidity: function getallContribute(address addr) view returns((address,string,string,uint256,bool,uint256,bool,uint256,uint256,uint256)[])
func (_Calldemo *CalldemoCaller) GetallContribute(opts *bind.CallOpts, addr common.Address) ([]CharityFunding, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "getallContribute", addr)

	if err != nil {
		return *new([]CharityFunding), err
	}

	out0 := *abi.ConvertType(out[0], new([]CharityFunding)).(*[]CharityFunding)

	return out0, err

}

// GetallContribute is a free data retrieval call binding the contract method 0x566d4516.
//
// Solidity: function getallContribute(address addr) view returns((address,string,string,uint256,bool,uint256,bool,uint256,uint256,uint256)[])
func (_Calldemo *CalldemoSession) GetallContribute(addr common.Address) ([]CharityFunding, error) {
	return _Calldemo.Contract.GetallContribute(&_Calldemo.CallOpts, addr)
}

// GetallContribute is a free data retrieval call binding the contract method 0x566d4516.
//
// Solidity: function getallContribute(address addr) view returns((address,string,string,uint256,bool,uint256,bool,uint256,uint256,uint256)[])
func (_Calldemo *CalldemoCallerSession) GetallContribute(addr common.Address) ([]CharityFunding, error) {
	return _Calldemo.Contract.GetallContribute(&_Calldemo.CallOpts, addr)
}

// Host is a free data retrieval call binding the contract method 0xf437bc59.
//
// Solidity: function host() view returns(address)
func (_Calldemo *CalldemoCaller) Host(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "host")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Host is a free data retrieval call binding the contract method 0xf437bc59.
//
// Solidity: function host() view returns(address)
func (_Calldemo *CalldemoSession) Host() (common.Address, error) {
	return _Calldemo.Contract.Host(&_Calldemo.CallOpts)
}

// Host is a free data retrieval call binding the contract method 0xf437bc59.
//
// Solidity: function host() view returns(address)
func (_Calldemo *CalldemoCallerSession) Host() (common.Address, error) {
	return _Calldemo.Contract.Host(&_Calldemo.CallOpts)
}

// MsgFunder is a free data retrieval call binding the contract method 0xbf8405c2.
//
// Solidity: function msgFunder(address , uint256 ) view returns(address addr, uint256 amount)
func (_Calldemo *CalldemoCaller) MsgFunder(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Addr   common.Address
	Amount *big.Int
}, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "msgFunder", arg0, arg1)

	outstruct := new(struct {
		Addr   common.Address
		Amount *big.Int
	})

	outstruct.Addr = out[0].(common.Address)
	outstruct.Amount = out[1].(*big.Int)

	return *outstruct, err

}

// MsgFunder is a free data retrieval call binding the contract method 0xbf8405c2.
//
// Solidity: function msgFunder(address , uint256 ) view returns(address addr, uint256 amount)
func (_Calldemo *CalldemoSession) MsgFunder(arg0 common.Address, arg1 *big.Int) (struct {
	Addr   common.Address
	Amount *big.Int
}, error) {
	return _Calldemo.Contract.MsgFunder(&_Calldemo.CallOpts, arg0, arg1)
}

// MsgFunder is a free data retrieval call binding the contract method 0xbf8405c2.
//
// Solidity: function msgFunder(address , uint256 ) view returns(address addr, uint256 amount)
func (_Calldemo *CalldemoCallerSession) MsgFunder(arg0 common.Address, arg1 *big.Int) (struct {
	Addr   common.Address
	Amount *big.Int
}, error) {
	return _Calldemo.Contract.MsgFunder(&_Calldemo.CallOpts, arg0, arg1)
}

// NumCollect is a free data retrieval call binding the contract method 0x8f838cd1.
//
// Solidity: function numCollect() view returns(uint256)
func (_Calldemo *CalldemoCaller) NumCollect(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "numCollect")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumCollect is a free data retrieval call binding the contract method 0x8f838cd1.
//
// Solidity: function numCollect() view returns(uint256)
func (_Calldemo *CalldemoSession) NumCollect() (*big.Int, error) {
	return _Calldemo.Contract.NumCollect(&_Calldemo.CallOpts)
}

// NumCollect is a free data retrieval call binding the contract method 0x8f838cd1.
//
// Solidity: function numCollect() view returns(uint256)
func (_Calldemo *CalldemoCallerSession) NumCollect() (*big.Int, error) {
	return _Calldemo.Contract.NumCollect(&_Calldemo.CallOpts)
}

// NumContributeFundings is a free data retrieval call binding the contract method 0xdd6a29bd.
//
// Solidity: function numContributeFundings() view returns(uint256)
func (_Calldemo *CalldemoCaller) NumContributeFundings(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "numContributeFundings")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumContributeFundings is a free data retrieval call binding the contract method 0xdd6a29bd.
//
// Solidity: function numContributeFundings() view returns(uint256)
func (_Calldemo *CalldemoSession) NumContributeFundings() (*big.Int, error) {
	return _Calldemo.Contract.NumContributeFundings(&_Calldemo.CallOpts)
}

// NumContributeFundings is a free data retrieval call binding the contract method 0xdd6a29bd.
//
// Solidity: function numContributeFundings() view returns(uint256)
func (_Calldemo *CalldemoCallerSession) NumContributeFundings() (*big.Int, error) {
	return _Calldemo.Contract.NumContributeFundings(&_Calldemo.CallOpts)
}

// NumFundings is a free data retrieval call binding the contract method 0xc70a2d79.
//
// Solidity: function numFundings() view returns(uint256)
func (_Calldemo *CalldemoCaller) NumFundings(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "numFundings")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// NumFundings is a free data retrieval call binding the contract method 0xc70a2d79.
//
// Solidity: function numFundings() view returns(uint256)
func (_Calldemo *CalldemoSession) NumFundings() (*big.Int, error) {
	return _Calldemo.Contract.NumFundings(&_Calldemo.CallOpts)
}

// NumFundings is a free data retrieval call binding the contract method 0xc70a2d79.
//
// Solidity: function numFundings() view returns(uint256)
func (_Calldemo *CalldemoCallerSession) NumFundings() (*big.Int, error) {
	return _Calldemo.Contract.NumFundings(&_Calldemo.CallOpts)
}

// TfCollectFunding is a free data retrieval call binding the contract method 0x6128666b.
//
// Solidity: function tfCollectFunding(address , uint256 ) view returns(address initiator, string title, string info, uint256 goal, bool exist, uint256 id, bool success, uint256 amount, uint256 numFunders, uint256 numUses)
func (_Calldemo *CalldemoCaller) TfCollectFunding(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (struct {
	Initiator  common.Address
	Title      string
	Info       string
	Goal       *big.Int
	Exist      bool
	Id         *big.Int
	Success    bool
	Amount     *big.Int
	NumFunders *big.Int
	NumUses    *big.Int
}, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "tfCollectFunding", arg0, arg1)

	outstruct := new(struct {
		Initiator  common.Address
		Title      string
		Info       string
		Goal       *big.Int
		Exist      bool
		Id         *big.Int
		Success    bool
		Amount     *big.Int
		NumFunders *big.Int
		NumUses    *big.Int
	})

	outstruct.Initiator = out[0].(common.Address)
	outstruct.Title = out[1].(string)
	outstruct.Info = out[2].(string)
	outstruct.Goal = out[3].(*big.Int)
	outstruct.Exist = out[4].(bool)
	outstruct.Id = out[5].(*big.Int)
	outstruct.Success = out[6].(bool)
	outstruct.Amount = out[7].(*big.Int)
	outstruct.NumFunders = out[8].(*big.Int)
	outstruct.NumUses = out[9].(*big.Int)

	return *outstruct, err

}

// TfCollectFunding is a free data retrieval call binding the contract method 0x6128666b.
//
// Solidity: function tfCollectFunding(address , uint256 ) view returns(address initiator, string title, string info, uint256 goal, bool exist, uint256 id, bool success, uint256 amount, uint256 numFunders, uint256 numUses)
func (_Calldemo *CalldemoSession) TfCollectFunding(arg0 common.Address, arg1 *big.Int) (struct {
	Initiator  common.Address
	Title      string
	Info       string
	Goal       *big.Int
	Exist      bool
	Id         *big.Int
	Success    bool
	Amount     *big.Int
	NumFunders *big.Int
	NumUses    *big.Int
}, error) {
	return _Calldemo.Contract.TfCollectFunding(&_Calldemo.CallOpts, arg0, arg1)
}

// TfCollectFunding is a free data retrieval call binding the contract method 0x6128666b.
//
// Solidity: function tfCollectFunding(address , uint256 ) view returns(address initiator, string title, string info, uint256 goal, bool exist, uint256 id, bool success, uint256 amount, uint256 numFunders, uint256 numUses)
func (_Calldemo *CalldemoCallerSession) TfCollectFunding(arg0 common.Address, arg1 *big.Int) (struct {
	Initiator  common.Address
	Title      string
	Info       string
	Goal       *big.Int
	Exist      bool
	Id         *big.Int
	Success    bool
	Amount     *big.Int
	NumFunders *big.Int
	NumUses    *big.Int
}, error) {
	return _Calldemo.Contract.TfCollectFunding(&_Calldemo.CallOpts, arg0, arg1)
}

// TfContribute is a free data retrieval call binding the contract method 0x3461c378.
//
// Solidity: function tfContribute(address , uint256 ) view returns(bool)
func (_Calldemo *CalldemoCaller) TfContribute(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (bool, error) {
	var out []interface{}
	err := _Calldemo.contract.Call(opts, &out, "tfContribute", arg0, arg1)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// TfContribute is a free data retrieval call binding the contract method 0x3461c378.
//
// Solidity: function tfContribute(address , uint256 ) view returns(bool)
func (_Calldemo *CalldemoSession) TfContribute(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Calldemo.Contract.TfContribute(&_Calldemo.CallOpts, arg0, arg1)
}

// TfContribute is a free data retrieval call binding the contract method 0x3461c378.
//
// Solidity: function tfContribute(address , uint256 ) view returns(bool)
func (_Calldemo *CalldemoCallerSession) TfContribute(arg0 common.Address, arg1 *big.Int) (bool, error) {
	return _Calldemo.Contract.TfContribute(&_Calldemo.CallOpts, arg0, arg1)
}

// InsCollectFunding is a paid mutator transaction binding the contract method 0xb778be69.
//
// Solidity: function InsCollectFunding(uint256 key) returns(uint256)
func (_Calldemo *CalldemoTransactor) InsCollectFunding(opts *bind.TransactOpts, key *big.Int) (*types.Transaction, error) {
	return _Calldemo.contract.Transact(opts, "InsCollectFunding", key)
}

// InsCollectFunding is a paid mutator transaction binding the contract method 0xb778be69.
//
// Solidity: function InsCollectFunding(uint256 key) returns(uint256)
func (_Calldemo *CalldemoSession) InsCollectFunding(key *big.Int) (*types.Transaction, error) {
	return _Calldemo.Contract.InsCollectFunding(&_Calldemo.TransactOpts, key)
}

// InsCollectFunding is a paid mutator transaction binding the contract method 0xb778be69.
//
// Solidity: function InsCollectFunding(uint256 key) returns(uint256)
func (_Calldemo *CalldemoTransactorSession) InsCollectFunding(key *big.Int) (*types.Transaction, error) {
	return _Calldemo.Contract.InsCollectFunding(&_Calldemo.TransactOpts, key)
}

// Contribute is a paid mutator transaction binding the contract method 0xc1cbbca7.
//
// Solidity: function contribute(uint256 ID) payable returns()
func (_Calldemo *CalldemoTransactor) Contribute(opts *bind.TransactOpts, ID *big.Int) (*types.Transaction, error) {
	return _Calldemo.contract.Transact(opts, "contribute", ID)
}

// Contribute is a paid mutator transaction binding the contract method 0xc1cbbca7.
//
// Solidity: function contribute(uint256 ID) payable returns()
func (_Calldemo *CalldemoSession) Contribute(ID *big.Int) (*types.Transaction, error) {
	return _Calldemo.Contract.Contribute(&_Calldemo.TransactOpts, ID)
}

// Contribute is a paid mutator transaction binding the contract method 0xc1cbbca7.
//
// Solidity: function contribute(uint256 ID) payable returns()
func (_Calldemo *CalldemoTransactorSession) Contribute(ID *big.Int) (*types.Transaction, error) {
	return _Calldemo.Contract.Contribute(&_Calldemo.TransactOpts, ID)
}

// DelOneCollectFunding is a paid mutator transaction binding the contract method 0x3cff8f7c.
//
// Solidity: function delOneCollectFunding(uint256 index) returns(bool)
func (_Calldemo *CalldemoTransactor) DelOneCollectFunding(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Calldemo.contract.Transact(opts, "delOneCollectFunding", index)
}

// DelOneCollectFunding is a paid mutator transaction binding the contract method 0x3cff8f7c.
//
// Solidity: function delOneCollectFunding(uint256 index) returns(bool)
func (_Calldemo *CalldemoSession) DelOneCollectFunding(index *big.Int) (*types.Transaction, error) {
	return _Calldemo.Contract.DelOneCollectFunding(&_Calldemo.TransactOpts, index)
}

// DelOneCollectFunding is a paid mutator transaction binding the contract method 0x3cff8f7c.
//
// Solidity: function delOneCollectFunding(uint256 index) returns(bool)
func (_Calldemo *CalldemoTransactorSession) DelOneCollectFunding(index *big.Int) (*types.Transaction, error) {
	return _Calldemo.Contract.DelOneCollectFunding(&_Calldemo.TransactOpts, index)
}

// DeleteFunding is a paid mutator transaction binding the contract method 0x816a63ca.
//
// Solidity: function deleteFunding(uint256 index) returns(uint256)
func (_Calldemo *CalldemoTransactor) DeleteFunding(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Calldemo.contract.Transact(opts, "deleteFunding", index)
}

// DeleteFunding is a paid mutator transaction binding the contract method 0x816a63ca.
//
// Solidity: function deleteFunding(uint256 index) returns(uint256)
func (_Calldemo *CalldemoSession) DeleteFunding(index *big.Int) (*types.Transaction, error) {
	return _Calldemo.Contract.DeleteFunding(&_Calldemo.TransactOpts, index)
}

// DeleteFunding is a paid mutator transaction binding the contract method 0x816a63ca.
//
// Solidity: function deleteFunding(uint256 index) returns(uint256)
func (_Calldemo *CalldemoTransactorSession) DeleteFunding(index *big.Int) (*types.Transaction, error) {
	return _Calldemo.Contract.DeleteFunding(&_Calldemo.TransactOpts, index)
}

// NewFunding is a paid mutator transaction binding the contract method 0x4147c04c.
//
// Solidity: function newFunding(address initiator, string title, string info, uint256 goal) returns(uint256)
func (_Calldemo *CalldemoTransactor) NewFunding(opts *bind.TransactOpts, initiator common.Address, title string, info string, goal *big.Int) (*types.Transaction, error) {
	return _Calldemo.contract.Transact(opts, "newFunding", initiator, title, info, goal)
}

// NewFunding is a paid mutator transaction binding the contract method 0x4147c04c.
//
// Solidity: function newFunding(address initiator, string title, string info, uint256 goal) returns(uint256)
func (_Calldemo *CalldemoSession) NewFunding(initiator common.Address, title string, info string, goal *big.Int) (*types.Transaction, error) {
	return _Calldemo.Contract.NewFunding(&_Calldemo.TransactOpts, initiator, title, info, goal)
}

// NewFunding is a paid mutator transaction binding the contract method 0x4147c04c.
//
// Solidity: function newFunding(address initiator, string title, string info, uint256 goal) returns(uint256)
func (_Calldemo *CalldemoTransactorSession) NewFunding(initiator common.Address, title string, info string, goal *big.Int) (*types.Transaction, error) {
	return _Calldemo.Contract.NewFunding(&_Calldemo.TransactOpts, initiator, title, info, goal)
}

// ReturnMoney is a paid mutator transaction binding the contract method 0x3785fc1a.
//
// Solidity: function returnMoney(uint256 ID) returns()
func (_Calldemo *CalldemoTransactor) ReturnMoney(opts *bind.TransactOpts, ID *big.Int) (*types.Transaction, error) {
	return _Calldemo.contract.Transact(opts, "returnMoney", ID)
}

// ReturnMoney is a paid mutator transaction binding the contract method 0x3785fc1a.
//
// Solidity: function returnMoney(uint256 ID) returns()
func (_Calldemo *CalldemoSession) ReturnMoney(ID *big.Int) (*types.Transaction, error) {
	return _Calldemo.Contract.ReturnMoney(&_Calldemo.TransactOpts, ID)
}

// ReturnMoney is a paid mutator transaction binding the contract method 0x3785fc1a.
//
// Solidity: function returnMoney(uint256 ID) returns()
func (_Calldemo *CalldemoTransactorSession) ReturnMoney(ID *big.Int) (*types.Transaction, error) {
	return _Calldemo.Contract.ReturnMoney(&_Calldemo.TransactOpts, ID)
}
