package token

import "go.mongodb.org/mongo-driver/bson/primitive"

type Token struct {
	Address  string               `json:"address" bson:"address"`
	Name     string               `json:"name" bson:"name"`
	Symbol   string               `json:"symbol" bson:"symbol"`
	Decimals int64                `json:"decimals" bson:"decimals"`
	Value    primitive.Decimal128 `json:"value" bson:"value"`
}
