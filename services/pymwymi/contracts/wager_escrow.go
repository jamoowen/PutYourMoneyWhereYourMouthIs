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
	_ = abi.ConvertType
)

// WagerEscrowMetaData contains all meta data concerning the WagerEscrow contract.
var WagerEscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_basisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"COMMISSION_BASIS_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PARTICIPANTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptWager\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimRefund\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimWinnings\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"commissionBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createWager\",\"inputs\":[{\"name\":\"_participants\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deleteCompletedWagers\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getAllWagers\",\"inputs\":[],\"outputs\":[{\"name\":\"ids\",\"type\":\"uint256[]\",\"internalType\":\"uint256[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParticipant\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hasClaimed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getSupportedTokens\",\"inputs\":[],\"outputs\":[{\"name\":\"tokens\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"erc20Contracts\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getWager\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWagerEscrow.WagerStatus\"},{\"name\":\"winner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"participantAddresses\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setCommissionBasisPoints\",\"inputs\":[{\"name\":\"_basisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWagerStatus\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumWagerEscrow.WagerStatus\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWinner\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"winner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportedTokensArray\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawCommission\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"Claim\",\"inputs\":[{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CompletedWagersDeleted\",\"inputs\":[{\"name\":\"ids\",\"type\":\"uint256[]\",\"indexed\":true,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WagerAccepted\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WagerCreated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WagerDeleted\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WagerStatusUpdated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumWagerEscrow.WagerStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WinnerSet\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"winner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// WagerEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use WagerEscrowMetaData.ABI instead.
var WagerEscrowABI = WagerEscrowMetaData.ABI

// WagerEscrow is an auto generated Go binding around an Ethereum contract.
type WagerEscrow struct {
	WagerEscrowCaller     // Read-only binding to the contract
	WagerEscrowTransactor // Write-only binding to the contract
	WagerEscrowFilterer   // Log filterer for contract events
}

// WagerEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type WagerEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WagerEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type WagerEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WagerEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type WagerEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// WagerEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type WagerEscrowSession struct {
	Contract     *WagerEscrow      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// WagerEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type WagerEscrowCallerSession struct {
	Contract *WagerEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// WagerEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type WagerEscrowTransactorSession struct {
	Contract     *WagerEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// WagerEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type WagerEscrowRaw struct {
	Contract *WagerEscrow // Generic contract binding to access the raw methods on
}

// WagerEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type WagerEscrowCallerRaw struct {
	Contract *WagerEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// WagerEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type WagerEscrowTransactorRaw struct {
	Contract *WagerEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewWagerEscrow creates a new instance of WagerEscrow, bound to a specific deployed contract.
func NewWagerEscrow(address common.Address, backend bind.ContractBackend) (*WagerEscrow, error) {
	contract, err := bindWagerEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &WagerEscrow{WagerEscrowCaller: WagerEscrowCaller{contract: contract}, WagerEscrowTransactor: WagerEscrowTransactor{contract: contract}, WagerEscrowFilterer: WagerEscrowFilterer{contract: contract}}, nil
}

// NewWagerEscrowCaller creates a new read-only instance of WagerEscrow, bound to a specific deployed contract.
func NewWagerEscrowCaller(address common.Address, caller bind.ContractCaller) (*WagerEscrowCaller, error) {
	contract, err := bindWagerEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &WagerEscrowCaller{contract: contract}, nil
}

// NewWagerEscrowTransactor creates a new write-only instance of WagerEscrow, bound to a specific deployed contract.
func NewWagerEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*WagerEscrowTransactor, error) {
	contract, err := bindWagerEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &WagerEscrowTransactor{contract: contract}, nil
}

// NewWagerEscrowFilterer creates a new log filterer instance of WagerEscrow, bound to a specific deployed contract.
func NewWagerEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*WagerEscrowFilterer, error) {
	contract, err := bindWagerEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &WagerEscrowFilterer{contract: contract}, nil
}

// bindWagerEscrow binds a generic wrapper to an already deployed contract.
func bindWagerEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := WagerEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WagerEscrow *WagerEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WagerEscrow.Contract.WagerEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WagerEscrow *WagerEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WagerEscrow.Contract.WagerEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WagerEscrow *WagerEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WagerEscrow.Contract.WagerEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_WagerEscrow *WagerEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _WagerEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_WagerEscrow *WagerEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WagerEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_WagerEscrow *WagerEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _WagerEscrow.Contract.contract.Transact(opts, method, params...)
}

