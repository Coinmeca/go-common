package account

import "go.mongodb.org/mongo-driver/bson/primitive"

type Count struct {
	Buy   int64 `json:"buy" bson:"buy"`
	Sell  int64 `json:"sell" bson:"sell"`
	Order int64 `json:"order" bson:"order"`
	Long  int64 `json:"long" bson:"long"`
	Short int64 `json:"short" bson:"short"`
}

type Volume struct {
	Amount primitive.Decimal128 `json:"amount" bson:"amount"`
	Value  primitive.Decimal128 `json:"value" bson:"value"`
}

type Position struct {
	Asset string               `json:"asset" bson:"asset"`
	Size  primitive.Decimal128 `json:"size" bson:"size"`
}

type Asset struct {
	Account           string               `json:"account" bson:"account"`
	ChainId           string               `json:"chainId" bson:"chainId"`
	Asset             string               `json:"asset" bson:"asset"`
	Count             Count                `json:"count" bson:"count"`
	Order             primitive.Decimal128 `json:"order" bson:"order"`
	Position          []Position           `json:"position" bson:"position"`
	Leverage          primitive.Decimal128 `json:"leverage" bson:"leverage"`
	Pnl               primitive.Decimal128 `json:"pnl" bson:"pnl"`
	TotalBuy          Volume               `json:"totalBuy" bson:"totalBuy"`
	TotalSell         Volume               `json:"totalSell" bson:"totalSell"`
	TotalReturn       Volume               `json:"totalReturn" bson:"totalReturn"`
	TotalReturnRate   Volume               `json:"totalReturnRate" bson:"totalReturnRate"`
	AverageBuy        Volume               `json:"averageBuy" bson:"averageBuy"`
	AverageSell       Volume               `json:"averageSell" bson:"averageSell"`
	AverageReturn     Volume               `json:"averageReturn" bson:"averageReturn"`
	AverageReturnRate Volume               `json:"averageReturnRate" bson:"averageReturnRate"`
}
