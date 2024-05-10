package treasury

import "go.mongodb.org/mongo-driver/bson/primitive"

type Info struct {
	ChainId				int64					`json:"chainId" bson:"chainId"`
	TVL					primitive.Decimal128	`json:"tvl" bson:"tvl"`
	TV					primitive.Decimal128	`json:"tv" bson:"tv"`
	TW					primitive.Decimal128	`json:"tw" bson:"tw"`
}

type Chart struct {
	ChainId				int64					`json:"chainId" bson:"chainId"`
	Time				int64					`json:"time" bson:"time"`
	TVL					primitive.Decimal128	`json:"tvl" bson:"tvl"`
	TV					primitive.Decimal128	`json:"tv" bson:"tv"`
	TW					primitive.Decimal128	`json:"tw" bson:"tw"`
}