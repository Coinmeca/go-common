package market

import (
	"github.com/coinmeca/go-common/commonmethod/token"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Recent struct {
	Time				int64					`json:"time" bson:"time"`
	Type				int						`json:"type" bson:"type"`
	Market				string					`json:"market" bson:"market"`
	User				string					`json:"user" bson:"user"`
	Sell				string					`json:"sell" bson:"sell"`
	Price				primitive.Decimal128	`json:"symbol" bson:"symbol"`
	Amount				primitive.Decimal128	`json:"amount" bson:"amount"`
	Buy					string					`json:"buy" bson:"buy"`
	Quantity			primitive.Decimal128	`json:"quantity" bson:"quantity"`
	TxHash				string					`json:"txHash" bson:"txHash"`
}

type Tick struct {
	Price				primitive.Decimal128	`json:"price" bson:"price"`
	Balance				primitive.Decimal128	`json:"amount" bson:"amount"`
}

type Liquidity struct {
	Base				primitive.Decimal128	`json:"base" bson:"base"`
	Quote				primitive.Decimal128	`json:"quote" bson:"quote"`
}

type Volume struct {
	Base				primitive.Decimal128	`json:"base" bson:"base"`
	Quote				primitive.Decimal128	`json:"quote" bson:"quote"`
}

type Chart struct {
	Time				int64					`json:"time" bson:"time"`
	Open				primitive.Decimal128	`json:"open" bson:"open"`
	High				primitive.Decimal128	`json:"high" bson:"high"`
	Low					primitive.Decimal128	`json:"low" bson:"low"`
	Close				primitive.Decimal128	`json:"close" bson:"close"`
	Volume				Volume					`json:"volume" bson:"volume"`
}

type Orderbook struct {
	Asks				[]Tick					`json:"asks" bson:"asks"`
	Bids				[]Tick					`json:"bids" bson:"bids"`
}

type Last struct {
	Price				primitive.Decimal128	`json:"price" bson:"price"`
	High				primitive.Decimal128	`json:"high" bson:"high"`
	Low					primitive.Decimal128	`json:"low" bson:"low"`
	Volume				Volume					`json:"volume" bson:"volume"`
	Chart				Chart					`json:"chart" bson:"chart"`
	Recent				Recent					`json:"recent" bson:"recent"`
}

type Market struct {
	Address				string					`json:"address" bson:"address"`
	Name				string					`json:"name" bson:"name"`
	Symbol				string					`json:"symbol" bson:"symbol"`
	Base				token.Token				`json:"base" bson:"base"`
	Quote				token.Token				`json:"quote" bson:"quote"`
	Price				primitive.Decimal128	`json:"price" bson:"price"`
	Liquidity			Liquidity				`json:"liquidity" bson:"liquidity"`
	Volume				Volume					`json:"volume" bson:"volume"`
	Tick				primitive.Decimal128	`json:"tick" bson:"tick"`
	Fee					primitive.Decimal128	`json:"fee" bson:"fee"`
	Lock				bool					`json:"lock" bson:"lock"`
	Threshold			primitive.Decimal128	`json:"threshold" bson:"threshold"`
	Orderbook			Orderbook				`json:"orderbook" bson:"orderbook"`
	Recents				[]Recent				`json:"recents" bson:"recents"`
	Last				Last					`json:"last" bson:"last"`
}