package farm

import (
	"github.com/coinmeca/go-common/commonmethod/token"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recent struct {
	Time				int64					`json:"time" bson:"time"`
	Type				int						`json:"type" bson:"type"`
	Farm				string					`json:"farm" bson:"farm"`
	User				string					`json:"owner" bson:"owner"`
	Amount				primitive.Decimal128	`json:"amount" bson:"amount"`
	Share				primitive.Decimal128	`json:"share" bson:"share"`
	TxHash				string					`json:"txHash" bson:"txHash"`
}


type Last struct {
	Staking				primitive.Decimal128	`json:"staking" bson:"staking"`
	Interest			primitive.Decimal128	`json:"interest" bson:"interest"`
	Staking24h			primitive.Decimal128	`json:"staking24h" bson:"staking24h"`
	Unstaking24h		primitive.Decimal128	`json:"unstaking24h" bson:"unstaking24h"`
	Interest24h			primitive.Decimal128	`json:"interest24h" bson:"interest24h"`
	ValueLocked			primitive.Decimal128	`json:"valueLocked" bson:"valueLocked"`
	Recent				Recent					`json:"recent" bson:"recent"`
}

type Farm struct {
	Address				string					`json:"address" bson:"address"`
	Id					string					`json:"id" bson:"id"`
	Name				string					`json:"name" bson:"name"`
	Master				string					`json:"master" bson:"master"`
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