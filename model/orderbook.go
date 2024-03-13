package model

import (
	"github.com/shopspring/decimal"
)

type MarketOrderbook struct {
	//OrderbookPair string
	OrderType  string
	OrderIndex uint8
	Price      decimal.Decimal
	Balance    decimal.Decimal
	Address    string
}

type OrderbookPriceApi struct {
	Asks []OrderbookPriceItem `json:"asks"`
	Bids []OrderbookPriceItem `json:"bids"`
}

type OrderbookPriceItem struct {
	Price   float64 `json:"price"`
	Balance float64 `json:"balance"`
}
