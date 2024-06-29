package vault

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recent struct {
	ChainId  string               `json:"chainId" bson:"chainId"`
	Address  string               `json:"address" bson:"address"`
	Time     int64                `json:"time" bson:"time"`
	Type     int64                `json:"type" bson:"type"`
	User     string               `json:"user" bson:"user"`
	Amount   primitive.Decimal128 `json:"volume" bson:"volume"`
	Meca     primitive.Decimal128 `json:"meca" bson:"meca"`
	Share    primitive.Decimal128 `json:"share" bson:"share"`
	TxHash   string               `json:"txHash" bson:"txHash"`
	UpdateAt string               `json:"updateAt" bson:"updateAt"`
}

type Last struct {
	ChainId     string               `json:"chainId" bson:"chainId"`
	Address     string               `json:"address" bson:"address"`
	Rate        primitive.Decimal128 `json:"rate" bson:"rate"`
	Deposit     primitive.Decimal128 `json:"deposit" bson:"deposit"`
	Withdraw    primitive.Decimal128 `json:"withdraw" bson:"withdraw"`
	Locked      primitive.Decimal128 `json:"locked" bson:"locked"`
	ValueLocked primitive.Decimal128 `json:"valueLocked" bson:"valueLocked"`
	Value       primitive.Decimal128 `json:"value" bson:"value"`
	Weight      primitive.Decimal128 `json:"weight" bson:"weight"`
	Mint        primitive.Decimal128 `json:"mint" bson:"mint"`
	Burn        primitive.Decimal128 `json:"burn" bson:"burn"`
	Chart       Chart                `json:"chart" bson:"chart"`
	ChartSub    ChartSub             `json:"chartSub" bson:"chartSub"`
	Recent      Recent               `json:"recent" bson:"Recent"`
}

type Vault struct {
	ChainId     string               `json:"chainId" bson:"chainId"`
	Address     string               `json:"address" bson:"address"`
	Key         bool                 `json:"key" bson:"key"`
	Name        string               `json:"name" bson:"name"`
	Symbol      string               `json:"symbol" bson:"symbol"`
	Decimals    int64                `json:"decimals" bson:"decimals"`
	Rate        primitive.Decimal128 `json:"rate" bson:"rate"`
	Ratio       primitive.Decimal128 `json:"ratio" bson:"ratio"`
	Locked      primitive.Decimal128 `json:"locked" bson:"locked"`
	ValueLocked primitive.Decimal128 `json:"valueLocked" bson:"valueLocked"`
	Weight      primitive.Decimal128 `json:"weight" bson:"weight"`
	Need        primitive.Decimal128 `json:"need" bson:"need"`
	Require     primitive.Decimal128 `json:"require" bson:"require"`
	Deposit     primitive.Decimal128 `json:"deposit" bson:"deposit"`
	Deposit24h  primitive.Decimal128 `json:"deposit24h" bson:"deposit24h"`
	Withdraw    primitive.Decimal128 `json:"withdraw" bson:"withdraw"`
	Withdraw24h primitive.Decimal128 `json:"withdraw24h" bson:"withdraw24h"`
	Mint        primitive.Decimal128 `json:"mint" bson:"mint"`
	Burn        primitive.Decimal128 `json:"burn" bson:"burn"`
	Value       primitive.Decimal128 `json:"value" bson:"value"`
	Chart       []Chart              `json:"chart" bson:"chart"`
	ChartSub    []ChartSub           `json:"chartSub" bson:"chartSub"`
	Recents     []Recent             `json:"recents" bson:"recents"`
	Last        Last                 `json:"last" bson:"last"`
}

type Chart struct {
	ChainId string               `json:"chainId" bson:"chainId"`
	Address string               `json:"market" bson:"market"`
	Time    int64                `json:"time" bson:"time"`
	Open    primitive.Decimal128 `json:"open" bson:"open"`
	High    primitive.Decimal128 `json:"high" bson:"high"`
	Low     primitive.Decimal128 `json:"low" bson:"low"`
	Close   primitive.Decimal128 `json:"close" bson:"close"`
	Volume  primitive.Decimal128 `json:"volume" bson:"volume"`
}

type ChartSub struct {
	ChainId     string               `json:"chainId" bson:"chainId"`
	Address     string               `json:"address" bson:"address"`
	Time        int64                `json:"time" bson:"time"`
	Weight      primitive.Decimal128 `json:"weight" bson:"weight"`
	Locked      primitive.Decimal128 `json:"locked" bson:"locked"`
	Value       primitive.Decimal128 `json:"value" bson:"value"`
	ValueLocked primitive.Decimal128 `json:"ValueLocked" bson:"ValueLocked"`
}
