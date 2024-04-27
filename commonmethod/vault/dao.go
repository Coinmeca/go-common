package vault

import "time"

type Recent struct {
	Type     int       `json:"type" bson:"type"`
	Symbol   string    `json:"symbol" bson:"symbol"`
	From     string    `json:"from" bson:"from"`
	Address  string    `json:"address" bson:"address"`
	Time     string    `json:"time" bson:"time"`
	Volume   string    `json:"volume" bson:"volume"`
	Meca     string    `json:"meca" bson:"meca"`
	Share    float64   `json:"share" bson:"share"`
	TxHash   string    `json:"txHash" bson:"txHash"`
	UpdateAt time.Time `json:"updateAt" bson:"updateAt"`
}

type Vault struct {
	Key                   bool    `json:"key" bson:"key"`
	Address               string  `json:"address" bson:"address"`
	Symbol                string  `json:"symbol" bson:"symbol"`
	Name                  string  `json:"name" bson:"name"`
	Decimals              int     `json:"decimals" bson:"decimals"`
	Exchange              float64 `json:"exchange" bson:"exchange"`
	ExchangeChange24h     float64 `json:"exchangeChange24h" bson:"exchangeChange24h"`
	ExchangeChangeRate24h float64 `json:"exchangeChangeRate24h" bson:"exchangeChangeRate24h"`
	Tl                    float64 `json:"tl" bson:"tl"`
	TlChange              float64 `json:"tlChange" bson:"tlChange"`
	Tvl                   float64 `json:"tvl" bson:"tvl"`
	TvlChange             float64 `json:"tvlChange" bson:"tvlChange"`
	Weight                float64 `json:"weight" bson:"weight"`
	WeightChange          float64 `json:"weightChange" bson:"weightChange"`
	Deposit               float64 `json:"deposit" bson:"deposit"`
	Deposit24h            float64 `json:"deposit24h" bson:"deposit24h"`
	Withdraw              float64 `json:"withdraw" bson:"withdraw"`
	Withdraw24h           float64 `json:"withdraw24h" bson:"withdraw24h"`
	PerToken              float64 `json:"perToken" bson:"perToken"`
	TokenPer              float64 `json:"tokenPer" bson:"tokenPer"`
	Mint                  float64 `json:"mint" bson:"mint"`
	Burn                  float64 `json:"burn" bson:"burn"`
	//Chart
	//Recent
}
