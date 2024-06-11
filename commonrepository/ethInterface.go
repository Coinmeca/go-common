package repository

import (
	"context"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type EthClientRepository interface {
	BlockByNumber(number *big.Int) (*types.Block, error)
	LatestBlockNumber() (*big.Int, error)
	TxByHash(txHash string) (*types.Transaction, bool, error)
	TxReceipt(txHash string) (*types.Receipt, error)
	BalanceAt(account common.Address) (*big.Int, error)
	CallContract(contractAddr common.Address, msg string, tokenType string, params ...interface{}) ([]interface{}, error)
	TransactContract(txOpts *bind.TransactOpts, contractAddr common.Address, msg string, tokenType string, params ...interface{}) (*types.Transaction, error)
	SendTx(tx *types.Transaction) error
	GetNodeChainID(ctx context.Context) (*big.Int, error)
	//
	GetSender(ctx context.Context, tx *types.Transaction) (common.Address, error)
	//
	PendingNonceAt(account common.Address) (uint64, error)
	//
	EstimateGas(msg ethereum.CallMsg) (uint64, error)
	//
	GetContractABI(abiType string) (abi.ABI, error)
	//
	Call(result interface{}, method string, args ...interface{}) error
	//
	EthCallContract(msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error)
	//
	CallContext(result interface{}, method string, args ...interface{}) error
	//
	TransactionReceipt(txHash string) (*types.Receipt, error)
}
