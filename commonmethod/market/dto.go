package market

type MarketTokenInfo struct {
	Name      string  `json:"name" bson:"name"`
	Symbol    string  `json:"symbol" bson:"symbol"`
	Decimals  int     `json:"decimals" bson:"decimals"`
	Address   string  `json:"address" bson:"address"`
	Liquidity float64 `json:"liquidity" bson:"liquidity"`
}
