package account

import "go.mongodb.org/mongo-driver/bson/primitive"

type OpenPosition struct {
	Account  string               `json:"account" bson:"account"`
	ChainId  string               `json:"chainId" bson:"chainId"`
	Pay      string               `json:"pay" bson:"pay"`
	Leverage primitive.Decimal128 `json:"leverage" bson:"leverage"`
	Item     string               `json:"item" bson:"item"`

	Size primitive.Decimal128 `json:"size" bson:"size"`
	Pnl  primitive.Decimal128 `json:"pnl" bson:"pnl"`
}
