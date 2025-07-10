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

// ChallengeEscrowMetaData contains all meta data concerning the ChallengeEscrow contract.
var ChallengeEscrowMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_basisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"receive\",\"stateMutability\":\"payable\"},{\"type\":\"function\",\"name\":\"COMMISSION_BASIS_POINTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"MAX_PARTICIPANTS\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"acceptChallenge\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"addToken\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimRefund\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"claimWinnings\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"commissionBalances\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"createChallenge\",\"inputs\":[{\"name\":\"_participants\",\"type\":\"address[]\",\"internalType\":\"address[]\"},{\"name\":\"_stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"deleteCompletedChallenges\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getChallenge\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumChallengeEscrow.ChallengeStatus\"},{\"name\":\"winner\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"totalStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"requiredStake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"participantArray\",\"type\":\"address[]\",\"internalType\":\"address[]\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getParticipant\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"user\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"walletAddress\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"stake\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"hasClaimed\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setChallengeStatus\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"internalType\":\"enumChallengeEscrow.ChallengeStatus\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setCommissionBasisPoints\",\"inputs\":[{\"name\":\"_basisPoints\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setWinner\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"winner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"supportedTokens\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"supportedTokensArray\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"withdrawCommission\",\"inputs\":[{\"name\":\"_token\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ChallengeAccepted\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChallengeCreated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"creator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChallengeDeleted\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ChallengeStatusUpdated\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"status\",\"type\":\"uint8\",\"indexed\":false,\"internalType\":\"enumChallengeEscrow.ChallengeStatus\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Claim\",\"inputs\":[{\"name\":\"participant\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"CompletedChallengesDeleted\",\"inputs\":[{\"name\":\"ids\",\"type\":\"uint256[]\",\"indexed\":true,\"internalType\":\"uint256[]\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"WinnerSet\",\"inputs\":[{\"name\":\"id\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"winner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"AlreadyAccepted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"AlreadyClaimed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ChallengeNotActive\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ClaimsNotAllowed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidToken\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MaxParticipantsExceeded\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotAccepted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCancelled\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotCompleted\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotOwner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotParticipant\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"NotWinner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]}]",
}

// ChallengeEscrowABI is the input ABI used to generate the binding from.
// Deprecated: Use ChallengeEscrowMetaData.ABI instead.
var ChallengeEscrowABI = ChallengeEscrowMetaData.ABI

// ChallengeEscrow is an auto generated Go binding around an Ethereum contract.
type ChallengeEscrow struct {
	ChallengeEscrowCaller     // Read-only binding to the contract
	ChallengeEscrowTransactor // Write-only binding to the contract
	ChallengeEscrowFilterer   // Log filterer for contract events
}

// ChallengeEscrowCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChallengeEscrowCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeEscrowTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChallengeEscrowTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeEscrowFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChallengeEscrowFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChallengeEscrowSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChallengeEscrowSession struct {
	Contract     *ChallengeEscrow  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChallengeEscrowCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChallengeEscrowCallerSession struct {
	Contract *ChallengeEscrowCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// ChallengeEscrowTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChallengeEscrowTransactorSession struct {
	Contract     *ChallengeEscrowTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// ChallengeEscrowRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChallengeEscrowRaw struct {
	Contract *ChallengeEscrow // Generic contract binding to access the raw methods on
}

// ChallengeEscrowCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChallengeEscrowCallerRaw struct {
	Contract *ChallengeEscrowCaller // Generic read-only contract binding to access the raw methods on
}

// ChallengeEscrowTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChallengeEscrowTransactorRaw struct {
	Contract *ChallengeEscrowTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChallengeEscrow creates a new instance of ChallengeEscrow, bound to a specific deployed contract.
