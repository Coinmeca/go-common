package token

import "go.mongodb.org/mongo-driver/bson/primitive"

type Token struct {
	Address		string               `json:"address" bson:"address"`
	Name		string               `json:"name" bson:"name"`
	Symbol		string               `json:"symbol" bson:"symbol"`
	Decimal		primitive.Decimal128 `json:"decimal" bson:"decimal"`
	Liquidity	primitive.Decimal128 `json:"decimal" bson:"decimal"`
}