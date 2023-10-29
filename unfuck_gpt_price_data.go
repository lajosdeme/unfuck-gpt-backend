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

// IFtsoRegistryPriceInfo is an auto generated low-level Go binding around an user-defined struct.
type IFtsoRegistryPriceInfo struct {
	FtsoIndex *big.Int
	Price     *big.Int
	Decimals  *big.Int
	Timestamp *big.Int
}

// UnfuckPriceDataABI is the input ABI used to generate the binding from.
const UnfuckPriceDataABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"currentPriceInfo\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ftsoIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_symbol\",\"type\":\"string\"}],\"name\":\"latestPriceForToken\",\"outputs\":[{\"components\":[{\"internalType\":\"uint256\",\"name\":\"ftsoIndex\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"decimals\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"internalType\":\"structIFtsoRegistry.PriceInfo\",\"name\":\"\",\"type\":\"tuple\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"priceInfosCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"refreshPrices\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// UnfuckPriceDataBin is the compiled bytecode used for deploying new contracts.
var UnfuckPriceDataBin = "0x6080604052600180546001600160a01b03191673ad67fe66660fb8dfe9d6b1b4240d8650e30f601917905534801561003657600080fd5b5061004033610045565b610095565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6108ee806100a46000396000f3fe608060405234801561001057600080fd5b506004361061007d5760003560e01c80633d44abf11161005b5780633d44abf1146100ff578063715018a6146101455780638da5cb5b1461014d578063f2fde38b1461016857600080fd5b8063068b6bd51461008257806316638d591461008c57806334ece79d146100a8575b600080fd5b61008a61017b565b005b61009560035481565b6040519081526020015b60405180910390f35b6100df6100b63660046105bf565b600260208190526000918252604090912080546001820154928201546003909201549092919084565b60408051948552602085019390935291830152606082015260800161009f565b61011261010d366004610648565b610317565b60405161009f91908151815260208083015190820152604080830151908201526060918201519181019190915260800190565b61008a610483565b6000546040516001600160a01b03909116815260200161009f565b61008a6101763660046106f2565b610497565b610183610515565b60015460405163413b07e560e11b815260206004820152600c60248201526b4674736f526567697374727960a01b60448201526000916001600160a01b0316906382760fca90606401602060405180830381865afa1580156101e9573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061020d9190610716565b90506000816001600160a01b03166358f9296f6040518163ffffffff1660e01b8152600401600060405180830381865afa15801561024f573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526102779190810190610733565b8051600355905060005b81518110156103125781818151811061029c5761029c610814565b6020026020010151600260008484815181106102ba576102ba610814565b602002602001015160000151815260200190815260200160002060008201518160000155602082015181600101556040820151816002015560608201518160030155905050808061030a9061082a565b915050610281565b505050565b6103426040518060800160405280600081526020016000815260200160008152602001600081525090565b60015460405163413b07e560e11b815260206004820152600c60248201526b4674736f526567697374727960a01b60448201526000916001600160a01b0316906382760fca90606401602060405180830381865afa1580156103a8573d6000803e3d6000fd5b505050506040513d601f19601f820116820180604052508101906103cc9190610716565b90506000816001600160a01b031663e848da30856040518263ffffffff1660e01b81526004016103fc9190610851565b602060405180830381865afa158015610419573d6000803e3d6000fd5b505050506040513d601f19601f8201168201806040525081019061043d919061089f565b6000908152600260208181526040928390208351608081018552815481526001820154928101929092529182015492810192909252600301546060820152949350505050565b61048b610515565b610495600061056f565b565b61049f610515565b6001600160a01b0381166105095760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084015b60405180910390fd5b6105128161056f565b50565b6000546001600160a01b031633146104955760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e65726044820152606401610500565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156105d157600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b6040516080810167ffffffffffffffff81118282101715610611576106116105d8565b60405290565b604051601f8201601f1916810167ffffffffffffffff81118282101715610640576106406105d8565b604052919050565b6000602080838503121561065b57600080fd5b823567ffffffffffffffff8082111561067357600080fd5b818501915085601f83011261068757600080fd5b813581811115610699576106996105d8565b6106ab601f8201601f19168501610617565b915080825286848285010111156106c157600080fd5b8084840185840137600090820190930192909252509392505050565b6001600160a01b038116811461051257600080fd5b60006020828403121561070457600080fd5b813561070f816106dd565b9392505050565b60006020828403121561072857600080fd5b815161070f816106dd565b6000602080838503121561074657600080fd5b825167ffffffffffffffff8082111561075e57600080fd5b818501915085601f83011261077257600080fd5b815181811115610784576107846105d8565b610792848260051b01610617565b818152848101925060079190911b8301840190878211156107b257600080fd5b928401925b8184101561080957608084890312156107d05760008081fd5b6107d86105ee565b84518152858501518682015260408086015190820152606080860151908201528352608090930192918401916107b7565b979650505050505050565b634e487b7160e01b600052603260045260246000fd5b60006001820161084a57634e487b7160e01b600052601160045260246000fd5b5060010190565b600060208083528351808285015260005b8181101561087e57858101830151858201604001528201610862565b506000604082860101526040601f19601f8301168501019250505092915050565b6000602082840312156108b157600080fd5b505191905056fea2646970667358221220c7212c880275f541b680db35d91f2c12d687e943e71c88b765223f045f80372964736f6c63430008140033"