func NewChallengeEscrow(address common.Address, backend bind.ContractBackend) (*ChallengeEscrow, error) {
	contract, err := bindChallengeEscrow(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChallengeEscrow{ChallengeEscrowCaller: ChallengeEscrowCaller{contract: contract}, ChallengeEscrowTransactor: ChallengeEscrowTransactor{contract: contract}, ChallengeEscrowFilterer: ChallengeEscrowFilterer{contract: contract}}, nil
}

// NewChallengeEscrowCaller creates a new read-only instance of ChallengeEscrow, bound to a specific deployed contract.
func NewChallengeEscrowCaller(address common.Address, caller bind.ContractCaller) (*ChallengeEscrowCaller, error) {
	contract, err := bindChallengeEscrow(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeEscrowCaller{contract: contract}, nil
}

// NewChallengeEscrowTransactor creates a new write-only instance of ChallengeEscrow, bound to a specific deployed contract.
func NewChallengeEscrowTransactor(address common.Address, transactor bind.ContractTransactor) (*ChallengeEscrowTransactor, error) {
	contract, err := bindChallengeEscrow(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChallengeEscrowTransactor{contract: contract}, nil
}

// NewChallengeEscrowFilterer creates a new log filterer instance of ChallengeEscrow, bound to a specific deployed contract.
func NewChallengeEscrowFilterer(address common.Address, filterer bind.ContractFilterer) (*ChallengeEscrowFilterer, error) {
	contract, err := bindChallengeEscrow(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChallengeEscrowFilterer{contract: contract}, nil
}

// bindChallengeEscrow binds a generic wrapper to an already deployed contract.
func bindChallengeEscrow(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChallengeEscrowMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeEscrow *ChallengeEscrowRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeEscrow.Contract.ChallengeEscrowCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeEscrow *ChallengeEscrowRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.ChallengeEscrowTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeEscrow *ChallengeEscrowRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.ChallengeEscrowTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChallengeEscrow *ChallengeEscrowCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChallengeEscrow.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChallengeEscrow *ChallengeEscrowTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChallengeEscrow *ChallengeEscrowTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.contract.Transact(opts, method, params...)
}

// COMMISSIONBASISPOINTS is a free data retrieval call binding the contract method 0xde1f6ea8.
//
// Solidity: function COMMISSION_BASIS_POINTS() view returns(uint256)
func (_ChallengeEscrow *ChallengeEscrowCaller) COMMISSIONBASISPOINTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChallengeEscrow.contract.Call(opts, &out, "COMMISSION_BASIS_POINTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// COMMISSIONBASISPOINTS is a free data retrieval call binding the contract method 0xde1f6ea8.
//
// Solidity: function COMMISSION_BASIS_POINTS() view returns(uint256)
func (_ChallengeEscrow *ChallengeEscrowSession) COMMISSIONBASISPOINTS() (*big.Int, error) {
	return _ChallengeEscrow.Contract.COMMISSIONBASISPOINTS(&_ChallengeEscrow.CallOpts)
}

// COMMISSIONBASISPOINTS is a free data retrieval call binding the contract method 0xde1f6ea8.
//
// Solidity: function COMMISSION_BASIS_POINTS() view returns(uint256)
func (_ChallengeEscrow *ChallengeEscrowCallerSession) COMMISSIONBASISPOINTS() (*big.Int, error) {
	return _ChallengeEscrow.Contract.COMMISSIONBASISPOINTS(&_ChallengeEscrow.CallOpts)
}

// MAXPARTICIPANTS is a free data retrieval call binding the contract method 0xf3baf070.
//
// Solidity: function MAX_PARTICIPANTS() view returns(uint256)
func (_ChallengeEscrow *ChallengeEscrowCaller) MAXPARTICIPANTS(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChallengeEscrow.contract.Call(opts, &out, "MAX_PARTICIPANTS")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MAXPARTICIPANTS is a free data retrieval call binding the contract method 0xf3baf070.
//
// Solidity: function MAX_PARTICIPANTS() view returns(uint256)
func (_ChallengeEscrow *ChallengeEscrowSession) MAXPARTICIPANTS() (*big.Int, error) {
	return _ChallengeEscrow.Contract.MAXPARTICIPANTS(&_ChallengeEscrow.CallOpts)
}

// MAXPARTICIPANTS is a free data retrieval call binding the contract method 0xf3baf070.
//
// Solidity: function MAX_PARTICIPANTS() view returns(uint256)
func (_ChallengeEscrow *ChallengeEscrowCallerSession) MAXPARTICIPANTS() (*big.Int, error) {
	return _ChallengeEscrow.Contract.MAXPARTICIPANTS(&_ChallengeEscrow.CallOpts)
}

// CommissionBalances is a free data retrieval call binding the contract method 0xb2642987.
//
// Solidity: function commissionBalances(address ) view returns(uint256)
func (_ChallengeEscrow *ChallengeEscrowCaller) CommissionBalances(opts *bind.CallOpts, arg0 common.Address) (*big.Int, error) {
	var out []interface{}
	err := _ChallengeEscrow.contract.Call(opts, &out, "commissionBalances", arg0)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// CommissionBalances is a free data retrieval call binding the contract method 0xb2642987.
//
// Solidity: function commissionBalances(address ) view returns(uint256)
func (_ChallengeEscrow *ChallengeEscrowSession) CommissionBalances(arg0 common.Address) (*big.Int, error) {
	return _ChallengeEscrow.Contract.CommissionBalances(&_ChallengeEscrow.CallOpts, arg0)
}

// CommissionBalances is a free data retrieval call binding the contract method 0xb2642987.
//
// Solidity: function commissionBalances(address ) view returns(uint256)
func (_ChallengeEscrow *ChallengeEscrowCallerSession) CommissionBalances(arg0 common.Address) (*big.Int, error) {
	return _ChallengeEscrow.Contract.CommissionBalances(&_ChallengeEscrow.CallOpts, arg0)
}

// GetChallenge is a free data retrieval call binding the contract method 0x1bdd4b74.
//
// Solidity: function getChallenge(uint256 id) view returns(uint8 status, address winner, uint256 totalStake, uint256 requiredStake, address[] participantArray)
func (_ChallengeEscrow *ChallengeEscrowCaller) GetChallenge(opts *bind.CallOpts, id *big.Int) (struct {
	Status           uint8
	Winner           common.Address
	TotalStake       *big.Int
	RequiredStake    *big.Int
	ParticipantArray []common.Address
}, error) {
	var out []interface{}
	err := _ChallengeEscrow.contract.Call(opts, &out, "getChallenge", id)

	outstruct := new(struct {
		Status           uint8
		Winner           common.Address
		TotalStake       *big.Int
		RequiredStake    *big.Int
		ParticipantArray []common.Address
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Status = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.Winner = *abi.ConvertType(out[1], new(common.Address)).(*common.Address)
	outstruct.TotalStake = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.RequiredStake = *abi.ConvertType(out[3], new(*big.Int)).(**big.Int)
	outstruct.ParticipantArray = *abi.ConvertType(out[4], new([]common.Address)).(*[]common.Address)

	return *outstruct, err

}

// GetChallenge is a free data retrieval call binding the contract method 0x1bdd4b74.
//
// Solidity: function getChallenge(uint256 id) view returns(uint8 status, address winner, uint256 totalStake, uint256 requiredStake, address[] participantArray)
func (_ChallengeEscrow *ChallengeEscrowSession) GetChallenge(id *big.Int) (struct {
	Status           uint8
	Winner           common.Address
	TotalStake       *big.Int
	RequiredStake    *big.Int
	ParticipantArray []common.Address
}, error) {
	return _ChallengeEscrow.Contract.GetChallenge(&_ChallengeEscrow.CallOpts, id)
}

// GetChallenge is a free data retrieval call binding the contract method 0x1bdd4b74.
//
// Solidity: function getChallenge(uint256 id) view returns(uint8 status, address winner, uint256 totalStake, uint256 requiredStake, address[] participantArray)
func (_ChallengeEscrow *ChallengeEscrowCallerSession) GetChallenge(id *big.Int) (struct {
	Status           uint8
	Winner           common.Address
	TotalStake       *big.Int
	RequiredStake    *big.Int
	ParticipantArray []common.Address
}, error) {
	return _ChallengeEscrow.Contract.GetChallenge(&_ChallengeEscrow.CallOpts, id)
}

// GetParticipant is a free data retrieval call binding the contract method 0x35f3ad7a.
//
// Solidity: function getParticipant(uint256 id, address user) view returns(address walletAddress, uint256 stake, bool hasClaimed)
func (_ChallengeEscrow *ChallengeEscrowCaller) GetParticipant(opts *bind.CallOpts, id *big.Int, user common.Address) (struct {
	WalletAddress common.Address
	Stake         *big.Int
	HasClaimed    bool
}, error) {
	var out []interface{}
	err := _ChallengeEscrow.contract.Call(opts, &out, "getParticipant", id, user)

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
func (_ChallengeEscrow *ChallengeEscrowSession) GetParticipant(id *big.Int, user common.Address) (struct {
	WalletAddress common.Address
	Stake         *big.Int
	HasClaimed    bool
}, error) {
	return _ChallengeEscrow.Contract.GetParticipant(&_ChallengeEscrow.CallOpts, id, user)
}

// GetParticipant is a free data retrieval call binding the contract method 0x35f3ad7a.
//
// Solidity: function getParticipant(uint256 id, address user) view returns(address walletAddress, uint256 stake, bool hasClaimed)
func (_ChallengeEscrow *ChallengeEscrowCallerSession) GetParticipant(id *big.Int, user common.Address) (struct {
	WalletAddress common.Address
	Stake         *big.Int
	HasClaimed    bool
}, error) {
	return _ChallengeEscrow.Contract.GetParticipant(&_ChallengeEscrow.CallOpts, id, user)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChallengeEscrow *ChallengeEscrowCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ChallengeEscrow.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChallengeEscrow *ChallengeEscrowSession) Owner() (common.Address, error) {
	return _ChallengeEscrow.Contract.Owner(&_ChallengeEscrow.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ChallengeEscrow *ChallengeEscrowCallerSession) Owner() (common.Address, error) {
	return _ChallengeEscrow.Contract.Owner(&_ChallengeEscrow.CallOpts)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(address)
func (_ChallengeEscrow *ChallengeEscrowCaller) SupportedTokens(opts *bind.CallOpts, arg0 common.Address) (common.Address, error) {
	var out []interface{}
	err := _ChallengeEscrow.contract.Call(opts, &out, "supportedTokens", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(address)
func (_ChallengeEscrow *ChallengeEscrowSession) SupportedTokens(arg0 common.Address) (common.Address, error) {
	return _ChallengeEscrow.Contract.SupportedTokens(&_ChallengeEscrow.CallOpts, arg0)
}

// SupportedTokens is a free data retrieval call binding the contract method 0x68c4ac26.
//
// Solidity: function supportedTokens(address ) view returns(address)
func (_ChallengeEscrow *ChallengeEscrowCallerSession) SupportedTokens(arg0 common.Address) (common.Address, error) {
	return _ChallengeEscrow.Contract.SupportedTokens(&_ChallengeEscrow.CallOpts, arg0)
}

// SupportedTokensArray is a free data retrieval call binding the contract method 0x2e9b3201.
//
// Solidity: function supportedTokensArray(uint256 ) view returns(address)
func (_ChallengeEscrow *ChallengeEscrowCaller) SupportedTokensArray(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _ChallengeEscrow.contract.Call(opts, &out, "supportedTokensArray", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// SupportedTokensArray is a free data retrieval call binding the contract method 0x2e9b3201.
//
// Solidity: function supportedTokensArray(uint256 ) view returns(address)
func (_ChallengeEscrow *ChallengeEscrowSession) SupportedTokensArray(arg0 *big.Int) (common.Address, error) {
	return _ChallengeEscrow.Contract.SupportedTokensArray(&_ChallengeEscrow.CallOpts, arg0)
}

// SupportedTokensArray is a free data retrieval call binding the contract method 0x2e9b3201.
//
// Solidity: function supportedTokensArray(uint256 ) view returns(address)
func (_ChallengeEscrow *ChallengeEscrowCallerSession) SupportedTokensArray(arg0 *big.Int) (common.Address, error) {
	return _ChallengeEscrow.Contract.SupportedTokensArray(&_ChallengeEscrow.CallOpts, arg0)
}

// AcceptChallenge is a paid mutator transaction binding the contract method 0x755923ff.
//
// Solidity: function acceptChallenge(uint256 id, uint256 _stake, address _token) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactor) AcceptChallenge(opts *bind.TransactOpts, id *big.Int, _stake *big.Int, _token common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.contract.Transact(opts, "acceptChallenge", id, _stake, _token)
}

// AcceptChallenge is a paid mutator transaction binding the contract method 0x755923ff.
//
// Solidity: function acceptChallenge(uint256 id, uint256 _stake, address _token) returns()
func (_ChallengeEscrow *ChallengeEscrowSession) AcceptChallenge(id *big.Int, _stake *big.Int, _token common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.AcceptChallenge(&_ChallengeEscrow.TransactOpts, id, _stake, _token)
}

// AcceptChallenge is a paid mutator transaction binding the contract method 0x755923ff.
//
// Solidity: function acceptChallenge(uint256 id, uint256 _stake, address _token) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactorSession) AcceptChallenge(id *big.Int, _stake *big.Int, _token common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.AcceptChallenge(&_ChallengeEscrow.TransactOpts, id, _stake, _token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address _token) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactor) AddToken(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.contract.Transact(opts, "addToken", _token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address _token) returns()
func (_ChallengeEscrow *ChallengeEscrowSession) AddToken(_token common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.AddToken(&_ChallengeEscrow.TransactOpts, _token)
}

// AddToken is a paid mutator transaction binding the contract method 0xd48bfca7.
//
// Solidity: function addToken(address _token) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactorSession) AddToken(_token common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.AddToken(&_ChallengeEscrow.TransactOpts, _token)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 id) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactor) ClaimRefund(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _ChallengeEscrow.contract.Transact(opts, "claimRefund", id)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 id) returns()
func (_ChallengeEscrow *ChallengeEscrowSession) ClaimRefund(id *big.Int) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.ClaimRefund(&_ChallengeEscrow.TransactOpts, id)
}

// ClaimRefund is a paid mutator transaction binding the contract method 0x5b7baf64.
//
// Solidity: function claimRefund(uint256 id) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactorSession) ClaimRefund(id *big.Int) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.ClaimRefund(&_ChallengeEscrow.TransactOpts, id)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x677bd9ff.
//
// Solidity: function claimWinnings(uint256 id) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactor) ClaimWinnings(opts *bind.TransactOpts, id *big.Int) (*types.Transaction, error) {
	return _ChallengeEscrow.contract.Transact(opts, "claimWinnings", id)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x677bd9ff.
//
// Solidity: function claimWinnings(uint256 id) returns()
func (_ChallengeEscrow *ChallengeEscrowSession) ClaimWinnings(id *big.Int) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.ClaimWinnings(&_ChallengeEscrow.TransactOpts, id)
}

// ClaimWinnings is a paid mutator transaction binding the contract method 0x677bd9ff.
//
// Solidity: function claimWinnings(uint256 id) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactorSession) ClaimWinnings(id *big.Int) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.ClaimWinnings(&_ChallengeEscrow.TransactOpts, id)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xe7cd6ca2.
//
// Solidity: function createChallenge(address[] _participants, uint256 _stake, address _token) returns(uint256)
func (_ChallengeEscrow *ChallengeEscrowTransactor) CreateChallenge(opts *bind.TransactOpts, _participants []common.Address, _stake *big.Int, _token common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.contract.Transact(opts, "createChallenge", _participants, _stake, _token)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xe7cd6ca2.
//
// Solidity: function createChallenge(address[] _participants, uint256 _stake, address _token) returns(uint256)
func (_ChallengeEscrow *ChallengeEscrowSession) CreateChallenge(_participants []common.Address, _stake *big.Int, _token common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.CreateChallenge(&_ChallengeEscrow.TransactOpts, _participants, _stake, _token)
}

// CreateChallenge is a paid mutator transaction binding the contract method 0xe7cd6ca2.
//
// Solidity: function createChallenge(address[] _participants, uint256 _stake, address _token) returns(uint256)
func (_ChallengeEscrow *ChallengeEscrowTransactorSession) CreateChallenge(_participants []common.Address, _stake *big.Int, _token common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.CreateChallenge(&_ChallengeEscrow.TransactOpts, _participants, _stake, _token)
}

// DeleteCompletedChallenges is a paid mutator transaction binding the contract method 0xdc62f9ef.
//
// Solidity: function deleteCompletedChallenges() returns()
func (_ChallengeEscrow *ChallengeEscrowTransactor) DeleteCompletedChallenges(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeEscrow.contract.Transact(opts, "deleteCompletedChallenges")
}

// DeleteCompletedChallenges is a paid mutator transaction binding the contract method 0xdc62f9ef.
//
// Solidity: function deleteCompletedChallenges() returns()
func (_ChallengeEscrow *ChallengeEscrowSession) DeleteCompletedChallenges() (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.DeleteCompletedChallenges(&_ChallengeEscrow.TransactOpts)
}

// DeleteCompletedChallenges is a paid mutator transaction binding the contract method 0xdc62f9ef.
//
// Solidity: function deleteCompletedChallenges() returns()
func (_ChallengeEscrow *ChallengeEscrowTransactorSession) DeleteCompletedChallenges() (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.DeleteCompletedChallenges(&_ChallengeEscrow.TransactOpts)
}

// SetChallengeStatus is a paid mutator transaction binding the contract method 0x80ce1d60.
//
// Solidity: function setChallengeStatus(uint256 id, uint8 status) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactor) SetChallengeStatus(opts *bind.TransactOpts, id *big.Int, status uint8) (*types.Transaction, error) {
	return _ChallengeEscrow.contract.Transact(opts, "setChallengeStatus", id, status)
}

// SetChallengeStatus is a paid mutator transaction binding the contract method 0x80ce1d60.
//
// Solidity: function setChallengeStatus(uint256 id, uint8 status) returns()
func (_ChallengeEscrow *ChallengeEscrowSession) SetChallengeStatus(id *big.Int, status uint8) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.SetChallengeStatus(&_ChallengeEscrow.TransactOpts, id, status)
}

// SetChallengeStatus is a paid mutator transaction binding the contract method 0x80ce1d60.
//
// Solidity: function setChallengeStatus(uint256 id, uint8 status) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactorSession) SetChallengeStatus(id *big.Int, status uint8) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.SetChallengeStatus(&_ChallengeEscrow.TransactOpts, id, status)
}

// SetCommissionBasisPoints is a paid mutator transaction binding the contract method 0xf00f6921.
//
// Solidity: function setCommissionBasisPoints(uint256 _basisPoints) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactor) SetCommissionBasisPoints(opts *bind.TransactOpts, _basisPoints *big.Int) (*types.Transaction, error) {
	return _ChallengeEscrow.contract.Transact(opts, "setCommissionBasisPoints", _basisPoints)
}

// SetCommissionBasisPoints is a paid mutator transaction binding the contract method 0xf00f6921.
//
// Solidity: function setCommissionBasisPoints(uint256 _basisPoints) returns()
func (_ChallengeEscrow *ChallengeEscrowSession) SetCommissionBasisPoints(_basisPoints *big.Int) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.SetCommissionBasisPoints(&_ChallengeEscrow.TransactOpts, _basisPoints)
}

// SetCommissionBasisPoints is a paid mutator transaction binding the contract method 0xf00f6921.
//
// Solidity: function setCommissionBasisPoints(uint256 _basisPoints) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactorSession) SetCommissionBasisPoints(_basisPoints *big.Int) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.SetCommissionBasisPoints(&_ChallengeEscrow.TransactOpts, _basisPoints)
}

// SetWinner is a paid mutator transaction binding the contract method 0x9c623683.
//
// Solidity: function setWinner(uint256 id, address winner) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactor) SetWinner(opts *bind.TransactOpts, id *big.Int, winner common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.contract.Transact(opts, "setWinner", id, winner)
}

// SetWinner is a paid mutator transaction binding the contract method 0x9c623683.
//
// Solidity: function setWinner(uint256 id, address winner) returns()
func (_ChallengeEscrow *ChallengeEscrowSession) SetWinner(id *big.Int, winner common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.SetWinner(&_ChallengeEscrow.TransactOpts, id, winner)
}

// SetWinner is a paid mutator transaction binding the contract method 0x9c623683.
//
// Solidity: function setWinner(uint256 id, address winner) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactorSession) SetWinner(id *big.Int, winner common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.SetWinner(&_ChallengeEscrow.TransactOpts, id, winner)
}

// WithdrawCommission is a paid mutator transaction binding the contract method 0x16c58d04.
//
// Solidity: function withdrawCommission(address _token) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactor) WithdrawCommission(opts *bind.TransactOpts, _token common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.contract.Transact(opts, "withdrawCommission", _token)
}

// WithdrawCommission is a paid mutator transaction binding the contract method 0x16c58d04.
//
// Solidity: function withdrawCommission(address _token) returns()
func (_ChallengeEscrow *ChallengeEscrowSession) WithdrawCommission(_token common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.WithdrawCommission(&_ChallengeEscrow.TransactOpts, _token)
}

// WithdrawCommission is a paid mutator transaction binding the contract method 0x16c58d04.
//
// Solidity: function withdrawCommission(address _token) returns()
func (_ChallengeEscrow *ChallengeEscrowTransactorSession) WithdrawCommission(_token common.Address) (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.WithdrawCommission(&_ChallengeEscrow.TransactOpts, _token)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ChallengeEscrow *ChallengeEscrowTransactor) Receive(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChallengeEscrow.contract.RawTransact(opts, nil) // calldata is disallowed for receive function
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ChallengeEscrow *ChallengeEscrowSession) Receive() (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.Receive(&_ChallengeEscrow.TransactOpts)
}

// Receive is a paid mutator transaction binding the contract receive function.
//
// Solidity: receive() payable returns()
func (_ChallengeEscrow *ChallengeEscrowTransactorSession) Receive() (*types.Transaction, error) {
	return _ChallengeEscrow.Contract.Receive(&_ChallengeEscrow.TransactOpts)
}

// ChallengeEscrowChallengeAcceptedIterator is returned from FilterChallengeAccepted and is used to iterate over the raw logs and unpacked data for ChallengeAccepted events raised by the ChallengeEscrow contract.
type ChallengeEscrowChallengeAcceptedIterator struct {
	Event *ChallengeEscrowChallengeAccepted // Event containing the contract specifics and raw log

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
func (it *ChallengeEscrowChallengeAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeEscrowChallengeAccepted)
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
		it.Event = new(ChallengeEscrowChallengeAccepted)
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
func (it *ChallengeEscrowChallengeAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeEscrowChallengeAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeEscrowChallengeAccepted represents a ChallengeAccepted event raised by the ChallengeEscrow contract.
type ChallengeEscrowChallengeAccepted struct {
	Id          *big.Int
	Participant common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterChallengeAccepted is a free log retrieval operation binding the contract event 0x92908348abc2acb40e8d36acdad97909db4110254e1625c0509520a68bd999c9.
//
// Solidity: event ChallengeAccepted(uint256 indexed id, address indexed participant, uint256 amount)
func (_ChallengeEscrow *ChallengeEscrowFilterer) FilterChallengeAccepted(opts *bind.FilterOpts, id []*big.Int, participant []common.Address) (*ChallengeEscrowChallengeAcceptedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.FilterLogs(opts, "ChallengeAccepted", idRule, participantRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeEscrowChallengeAcceptedIterator{contract: _ChallengeEscrow.contract, event: "ChallengeAccepted", logs: logs, sub: sub}, nil
}

// WatchChallengeAccepted is a free log subscription operation binding the contract event 0x92908348abc2acb40e8d36acdad97909db4110254e1625c0509520a68bd999c9.
//
// Solidity: event ChallengeAccepted(uint256 indexed id, address indexed participant, uint256 amount)
func (_ChallengeEscrow *ChallengeEscrowFilterer) WatchChallengeAccepted(opts *bind.WatchOpts, sink chan<- *ChallengeEscrowChallengeAccepted, id []*big.Int, participant []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.WatchLogs(opts, "ChallengeAccepted", idRule, participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeEscrowChallengeAccepted)
				if err := _ChallengeEscrow.contract.UnpackLog(event, "ChallengeAccepted", log); err != nil {
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

// ParseChallengeAccepted is a log parse operation binding the contract event 0x92908348abc2acb40e8d36acdad97909db4110254e1625c0509520a68bd999c9.
//
// Solidity: event ChallengeAccepted(uint256 indexed id, address indexed participant, uint256 amount)
func (_ChallengeEscrow *ChallengeEscrowFilterer) ParseChallengeAccepted(log types.Log) (*ChallengeEscrowChallengeAccepted, error) {
	event := new(ChallengeEscrowChallengeAccepted)
	if err := _ChallengeEscrow.contract.UnpackLog(event, "ChallengeAccepted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeEscrowChallengeCreatedIterator is returned from FilterChallengeCreated and is used to iterate over the raw logs and unpacked data for ChallengeCreated events raised by the ChallengeEscrow contract.
type ChallengeEscrowChallengeCreatedIterator struct {
	Event *ChallengeEscrowChallengeCreated // Event containing the contract specifics and raw log

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
func (it *ChallengeEscrowChallengeCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeEscrowChallengeCreated)
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
		it.Event = new(ChallengeEscrowChallengeCreated)
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
func (it *ChallengeEscrowChallengeCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeEscrowChallengeCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeEscrowChallengeCreated represents a ChallengeCreated event raised by the ChallengeEscrow contract.
type ChallengeEscrowChallengeCreated struct {
	Id      *big.Int
	Creator common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterChallengeCreated is a free log retrieval operation binding the contract event 0x83a97ccca8e136a4d58fd664739225fcfe9db9a5b0bebf35df44614f2aa9b63a.
//
// Solidity: event ChallengeCreated(uint256 indexed id, address indexed creator)
func (_ChallengeEscrow *ChallengeEscrowFilterer) FilterChallengeCreated(opts *bind.FilterOpts, id []*big.Int, creator []common.Address) (*ChallengeEscrowChallengeCreatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.FilterLogs(opts, "ChallengeCreated", idRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeEscrowChallengeCreatedIterator{contract: _ChallengeEscrow.contract, event: "ChallengeCreated", logs: logs, sub: sub}, nil
}

// WatchChallengeCreated is a free log subscription operation binding the contract event 0x83a97ccca8e136a4d58fd664739225fcfe9db9a5b0bebf35df44614f2aa9b63a.
//
// Solidity: event ChallengeCreated(uint256 indexed id, address indexed creator)
func (_ChallengeEscrow *ChallengeEscrowFilterer) WatchChallengeCreated(opts *bind.WatchOpts, sink chan<- *ChallengeEscrowChallengeCreated, id []*big.Int, creator []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.WatchLogs(opts, "ChallengeCreated", idRule, creatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeEscrowChallengeCreated)
				if err := _ChallengeEscrow.contract.UnpackLog(event, "ChallengeCreated", log); err != nil {
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

// ParseChallengeCreated is a log parse operation binding the contract event 0x83a97ccca8e136a4d58fd664739225fcfe9db9a5b0bebf35df44614f2aa9b63a.
//
// Solidity: event ChallengeCreated(uint256 indexed id, address indexed creator)
func (_ChallengeEscrow *ChallengeEscrowFilterer) ParseChallengeCreated(log types.Log) (*ChallengeEscrowChallengeCreated, error) {
	event := new(ChallengeEscrowChallengeCreated)
	if err := _ChallengeEscrow.contract.UnpackLog(event, "ChallengeCreated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeEscrowChallengeDeletedIterator is returned from FilterChallengeDeleted and is used to iterate over the raw logs and unpacked data for ChallengeDeleted events raised by the ChallengeEscrow contract.
type ChallengeEscrowChallengeDeletedIterator struct {
	Event *ChallengeEscrowChallengeDeleted // Event containing the contract specifics and raw log

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
func (it *ChallengeEscrowChallengeDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeEscrowChallengeDeleted)
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
		it.Event = new(ChallengeEscrowChallengeDeleted)
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
func (it *ChallengeEscrowChallengeDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeEscrowChallengeDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeEscrowChallengeDeleted represents a ChallengeDeleted event raised by the ChallengeEscrow contract.
type ChallengeEscrowChallengeDeleted struct {
	Id  *big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterChallengeDeleted is a free log retrieval operation binding the contract event 0x3634596e4e000073bb0d932cc30db6d8e93bc7ee56f244a032fce06c0bcb2421.
//
// Solidity: event ChallengeDeleted(uint256 indexed id)
func (_ChallengeEscrow *ChallengeEscrowFilterer) FilterChallengeDeleted(opts *bind.FilterOpts, id []*big.Int) (*ChallengeEscrowChallengeDeletedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.FilterLogs(opts, "ChallengeDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeEscrowChallengeDeletedIterator{contract: _ChallengeEscrow.contract, event: "ChallengeDeleted", logs: logs, sub: sub}, nil
}

// WatchChallengeDeleted is a free log subscription operation binding the contract event 0x3634596e4e000073bb0d932cc30db6d8e93bc7ee56f244a032fce06c0bcb2421.
//
// Solidity: event ChallengeDeleted(uint256 indexed id)
func (_ChallengeEscrow *ChallengeEscrowFilterer) WatchChallengeDeleted(opts *bind.WatchOpts, sink chan<- *ChallengeEscrowChallengeDeleted, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.WatchLogs(opts, "ChallengeDeleted", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeEscrowChallengeDeleted)
				if err := _ChallengeEscrow.contract.UnpackLog(event, "ChallengeDeleted", log); err != nil {
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

// ParseChallengeDeleted is a log parse operation binding the contract event 0x3634596e4e000073bb0d932cc30db6d8e93bc7ee56f244a032fce06c0bcb2421.
//
// Solidity: event ChallengeDeleted(uint256 indexed id)
func (_ChallengeEscrow *ChallengeEscrowFilterer) ParseChallengeDeleted(log types.Log) (*ChallengeEscrowChallengeDeleted, error) {
	event := new(ChallengeEscrowChallengeDeleted)
	if err := _ChallengeEscrow.contract.UnpackLog(event, "ChallengeDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeEscrowChallengeStatusUpdatedIterator is returned from FilterChallengeStatusUpdated and is used to iterate over the raw logs and unpacked data for ChallengeStatusUpdated events raised by the ChallengeEscrow contract.
type ChallengeEscrowChallengeStatusUpdatedIterator struct {
	Event *ChallengeEscrowChallengeStatusUpdated // Event containing the contract specifics and raw log

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
func (it *ChallengeEscrowChallengeStatusUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeEscrowChallengeStatusUpdated)
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
		it.Event = new(ChallengeEscrowChallengeStatusUpdated)
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
func (it *ChallengeEscrowChallengeStatusUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeEscrowChallengeStatusUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeEscrowChallengeStatusUpdated represents a ChallengeStatusUpdated event raised by the ChallengeEscrow contract.
type ChallengeEscrowChallengeStatusUpdated struct {
	Id     *big.Int
	Status uint8
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterChallengeStatusUpdated is a free log retrieval operation binding the contract event 0x30c74fda3a9f1359aab44880f5e640842f90cbed4e11d9fbf0ca69a8e6a84a80.
//
// Solidity: event ChallengeStatusUpdated(uint256 indexed id, uint8 status)
func (_ChallengeEscrow *ChallengeEscrowFilterer) FilterChallengeStatusUpdated(opts *bind.FilterOpts, id []*big.Int) (*ChallengeEscrowChallengeStatusUpdatedIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.FilterLogs(opts, "ChallengeStatusUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeEscrowChallengeStatusUpdatedIterator{contract: _ChallengeEscrow.contract, event: "ChallengeStatusUpdated", logs: logs, sub: sub}, nil
}

// WatchChallengeStatusUpdated is a free log subscription operation binding the contract event 0x30c74fda3a9f1359aab44880f5e640842f90cbed4e11d9fbf0ca69a8e6a84a80.
//
// Solidity: event ChallengeStatusUpdated(uint256 indexed id, uint8 status)
func (_ChallengeEscrow *ChallengeEscrowFilterer) WatchChallengeStatusUpdated(opts *bind.WatchOpts, sink chan<- *ChallengeEscrowChallengeStatusUpdated, id []*big.Int) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.WatchLogs(opts, "ChallengeStatusUpdated", idRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeEscrowChallengeStatusUpdated)
				if err := _ChallengeEscrow.contract.UnpackLog(event, "ChallengeStatusUpdated", log); err != nil {
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

// ParseChallengeStatusUpdated is a log parse operation binding the contract event 0x30c74fda3a9f1359aab44880f5e640842f90cbed4e11d9fbf0ca69a8e6a84a80.
//
// Solidity: event ChallengeStatusUpdated(uint256 indexed id, uint8 status)
func (_ChallengeEscrow *ChallengeEscrowFilterer) ParseChallengeStatusUpdated(log types.Log) (*ChallengeEscrowChallengeStatusUpdated, error) {
	event := new(ChallengeEscrowChallengeStatusUpdated)
	if err := _ChallengeEscrow.contract.UnpackLog(event, "ChallengeStatusUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeEscrowClaimIterator is returned from FilterClaim and is used to iterate over the raw logs and unpacked data for Claim events raised by the ChallengeEscrow contract.
type ChallengeEscrowClaimIterator struct {
	Event *ChallengeEscrowClaim // Event containing the contract specifics and raw log

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
func (it *ChallengeEscrowClaimIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeEscrowClaim)
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
		it.Event = new(ChallengeEscrowClaim)
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
func (it *ChallengeEscrowClaimIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeEscrowClaimIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeEscrowClaim represents a Claim event raised by the ChallengeEscrow contract.
type ChallengeEscrowClaim struct {
	Participant common.Address
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterClaim is a free log retrieval operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed participant, uint256 amount)
func (_ChallengeEscrow *ChallengeEscrowFilterer) FilterClaim(opts *bind.FilterOpts, participant []common.Address) (*ChallengeEscrowClaimIterator, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.FilterLogs(opts, "Claim", participantRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeEscrowClaimIterator{contract: _ChallengeEscrow.contract, event: "Claim", logs: logs, sub: sub}, nil
}

// WatchClaim is a free log subscription operation binding the contract event 0x47cee97cb7acd717b3c0aa1435d004cd5b3c8c57d70dbceb4e4458bbd60e39d4.
//
// Solidity: event Claim(address indexed participant, uint256 amount)
func (_ChallengeEscrow *ChallengeEscrowFilterer) WatchClaim(opts *bind.WatchOpts, sink chan<- *ChallengeEscrowClaim, participant []common.Address) (event.Subscription, error) {

	var participantRule []interface{}
	for _, participantItem := range participant {
		participantRule = append(participantRule, participantItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.WatchLogs(opts, "Claim", participantRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeEscrowClaim)
				if err := _ChallengeEscrow.contract.UnpackLog(event, "Claim", log); err != nil {
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
func (_ChallengeEscrow *ChallengeEscrowFilterer) ParseClaim(log types.Log) (*ChallengeEscrowClaim, error) {
	event := new(ChallengeEscrowClaim)
	if err := _ChallengeEscrow.contract.UnpackLog(event, "Claim", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeEscrowCompletedChallengesDeletedIterator is returned from FilterCompletedChallengesDeleted and is used to iterate over the raw logs and unpacked data for CompletedChallengesDeleted events raised by the ChallengeEscrow contract.
type ChallengeEscrowCompletedChallengesDeletedIterator struct {
	Event *ChallengeEscrowCompletedChallengesDeleted // Event containing the contract specifics and raw log

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
func (it *ChallengeEscrowCompletedChallengesDeletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeEscrowCompletedChallengesDeleted)
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
		it.Event = new(ChallengeEscrowCompletedChallengesDeleted)
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
func (it *ChallengeEscrowCompletedChallengesDeletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeEscrowCompletedChallengesDeletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeEscrowCompletedChallengesDeleted represents a CompletedChallengesDeleted event raised by the ChallengeEscrow contract.
type ChallengeEscrowCompletedChallengesDeleted struct {
	Ids []*big.Int
	Raw types.Log // Blockchain specific contextual infos
}

// FilterCompletedChallengesDeleted is a free log retrieval operation binding the contract event 0x21468011b8061ecf9fd001704275feb14aae85c7961408d34ab4e694a18a3cc6.
//
// Solidity: event CompletedChallengesDeleted(uint256[] indexed ids)
func (_ChallengeEscrow *ChallengeEscrowFilterer) FilterCompletedChallengesDeleted(opts *bind.FilterOpts, ids [][]*big.Int) (*ChallengeEscrowCompletedChallengesDeletedIterator, error) {

	var idsRule []interface{}
	for _, idsItem := range ids {
		idsRule = append(idsRule, idsItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.FilterLogs(opts, "CompletedChallengesDeleted", idsRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeEscrowCompletedChallengesDeletedIterator{contract: _ChallengeEscrow.contract, event: "CompletedChallengesDeleted", logs: logs, sub: sub}, nil
}

// WatchCompletedChallengesDeleted is a free log subscription operation binding the contract event 0x21468011b8061ecf9fd001704275feb14aae85c7961408d34ab4e694a18a3cc6.
//
// Solidity: event CompletedChallengesDeleted(uint256[] indexed ids)
func (_ChallengeEscrow *ChallengeEscrowFilterer) WatchCompletedChallengesDeleted(opts *bind.WatchOpts, sink chan<- *ChallengeEscrowCompletedChallengesDeleted, ids [][]*big.Int) (event.Subscription, error) {

	var idsRule []interface{}
	for _, idsItem := range ids {
		idsRule = append(idsRule, idsItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.WatchLogs(opts, "CompletedChallengesDeleted", idsRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeEscrowCompletedChallengesDeleted)
				if err := _ChallengeEscrow.contract.UnpackLog(event, "CompletedChallengesDeleted", log); err != nil {
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

// ParseCompletedChallengesDeleted is a log parse operation binding the contract event 0x21468011b8061ecf9fd001704275feb14aae85c7961408d34ab4e694a18a3cc6.
//
// Solidity: event CompletedChallengesDeleted(uint256[] indexed ids)
func (_ChallengeEscrow *ChallengeEscrowFilterer) ParseCompletedChallengesDeleted(log types.Log) (*ChallengeEscrowCompletedChallengesDeleted, error) {
	event := new(ChallengeEscrowCompletedChallengesDeleted)
	if err := _ChallengeEscrow.contract.UnpackLog(event, "CompletedChallengesDeleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChallengeEscrowWinnerSetIterator is returned from FilterWinnerSet and is used to iterate over the raw logs and unpacked data for WinnerSet events raised by the ChallengeEscrow contract.
type ChallengeEscrowWinnerSetIterator struct {
	Event *ChallengeEscrowWinnerSet // Event containing the contract specifics and raw log

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
func (it *ChallengeEscrowWinnerSetIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChallengeEscrowWinnerSet)
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
		it.Event = new(ChallengeEscrowWinnerSet)
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
func (it *ChallengeEscrowWinnerSetIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChallengeEscrowWinnerSetIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChallengeEscrowWinnerSet represents a WinnerSet event raised by the ChallengeEscrow contract.
type ChallengeEscrowWinnerSet struct {
	Id     *big.Int
	Winner common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterWinnerSet is a free log retrieval operation binding the contract event 0x44fd60df9b05a22172cc43446c2ff1abb74d6c6ce938acd70bd3ee1ea2bf595d.
//
// Solidity: event WinnerSet(uint256 indexed id, address indexed winner)
func (_ChallengeEscrow *ChallengeEscrowFilterer) FilterWinnerSet(opts *bind.FilterOpts, id []*big.Int, winner []common.Address) (*ChallengeEscrowWinnerSetIterator, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.FilterLogs(opts, "WinnerSet", idRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return &ChallengeEscrowWinnerSetIterator{contract: _ChallengeEscrow.contract, event: "WinnerSet", logs: logs, sub: sub}, nil
}

// WatchWinnerSet is a free log subscription operation binding the contract event 0x44fd60df9b05a22172cc43446c2ff1abb74d6c6ce938acd70bd3ee1ea2bf595d.
//
// Solidity: event WinnerSet(uint256 indexed id, address indexed winner)
func (_ChallengeEscrow *ChallengeEscrowFilterer) WatchWinnerSet(opts *bind.WatchOpts, sink chan<- *ChallengeEscrowWinnerSet, id []*big.Int, winner []common.Address) (event.Subscription, error) {

	var idRule []interface{}
	for _, idItem := range id {
		idRule = append(idRule, idItem)
	}
	var winnerRule []interface{}
	for _, winnerItem := range winner {
		winnerRule = append(winnerRule, winnerItem)
	}

	logs, sub, err := _ChallengeEscrow.contract.WatchLogs(opts, "WinnerSet", idRule, winnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChallengeEscrowWinnerSet)
				if err := _ChallengeEscrow.contract.UnpackLog(event, "WinnerSet", log); err != nil {
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
func (_ChallengeEscrow *ChallengeEscrowFilterer) ParseWinnerSet(log types.Log) (*ChallengeEscrowWinnerSet, error) {
	event := new(ChallengeEscrowWinnerSet)
	if err := _ChallengeEscrow.contract.UnpackLog(event, "WinnerSet", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
