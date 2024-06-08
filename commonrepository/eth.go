package repository

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
)

type EthRepository struct {
	ethClient   *ethclient.Client
	contractABI map[string]abi.ABI
	//contractInfo map[string]commondatabase.Contract
}

func NewEthRepository(alchemyURL string) *EthRepository {
	client, err := ethclient.Dial(alchemyURL)
	if err != nil {
		log.Fatalf("Failed to connect to the ethereum client: %v", err)
		return nil
	}

	return &EthRepository{
		ethClient: client,
	}
}

//func NewEth() (map[string]EthClientRepository, error) {
//	client, err := rpc.Dial()
//	if err != nil {
//		return nil, fmt.Errorf("Client connect failed : %s", err.Error())
//	}
//
//}

func (ec *EthRepository) GetContractABI(abiType string) (abi.ABI, error) {
	abi, ok := ec.contractABI[abiType]
	if !ok {
		return abi, errors.New("Contract not found")
	}
	return abi, nil
}

func (ec *EthRepository) EstimateGas(msg ethereum.CallMsg) (uint64, error) {
	result, err := ec.ethClient.EstimateGas(context.Background(), msg)
	return result, errors.Unwrap(err)
}

func (ec *EthRepository) SuggestGasPrice() (*big.Int, error) {
	result, err := ec.ethClient.SuggestGasPrice(context.Background())
	return result, errors.Unwrap(err)
}

//func (ec *EthRepository) GetContractABIType(contractType string) (string, error) {
//	contractInfo, ok := ec.contractInfo[contractType]
//	if !ok {
//		return contractInfo.ABIType, errors.New("Contract not found")
//	}
//	return contractInfo.ABIType, nil
//}

func (ec *EthRepository) GetNodeChainID(ctx context.Context) (*big.Int, error) {
	chainId, err := ec.ethClient.ChainID(ctx)
	if err != nil {
		// logger.LogError("Failed to retrieve chain Id: " + err.Error())
		return chainId, errors.Wrap(err, "Failed to ChainID")
	}
	return chainId, nil
}

func (ec *EthRepository) LatestBlockByNumber() (*big.Int, error) {
	header, err := ec.ethClient.HeaderByNumber(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
		return big.NewInt(0), errors.Wrap(err, "LatestBlockByNumber Fail")
	}
	blockNumber := header.Number
	return blockNumber, nil
}

func (ec *EthRepository) BlockByNumber(number *big.Int) (*types.Block, error) {
	result, err := ec.ethClient.BlockByNumber(context.Background(), number)
	return result, errors.Unwrap(err)
}

func (ec *EthRepository) SendTx(tx *types.Transaction) error {
	err := ec.ethClient.SendTransaction(context.Background(), tx)
	return errors.Unwrap(err)
}

func (ec *EthRepository) PendingNonceAt(account common.Address) (uint64, error) {
	result, err := ec.ethClient.PendingNonceAt(context.Background(), account)
	return result, errors.Unwrap(err)
}

func (ec *EthRepository) GetSender(ctx context.Context, tx *types.Transaction) (common.Address, error) {
	chainId, err := ec.GetNodeChainID(ctx)
	if err != nil {
		return common.Address{}, errors.Wrap(err, "Failed to GetNodeChainID")
	}
	result, err := types.Sender(types.LatestSignerForChainID(chainId), tx)
	return result, errors.Unwrap(err)
}

func (ec *EthRepository) CallContract(contractAddr common.Address, msg string, abiType string, params ...interface{}) ([]interface{}, error) {
	contract := bind.NewBoundContract(contractAddr, ec.contractABI[abiType], ec.ethClient, ec.ethClient, ec.ethClient)

	var result []interface{}

	err := contract.Call(nil, &result, msg, params...)
	return result, errors.Wrap(err, "Failed to CallContract")
}

func (ec *EthRepository) TransactContract(txOpts *bind.TransactOpts, contractAddr common.Address, msg string, abiType string, params ...interface{}) (*types.Transaction, error) {
	contract := bind.NewBoundContract(contractAddr, ec.contractABI[abiType], ec.ethClient, ec.ethClient, ec.ethClient)

	tx, err := contract.Transact(txOpts, msg, params...)
	return tx, errors.Wrap(err, "Failed to TransactContract")
}

func (ec *EthRepository) TxReceipt(txHash string) (*types.Receipt, error) {
	result, err := ec.ethClient.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	return result, errors.Unwrap(err)
}

func (ec *EthRepository) TxByHash(txHash string) (*types.Transaction, bool, error) {
	result, isPending, err := ec.ethClient.TransactionByHash(context.Background(), common.HexToHash(txHash))
	return result, isPending, errors.Unwrap(err)
}

func (ec *EthRepository) BalanceAt(account common.Address) (*big.Int, error) {
	result, err := ec.ethClient.BalanceAt(context.Background(), account, nil)
	return result, errors.Unwrap(err)
}

func GetEthRepositories(ethRepos map[string]EthClientRepository, keys ...string) (map[string]*EthRepository, error) {
	result := make(map[string]*EthRepository)
	for _, key := range keys {
		repo, exists := ethRepos[key]
		if !exists {
			return nil, fmt.Errorf("repository not found for key: %s", key)
		}

		ethRepo, ok := repo.(*EthRepository)
		if !ok {
			return nil, fmt.Errorf("could not cast to EthRepository for key: %s", key)
		}

		result[key] = ethRepo
	}
	return result, nil
}

func GetAllEthRepositories(ethRepos map[string]EthClientRepository) (map[string]*EthRepository, error) {
	result := make(map[string]*EthRepository)
	for key, repo := range ethRepos {
		ethRepo, ok := repo.(*EthRepository)
		if !ok {
			return nil, fmt.Errorf("could not cast to EthRepository for key: %s", key)
		}

		result[key] = ethRepo
	}
	return result, nil
}

func (ec *EthRepository) WaitMined(signedTx *types.Transaction) (*types.Receipt, error) {
	receipt, err := bind.WaitMined(context.Background(), ec.ethClient, signedTx)
	return receipt, errors.Unwrap(err)
}

func (ec *EthRepository) Call(result interface{}, method string, args ...interface{}) error {
	err := ec.ethClient.Client().Call(result, method, args...)
	return errors.Unwrap(err)
}

func (ec *EthRepository) EthCallContract(msg ethereum.CallMsg, blockNumber *big.Int) ([]byte, error) {
	result, err := ec.ethClient.CallContract(context.Background(), msg, blockNumber)
	if err != nil {
		return nil, errors.Unwrap(err)
	}
	return result, nil
}

func (ec *EthRepository) CallContext(result interface{}, method string, args ...interface{}) error {
	err := ec.ethClient.Client().CallContext(context.Background(), result, method, args...)
	return errors.Unwrap(err)
}

func (ec *EthRepository) TransactionReceipt(txHash string) (*types.Receipt, error) {
	result, err := ec.ethClient.TransactionReceipt(context.Background(), common.HexToHash(txHash))
	return result, errors.Unwrap(err)
}

func (ec *EthRepository) GetEthClient() *ethclient.Client {
	return ec.ethClient
}
