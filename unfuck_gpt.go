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

// UnfuckCoreABI is the input ABI used to generate the binding from.
const UnfuckCoreABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_wormholeRelayer\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"MembershipChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"members\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_to\",\"type\":\"address\"}],\"name\":\"membershipChangeOnHostChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"nftContracts\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"payload\",\"type\":\"bytes\"},{\"internalType\":\"bytes[]\",\"name\":\"\",\"type\":\"bytes[]\"},{\"internalType\":\"bytes32\",\"name\":\"sourceAddress\",\"type\":\"bytes32\"},{\"internalType\":\"uint16\",\"name\":\"sourceChain\",\"type\":\"uint16\"},{\"internalType\":\"bytes32\",\"name\":\"deliveryHash\",\"type\":\"bytes32\"}],\"name\":\"receiveWormholeMessages\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"name\":\"seenDeliveryVaaHashes\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_chainId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_contract\",\"type\":\"address\"}],\"name\":\"setNftContractAddressOnChain\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"wormholeRelayer\",\"outputs\":[{\"internalType\":\"contractIWormholeRelayer\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// UnfuckCoreBin is the compiled bytecode used for deploying new contracts.
var UnfuckCoreBin = "0x60a060405234801561001057600080fd5b50604051610a61380380610a6183398101604081905261002f91610099565b61003833610049565b6001600160a01b03166080526100c9565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6000602082840312156100ab57600080fd5b81516001600160a01b03811681146100c257600080fd5b9392505050565b6080516109766100eb600039600081816101d3015261024001526109766000f3fe6080604052600436106100915760003560e01c80638da5cb5b116100595780638da5cb5b146101835780639b4e17ef146101a1578063da25b725146101c1578063f2fde38b146101f5578063ff9910ea1461021557600080fd5b806308ae4b0c146100965780630bb4d82a146100db578063180f6cc214610129578063529dca3214610159578063715018a61461016e575b600080fd5b3480156100a257600080fd5b506100c66100b136600461069f565b60026020526000908152604090205460ff1681565b60405190151581526020015b60405180910390f35b3480156100e757600080fd5b506101116100f63660046106c3565b6003602052600090815260409020546001600160a01b031681565b6040516001600160a01b0390911681526020016100d2565b34801561013557600080fd5b506100c66101443660046106c3565b60016020526000908152604090205460ff1681565b61016c6101673660046107aa565b610235565b005b34801561017a57600080fd5b5061016c6103d8565b34801561018f57600080fd5b506000546001600160a01b0316610111565b3480156101ad57600080fd5b5061016c6101bc3660046108b3565b6103ec565b3480156101cd57600080fd5b506101117f000000000000000000000000000000000000000000000000000000000000000081565b34801561020157600080fd5b5061016c61021036600461069f565b6104cf565b34801561022157600080fd5b5061016c6102303660046108ec565b610548565b336001600160a01b037f000000000000000000000000000000000000000000000000000000000000000016146102a95760405162461bcd60e51b815260206004820152601460248201527313db9b1e481c995b185e595c88185b1b1bddd95960621b60448201526064015b60405180910390fd5b60008181526001602052604090205460ff16156103085760405162461bcd60e51b815260206004820152601960248201527f4d65737361676520616c72656164792070726f6365737365640000000000000060448201526064016102a0565b6000818152600160208181526040808420805460ff191690931790925561ffff85168352600390529020546001600160a01b0316836001600160a01b0316146103aa5760405162461bcd60e51b815260206004820152602e60248201527f4f6e6c792074686520636f6e66696775726564204e465420636f6e747261637460448201526d2063616e2073656e64206d73677360901b60648201526084016102a0565b600080868060200190518101906103c19190610911565b915091506103cf828261057e565b50505050505050565b6103e06105e0565b6103ea600061063a565b565b46620138811461043e5760405162461bcd60e51b815260206004820152601c60248201527f486f737420636861696e20697320706f6c79676f6e206d756d6261690000000060448201526064016102a0565b600560005260036020527f405aad32e1adbac89bb7f176e338b8fc6e994ca210c9bb7bdca249b465942250546001600160a01b031633146104c15760405162461bcd60e51b815260206004820152601760248201527f4e6f742066726f6d20686f737420636861696e206e667400000000000000000060448201526064016102a0565b6104cb828261057e565b5050565b6104d76105e0565b6001600160a01b03811661053c5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b60648201526084016102a0565b6105458161063a565b50565b6105506105e0565b60009182526003602052604090912080546001600160a01b0319166001600160a01b03909216919091179055565b6001600160a01b03808316600081815260026020526040808220805460ff199081169091559385168083528183208054909516600117909455517f3331f695d86f449bec7d3bdf70eefa8b100f9dd3b1699e3ad142d67bd846840b9190a35050565b6000546001600160a01b031633146103ea5760405162461bcd60e51b815260206004820181905260248201527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e657260448201526064016102a0565b600080546001600160a01b038381166001600160a01b0319831681178455604051919092169283917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e09190a35050565b6001600160a01b038116811461054557600080fd5b6000602082840312156106b157600080fd5b81356106bc8161068a565b9392505050565b6000602082840312156106d557600080fd5b5035919050565b634e487b7160e01b600052604160045260246000fd5b604051601f8201601f1916810167ffffffffffffffff8111828210171561071b5761071b6106dc565b604052919050565b600082601f83011261073457600080fd5b813567ffffffffffffffff81111561074e5761074e6106dc565b610761601f8201601f19166020016106f2565b81815284602083860101111561077657600080fd5b816020850160208301376000918101602001919091529392505050565b803561ffff811681146107a557600080fd5b919050565b600080600080600060a086880312156107c257600080fd5b853567ffffffffffffffff808211156107da57600080fd5b6107e689838a01610723565b96506020915081880135818111156107fd57600080fd5b8801601f81018a1361080e57600080fd5b803582811115610820576108206106dc565b8060051b61082f8582016106f2565b918252828101850191858101908d84111561084957600080fd5b86850192505b83831015610885578235868111156108675760008081fd5b6108758f8983890101610723565b835250918601919086019061084f565b809a5050505050505050604086013592506108a260608701610793565b949793965091946080013592915050565b600080604083850312156108c657600080fd5b82356108d18161068a565b915060208301356108e18161068a565b809150509250929050565b600080604083850312156108ff57600080fd5b8235915060208301356108e18161068a565b6000806040838503121561092457600080fd5b825161092f8161068a565b60208401519092506108e18161068a56fea26469706673582212207cfbe3b036d76bfd5b39fdc2e7adf3a15cb59fb4e9f68fa611aef84f780e928a64736f6c63430008140033"