// COMMISSIONBASISPOINTS is a free data retrieval call binding the contract method 0xde1f6ea8.
//
// Solidity: function COMMISSION_BASIS_POINTS() view returns(uint256)
func (_WagerEscrow *WagerEscrowCaller) COMMISSIONBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WagerEscrow.contract.Call(opts, &out, "COMMISSION_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMISSIONBASISPOINTS is a free data retrieval call binding the contract method 0xde1f6ea8.
//
// Solidity: function COMMISSION_BASIS_POINTS() view returns(uint256)
func (_WagerEscrow *WagerEscrowSession) COMMISSIONBASISPOINTS() (*big.Int, error) {
	return _WagerEscrow.Contract.COMMISSIONBASISPOINTS(&_WagerEscrow.CallOpts)
}

// COMMISSIONBASISPOINTS is a free data retrieval call binding the contract method 0xde1f6ea8.
//
// Solidity: function COMMISSION_BASIS_POINTS() view returns(uint256)
func (_WagerEscrow *WagerEscrowCallerSession) COMMISSIONBASISPOINTS() (*big.Int, error) {
	return _WagerEscrow.Contract.COMMISSIONBASISPOINTS(&_WagerEscrow.CallOpts)
}

// MAXPARTICIPANTS is a free data retrieval call binding the contract method 0xf3baf070.
//
// Solidity: function MAX_PARTICIPANTS() view returns(uint256)
func (_WagerEscrow *WagerEscrowCaller) MAXPARTICIPANTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _WagerEscrow.contract.Call(opts, &out, "MAX_PARTICIPANTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPARTICIPANTS is a free data retrieval call binding the contract method 0xf3baf070.
//
// Solidity: function MAX_PARTICIPANTS() view returns(uint256)
func (_WagerEscrow *WagerEscrowSession) MAXPARTICIPANTS() (*big.Int, error) {
	return _WagerEscrow.Contract.MAXPARTICIPANTS(&_WagerEscrow.CallOpts)
}

// MAXPARTICIPANTS is a free data retrieval call binding the contract method 0xf3baf070.
//
// Solidity: function MAX_PARTICIPANTS() view returns(uint256)
func (_WagerEscrow *WagerEscrowCallerSession) MAXPARTICIPANTS() (*big.Int, error) {
	return _WagerEscrow.Contract.MAXPARTICIPANTS(&_WagerEscrow.CallOpts)
}

// CommissionBalances is a free data retrieval call binding the contract method 0xb2642987.
//
// Solidity: function commissionBalances(address ) view returns(uint256)
func (_WagerEscrow *WagerEscrowCaller) CommissionBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _WagerEscrow.contract.Call(opts, &out, "commissionBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommissionBalances is a free data retrieval call binding the contract method 0xb2642987.
//
// Solidity: function commissionBalances(address ) view returns(uint256)
func (_WagerEscrow *WagerEscrowSession) CommissionBalances(arg0 common.Address) (*big.Int, error) {
	return _WagerEscrow.Contract.CommissionBalances(&_WagerEscrow.CallOpts, arg0)
}

// CommissionBalances is a free data retrieval call binding the contract method 0xb2642987.
//
// Solidity: function commissionBalances(address ) view returns(uint256)
func (_WagerEscrow *WagerEscrowCallerSession) CommissionBalances(arg0 common.Address) (*big.Int, error) {
	return _WagerEscrow.Contract.CommissionBalances(&_WagerEscrow.CallOpts, arg0)
}

// GetAllWagers is a free data retrieval call binding the contract method 0x7fed58c7.
//
// Solidity: function getAllWagers() view returns(uint256[] ids)
func (_WagerEscrow *WagerEscrowCaller) GetAllWagers(opts *bind.CallOpts) ([]*big.Int, error) {
	var out []interface{}
	err := _WagerEscrow.contract.Call(opts, &out, "getAllWagers")

	if err != nil {
		return *new([]*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new([]*big.Int)).(*[]*big.Int)

	return out0, err

}

// GetAllWagers is a free data retrieval call binding the contract method 0x7fed58c7.
//
// Solidity: function getAllWagers() view returns(uint256[] ids)
func (_WagerEscrow *WagerEscrowSession) GetAllWagers() ([]*big.Int, error) {
	return _WagerEscrow.Contract.GetAllWagers(&_WagerEscrow.CallOpts)
}

// GetAllWagers is a free data retrieval call binding the contract method 0x7fed58c7.
//
// Solidity: function getAllWagers() view returns(uint256[] ids)
func (_WagerEscrow *WagerEscrowCallerSession) GetAllWagers() ([]*big.Int, error) {
	return _WagerEscrow.Contract.GetAllWagers(&_WagerEscrow.CallOpts)
}

// GetParticipant is a free data retrieval call binding the contract method 0x35f3ad7a.
//
// Solidity: function getParticipant(uint256 id, address user) view returns(address walletAddress, uint256 stake, bool hasClaimed)
func (_WagerEscrow *WagerEscrowCaller) GetParticipant(opts *bind.CallOpts, id *big.Int, user common.Address) (struct {
	WalletAddress common.Address
	Stake         *big.Int
	HasClaimed    bool
}, error) {
	var out []interface{}
	err := _WagerEscrow.contract.Call(opts, &out, "getParticipant", id, user)

	outstruct := new(struct {
		WalletAddress common.Address
		Stake         *big.Int
		HasClaimed    bool
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.WalletAddress = *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	outstruct.Stake = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.HasClaimed = *abi.ConvertType(out[2], new(bool)).(*bool)

	return *outstruct, err

}

// GetParticipant is a free data retrieval call binding the contract method 0x35f3ad7a.
//
// Solidity: function getParticipant(uint256 id, address user) view returns(address walletAddress, uint256 stake, bool hasClaimed)
func (_WagerEscrow *WagerEscrowSession) GetParticipant(id *big.Int, user common.Address) (struct {
	WalletAddress common.Address
	Stake         *big.Int
	HasClaimed    bool
}, error) {
	return _WagerEscrow.Contract.GetParticipant(&_WagerEscrow.CallOpts, id, user)
}

// GetParticipant is a free data retrieval call binding the contract method 0x35f3ad7a.
//
// Solidity: function getParticipant(uint256 id, address user) view returns(address walletAddress, uint256 stake, bool hasClaimed)
func (_WagerEscrow *WagerEscrowCallerSession) GetParticipant(id *big.Int, user common.Address) (struct {
	WalletAddress common.Address
	Stake         *big.Int
	HasClaimed    bool
}, error) {
	return _WagerEscrow.Contract.GetParticipant(&_WagerEscrow.CallOpts, id, user)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens, address[] erc20Contracts)
func (_WagerEscrow *WagerEscrowCaller) GetSupportedTokens(opts *bind.CallOpts) (struct {
	Tokens         []common.Address
	Erc20Contracts []common.Address
}, error) {
	var out []interface{}
	err := _WagerEscrow.contract.Call(opts, &out, "getSupportedTokens")

	outstruct := new(struct {
		Tokens         []common.Address
		Erc20Contracts []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Tokens = *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)
	outstruct.Erc20Contracts = *abi.ConvertType(out[1], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens, address[] erc20Contracts)
func (_WagerEscrow *WagerEscrowSession) GetSupportedTokens() (struct {
	Tokens         []common.Address
	Erc20Contracts []common.Address
}, error) {
	return _WagerEscrow.Contract.GetSupportedTokens(&_WagerEscrow.CallOpts)
}

// GetSupportedTokens is a free data retrieval call binding the contract method 0xd3c7c2c7.
//
// Solidity: function getSupportedTokens() view returns(address[] tokens, address[] erc20Contracts)
func (_WagerEscrow *WagerEscrowCallerSession) GetSupportedTokens() (struct {
	Tokens         []common.Address
	Erc20Contracts []common.Address
}, error) {
	return _WagerEscrow.Contract.GetSupportedTokens(&_WagerEscrow.CallOpts)
}

// GetWager is a free data retrieval call binding the contract method 0x7a2756f2.
//
// Solidity: function getWager(uint256 id) view returns(uint8 status, address winner, uint256 totalStake, uint256 requiredStake, address[] participantAddresses)
func (_WagerEscrow *WagerEscrowCaller) GetWager(opts *bind.CallOpts, id *big.Int) (struct {
	Status               uint8
	Winner               common.Address
	TotalStake           *big.Int
	RequiredStake        *big.Int
	ParticipantAddresses []common.Address
}, error) {
	var out []interface{}
	err := _WagerEscrow.contract.Call(opts, &out, "getWager", id)

	outstruct := new(struct {
		Status               uint8
		Winner               common.Address
		TotalStake           *big.Int
		RequiredStake        *big.Int
		ParticipantAddresses []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Winner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TotalStake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RequiredStake = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ParticipantAddresses = *abi.ConvertType(out[4], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetWager is a free data retrieval call binding the contract method 0x7a2756f2.
//
// Solidity: function getWager(uint256 id) view returns(uint8 status, address winner, uint256 totalStake, uint256 requiredStake, address[] participantAddresses)
func (_WagerEscrow *WagerEscrowSession) GetWager(id *big.Int) (struct {
	Status               uint8
	Winner               common.Address
	TotalStake           *big.Int
	RequiredStake        *big.Int
	ParticipantAddresses []common.Address
}, error) {
	return _WagerEscrow.Contract.GetWager(&_WagerEscrow.CallOpts, id)
}

// GetWager is a free data retrieval call binding the contract method 0x7a2756f2.
//
// Solidity: function getWager(uint256 id) view returns(uint8 status, address winner, uint256 totalStake, uint256 requiredStake, address[] participantAddresses)
func (_WagerEscrow *WagerEscrowCallerSession) GetWager(id *big.Int) (struct {
	Status               uint8
	Winner               common.Address
	TotalStake           *big.Int
	RequiredStake        *big.Int
	ParticipantAddresses []common.Address
}, error) {
	return _WagerEscrow.Contract.GetWager(&_WagerEscrow.CallOpts, id)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WagerEscrow *WagerEscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _WagerEscrow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WagerEscrow *WagerEscrowSession) Owner() (common.Address, error) {
	return _WagerEscrow.Contract.Owner(&_WagerEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_WagerEscrow *WagerEscrowCallerSession) Owner() (common.Address, error) {
	return _WagerEscrow.Contract.Owner(&_WagerEscrow.CallOpts)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(address)
func (_WagerEscrow *WagerEscrowCaller) SupportedTokens(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _WagerEscrow.contract.Call(opts, &out, "supportedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(address)
func (_WagerEscrow *WagerEscrowSession) SupportedTokens(arg0 common.Address) (common.Address, error) {
	return _WagerEscrow.Contract.SupportedTokens(&_WagerEscrow.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(address)
func (_WagerEscrow *WagerEscrowCallerSession) SupportedTokens(arg0 common.Address) (common.Address, error) {
	return _WagerEscrow.Contract.SupportedTokens(&_WagerEscrow.CallOpts, arg0)
}

// SupportedTokensArray is a free data retrieval call binding the contract method 0x2e9b3201.
//
// Solidity: function supportedTokensArray(uint256 ) view returns(address)
func (_WagerEscrow *WagerEscrowCaller) SupportedTokensArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _WagerEscrow.contract.Call(opts, &out, "supportedTokensArray", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportedTokensArray is a free data retrieval call binding the contract method 0x2e9b3201.
//
// Solidity: function supportedTokensArray(uint256 ) view returns(address)
func (_WagerEscrow *WagerEscrowSession) SupportedTokensArray(arg0 *big.Int) (common.Address, error) {
	return _WagerEscrow.Contract.SupportedTokensArray(&_WagerEscrow.CallOpts, arg0)
}

// SupportedTokensArray is a free data retrieval call binding the contract method 0x2e9b3201.
//
// Solidity: function supportedTokensArray(uint256 ) view returns(address)
func (_WagerEscrow *WagerEscrowCallerSession) SupportedTokensArray(arg0 *big.Int) (common.Address, error) {
	return _WagerEscrow.Contract.SupportedTokensArray(&_WagerEscrow.CallOpts, arg0)
}

// AcceptWager is a paid mutator transaction binding the contract method 0xb12b3f91.
//
// Solidity: function acceptWager(uint256 id, uint256 _stake, address _token) returns()
func (_WagerEscrow *WagerEscrowTransactor) AcceptWager(opts *bind.TransactOpts, id *big.Int, _stake *big.Int, _token common.Address) (*types.Transaction, error) {
	return _WagerEscrow.contract.Transact(opts, "acceptWager", id, _stake, _token)
}

// AcceptWager is a paid mutator transaction binding the contract method 0xb12b3f91.
//
// Solidity: function acceptWager(uint256 id, uint256 _stake, address _token) returns()
func (_WagerEscrow *WagerEscrowSession) AcceptWager(id *big.Int, _stake *big.Int, _token common.Address) (*types.Transaction, error) {
	return _WagerEscrow.Contract.AcceptWager(&_WagerEscrow.TransactOpts, id, _stake, _token)
}

// AcceptWager is a paid mutator transaction binding the contract method 0xb12b3f91.
//
// Solidity: function acceptWager(uint256 id, uint256 _stake, address _token) returns()
func (_WagerEscrow *WagerEscrowTransactorSession) AcceptWager(id *big.Int, _stake *big.Int, _token common.Address) (*types.Transaction, error) {
	return _WagerEscrow.Contract.AcceptWager(&_WagerEscrow.TransactOpts, id, _stake, _token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address _token) returns()
func (_WagerEscrow *WagerEscrowTransactor) AddToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _WagerEscrow.contract.Transact(opts, "addToken", _token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address _token) returns()
func (_WagerEscrow *WagerEscrowSession) AddToken(_token common.Address) (*types.Transaction, error) {
	return _WagerEscrow.Contract.AddToken(&_WagerEscrow.TransactOpts, _token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address _token) returns()
func (_WagerEscrow *WagerEscrowTransactorSession) AddToken(_token common.Address) (*types.Transaction, error) {
	return _WagerEscrow.Contract.AddToken(&_WagerEscrow.TransactOpts, _token)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 id) returns()
func (_WagerEscrow *WagerEscrowTransactor) ClaimRefund(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _WagerEscrow.contract.Transact(opts, "claimRefund", id)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 id) returns()
func (_WagerEscrow *WagerEscrowSession) ClaimRefund(id *big.Int) (*types.Transaction, error) {
	return _WagerEscrow.Contract.ClaimRefund(&_WagerEscrow.TransactOpts, id)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 id) returns()
func (_WagerEscrow *WagerEscrowTransactorSession) ClaimRefund(id *big.Int) (*types.Transaction, error) {
	return _WagerEscrow.Contract.ClaimRefund(&_WagerEscrow.TransactOpts, id)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x677bd9ff.
//
// Solidity: function claimWinnings(uint256 id) returns()
func (_WagerEscrow *WagerEscrowTransactor) ClaimWinnings(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _WagerEscrow.contract.Transact(opts, "claimWinnings", id)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x677bd9ff.
//
// Solidity: function claimWinnings(uint256 id) returns()
func (_WagerEscrow *WagerEscrowSession) ClaimWinnings(id *big.Int) (*types.Transaction, error) {
	return _WagerEscrow.Contract.ClaimWinnings(&_WagerEscrow.TransactOpts, id)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x677bd9ff.
//
// Solidity: function claimWinnings(uint256 id) returns()
func (_WagerEscrow *WagerEscrowTransactorSession) ClaimWinnings(id *big.Int) (*types.Transaction, error) {
	return _WagerEscrow.Contract.ClaimWinnings(&_WagerEscrow.TransactOpts, id)
}

// CreateWager is a paid mutator transaction binding the contract method 0x05385fe9.
//
// Solidity: function createWager(address[] _participants, uint256 _stake, address _token) returns(uint256)
func (_WagerEscrow *WagerEscrowTransactor) CreateWager(opts *bind.TransactOpts, _participants []common.Address, _stake *big.Int, _token common.Address) (*types.Transaction, error) {
	return _WagerEscrow.contract.Transact(opts, "createWager", _participants, _stake, _token)
}

// CreateWager is a paid mutator transaction binding the contract method 0x05385fe9.
//
// Solidity: function createWager(address[] _participants, uint256 _stake, address _token) returns(uint256)
func (_WagerEscrow *WagerEscrowSession) CreateWager(_participants []common.Address, _stake *big.Int, _token common.Address) (*types.Transaction, error) {
	return _WagerEscrow.Contract.CreateWager(&_WagerEscrow.TransactOpts, _participants, _stake, _token)
}

// CreateWager is a paid mutator transaction binding the contract method 0x05385fe9.
//
// Solidity: function createWager(address[] _participants, uint256 _stake, address _token) returns(uint256)
func (_WagerEscrow *WagerEscrowTransactorSession) CreateWager(_participants []common.Address, _stake *big.Int, _token common.Address) (*types.Transaction, error) {
	return _WagerEscrow.Contract.CreateWager(&_WagerEscrow.TransactOpts, _participants, _stake, _token)
}

// DeleteCompletedWagers is a paid mutator transaction binding the contract method 0x3dc44c0b.
//
// Solidity: function deleteCompletedWagers() returns()
func (_WagerEscrow *WagerEscrowTransactor) DeleteCompletedWagers(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WagerEscrow.contract.Transact(opts, "deleteCompletedWagers")
}

// DeleteCompletedWagers is a paid mutator transaction binding the contract method 0x3dc44c0b.
//
// Solidity: function deleteCompletedWagers() returns()
func (_WagerEscrow *WagerEscrowSession) DeleteCompletedWagers() (*types.Transaction, error) {
	return _WagerEscrow.Contract.DeleteCompletedWagers(&_WagerEscrow.TransactOpts)
}

// DeleteCompletedWagers is a paid mutator transaction binding the contract method 0x3dc44c0b.
//
// Solidity: function deleteCompletedWagers() returns()
func (_WagerEscrow *WagerEscrowTransactorSession) DeleteCompletedWagers() (*types.Transaction, error) {
	return _WagerEscrow.Contract.DeleteCompletedWagers(&_WagerEscrow.TransactOpts)
}

// SetCommissionBasisPoints is a paid mutator transaction binding the contract method 0xf00f6921.
//
// Solidity: function setCommissionBasisPoints(uint256 _basisPoints) returns()
func (_WagerEscrow *WagerEscrowTransactor) SetCommissionBasisPoints(opts *bind.TransactOpts, _basisPoints *big.Int) (*types.Transaction, error) {
	return _WagerEscrow.contract.Transact(opts, "setCommissionBasisPoints", _basisPoints)
}

// SetCommissionBasisPoints is a paid mutator transaction binding the contract method 0xf00f6921.
//
// Solidity: function setCommissionBasisPoints(uint256 _basisPoints) returns()
func (_WagerEscrow *WagerEscrowSession) SetCommissionBasisPoints(_basisPoints *big.Int) (*types.Transaction, error) {
	return _WagerEscrow.Contract.SetCommissionBasisPoints(&_WagerEscrow.TransactOpts, _basisPoints)
}

// SetCommissionBasisPoints is a paid mutator transaction binding the contract method 0xf00f6921.
//
// Solidity: function setCommissionBasisPoints(uint256 _basisPoints) returns()
func (_WagerEscrow *WagerEscrowTransactorSession) SetCommissionBasisPoints(_basisPoints *big.Int) (*types.Transaction, error) {
	return _WagerEscrow.Contract.SetCommissionBasisPoints(&_WagerEscrow.TransactOpts, _basisPoints)
}

// SetWagerStatus is a paid mutator transaction binding the contract method 0xf84cc579.
//
// Solidity: function setWagerStatus(uint256 id, uint8 status) returns()
func (_WagerEscrow *WagerEscrowTransactor) SetWagerStatus(opts *bind.TransactOpts, id *big.Int, status uint8) (*types.Transaction, error) {
	return _WagerEscrow.contract.Transact(opts, "setWagerStatus", id, status)
}

// SetWagerStatus is a paid mutator transaction binding the contract method 0xf84cc579.
//
// Solidity: function setWagerStatus(uint256 id, uint8 status) returns()
func (_WagerEscrow *WagerEscrowSession) SetWagerStatus(id *big.Int, status uint8) (*types.Transaction, error) {
	return _WagerEscrow.Contract.SetWagerStatus(&_WagerEscrow.TransactOpts, id, status)
}

// SetWagerStatus is a paid mutator transaction binding the contract method 0xf84cc579.
//
// Solidity: function setWagerStatus(uint256 id, uint8 status) returns()
func (_WagerEscrow *WagerEscrowTransactorSession) SetWagerStatus(id *big.Int, status uint8) (*types.Transaction, error) {
	return _WagerEscrow.Contract.SetWagerStatus(&_WagerEscrow.TransactOpts, id, status)
}

// SetWinner is a paid mutator transaction binding the contract method 0x9c623683.
//
// Solidity: function setWinner(uint256 id, address winner) returns()
func (_WagerEscrow *WagerEscrowTransactor) SetWinner(opts *bind.TransactOpts, id *big.Int, winner common.Address) (*types.Transaction, error) {
	return _WagerEscrow.contract.Transact(opts, "setWinner", id, winner)
}

// SetWinner is a paid mutator transaction binding the contract method 0x9c623683.
//
// Solidity: function setWinner(uint256 id, address winner) returns()
func (_WagerEscrow *WagerEscrowSession) SetWinner(id *big.Int, winner common.Address) (*types.Transaction, error) {
	return _WagerEscrow.Contract.SetWinner(&_WagerEscrow.TransactOpts, id, winner)
}

// SetWinner is a paid mutator transaction binding the contract method 0x9c623683.
//
// Solidity: function setWinner(uint256 id, address winner) returns()
func (_WagerEscrow *WagerEscrowTransactorSession) SetWinner(id *big.Int, winner common.Address) (*types.Transaction, error) {
	return _WagerEscrow.Contract.SetWinner(&_WagerEscrow.TransactOpts, id, winner)
}

// WithdrawCommission is a paid mutator transaction binding the contract method 0x16c58d04.
//
// Solidity: function withdrawCommission(address _token) returns()
func (_WagerEscrow *WagerEscrowTransactor) WithdrawCommission(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _WagerEscrow.contract.Transact(opts, "withdrawCommission", _token)
}

// WithdrawCommission is a paid mutator transaction binding the contract method 0x16c58d04.
//
// Solidity: function withdrawCommission(address _token) returns()
func (_WagerEscrow *WagerEscrowSession) WithdrawCommission(_token common.Address) (*types.Transaction, error) {
	return _WagerEscrow.Contract.WithdrawCommission(&_WagerEscrow.TransactOpts, _token)
}

// WithdrawCommission is a paid mutator transaction binding the contract method 0x16c58d04.
//
// Solidity: function withdrawCommission(address _token) returns()
func (_WagerEscrow *WagerEscrowTransactorSession) WithdrawCommission(_token common.Address) (*types.Transaction, error) {
	return _WagerEscrow.Contract.WithdrawCommission(&_WagerEscrow.TransactOpts, _token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WagerEscrow *WagerEscrowTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _WagerEscrow.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WagerEscrow *WagerEscrowSession) Receive() (*types.Transaction, error) {
	return _WagerEscrow.Contract.Receive(&_WagerEscrow.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_WagerEscrow *WagerEscrowTransactorSession) Receive() (*types.Transaction, error) {
	return _WagerEscrow.Contract.Receive(&_WagerEscrow.TransactOpts)
}

// WagerEscrowClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the WagerEscrow contract.
type WagerEscrowClaimIterator struct {
	Event *WagerEscrowClaim // Event containing the contract specifics and raw log

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
func (it *WagerEscrowClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WagerEscrowClaim)
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
		it.Event = new(WagerEscrowClaim)
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
func (it *WagerEscrowClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WagerEscrowClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WagerEscrowClaim represents a Claim event raised by the WagerEscrow contract.
type WagerEscrowClaim struct {
	Participant common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed participant, uint256 amount)
func (_WagerEscrow *WagerEscrowFilterer) FilterClaim(opts *bind.FilterOpts, participant []common.Address) (*WagerEscrowClaimIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _WagerEscrow.contract.FilterLogs(opts, "Claim", participantRule)
	if err != nil {
		return nil, err
	}
	return &WagerEscrowClaimIterator{contract: _WagerEscrow.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed participant, uint256 amount)
func (_WagerEscrow *WagerEscrowFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *WagerEscrowClaim, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _WagerEscrow.contract.WatchLogs(opts, "Claim", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WagerEscrowClaim)
				if err := _WagerEscrow.contract.UnpackLog(event, "Claim", log); err != nil {
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

// ParseClaim is a log parse operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed participant, uint256 amount)
func (_WagerEscrow *WagerEscrowFilterer) ParseClaim(log types.Log) (*WagerEscrowClaim, error) {
	event := new(WagerEscrowClaim)
	if err := _WagerEscrow.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WagerEscrowCompletedWagersDeletedIterator is returned from FilterCompletedWagersDeleted and is used to iterate over the raw logs and unpacked data for CompletedWagersDeleted events raised by the WagerEscrow contract.
type WagerEscrowCompletedWagersDeletedIterator struct {
	Event *WagerEscrowCompletedWagersDeleted // Event containing the contract specifics and raw log

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
func (it *WagerEscrowCompletedWagersDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WagerEscrowCompletedWagersDeleted)
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
		it.Event = new(WagerEscrowCompletedWagersDeleted)
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
func (it *WagerEscrowCompletedWagersDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WagerEscrowCompletedWagersDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WagerEscrowCompletedWagersDeleted represents a CompletedWagersDeleted event raised by the WagerEscrow contract.
type WagerEscrowCompletedWagersDeleted struct {
	Ids []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCompletedWagersDeleted is a free log retrieval operation binding the contract event 0x23f742627f497d097f0d54e28cb5c46c4e0d78818c0ba173cce3abf9c1119a74.
//
// Solidity: event CompletedWagersDeleted(uint256[] indexed ids)
func (_WagerEscrow *WagerEscrowFilterer) FilterCompletedWagersDeleted(opts *bind.FilterOpts, ids [][]*big.Int) (*WagerEscrowCompletedWagersDeletedIterator, error) {

	var idsRule []interface{}
	for _, idsItem := range ids {
		idsRule = append(idsRule, idsItem)
	}

	logs, sub, err := _WagerEscrow.contract.FilterLogs(opts, "CompletedWagersDeleted", idsRule)
	if err != nil {
		return nil, err
	}
	return &WagerEscrowCompletedWagersDeletedIterator{contract: _WagerEscrow.contract, event: "CompletedWagersDeleted", logs: logs, sub: sub}, nil
}

// WatchCompletedWagersDeleted is a free log subscription operation binding the contract event 0x23f742627f497d097f0d54e28cb5c46c4e0d78818c0ba173cce3abf9c1119a74.
//
// Solidity: event CompletedWagersDeleted(uint256[] indexed ids)
func (_WagerEscrow *WagerEscrowFilterer) WatchCompletedWagersDeleted(opts *bind.WatchOpts, sink chan<- *WagerEscrowCompletedWagersDeleted, ids [][]*big.Int) (event.Subscription, error) {

	var idsRule []interface{}
	for _, idsItem := range ids {
		idsRule = append(idsRule, idsItem)
	}

	logs, sub, err := _WagerEscrow.contract.WatchLogs(opts, "CompletedWagersDeleted", idsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WagerEscrowCompletedWagersDeleted)
				if err := _WagerEscrow.contract.UnpackLog(event, "CompletedWagersDeleted", log); err != nil {
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

// ParseCompletedWagersDeleted is a log parse operation binding the contract event 0x23f742627f497d097f0d54e28cb5c46c4e0d78818c0ba173cce3abf9c1119a74.
//
// Solidity: event CompletedWagersDeleted(uint256[] indexed ids)
func (_WagerEscrow *WagerEscrowFilterer) ParseCompletedWagersDeleted(log types.Log) (*WagerEscrowCompletedWagersDeleted, error) {
	event := new(WagerEscrowCompletedWagersDeleted)
	if err := _WagerEscrow.contract.UnpackLog(event, "CompletedWagersDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WagerEscrowWagerAcceptedIterator is returned from FilterWagerAccepted and is used to iterate over the raw logs and unpacked data for WagerAccepted events raised by the WagerEscrow contract.
type WagerEscrowWagerAcceptedIterator struct {
	Event *WagerEscrowWagerAccepted // Event containing the contract specifics and raw log

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
func (it *WagerEscrowWagerAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WagerEscrowWagerAccepted)
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
		it.Event = new(WagerEscrowWagerAccepted)
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
func (it *WagerEscrowWagerAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WagerEscrowWagerAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WagerEscrowWagerAccepted represents a WagerAccepted event raised by the WagerEscrow contract.
type WagerEscrowWagerAccepted struct {
	Id          *big.Int
	Participant common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterWagerAccepted is a free log retrieval operation binding the contract event 0xb588749a0bc7dad21f5e6d249c0754fdb7b258080ebfeecb29265842c8649555.
//
// Solidity: event WagerAccepted(uint256 indexed id, address indexed participant, uint256 amount)
func (_WagerEscrow *WagerEscrowFilterer) FilterWagerAccepted(opts *bind.FilterOpts, id []*big.Int, participant []common.Address) (*WagerEscrowWagerAcceptedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _WagerEscrow.contract.FilterLogs(opts, "WagerAccepted", idRule, participantRule)
	if err != nil {
		return nil, err
	}
	return &WagerEscrowWagerAcceptedIterator{contract: _WagerEscrow.contract, event: "WagerAccepted", logs: logs, sub: sub}, nil
}

// WatchWagerAccepted is a free log subscription operation binding the contract event 0xb588749a0bc7dad21f5e6d249c0754fdb7b258080ebfeecb29265842c8649555.
//
// Solidity: event WagerAccepted(uint256 indexed id, address indexed participant, uint256 amount)
func (_WagerEscrow *WagerEscrowFilterer) WatchWagerAccepted(opts *bind.WatchOpts, sink chan<- *WagerEscrowWagerAccepted, id []*big.Int, participant []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _WagerEscrow.contract.WatchLogs(opts, "WagerAccepted", idRule, participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WagerEscrowWagerAccepted)
				if err := _WagerEscrow.contract.UnpackLog(event, "WagerAccepted", log); err != nil {
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

// ParseWagerAccepted is a log parse operation binding the contract event 0xb588749a0bc7dad21f5e6d249c0754fdb7b258080ebfeecb29265842c8649555.
//
// Solidity: event WagerAccepted(uint256 indexed id, address indexed participant, uint256 amount)
func (_WagerEscrow *WagerEscrowFilterer) ParseWagerAccepted(log types.Log) (*WagerEscrowWagerAccepted, error) {
	event := new(WagerEscrowWagerAccepted)
	if err := _WagerEscrow.contract.UnpackLog(event, "WagerAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WagerEscrowWagerCreatedIterator is returned from FilterWagerCreated and is used to iterate over the raw logs and unpacked data for WagerCreated events raised by the WagerEscrow contract.
type WagerEscrowWagerCreatedIterator struct {
	Event *WagerEscrowWagerCreated // Event containing the contract specifics and raw log

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
func (it *WagerEscrowWagerCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WagerEscrowWagerCreated)
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
		it.Event = new(WagerEscrowWagerCreated)
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
func (it *WagerEscrowWagerCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WagerEscrowWagerCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WagerEscrowWagerCreated represents a WagerCreated event raised by the WagerEscrow contract.
type WagerEscrowWagerCreated struct {
	Id      *big.Int
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterWagerCreated is a free log retrieval operation binding the contract event 0xec930341041c5ff4a5017383755eb9783d4ad8bf291b661b7af9595a4e11e4df.
//
// Solidity: event WagerCreated(uint256 indexed id, address indexed creator)
func (_WagerEscrow *WagerEscrowFilterer) FilterWagerCreated(opts *bind.FilterOpts, id []*big.Int, creator []common.Address) (*WagerEscrowWagerCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WagerEscrow.contract.FilterLogs(opts, "WagerCreated", idRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &WagerEscrowWagerCreatedIterator{contract: _WagerEscrow.contract, event: "WagerCreated", logs: logs, sub: sub}, nil
}

// WatchWagerCreated is a free log subscription operation binding the contract event 0xec930341041c5ff4a5017383755eb9783d4ad8bf291b661b7af9595a4e11e4df.
//
// Solidity: event WagerCreated(uint256 indexed id, address indexed creator)
func (_WagerEscrow *WagerEscrowFilterer) WatchWagerCreated(opts *bind.WatchOpts, sink chan<- *WagerEscrowWagerCreated, id []*big.Int, creator []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _WagerEscrow.contract.WatchLogs(opts, "WagerCreated", idRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WagerEscrowWagerCreated)
				if err := _WagerEscrow.contract.UnpackLog(event, "WagerCreated", log); err != nil {
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

// ParseWagerCreated is a log parse operation binding the contract event 0xec930341041c5ff4a5017383755eb9783d4ad8bf291b661b7af9595a4e11e4df.
//
// Solidity: event WagerCreated(uint256 indexed id, address indexed creator)
func (_WagerEscrow *WagerEscrowFilterer) ParseWagerCreated(log types.Log) (*WagerEscrowWagerCreated, error) {
	event := new(WagerEscrowWagerCreated)
	if err := _WagerEscrow.contract.UnpackLog(event, "WagerCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WagerEscrowWagerDeletedIterator is returned from FilterWagerDeleted and is used to iterate over the raw logs and unpacked data for WagerDeleted events raised by the WagerEscrow contract.
type WagerEscrowWagerDeletedIterator struct {
	Event *WagerEscrowWagerDeleted // Event containing the contract specifics and raw log

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
func (it *WagerEscrowWagerDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WagerEscrowWagerDeleted)
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
		it.Event = new(WagerEscrowWagerDeleted)
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
func (it *WagerEscrowWagerDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WagerEscrowWagerDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WagerEscrowWagerDeleted represents a WagerDeleted event raised by the WagerEscrow contract.
type WagerEscrowWagerDeleted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterWagerDeleted is a free log retrieval operation binding the contract event 0xeb411d22e3431c2d584b53e3697619caf3965d451ca3f1eddc5a46f66c257ba5.
//
// Solidity: event WagerDeleted(uint256 indexed id)
func (_WagerEscrow *WagerEscrowFilterer) FilterWagerDeleted(opts *bind.FilterOpts, id []*big.Int) (*WagerEscrowWagerDeletedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _WagerEscrow.contract.FilterLogs(opts, "WagerDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return &WagerEscrowWagerDeletedIterator{contract: _WagerEscrow.contract, event: "WagerDeleted", logs: logs, sub: sub}, nil
}

// WatchWagerDeleted is a free log subscription operation binding the contract event 0xeb411d22e3431c2d584b53e3697619caf3965d451ca3f1eddc5a46f66c257ba5.
//
// Solidity: event WagerDeleted(uint256 indexed id)
func (_WagerEscrow *WagerEscrowFilterer) WatchWagerDeleted(opts *bind.WatchOpts, sink chan<- *WagerEscrowWagerDeleted, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _WagerEscrow.contract.WatchLogs(opts, "WagerDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WagerEscrowWagerDeleted)
				if err := _WagerEscrow.contract.UnpackLog(event, "WagerDeleted", log); err != nil {
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

// ParseWagerDeleted is a log parse operation binding the contract event 0xeb411d22e3431c2d584b53e3697619caf3965d451ca3f1eddc5a46f66c257ba5.
//
// Solidity: event WagerDeleted(uint256 indexed id)
func (_WagerEscrow *WagerEscrowFilterer) ParseWagerDeleted(log types.Log) (*WagerEscrowWagerDeleted, error) {
	event := new(WagerEscrowWagerDeleted)
	if err := _WagerEscrow.contract.UnpackLog(event, "WagerDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WagerEscrowWagerStatusUpdatedIterator is returned from FilterWagerStatusUpdated and is used to iterate over the raw logs and unpacked data for WagerStatusUpdated events raised by the WagerEscrow contract.
type WagerEscrowWagerStatusUpdatedIterator struct {
	Event *WagerEscrowWagerStatusUpdated // Event containing the contract specifics and raw log

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
func (it *WagerEscrowWagerStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WagerEscrowWagerStatusUpdated)
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
		it.Event = new(WagerEscrowWagerStatusUpdated)
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
func (it *WagerEscrowWagerStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WagerEscrowWagerStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WagerEscrowWagerStatusUpdated represents a WagerStatusUpdated event raised by the WagerEscrow contract.
type WagerEscrowWagerStatusUpdated struct {
	Id     *big.Int
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWagerStatusUpdated is a free log retrieval operation binding the contract event 0x9bd6564623fb7c831af9798510036a9c2d1494d6a70fb5517e31544b00b01eb4.
//
// Solidity: event WagerStatusUpdated(uint256 indexed id, uint8 status)
func (_WagerEscrow *WagerEscrowFilterer) FilterWagerStatusUpdated(opts *bind.FilterOpts, id []*big.Int) (*WagerEscrowWagerStatusUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _WagerEscrow.contract.FilterLogs(opts, "WagerStatusUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &WagerEscrowWagerStatusUpdatedIterator{contract: _WagerEscrow.contract, event: "WagerStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchWagerStatusUpdated is a free log subscription operation binding the contract event 0x9bd6564623fb7c831af9798510036a9c2d1494d6a70fb5517e31544b00b01eb4.
//
// Solidity: event WagerStatusUpdated(uint256 indexed id, uint8 status)
func (_WagerEscrow *WagerEscrowFilterer) WatchWagerStatusUpdated(opts *bind.WatchOpts, sink chan<- *WagerEscrowWagerStatusUpdated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _WagerEscrow.contract.WatchLogs(opts, "WagerStatusUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WagerEscrowWagerStatusUpdated)
				if err := _WagerEscrow.contract.UnpackLog(event, "WagerStatusUpdated", log); err != nil {
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

// ParseWagerStatusUpdated is a log parse operation binding the contract event 0x9bd6564623fb7c831af9798510036a9c2d1494d6a70fb5517e31544b00b01eb4.
//
// Solidity: event WagerStatusUpdated(uint256 indexed id, uint8 status)
func (_WagerEscrow *WagerEscrowFilterer) ParseWagerStatusUpdated(log types.Log) (*WagerEscrowWagerStatusUpdated, error) {
	event := new(WagerEscrowWagerStatusUpdated)
	if err := _WagerEscrow.contract.UnpackLog(event, "WagerStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// WagerEscrowWinnerSetIterator is returned from FilterWinnerSet and is used to iterate over the raw logs and unpacked data for WinnerSet events raised by the WagerEscrow contract.
type WagerEscrowWinnerSetIterator struct {
	Event *WagerEscrowWinnerSet // Event containing the contract specifics and raw log

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
func (it *WagerEscrowWinnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(WagerEscrowWinnerSet)
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
		it.Event = new(WagerEscrowWinnerSet)
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
func (it *WagerEscrowWinnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *WagerEscrowWinnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// WagerEscrowWinnerSet represents a WinnerSet event raised by the WagerEscrow contract.
type WagerEscrowWinnerSet struct {
	Id     *big.Int
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWinnerSet is a free log retrieval operation binding the contract event 0x44fd60df9b05a22172cc43446c2ff1abb74d6c6ce938acd70bd3ee1ea2bf595d.
//
// Solidity: event WinnerSet(uint256 indexed id, address indexed winner)
func (_WagerEscrow *WagerEscrowFilterer) FilterWinnerSet(opts *bind.FilterOpts, id []*big.Int, winner []common.Address) (*WagerEscrowWinnerSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _WagerEscrow.contract.FilterLogs(opts, "WinnerSet", idRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &WagerEscrowWinnerSetIterator{contract: _WagerEscrow.contract, event: "WinnerSet", logs: logs, sub: sub}, nil
}

// WatchWinnerSet is a free log subscription operation binding the contract event 0x44fd60df9b05a22172cc43446c2ff1abb74d6c6ce938acd70bd3ee1ea2bf595d.
//
// Solidity: event WinnerSet(uint256 indexed id, address indexed winner)
func (_WagerEscrow *WagerEscrowFilterer) WatchWinnerSet(opts *bind.WatchOpts, sink chan<- *WagerEscrowWinnerSet, id []*big.Int, winner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _WagerEscrow.contract.WatchLogs(opts, "WinnerSet", idRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(WagerEscrowWinnerSet)
				if err := _WagerEscrow.contract.UnpackLog(event, "WinnerSet", log); err != nil {
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

// ParseWinnerSet is a log parse operation binding the contract event 0x44fd60df9b05a22172cc43446c2ff1abb74d6c6ce938acd70bd3ee1ea2bf595d.
//
// Solidity: event WinnerSet(uint256 indexed id, address indexed winner)
func (_WagerEscrow *WagerEscrowFilterer) ParseWinnerSet(log types.Log) (*WagerEscrowWinnerSet, error) {
	event := new(WagerEscrowWinnerSet)
	if err := _WagerEscrow.contract.UnpackLog(event, "WinnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
