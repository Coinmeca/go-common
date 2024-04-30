package task

import (
	ABI "github.com/coinmeca/go-common/abi"
	cv "github.com/coinmeca/go-common/chain"
	"github.com/coinmeca/go-common/logger"
	"github.com/coinmeca/go-common/model"
	repo "github.com/coinmeca/go-common/repository"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
	"math/big"
	"strings"
	"sync"
	"time"
)

func LoadVaultVolume() {
	blockNo, err := repo.LastBlockNoFromVaultVolume(CTX)
	if err != nil {
		fmt.Println("last block error")
		currentBlock, _ := EthHttpsClient.BlockNumber(context.Background())
		blockNo = int64(currentBlock) + 1
	} else {
		blockNo += 1
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(31812267), // mumbai: 35259091
		Topics:    [][]common.Hash{{TP4, TP5}},
		Addresses: []common.Address{CAxVAULT},
	}

	logs, err := EthHttpsClient.FilterLogs(context.Background(), query)
	if err != nil {
		logger.Error("LoadVaultVolume", "err", err)
		panic(err)
	}
	fmt.Printf("VAULT logs %v (last block %v)\n", len(logs), blockNo)

	for i, log := range logs {
		fmt.Printf("[%d] vault event: %+v\n", i+1, log)
		setVaultVolume(log)
	}
}

func LoadVaultTokens() {
	setVaultData(true)
}

func LoadVaultPrice() {
	setVaultData(false)
}

func setVaultData(isPreset bool) {
	t := time.Now()

	tokenList, err := vaultTokenList()
	if err != nil {
		logger.Error("setVaultData", "err", err)
		return
	}

	wg := sync.WaitGroup{}
	for _, token := range tokenList {
		wg.Add(1)
		//fmt.Printf("[%s-%d] %+v \n", token.Symbol, i+1, token)
		go func(c context.Context, v model.VaultAbiInfo) {
			if isPreset {
				setVaultTokenInfo(c, v)
			} else {
				setVaultPrice(c, v)
			}
			wg.Done()
		}(CTX, token)
	}
	wg.Wait()

	// todo: return struct { symbol, decimals }
	logger.Debug("setVaultData", "time", time.Since(t).Milliseconds())
}

func vaultTokenList() ([]model.VaultAbiInfo, error) {
	dataHex := fmt.Sprintf("0x%s", ABI.VaultSigMap["get_all"])
	contractAddr := common.HexToAddress(cv.MecaAddrMap[CHAINxID]["VAULT"])
	callMsg := ethereum.CallMsg{To: &contractAddr, Data: common.FromHex(dataHex)}

	res, err := EthHttpsClient.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		logger.Error("vaultTokenList", "err", err)
		return nil, err
	}

	var tokens []model.VaultAbiInfo
	err = VaultABI.UnpackIntoInterface(&tokens, "get_all", res)
	if err != nil {
		logger.Error("vaultTokenList", "err", err)
		return nil, err
	}
	//fmt.Println("unpack data...", tokens) // for monitoring

	return tokens, nil
}

func setVaultVolume(log types.Log) {
	r := model.VaultVolumeRow{
		BlockHash: strings.ToLower(log.BlockHash.Hex()),
		BlockNo:   log.BlockNumber,
		TxHash:    strings.ToLower(log.TxHash.Hex()),
		TxIndex:   uint32(log.TxIndex),
		Event:     log.Topics[0].String(),
	}

	switch log.Topics[0] {
	case TP4:
		r.EventType = "DEPOSIT"
	case TP5:
		r.EventType = "WITHDRAW"
	}

	r.Owner = fmt.Sprintf("0x%s", strings.ToLower(log.Topics[1].Hex()[26:]))
	r.Token = fmt.Sprintf("0x%s", strings.ToLower(log.Topics[2].Hex()[26:]))
	exp := TokenDecimals[r.Token] * -1
	r.Amount = decimal.NewFromBigInt(log.Topics[3].Big(), exp)
	r.Symbol = TokenSymbols[r.Token]

	if len(log.Data) > 0 {
		var e struct {
			//Amount   *big.Int
			Quantity *big.Int
		}
		err := VaultABI.UnpackIntoInterface(&e, "Deposit", log.Data)
		if err != nil {
			fmt.Errorf("decode error...%v\n", err)
		}

		//fmt.Printf("log data: %v\n", e)
		if e.Quantity != nil {
			r.MecaQuantity = decimal.NewFromBigInt(e.Quantity, exp)
		}
	}
	//fmt.Printf("(vault volume) %+v\n", r)
	repo.SetVaultVolume(CTX, r)
}

func setVaultPrice(ctx context.Context, v model.VaultAbiInfo) {
	t := time.Now().UTC()
	dc := -1 * int32(v.Decimals)

	usdPrice, _ := repo.LatestUsdPrice(ctx, v.Symbol)
	treasury := decimal.NewFromBigInt(v.Treasury, dc)
	priceRow := model.VaultPriceRow{
		Time:     fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", t.Year(), t.Month(), t.Day(), t.Hour(), (t.Minute()/5)*5),
		Address:  strings.ToLower(v.Addr.Hex()),
		Symbol:   v.Symbol,
		Treasury: treasury,
		Tvl:      treasury.Mul(usdPrice),
		Price:    decimal.NewFromBigInt(v.Rate, dc),
		Weight:   decimal.NewFromBigInt(v.Weight, dc),
		Need:     decimal.NewFromBigInt(v.Need, dc),
	}
	//fmt.Printf("(vault price) %+v\n", priceRow)
	repo.SetVaultPrice(ctx, priceRow)
}

func setVaultTokenInfo(ctx context.Context, t model.VaultAbiInfo) {
	tokenRow := model.VaultTokenRow{
		IsKey:    t.Key,
		Address:  strings.ToLower(t.Addr.Hex()),
		Name:     t.Name,
		Symbol:   t.Symbol,
		Decimals: t.Decimals,
	}
	//fmt.Printf("(vault token) %+v\n", tokenRow)
	repo.SetVaultTokenInfo(ctx, tokenRow)
}
