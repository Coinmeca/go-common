package task

import (
	ABI "coinmeca-go_common/abi"
	"context"
	//ABI "dex-server/internal/abi"
	etherchain "coinmeca-go_common/chain"
	repo "coinmeca-go_common/repository"
	etherrpc "coinmeca-go_common/rpc"

	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"strings"
)

var (
	EthHttpsClient *ethclient.Client

	CTX      context.Context
	CHAINxID int

	OrderbookABI abi.ABI
	MarketABI    abi.ABI
	VaultABI     abi.ABI
	FarmABI      abi.ABI

	BookDecimals  map[string]int32  // market address
	TokenDecimals map[string]int32  // token address
	TokenSymbols  map[string]string // token address
	CAxBOOKS      []common.Address
	CAxVAULT      common.Address
)
var (
	TP1 = common.HexToHash(ABI.BuyEventHash)
	TP2 = common.HexToHash(ABI.SellEventHash)
	TP4 = common.HexToHash(ABI.DepositEventHash)
	TP5 = common.HexToHash(ABI.WithdrawEventHash)
	TP6 = common.HexToHash(ABI.ListingEventHash)
)

func NewTaskInstance(chain string) {
	id, ok := etherchain.ChainNameMap[chain]
	if ok == false {
		panic(fmt.Sprintf("not support chain (%s)", chain))
	}

	var err error
	httpsClientURI := etherchain.HTTPSProvider[id]
	EthHttpsClient, err = etherrpc.NewClient(httpsClientURI)
	if err != nil {
		panic(err)
	}

	CTX = context.Background()
	repo.InitDB(CTX, chain)

	loadContractAddresses(id)

	OrderbookABI, err = abi.JSON(strings.NewReader(ABI.Orderbook))
	if err != nil {
		panic(err)
	}
	MarketABI, err = abi.JSON(strings.NewReader(ABI.Market))
	if err != nil {
		panic(err)
	}
	VaultABI, err = abi.JSON(strings.NewReader(ABI.Vault))
	if err != nil {
		panic(err)
	}
	FarmABI, err = abi.JSON(strings.NewReader(ABI.Farm))
	if err != nil {
		panic(err)
	}
}

func loadContractAddresses(id int) {
	CHAINxID = id

	books, err := repo.MarketTokenInfo(CTX)
	if err != nil {
		panic(err)
	}

	CAxBOOKS = nil
	BookDecimals = make(map[string]int32, len(books))
	for _, b := range books {
		BookDecimals[b.Address] = b.Decimals
		CAxBOOKS = append(CAxBOOKS, common.HexToAddress(b.Address))
	}

	tokens, err := repo.VaultTokenInfo(CTX)
	if err != nil {
		panic(err)
	}
	TokenSymbols = make(map[string]string, len(tokens))
	TokenDecimals = make(map[string]int32, len(tokens))
	for _, t := range tokens {
		TokenSymbols[t.Address] = t.Symbol
		TokenDecimals[t.Address] = t.Decimals
	}

	CAxVAULT = common.HexToAddress(etherchain.MecaAddrMap[id]["VAULT"])
}

func CloseTaskInstance() {
	repo.CloseDB()
}
