package task

import (
	//ABI "dex-server/internal/abi"
	ABI "coinmeca-go_common/abi"
	//cv "dex-server/internal/configs"
	cv "coinmeca-go_common/utils"
	"fmt"
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/shopspring/decimal"
)

func TestOrderbookEventLogs(t *testing.T) {
	NewTaskInstance("polyzkt")
	defer CloseTaskInstance()

	var buyHash, sellHash = common.HexToHash(ABI.BuyEventHash), common.HexToHash(ABI.SellEventHash)
	var ownerHash = common.HexToHash("0x5B38Da6a701c568545dCfcB03FcB875f56beddC4")
	var maticHash, usdtHash = common.HexToHash(cv.MATICAddressMumbai), common.HexToHash(cv.USDTAddressMumbai)
	var eventLog = types.Log{
		Address:     common.HexToAddress(cv.MATICxUSDTMumbai),
		TxHash:      common.HexToHash("0x448ea94872267a10c944bbabf9fc505d2522ec1ad879aa4db08109fe7f7a7246"),
		TxIndex:     14,
		BlockNumber: 33490373,
		BlockHash:   common.HexToHash("0xb1f96e9ca9043087ef3507ab247acb6f4b29da736bc268011f2f0d047751bc7f"),
		Index:       44}

	// #1. [BUY] MATIC, p: $10, a: $30, q: 3
	priceHash := common.BigToHash(big.NewInt(10))
	//topicData := fmt.Sprintf("%064s%064x%064x%064x", cv.ETHAddressMumbai[2:], 1000*1e18, 3000*1e18, 3*1e18)
	topicData := fmt.Sprintf("%064x%064x", 30, 3)
	eventLog.Data = hexutil.MustDecode(fmt.Sprintf("0x%s", topicData))
	eventLog.Topics = []common.Hash{buyHash, ownerHash, maticHash, priceHash}
	orderEventData(eventLog, 0)

	// #2. [SELL] MATIC, p: $15, a: 3, q: $45
	priceHash = common.BigToHash(big.NewInt(15))
	eventLog.BlockNumber += 1
	//topicData = fmt.Sprintf("%064s%064x%064x%064x", cv.ETHAddressMumbai[2:], 1500*1e18, 3*1e18, 4500*1e18)
	topicData = fmt.Sprintf("%064x%064x", 45, 3)
	eventLog.Data = hexutil.MustDecode(fmt.Sprintf("0x%s", topicData))
	eventLog.Topics = []common.Hash{sellHash, ownerHash, maticHash, priceHash}
	orderEventData(eventLog, 0)

	// #3. [BUY] MATIC, p: $12, a: $60, q: 5
	priceHash = common.BigToHash(big.NewInt(12))
	eventLog.BlockNumber += 1
	//topicData = fmt.Sprintf("%064s%064x%064x%064x", cv.ETHAddressMumbai[2:], 1200*1e18, 6000*1e18, 5*1e18)
	topicData = fmt.Sprintf("%064x%064x", 60, 5)
	eventLog.Data = hexutil.MustDecode(fmt.Sprintf("0x%s", topicData))
	eventLog.Topics = []common.Hash{buyHash, ownerHash, maticHash, priceHash}
	orderEventData(eventLog, 0)

	//b, s, v := GetMarketVolume(cv.ETHAddressMumbai)

	// #4. [BUY] USDT, p: $1, a: $25, q: 25
	priceHash = common.BigToHash(big.NewInt(1))
	eventLog.BlockNumber += 1
	//topicData = fmt.Sprintf("%064s%064x%064x%064x", cv.USDTAddressMumbai[2:], 1*1e18, 25*1e18, 25*1e18)
	topicData = fmt.Sprintf("%064x%064x", 25, 25)
	eventLog.Data = hexutil.MustDecode(fmt.Sprintf("0x%s", topicData))
	eventLog.Topics = []common.Hash{buyHash, ownerHash, usdtHash, priceHash}
	orderEventData(eventLog, 0)

	// #5. [SELL] USDT, p: $1, a: 45, q: $45
	priceHash = common.BigToHash(big.NewInt(1))
	eventLog.BlockNumber += 1
	//topicData = fmt.Sprintf("%064s%064x%064x%064x", cv.USDTAddressMumbai[2:], 1*1e18, 45*1e18, 45*1e18)
	topicData = fmt.Sprintf("%064x%064x", 45, 45)
	eventLog.Data = hexutil.MustDecode(fmt.Sprintf("0x%s", topicData))
	eventLog.Topics = []common.Hash{sellHash, ownerHash, usdtHash, priceHash}
	orderEventData(eventLog, 0)

	// #6. [BUY] USDT, p: $1, a: $53, q: 53
	priceHash = common.BigToHash(big.NewInt(1))
	eventLog.BlockNumber += 1
	//topicData = fmt.Sprintf("%064s%064x%064x%064x", cv.USDTAddressMumbai[2:], 1*1e18, 53*1e18, 53*1e18)
	topicData = fmt.Sprintf("%064x%064x", 53, 53)
	eventLog.Data = hexutil.MustDecode(fmt.Sprintf("0x%s", topicData))
	eventLog.Topics = []common.Hash{buyHash, ownerHash, usdtHash, priceHash}
	orderEventData(eventLog, 0)

	//b, s, v = GetMarketVolume(cv.USDTAddressMumbai)
}

func orderEventData(log types.Log, exp int32) {
	type OrderbookEventRow struct {
		Event     string
		EventType string
		Owner     string
		Address   string
		Token     string
		BlockHash string
		BlockNo   uint64
		TxHash    string
		TxIndex   uint32
		Price     decimal.Decimal
		Amount    decimal.Decimal
		Quantity  decimal.Decimal
	}

	r := OrderbookEventRow{
		BlockHash: strings.ToLower(log.BlockHash.Hex()),
		BlockNo:   log.BlockNumber,
		TxHash:    strings.ToLower(log.TxHash.Hex()),
		TxIndex:   uint32(log.TxIndex),
		Address:   strings.ToLower(log.Address.Hex()),
		Event:     log.Topics[0].String(),
	}

	switch log.Topics[0] {
	case TP1:
		r.EventType = "BUY"
	case TP2:
		r.EventType = "SELL"
	}
	r.Owner = fmt.Sprintf("0x%s", strings.ToLower(log.Topics[1].Hex()[26:]))
	r.Token = fmt.Sprintf("0x%s", strings.ToLower(log.Topics[2].Hex()[26:]))
	r.Price = decimal.NewFromBigInt(log.Topics[3].Big(), exp)

	type OrderEventData struct {
		Amount   *big.Int
		Quantity *big.Int
	}

	if len(log.Data) > 0 {
		var e OrderEventData
		err := OrderbookABI.UnpackIntoInterface(&e, "Buy", log.Data)
		if err != nil {
			fmt.Errorf("decode error...%v\n", err)
		}

		r.Amount = decimal.NewFromBigInt(e.Amount, exp)
		r.Quantity = decimal.NewFromBigInt(e.Quantity, exp)
		if r.EventType == "SELL" {
			r.Amount, r.Quantity = r.Quantity, r.Amount
		}
	}

	fmt.Printf("Event Row: %+v\n", r)
}
