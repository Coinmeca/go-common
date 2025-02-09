package task

import (
	"context"
	"fmt"
	"math/big"
	"strings"
	"sync"
	"time"

	ABI "github.com/coinmeca/go-common/abi"
	"github.com/coinmeca/go-common/logger"
	"github.com/coinmeca/go-common/model"
	rep "github.com/coinmeca/go-common/repository"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

func FetchOrderbook() {
	t := time.Now()

	wg := sync.WaitGroup{}
	for _, address := range CAxBOOKS {
		wg.Add(1)
		go func(a common.Address) {
			setMarketOrderbook(a)
			defer wg.Done()
		}(address)
	}
	wg.Wait()

	logger.Info("FetchOrderbook", "time", time.Since(t).Milliseconds())
}

func setMarketOrderbook(orderbookAddress common.Address) {
	sig, val := ABI.OrderbookSigMap["get_orderbook"], fmt.Sprintf("%064x", 30)
	dataHex := fmt.Sprintf("0x%s%s", sig, val)
	callMsg := ethereum.CallMsg{To: &orderbookAddress, Data: common.FromHex(dataHex)}

	res, err := EthHttpsClient.CallContract(context.Background(), callMsg, nil)
	if err != nil {
		logger.Error("setMarketOrderbook", "err", err)
		return
	}

	list, err := OrderbookABI.Unpack("get_orderbook", res)
	if err != nil {
		logger.Error("setMarketOrderbook", "err", err)
		return
	}
	orderList := list[0].(struct {
		Asks []struct {
			Price   *big.Int `json:"price"`
			Balance *big.Int `json:"balance"`
		} `json:"asks"`
		Bids []struct {
			Price   *big.Int `json:"price"`
			Balance *big.Int `json:"balance"`
		} `json:"bids"`
	})
	//fmt.Printf("[%s](get_orderbook) unpack data... asks is %v, bids is %v\n", orderbookPair, len(orderList.Asks), len(orderList.Bids))

	address := strings.ToLower(orderbookAddress.Hex())
	wg := sync.WaitGroup{}
	for i, item := range orderList.Asks {
		price, balance := decimal.NewFromBigInt(item.Price, -18), decimal.NewFromBigInt(item.Balance, -18) // todo: decimal from token
		//fmt.Printf("[ASK:%02d] price: %d, balance: %d \n", i+1, item.Price, item.Balance)
		row := model.MarketOrderbook{OrderType: "ASK", OrderIndex: uint8(i + 1), Price: price, Balance: balance, Address: address}
		wg.Add(1)
		go func(c context.Context, r model.MarketOrderbook) {
			defer wg.Done()
			rep.SetMarketOrderbook(c, r)
		}(CTX, row)
	}
	for i, item := range orderList.Bids {
		price, balance := decimal.NewFromBigInt(item.Price, -18), decimal.NewFromBigInt(item.Balance, -18)
		//fmt.Printf("[BID:%02d] price: %d, balance: %d \n", i+1, item.Price, item.Balance)
		row := model.MarketOrderbook{OrderType: "BID", OrderIndex: uint8(i + 1), Price: price, Balance: balance, Address: address}
		wg.Add(1)
		go func(c context.Context, r model.MarketOrderbook) {
			defer wg.Done()
			rep.SetMarketOrderbook(c, r)
		}(CTX, row)
	}
	wg.Wait()
}

////////////////////////
////////////////////////
////////////////////////

func LogOrderbook() {
	for _, address := range CAxBOOKS {
		sig, val := ABI.OrderbookSigMap["get_orderbook"], fmt.Sprintf("%064x", 30)
		dataHex := fmt.Sprintf("0x%s%s", sig, val)
		//contractAddr := common.HexToAddress(address)
		callMsg := ethereum.CallMsg{To: &address, Data: common.FromHex(dataHex)}

		res, err := EthHttpsClient.CallContract(context.Background(), callMsg, nil)
		if err != nil {
			logger.Error("LogOrderbook", "err", err)
			return
		}

		list, err := OrderbookABI.Unpack("get_orderbook", res)
		if err != nil {
			logger.Error("LogOrderbook", "err", err)
			return
		}
		orderList := list[0].(struct {
			Asks []struct {
				Price   *big.Int `json:"price"`
				Balance *big.Int `json:"balance"`
			} `json:"asks"`
			Bids []struct {
				Price   *big.Int `json:"price"`
				Balance *big.Int `json:"balance"`
			} `json:"bids"`
		})

		fmt.Printf("[%s](get_orderbook) unpack data... asks is %v, bids is %v\n", address.Hex(), len(orderList.Asks), len(orderList.Bids))
		for i, item := range orderList.Asks {
			price, balance := decimal.NewFromBigInt(item.Price, -18), decimal.NewFromBigInt(item.Balance, -18) // todo: decimal from token
			fmt.Printf("[ASK:%02d] price: %v, balance: %v \n", i+1, price, balance)
		}
		for i, item := range orderList.Bids {
			price, balance := decimal.NewFromBigInt(item.Price, -18), decimal.NewFromBigInt(item.Balance, -18)
			fmt.Printf("[BID:%02d] price: %v, balance: %v \n", i+1, price, balance)
		}
	}
}
