package account

import "go.mongodb.org/mongo-driver/bson/primitive"

type Count struct {
	Order    int64 `json:"order" bson:"order"`
	Buy      int64 `json:"buy" bson:"buy"`
	Sell     int64 `json:"sell" bson:"sell"`
	Long     int64 `json:"long" bson:"long"`
	Short    int64 `json:"short" bson:"short"`
	Deposit  int64 `json:"deposit" bson:"deposit"`
	Withdraw int64 `json:"withdraw" bson:"withdraw"`
	Stake    int64 `json:"stake" bson:"stake"`
	Unstake  int64 `json:"unstake" bson:"unstake"`
}

type Volume struct {
	Amount primitive.Decimal128 `json:"amount" bson:"amount"`
	Value  primitive.Decimal128 `json:"value" bson:"value"`
}

type Position struct {
	Asset string               `json:"asset" bson:"asset"`
	Size  primitive.Decimal128 `json:"size" bson:"size"`
}
type Metric struct {
	Buy        Volume `json:"buy" bson:"buy"`
	Sell       Volume `json:"sell" bson:"sell"`
	Return     Volume `json:"return" bson:"return"`
	ReturnRate Volume `json:"returnRate" bson:"returnRate"`
}

type Asset struct {
	Account  string               `json:"account" bson:"account"`
	ChainId  string               `json:"chainId" bson:"chainId"`
	Address  string               `json:"address" bson:"address"`
	Count    Count                `json:"count" bson:"count"`
	Order    primitive.Decimal128 `json:"order" bson:"order"`
	Using    primitive.Decimal128 `json:"order" bson:"order"`
	Position []Position           `json:"position" bson:"position"`
	Leverage primitive.Decimal128 `json:"leverage" bson:"leverage"`
	Pnl      primitive.Decimal128 `json:"pnl" bson:"pnl"`
	Total    Metric               `json:"total" bson:"total"`
	Average  Metric               `json:"average" bson:"average"`
}
