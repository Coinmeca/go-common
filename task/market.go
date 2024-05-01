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

func LoadMarketTokens() {
	setMarketData(true) // tokens
}

func LoadMarketPrice() {
	setMarketData(false) // price
}

func LoadMarketVolume() {
	blockNo, err := repo.LastBlockNoFromMarketVolume(CTX)
	if err != nil {
		fmt.Println("last block error")
		currentBlock, _ := EthHttpsClient.BlockNumber(context.Background())
		blockNo = int64(currentBlock) + 1
	} else {
		blockNo += 1
	}

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(33433017),
		Topics:    [][]common.Hash{{TP1, TP2}},
		Addresses: CAxBOOKS,
	}

	logs, err := EthHttpsClient.FilterLogs(context.Background(), query)
	if err != nil {
		logger.Error("LoadMarketVolume", "err", err, "block number", blockNo)
		panic(err)
	}
	fmt.Printf("MARKET logs %v (last block %v)\n", len(logs), blockNo)

	for i, log := range logs {
		fmt.Printf("[%d] market volume: %+v\n", i+1, log)
		setMarketVolume(log)
	}
}

func setMarketData(isPreset bool) {
	t := time.Now()

	markets, err := marketList()
	if err != nil {
		//logger.Error.Printf("market list error (%v)\n", err)
		return
	}

	wg := sync.WaitGroup{}
	for _, market := range markets {
		wg.Add(1)
		//fmt.Printf("[%s-%d] %+v \n", market.Symbol, i+1, market) // for monitoring
		go func(m model.MarketAbiInfo) {
			if isPreset {
				setMarketToken(m)
			} else {
				setMarketPrice(m)
			}
			wg.Done()
		}(market)
	}
	wg.Wait()

	logger.Debug("setMarketData", "time", time.Since(t).Milliseconds())
}

func marketList() ([]model.MarketAbiInfo, error) {
	dataHex := fmt.Sprintf("0x%s", ABI.MarketSigMap["get_all"])
	contractAddr := common.HexToAddress(cv.MecaAddrMap[CHAINxID]["MARKET"])
	callMsg := ethereum.CallMsg{To: &contractAddr, Data: common.FromHex(dataHex)}

	res, err := EthHttpsClient.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		logger.Error("marketList", "err", err)
		return nil, err
	}

	//fmt.Printf("call all market %v\n", len(res))
	var packed []model.MarketAbiInfo
	err = MarketABI.UnpackIntoInterface(&packed, "get_all", res)
	//fmt.Printf("packed...%+v\n", packed)
	if err != nil {
		logger.Error("marketList", "err", err)
		return nil, err
	}
	return packed, nil
}

func setMarketPrice(m model.MarketAbiInfo) {
	address := strings.ToLower(m.Orderbook.Hex())

	t := time.Now().UTC()
	priceRow := model.MarketPriceRow{
		Price:   decimal.NewFromBigInt(m.Price, -1*BookDecimals[address]),
		Address: address,
		Time:    fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", t.Year(), t.Month(), t.Day(), t.Hour(), (t.Minute()/5)*5),
	}

	repo.SetMarketPrice(CTX, priceRow)
}

func setMarketToken(m model.MarketAbiInfo) {
	address := strings.ToLower(m.Orderbook.Hex())

	tokenRow := model.MarketTokenRow{
		Symbol:    m.Symbol,
		Name:      m.Name,
		Decimals:  BookDecimals[address],
		Base:      strings.ToLower(m.Base.Hex()),
		Orderbook: address,
		NFT:       strings.ToLower(m.NFT.String()),
		Quote:     strings.ToLower(m.Quote.String()),
	}

	//fmt.Printf("[%v] market token %+v\n", CHAINxID, tokenRow)
	repo.SetMarketToken(CTX, tokenRow)
}

func setMarketVolume(log types.Log) {
	address := strings.ToLower(log.Address.Hex())

	r := model.MarketVolumeRow{
		BlockHash: strings.ToLower(log.BlockHash.Hex()),
		BlockNo:   log.BlockNumber,
		TxHash:    strings.ToLower(log.TxHash.Hex()),
		TxIndex:   uint32(log.TxIndex),
		Event:     log.Topics[0].String(),
		Address:   address,
	}

	switch log.Topics[0] {
	case TP1:
		r.EventType = "BUY"
	case TP2:
		r.EventType = "SELL"
	}

	r.Owner = fmt.Sprintf("0x%s", strings.ToLower(log.Topics[1].Hex()[26:]))
	//r.Token = fmt.Sprintf("0x%s", strings.ToLower(log.Topics[2].Hex()[26:]))
	exp := BookDecimals[address] * -1
	r.Price = decimal.NewFromBigInt(log.Topics[3].Big(), exp)

	if len(log.Data) > 0 {
		var e struct {
			Amount   *big.Int
			Quantity *big.Int
		}
		err := OrderbookABI.UnpackIntoInterface(&e, "Buy", log.Data)
		if err != nil {
			fmt.Errorf("decode error...%v\n", err)
		}
		r.Amount = decimal.NewFromBigInt(e.Amount, exp)
		r.Quantity = decimal.NewFromBigInt(e.Quantity, exp)
		if r.EventType == "SELL" {
			r.Amount, r.Quantity = r.Quantity, r.Amount
		}
		//fmt.Printf("(MV) %v %v %v\n", r.EventType, r.Quantity, r.Amount)
	}
	repo.SetMarketVolume(CTX, r)
}

///////////////////////////////////
///////////////////////////////////
///////////////////////////////////

func LogMarketList() {
	markets, err := marketList()
	if err != nil {
		logger.Error("LogMarketList", "err", err)
		return
	}

	for i, market := range markets {
		orderbook := strings.ToLower(market.Orderbook.Hex())
		fmt.Printf("[%s-%d] %v, %+v \n", market.Symbol, i+1, orderbook, market) // for monitoring
	}
}
