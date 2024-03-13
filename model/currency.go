package model

import "github.com/shopspring/decimal"

type CurrencyRate struct {
	Code string          `json:"code"`
	Unit string          `json:"currencyCode"`
	Rate decimal.Decimal `json:"basePrice"`
}

type StableCoinQuote struct {
	Symbol      string          `json:"symbol"`
	Price       decimal.Decimal `json:"price"`
	MarketCap   decimal.Decimal `json:"market_cap"`
	LastUpdated string          `json:"last_updated"`
}

type CMCLatestQuote struct {
	Data map[string]struct {
		Id     uint16                     `json:"id"`
		Symbol string                     `json:"symbol"`
		Quote  map[string]StableCoinQuote `json:"quote"`
	} `json:"data"`
}
