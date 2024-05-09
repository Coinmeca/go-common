package token

import "go.mongodb.org/mongo-driver/bson/primitive"

type Token struct {
	Address		string					`json:"address" bson:"address"`
	Name		string					`json:"name" bson:"name"`
	Symbol		string					`json:"symbol" bson:"symbol"`
	Decimals	int8					`json:"decimals" bson:"decimals"`
	Liquidity	primitive.Decimal128	`json:"liquidity" bson:"liquidity"`
}