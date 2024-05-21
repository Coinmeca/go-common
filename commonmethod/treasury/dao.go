package treasury

import "go.mongodb.org/mongo-driver/bson/primitive"

type Treasury struct {
	ChainId				string					`json:"chainId" bson:"chainId"`
	Tvl					primitive.Decimal128	`json:"tvl" bson:"tvl"`
	Tv					primitive.Decimal128	`json:"tv" bson:"tv"`
	Tw					primitive.Decimal128	`json:"tw" bson:"tw"`
}

type Chart struct {
	ChainId				string					`json:"chainId" bson:"chainId"`
	Time				int64					`json:"time" bson:"time"`
	Tvl					primitive.Decimal128	`json:"tvl" bson:"tvl"`
	Tv					primitive.Decimal128	`json:"tv" bson:"tv"`
	Tw					primitive.Decimal128	`json:"tw" bson:"tw"`
} 