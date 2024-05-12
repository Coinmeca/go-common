package farm

import (
	"github.com/coinmeca/go-common/commonmethod/token"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recent struct {
	ChainId				string					`json:"chainId" bson:"chainId"`
	Address				string					`json:"address" bson:"address"`
	Time				int64					`json:"time" bson:"time"`
	Type				int						`json:"type" bson:"type"`
	User				string					`json:"owner" bson:"owner"`
	Amount				primitive.Decimal128	`json:"amount" bson:"amount"`
	Share				primitive.Decimal128	`json:"share" bson:"share"`
	TxHash				string					`json:"txHash" bson:"txHash"`
	UpdateAt			string					`json:"updateAt" bson:"updateAt"`
}


type Last struct {
	ChainId				string					`json:"chainId" bson:"chainId"`
	Address				string					`json:"address" bson:"address"`
	Staking				primitive.Decimal128	`json:"staking" bson:"staking"`
	Interest			primitive.Decimal128	`json:"interest" bson:"interest"`
	Staking24h			primitive.Decimal128	`json:"staking24h" bson:"staking24h"`
	Unstaking24h		primitive.Decimal128	`json:"unstaking24h" bson:"unstaking24h"`
	Interest24h			primitive.Decimal128	`json:"interest24h" bson:"interest24h"`
	ValueLocked			primitive.Decimal128	`json:"valueLocked" bson:"valueLocked"`
	Recent				Recent					`json:"recent" bson:"recent"`
}

type Farm struct {
	ChainId				string					`json:"chainId" bson:"chainId"`
	Address				string					`json:"address" bson:"address"`
	Id					int64					`json:"id" bson:"id"`
	Master				string					`json:"master" bson:"master"`
	Name				string					`json:"name" bson:"name"`
	Symbol				string					`json:"symbol" bson:"symbol"`
	Decimals			int64					`json:"decimals" bson:"decimals"`
	Stake				token.Token				`json:"stake" bson:"stake"`
	Earn				token.Token				`json:"earn" bson:"earn"`
	Staking				primitive.Decimal128	`json:"staking" bson:"staking"`
	Interest			primitive.Decimal128	`json:"interest" bson:"interest"`
	Staking24h			primitive.Decimal128	`json:"staking24h" bson:"staking24h"`
	Unstaking24h		primitive.Decimal128	`json:"unstaking24h" bson:"unstaking24h"`
	Interest24h			primitive.Decimal128	`json:"interest24h" bson:"interest24h"`
	ValueLocked			primitive.Decimal128	`json:"valueLocked" bson:"valueLocked"`
	Recents				[]Recent				`json:"recents" bson:"recents"`
	Last				Last					`json:"last" bson:"bson"`
}