// DeployUnfuckCore deploys a new Ethereum contract, binding an instance of UnfuckCore to it.
func DeployUnfuckCore(auth *bind.TransactOpts, backend bind.ContractBackend, _wormholeRelayer common.Address) (common.Address, *types.Transaction, *UnfuckCore, error) {
	parsed, err := abi.JSON(strings.NewReader(UnfuckCoreABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(UnfuckCoreBin), backend, _wormholeRelayer)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &UnfuckCore{UnfuckCoreCaller: UnfuckCoreCaller{contract: contract}, UnfuckCoreTransactor: UnfuckCoreTransactor{contract: contract}, UnfuckCoreFilterer: UnfuckCoreFilterer{contract: contract}}, nil
}

// UnfuckCore is an auto generated Go binding around an Ethereum contract.
type UnfuckCore struct {
	UnfuckCoreCaller     // Read-only binding to the contract
	UnfuckCoreTransactor // Write-only binding to the contract
	UnfuckCoreFilterer   // Log filterer for contract events
}

// UnfuckCoreCaller is an auto generated read-only Go binding around an Ethereum contract.
type UnfuckCoreCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnfuckCoreTransactor is an auto generated write-only Go binding around an Ethereum contract.
type UnfuckCoreTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnfuckCoreFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type UnfuckCoreFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// UnfuckCoreSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type UnfuckCoreSession struct {
	Contract     *UnfuckCore       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// UnfuckCoreCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type UnfuckCoreCallerSession struct {
	Contract *UnfuckCoreCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// UnfuckCoreTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type UnfuckCoreTransactorSession struct {
	Contract     *UnfuckCoreTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// UnfuckCoreRaw is an auto generated low-level Go binding around an Ethereum contract.
type UnfuckCoreRaw struct {
	Contract *UnfuckCore // Generic contract binding to access the raw methods on
}

// UnfuckCoreCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type UnfuckCoreCallerRaw struct {
	Contract *UnfuckCoreCaller // Generic read-only contract binding to access the raw methods on
}

// UnfuckCoreTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type UnfuckCoreTransactorRaw struct {
	Contract *UnfuckCoreTransactor // Generic write-only contract binding to access the raw methods on
}

// NewUnfuckCore creates a new instance of UnfuckCore, bound to a specific deployed contract.
func NewUnfuckCore(address common.Address, backend bind.ContractBackend) (*UnfuckCore, error) {
	contract, err := bindUnfuckCore(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &UnfuckCore{UnfuckCoreCaller: UnfuckCoreCaller{contract: contract}, UnfuckCoreTransactor: UnfuckCoreTransactor{contract: contract}, UnfuckCoreFilterer: UnfuckCoreFilterer{contract: contract}}, nil
}

// NewUnfuckCoreCaller creates a new read-only instance of UnfuckCore, bound to a specific deployed contract.
func NewUnfuckCoreCaller(address common.Address, caller bind.ContractCaller) (*UnfuckCoreCaller, error) {
	contract, err := bindUnfuckCore(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &UnfuckCoreCaller{contract: contract}, nil
}

// NewUnfuckCoreTransactor creates a new write-only instance of UnfuckCore, bound to a specific deployed contract.
func NewUnfuckCoreTransactor(address common.Address, transactor bind.ContractTransactor) (*UnfuckCoreTransactor, error) {
	contract, err := bindUnfuckCore(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &UnfuckCoreTransactor{contract: contract}, nil
}

// NewUnfuckCoreFilterer creates a new log filterer instance of UnfuckCore, bound to a specific deployed contract.
func NewUnfuckCoreFilterer(address common.Address, filterer bind.ContractFilterer) (*UnfuckCoreFilterer, error) {
	contract, err := bindUnfuckCore(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &UnfuckCoreFilterer{contract: contract}, nil
}

// bindUnfuckCore binds a generic wrapper to an already deployed contract.
func bindUnfuckCore(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(UnfuckCoreABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnfuckCore *UnfuckCoreRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnfuckCore.Contract.UnfuckCoreCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnfuckCore *UnfuckCoreRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnfuckCore.Contract.UnfuckCoreTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnfuckCore *UnfuckCoreRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnfuckCore.Contract.UnfuckCoreTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_UnfuckCore *UnfuckCoreCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _UnfuckCore.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_UnfuckCore *UnfuckCoreTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnfuckCore.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_UnfuckCore *UnfuckCoreTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _UnfuckCore.Contract.contract.Transact(opts, method, params...)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bool)
func (_UnfuckCore *UnfuckCoreCaller) Members(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _UnfuckCore.contract.Call(opts, &out, "members", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bool)
func (_UnfuckCore *UnfuckCoreSession) Members(arg0 common.Address) (bool, error) {
	return _UnfuckCore.Contract.Members(&_UnfuckCore.CallOpts, arg0)
}

// Members is a free data retrieval call binding the contract method 0x08ae4b0c.
//
// Solidity: function members(address ) view returns(bool)
func (_UnfuckCore *UnfuckCoreCallerSession) Members(arg0 common.Address) (bool, error) {
	return _UnfuckCore.Contract.Members(&_UnfuckCore.CallOpts, arg0)
}

// NftContracts is a free data retrieval call binding the contract method 0x0bb4d82a.
//
// Solidity: function nftContracts(uint256 ) view returns(address)
func (_UnfuckCore *UnfuckCoreCaller) NftContracts(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _UnfuckCore.contract.Call(opts, &out, "nftContracts", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// NftContracts is a free data retrieval call binding the contract method 0x0bb4d82a.
//
// Solidity: function nftContracts(uint256 ) view returns(address)
func (_UnfuckCore *UnfuckCoreSession) NftContracts(arg0 *big.Int) (common.Address, error) {
	return _UnfuckCore.Contract.NftContracts(&_UnfuckCore.CallOpts, arg0)
}

// NftContracts is a free data retrieval call binding the contract method 0x0bb4d82a.
//
// Solidity: function nftContracts(uint256 ) view returns(address)
func (_UnfuckCore *UnfuckCoreCallerSession) NftContracts(arg0 *big.Int) (common.Address, error) {
	return _UnfuckCore.Contract.NftContracts(&_UnfuckCore.CallOpts, arg0)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnfuckCore *UnfuckCoreCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnfuckCore.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnfuckCore *UnfuckCoreSession) Owner() (common.Address, error) {
	return _UnfuckCore.Contract.Owner(&_UnfuckCore.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_UnfuckCore *UnfuckCoreCallerSession) Owner() (common.Address, error) {
	return _UnfuckCore.Contract.Owner(&_UnfuckCore.CallOpts)
}

// SeenDeliveryVaaHashes is a free data retrieval call binding the contract method 0x180f6cc2.
//
// Solidity: function seenDeliveryVaaHashes(bytes32 ) view returns(bool)
func (_UnfuckCore *UnfuckCoreCaller) SeenDeliveryVaaHashes(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _UnfuckCore.contract.Call(opts, &out, "seenDeliveryVaaHashes", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SeenDeliveryVaaHashes is a free data retrieval call binding the contract method 0x180f6cc2.
//
// Solidity: function seenDeliveryVaaHashes(bytes32 ) view returns(bool)
func (_UnfuckCore *UnfuckCoreSession) SeenDeliveryVaaHashes(arg0 [32]byte) (bool, error) {
	return _UnfuckCore.Contract.SeenDeliveryVaaHashes(&_UnfuckCore.CallOpts, arg0)
}

// SeenDeliveryVaaHashes is a free data retrieval call binding the contract method 0x180f6cc2.
//
// Solidity: function seenDeliveryVaaHashes(bytes32 ) view returns(bool)
func (_UnfuckCore *UnfuckCoreCallerSession) SeenDeliveryVaaHashes(arg0 [32]byte) (bool, error) {
	return _UnfuckCore.Contract.SeenDeliveryVaaHashes(&_UnfuckCore.CallOpts, arg0)
}

// WormholeRelayer is a free data retrieval call binding the contract method 0xda25b725.
//
// Solidity: function wormholeRelayer() view returns(address)
func (_UnfuckCore *UnfuckCoreCaller) WormholeRelayer(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _UnfuckCore.contract.Call(opts, &out, "wormholeRelayer")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// WormholeRelayer is a free data retrieval call binding the contract method 0xda25b725.
//
// Solidity: function wormholeRelayer() view returns(address)
func (_UnfuckCore *UnfuckCoreSession) WormholeRelayer() (common.Address, error) {
	return _UnfuckCore.Contract.WormholeRelayer(&_UnfuckCore.CallOpts)
}

// WormholeRelayer is a free data retrieval call binding the contract method 0xda25b725.
//
// Solidity: function wormholeRelayer() view returns(address)
func (_UnfuckCore *UnfuckCoreCallerSession) WormholeRelayer() (common.Address, error) {
	return _UnfuckCore.Contract.WormholeRelayer(&_UnfuckCore.CallOpts)
}

// MembershipChangeOnHostChain is a paid mutator transaction binding the contract method 0x9b4e17ef.
//
// Solidity: function membershipChangeOnHostChain(address _from, address _to) returns()
func (_UnfuckCore *UnfuckCoreTransactor) MembershipChangeOnHostChain(opts *bind.TransactOpts, _from common.Address, _to common.Address) (*types.Transaction, error) {
	return _UnfuckCore.contract.Transact(opts, "membershipChangeOnHostChain", _from, _to)
}

// MembershipChangeOnHostChain is a paid mutator transaction binding the contract method 0x9b4e17ef.
//
// Solidity: function membershipChangeOnHostChain(address _from, address _to) returns()
func (_UnfuckCore *UnfuckCoreSession) MembershipChangeOnHostChain(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _UnfuckCore.Contract.MembershipChangeOnHostChain(&_UnfuckCore.TransactOpts, _from, _to)
}

// MembershipChangeOnHostChain is a paid mutator transaction binding the contract method 0x9b4e17ef.
//
// Solidity: function membershipChangeOnHostChain(address _from, address _to) returns()
func (_UnfuckCore *UnfuckCoreTransactorSession) MembershipChangeOnHostChain(_from common.Address, _to common.Address) (*types.Transaction, error) {
	return _UnfuckCore.Contract.MembershipChangeOnHostChain(&_UnfuckCore.TransactOpts, _from, _to)
}

// ReceiveWormholeMessages is a paid mutator transaction binding the contract method 0x529dca32.
//
// Solidity: function receiveWormholeMessages(bytes payload, bytes[] , bytes32 sourceAddress, uint16 sourceChain, bytes32 deliveryHash) payable returns()
func (_UnfuckCore *UnfuckCoreTransactor) ReceiveWormholeMessages(opts *bind.TransactOpts, payload []byte, arg1 [][]byte, sourceAddress [32]byte, sourceChain uint16, deliveryHash [32]byte) (*types.Transaction, error) {
	return _UnfuckCore.contract.Transact(opts, "receiveWormholeMessages", payload, arg1, sourceAddress, sourceChain, deliveryHash)
}

// ReceiveWormholeMessages is a paid mutator transaction binding the contract method 0x529dca32.
//
// Solidity: function receiveWormholeMessages(bytes payload, bytes[] , bytes32 sourceAddress, uint16 sourceChain, bytes32 deliveryHash) payable returns()
func (_UnfuckCore *UnfuckCoreSession) ReceiveWormholeMessages(payload []byte, arg1 [][]byte, sourceAddress [32]byte, sourceChain uint16, deliveryHash [32]byte) (*types.Transaction, error) {
	return _UnfuckCore.Contract.ReceiveWormholeMessages(&_UnfuckCore.TransactOpts, payload, arg1, sourceAddress, sourceChain, deliveryHash)
}

// ReceiveWormholeMessages is a paid mutator transaction binding the contract method 0x529dca32.
//
// Solidity: function receiveWormholeMessages(bytes payload, bytes[] , bytes32 sourceAddress, uint16 sourceChain, bytes32 deliveryHash) payable returns()
func (_UnfuckCore *UnfuckCoreTransactorSession) ReceiveWormholeMessages(payload []byte, arg1 [][]byte, sourceAddress [32]byte, sourceChain uint16, deliveryHash [32]byte) (*types.Transaction, error) {
	return _UnfuckCore.Contract.ReceiveWormholeMessages(&_UnfuckCore.TransactOpts, payload, arg1, sourceAddress, sourceChain, deliveryHash)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnfuckCore *UnfuckCoreTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _UnfuckCore.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnfuckCore *UnfuckCoreSession) RenounceOwnership() (*types.Transaction, error) {
	return _UnfuckCore.Contract.RenounceOwnership(&_UnfuckCore.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_UnfuckCore *UnfuckCoreTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _UnfuckCore.Contract.RenounceOwnership(&_UnfuckCore.TransactOpts)
}

// SetNftContractAddressOnChain is a paid mutator transaction binding the contract method 0xff9910ea.
//
// Solidity: function setNftContractAddressOnChain(uint256 _chainId, address _contract) returns()
func (_UnfuckCore *UnfuckCoreTransactor) SetNftContractAddressOnChain(opts *bind.TransactOpts, _chainId *big.Int, _contract common.Address) (*types.Transaction, error) {
	return _UnfuckCore.contract.Transact(opts, "setNftContractAddressOnChain", _chainId, _contract)
}

// SetNftContractAddressOnChain is a paid mutator transaction binding the contract method 0xff9910ea.
//
// Solidity: function setNftContractAddressOnChain(uint256 _chainId, address _contract) returns()
func (_UnfuckCore *UnfuckCoreSession) SetNftContractAddressOnChain(_chainId *big.Int, _contract common.Address) (*types.Transaction, error) {
	return _UnfuckCore.Contract.SetNftContractAddressOnChain(&_UnfuckCore.TransactOpts, _chainId, _contract)
}

// SetNftContractAddressOnChain is a paid mutator transaction binding the contract method 0xff9910ea.
//
// Solidity: function setNftContractAddressOnChain(uint256 _chainId, address _contract) returns()
func (_UnfuckCore *UnfuckCoreTransactorSession) SetNftContractAddressOnChain(_chainId *big.Int, _contract common.Address) (*types.Transaction, error) {
	return _UnfuckCore.Contract.SetNftContractAddressOnChain(&_UnfuckCore.TransactOpts, _chainId, _contract)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnfuckCore *UnfuckCoreTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _UnfuckCore.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnfuckCore *UnfuckCoreSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnfuckCore.Contract.TransferOwnership(&_UnfuckCore.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_UnfuckCore *UnfuckCoreTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _UnfuckCore.Contract.TransferOwnership(&_UnfuckCore.TransactOpts, newOwner)
}

// UnfuckCoreMembershipChangedIterator is returned from FilterMembershipChanged and is used to iterate over the raw logs and unpacked data for MembershipChanged events raised by the UnfuckCore contract.
type UnfuckCoreMembershipChangedIterator struct {
	Event *UnfuckCoreMembershipChanged // Event containing the contract specifics and raw log

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
func (it *UnfuckCoreMembershipChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnfuckCoreMembershipChanged)
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
		it.Event = new(UnfuckCoreMembershipChanged)
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
func (it *UnfuckCoreMembershipChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnfuckCoreMembershipChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnfuckCoreMembershipChanged represents a MembershipChanged event raised by the UnfuckCore contract.
type UnfuckCoreMembershipChanged struct {
	From common.Address
	To   common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterMembershipChanged is a free log retrieval operation binding the contract event 0x3331f695d86f449bec7d3bdf70eefa8b100f9dd3b1699e3ad142d67bd846840b.
//
// Solidity: event MembershipChanged(address indexed from, address indexed to)
func (_UnfuckCore *UnfuckCoreFilterer) FilterMembershipChanged(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*UnfuckCoreMembershipChangedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UnfuckCore.contract.FilterLogs(opts, "MembershipChanged", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &UnfuckCoreMembershipChangedIterator{contract: _UnfuckCore.contract, event: "MembershipChanged", logs: logs, sub: sub}, nil
}

// WatchMembershipChanged is a free log subscription operation binding the contract event 0x3331f695d86f449bec7d3bdf70eefa8b100f9dd3b1699e3ad142d67bd846840b.
//
// Solidity: event MembershipChanged(address indexed from, address indexed to)
func (_UnfuckCore *UnfuckCoreFilterer) WatchMembershipChanged(opts *bind.WatchOpts, sink chan<- *UnfuckCoreMembershipChanged, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _UnfuckCore.contract.WatchLogs(opts, "MembershipChanged", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnfuckCoreMembershipChanged)
				if err := _UnfuckCore.contract.UnpackLog(event, "MembershipChanged", log); err != nil {
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

// ParseMembershipChanged is a log parse operation binding the contract event 0x3331f695d86f449bec7d3bdf70eefa8b100f9dd3b1699e3ad142d67bd846840b.
//
// Solidity: event MembershipChanged(address indexed from, address indexed to)
func (_UnfuckCore *UnfuckCoreFilterer) ParseMembershipChanged(log types.Log) (*UnfuckCoreMembershipChanged, error) {
	event := new(UnfuckCoreMembershipChanged)
	if err := _UnfuckCore.contract.UnpackLog(event, "MembershipChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// UnfuckCoreOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the UnfuckCore contract.
type UnfuckCoreOwnershipTransferredIterator struct {
	Event *UnfuckCoreOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *UnfuckCoreOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(UnfuckCoreOwnershipTransferred)
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
		it.Event = new(UnfuckCoreOwnershipTransferred)
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
func (it *UnfuckCoreOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *UnfuckCoreOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// UnfuckCoreOwnershipTransferred represents a OwnershipTransferred event raised by the UnfuckCore contract.
type UnfuckCoreOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnfuckCore *UnfuckCoreFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*UnfuckCoreOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UnfuckCore.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &UnfuckCoreOwnershipTransferredIterator{contract: _UnfuckCore.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_UnfuckCore *UnfuckCoreFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *UnfuckCoreOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _UnfuckCore.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(UnfuckCoreOwnershipTransferred)
				if err := _UnfuckCore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_UnfuckCore *UnfuckCoreFilterer) ParseOwnershipTransferred(log types.Log) (*UnfuckCoreOwnershipTransferred, error) {
	event := new(UnfuckCoreOwnershipTransferred)
	if err := _UnfuckCore.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
