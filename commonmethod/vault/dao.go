package vault

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recent struct {
	Time				int64					`json:"time" bson:"time"`
	Type				int						`json:"type" bson:"type"`
	User				string					`json:"user" bson:"user"`
	Token				string					`json:"token" bson:"token"`
	Amount				primitive.Decimal128	`json:"volume" bson:"volume"`
	Meca				primitive.Decimal128	`json:"meca" bson:"meca"`
	Share				primitive.Decimal128	`json:"share" bson:"share"`
	TxHash				string					`json:"txHash" bson:"txHash"`
	UpdateAt			time.Time				`json:"updateAt" bson:"updateAt"`
}

type Last struct {
	Exchange			primitive.Decimal128	`json:"exchange" bson:"exchange"`
	Weight				primitive.Decimal128	`json:"weight" bson:"weight"`
	Tl					primitive.Decimal128	`json:"tl" bson:"tl"`
	Tvl					primitive.Decimal128	`json:"tvl" bson:"tvl"`
	Value				primitive.Decimal128	`json:"value" bson:"value"`
	Recent				Recent					`json:"recent" bson:"Recent"`
}

type Vault struct {
	Key					bool					`json:"key" bson:"key"`
	Address				string					`json:"address" bson:"address"`
	Symbol				string					`json:"symbol" bson:"symbol"`
	Name				string					`json:"name" bson:"name"`
	Decimals			int						`json:"decimals" bson:"decimals"`
	Exchange			primitive.Decimal128	`json:"exchange" bson:"exchange"`
	Rate				primitive.Decimal128	`json:"rate" bson:"rate"`
	Tl					primitive.Decimal128	`json:"tl" bson:"tl"`
	Tvl					primitive.Decimal128	`json:"tvl" bson:"tvl"`
	Weight				primitive.Decimal128	`json:"weight" bson:"weight"`
	Deposit				primitive.Decimal128	`json:"deposit" bson:"deposit"`
	Deposit24h			primitive.Decimal128	`json:"deposit24h" bson:"deposit24h"`
	Withdraw			primitive.Decimal128	`json:"withdraw" bson:"withdraw"`
	Withdraw24h			primitive.Decimal128	`json:"withdraw24h" bson:"withdraw24h"`
	PerToken			primitive.Decimal128	`json:"perToken" bson:"perToken"`
	TokenPer			primitive.Decimal128	`json:"tokenPer" bson:"tokenPer"`
	Mint				primitive.Decimal128	`json:"mint" bson:"mint"`
	Burn				primitive.Decimal128	`json:"burn" bson:"burn"`
	Value				primitive.Decimal128	`json:"value" bson:"value"`
	Recents				[]Recent				`json:"recents" bson:"recents"`
	Last				Last					`json:"last" bson:"last"`
}