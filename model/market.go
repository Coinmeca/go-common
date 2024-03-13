package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type MarketAbiInfo struct { // buy, sell
	Orderbook common.Address
	NFT       common.Address
	Symbol    string
	Name      string
	Base      common.Address
	Quote     common.Address
	Price     *big.Int
	Tick      *big.Int
	Fee       uint8
	Lock      bool
}

type MarketPriceRow struct {
	Time    string
	Address string
	Price   decimal.Decimal
}

type MarketVolumeRow struct {
	Event     string
	EventType string
	Owner     string
	BlockHash string
	BlockNo   uint64
	TxHash    string
	TxIndex   uint32
	Address   string
	Price     decimal.Decimal
	Amount    decimal.Decimal
	Quantity  decimal.Decimal
}

type MarketTokenRow struct {
	Symbol    string
	Name      string
	Decimals  int32
	Base      string
	Orderbook string
	NFT       string
	Quote     string
}

type MarketTokenApi MarketTokenRow

type MarketStatApi struct {
	Symbol    string  `json:"symbol"`
	Name      string  `json:"name"`
	Base      string  `json:"base"`
	Decimals  int32   `json:"decimals"`
	Orderbook string  `json:"market"`
	NFT       string  `json:"nft"`
	Quote     string  `json:"quote"`
	Price     float64 `json:"price"`
	Change    float64 `json:"change_rate"`
	Volume    float64 `json:"volume"`
}

type MarketDashboardApi struct {
	Price       float64 `json:"price"`
	VolumeQuote float64 `json:"volume_quote"`
	VolumeBase  float64 `json:"volume_base"`
	//Volume      float64 `json:"volume"`
	High       float64 `json:"high"`
	Low        float64 `json:"low"`
	ChangeRate float64 `json:"change_rate"`
	Change     float64 `json:"change"`
}

type MarketHistoryApi struct {
	Time     string  `json:"time"`
	Type     string  `json:"type"`
	Quantity float64 `json:"quantity"`
	Price    float64 `json:"price"`
}