// DeployUnfuckPriceData deploys a new Ethereum contract, binding an instance of UnfuckPriceData to it.
func DeployUnfuckPriceData(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *UnfuckPriceData, error) {
	parsed, err := abi.JSON(strings.NewReader(UnfuckPriceDataABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnfuckPriceDataBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UnfuckPriceData{UnfuckPriceDataCaller: UnfuckPriceDataCaller{contract: contract}, UnfuckPriceDataTransactor: UnfuckPriceDataTransactor{contract: contract}, UnfuckPriceDataFilterer: UnfuckPriceDataFilterer{contract: contract}}, nil
}

// UnfuckPriceData is an auto generated Go binding around an Ethereum contract.
type UnfuckPriceData struct {
	UnfuckPriceDataCaller     // Read-only binding to the contract
	UnfuckPriceDataTransactor // Write-only binding to the contract
	UnfuckPriceDataFilterer   // Log filterer for contract events
}

// UnfuckPriceDataCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnfuckPriceDataCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnfuckPriceDataTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnfuckPriceDataTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnfuckPriceDataFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnfuckPriceDataFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnfuckPriceDataSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnfuckPriceDataSession struct {
	Contract     *UnfuckPriceData  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnfuckPriceDataCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnfuckPriceDataCallerSession struct {
	Contract *UnfuckPriceDataCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// UnfuckPriceDataTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnfuckPriceDataTransactorSession struct {
	Contract     *UnfuckPriceDataTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// UnfuckPriceDataRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnfuckPriceDataRaw struct {
	Contract *UnfuckPriceData // Generic contract binding to access the raw methods on
}

// UnfuckPriceDataCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnfuckPriceDataCallerRaw struct {
	Contract *UnfuckPriceDataCaller // Generic read-only contract binding to access the raw methods on
}

// UnfuckPriceDataTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnfuckPriceDataTransactorRaw struct {
	Contract *UnfuckPriceDataTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnfuckPriceData creates a new instance of UnfuckPriceData, bound to a specific deployed contract.
func NewUnfuckPriceData(address common.Address, backend bind.ContractBackend) (*UnfuckPriceData, error) {
	contract, err := bindUnfuckPriceData(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnfuckPriceData{UnfuckPriceDataCaller: UnfuckPriceDataCaller{contract: contract}, UnfuckPriceDataTransactor: UnfuckPriceDataTransactor{contract: contract}, UnfuckPriceDataFilterer: UnfuckPriceDataFilterer{contract: contract}}, nil
}

// NewUnfuckPriceDataCaller creates a new read-only instance of UnfuckPriceData, bound to a specific deployed contract.
func NewUnfuckPriceDataCaller(address common.Address, caller bind.ContractCaller) (*UnfuckPriceDataCaller, error) {
	contract, err := bindUnfuckPriceData(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnfuckPriceDataCaller{contract: contract}, nil
}

// NewUnfuckPriceDataTransactor creates a new write-only instance of UnfuckPriceData, bound to a specific deployed contract.
func NewUnfuckPriceDataTransactor(address common.Address, transactor bind.ContractTransactor) (*UnfuckPriceDataTransactor, error) {
	contract, err := bindUnfuckPriceData(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnfuckPriceDataTransactor{contract: contract}, nil
}

// NewUnfuckPriceDataFilterer creates a new log filterer instance of UnfuckPriceData, bound to a specific deployed contract.
func NewUnfuckPriceDataFilterer(address common.Address, filterer bind.ContractFilterer) (*UnfuckPriceDataFilterer, error) {
	contract, err := bindUnfuckPriceData(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnfuckPriceDataFilterer{contract: contract}, nil
}

// bindUnfuckPriceData binds a generic wrapper to an already deployed contract.
func bindUnfuckPriceData(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnfuckPriceDataABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnfuckPriceData *UnfuckPriceDataRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnfuckPriceData.Contract.UnfuckPriceDataCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnfuckPriceData *UnfuckPriceDataRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnfuckPriceData.Contract.UnfuckPriceDataTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnfuckPriceData *UnfuckPriceDataRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnfuckPriceData.Contract.UnfuckPriceDataTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnfuckPriceData *UnfuckPriceDataCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnfuckPriceData.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnfuckPriceData *UnfuckPriceDataTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnfuckPriceData.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnfuckPriceData *UnfuckPriceDataTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnfuckPriceData.Contract.contract.Transact(opts, method, params...)
}

// CurrentPriceInfo is a free data retrieval call binding the contract method 0x34ece79d.
//
// Solidity: function currentPriceInfo(uint256 ) view returns(uint256 ftsoIndex, uint256 price, uint256 decimals, uint256 timestamp)
func (_UnfuckPriceData *UnfuckPriceDataCaller) CurrentPriceInfo(opts *bind.CallOpts, arg0 *big.Int) (struct {
	FtsoIndex *big.Int
	Price     *big.Int
	Decimals  *big.Int
	Timestamp *big.Int
}, error) {
	var out []interface{}
	err := _UnfuckPriceData.contract.Call(opts, &out, "currentPriceInfo", arg0)

	outstruct := new(struct {
		FtsoIndex *big.Int
		Price     *big.Int
		Decimals  *big.Int
		Timestamp *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.FtsoIndex = *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	outstruct.Price = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Decimals = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Timestamp = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// CurrentPriceInfo is a free data retrieval call binding the contract method 0x34ece79d.
//
// Solidity: function currentPriceInfo(uint256 ) view returns(uint256 ftsoIndex, uint256 price, uint256 decimals, uint256 timestamp)
func (_UnfuckPriceData *UnfuckPriceDataSession) CurrentPriceInfo(arg0 *big.Int) (struct {
	FtsoIndex *big.Int
	Price     *big.Int
	Decimals  *big.Int
	Timestamp *big.Int
}, error) {
	return _UnfuckPriceData.Contract.CurrentPriceInfo(&_UnfuckPriceData.CallOpts, arg0)
}

// CurrentPriceInfo is a free data retrieval call binding the contract method 0x34ece79d.
//
// Solidity: function currentPriceInfo(uint256 ) view returns(uint256 ftsoIndex, uint256 price, uint256 decimals, uint256 timestamp)
func (_UnfuckPriceData *UnfuckPriceDataCallerSession) CurrentPriceInfo(arg0 *big.Int) (struct {
	FtsoIndex *big.Int
	Price     *big.Int
	Decimals  *big.Int
	Timestamp *big.Int
}, error) {
	return _UnfuckPriceData.Contract.CurrentPriceInfo(&_UnfuckPriceData.CallOpts, arg0)
}

// LatestPriceForToken is a free data retrieval call binding the contract method 0x3d44abf1.
//
// Solidity: function latestPriceForToken(string _symbol) view returns((uint256,uint256,uint256,uint256))
func (_UnfuckPriceData *UnfuckPriceDataCaller) LatestPriceForToken(opts *bind.CallOpts, _symbol string) (IFtsoRegistryPriceInfo, error) {
	var out []interface{}
	err := _UnfuckPriceData.contract.Call(opts, &out, "latestPriceForToken", _symbol)

	if err != nil {
		return *new(IFtsoRegistryPriceInfo), err
	}

	out0 := *abi.ConvertType(out[0], new(IFtsoRegistryPriceInfo)).(*IFtsoRegistryPriceInfo)

	return out0, err

}

// LatestPriceForToken is a free data retrieval call binding the contract method 0x3d44abf1.
//
// Solidity: function latestPriceForToken(string _symbol) view returns((uint256,uint256,uint256,uint256))
func (_UnfuckPriceData *UnfuckPriceDataSession) LatestPriceForToken(_symbol string) (IFtsoRegistryPriceInfo, error) {
	return _UnfuckPriceData.Contract.LatestPriceForToken(&_UnfuckPriceData.CallOpts, _symbol)
}

// LatestPriceForToken is a free data retrieval call binding the contract method 0x3d44abf1.
//
// Solidity: function latestPriceForToken(string _symbol) view returns((uint256,uint256,uint256,uint256))
func (_UnfuckPriceData *UnfuckPriceDataCallerSession) LatestPriceForToken(_symbol string) (IFtsoRegistryPriceInfo, error) {
	return _UnfuckPriceData.Contract.LatestPriceForToken(&_UnfuckPriceData.CallOpts, _symbol)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnfuckPriceData *UnfuckPriceDataCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnfuckPriceData.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnfuckPriceData *UnfuckPriceDataSession) Owner() (common.Address, error) {
	return _UnfuckPriceData.Contract.Owner(&_UnfuckPriceData.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnfuckPriceData *UnfuckPriceDataCallerSession) Owner() (common.Address, error) {
	return _UnfuckPriceData.Contract.Owner(&_UnfuckPriceData.CallOpts)
}

// PriceInfosCount is a free data retrieval call binding the contract method 0x16638d59.
//
// Solidity: function priceInfosCount() view returns(uint256)
func (_UnfuckPriceData *UnfuckPriceDataCaller) PriceInfosCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _UnfuckPriceData.contract.Call(opts, &out, "priceInfosCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// PriceInfosCount is a free data retrieval call binding the contract method 0x16638d59.
//
// Solidity: function priceInfosCount() view returns(uint256)
func (_UnfuckPriceData *UnfuckPriceDataSession) PriceInfosCount() (*big.Int, error) {
	return _UnfuckPriceData.Contract.PriceInfosCount(&_UnfuckPriceData.CallOpts)
}

// PriceInfosCount is a free data retrieval call binding the contract method 0x16638d59.
//
// Solidity: function priceInfosCount() view returns(uint256)
func (_UnfuckPriceData *UnfuckPriceDataCallerSession) PriceInfosCount() (*big.Int, error) {
	return _UnfuckPriceData.Contract.PriceInfosCount(&_UnfuckPriceData.CallOpts)
}

// RefreshPrices is a paid mutator transaction binding the contract method 0x068b6bd5.
//
// Solidity: function refreshPrices() returns()
func (_UnfuckPriceData *UnfuckPriceDataTransactor) RefreshPrices(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnfuckPriceData.contract.Transact(opts, "refreshPrices")
}

// RefreshPrices is a paid mutator transaction binding the contract method 0x068b6bd5.
//
// Solidity: function refreshPrices() returns()
func (_UnfuckPriceData *UnfuckPriceDataSession) RefreshPrices() (*types.Transaction, error) {
	return _UnfuckPriceData.Contract.RefreshPrices(&_UnfuckPriceData.TransactOpts)
}

// RefreshPrices is a paid mutator transaction binding the contract method 0x068b6bd5.
//
// Solidity: function refreshPrices() returns()
func (_UnfuckPriceData *UnfuckPriceDataTransactorSession) RefreshPrices() (*types.Transaction, error) {
	return _UnfuckPriceData.Contract.RefreshPrices(&_UnfuckPriceData.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnfuckPriceData *UnfuckPriceDataTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnfuckPriceData.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnfuckPriceData *UnfuckPriceDataSession) RenounceOwnership() (*types.Transaction, error) {
	return _UnfuckPriceData.Contract.RenounceOwnership(&_UnfuckPriceData.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnfuckPriceData *UnfuckPriceDataTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UnfuckPriceData.Contract.RenounceOwnership(&_UnfuckPriceData.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnfuckPriceData *UnfuckPriceDataTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UnfuckPriceData.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnfuckPriceData *UnfuckPriceDataSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnfuckPriceData.Contract.TransferOwnership(&_UnfuckPriceData.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnfuckPriceData *UnfuckPriceDataTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnfuckPriceData.Contract.TransferOwnership(&_UnfuckPriceData.TransactOpts, newOwner)
}

// UnfuckPriceDataOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UnfuckPriceData contract.
type UnfuckPriceDataOwnershipTransferredIterator struct {
	Event *UnfuckPriceDataOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UnfuckPriceDataOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnfuckPriceDataOwnershipTransferred)
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
		it.Event = new(UnfuckPriceDataOwnershipTransferred)
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
func (it *UnfuckPriceDataOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnfuckPriceDataOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnfuckPriceDataOwnershipTransferred represents a OwnershipTransferred event raised by the UnfuckPriceData contract.
type UnfuckPriceDataOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnfuckPriceData *UnfuckPriceDataFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UnfuckPriceDataOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UnfuckPriceData.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UnfuckPriceDataOwnershipTransferredIterator{contract: _UnfuckPriceData.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnfuckPriceData *UnfuckPriceDataFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UnfuckPriceDataOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UnfuckPriceData.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnfuckPriceDataOwnershipTransferred)
				if err := _UnfuckPriceData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UnfuckPriceData *UnfuckPriceDataFilterer) ParseOwnershipTransferred(log types.Log) (*UnfuckPriceDataOwnershipTransferred, error) {
	event := new(UnfuckPriceDataOwnershipTransferred)
	if err := _UnfuckPriceData.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
