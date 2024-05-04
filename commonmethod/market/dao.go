package market

type Recent struct {
	Time				string			`json:"time" bson:"time"`
	Type				int				`json:"type" bson:"type"`
	Market				string			`json:"market" bson:"market"`
	Owner				string			`json:"owner" bson:"owner"`
	Sell				string			`json:"sell" bson:"sell"`
	Price				string			`json:"symbol" bson:"symbol"`
	Amount				string			`json:"amount" bson:"amount"`
	Buy					string			`json:"buy" bson:"buy"`
	Quantity			string			`json:"quantity" bson:"quantity"`
	TxHash				string			`json:"txHash" bson:"txHash"`
}

type Tick struct {
	Price				string			`json:"price" bson:"price"`
	Balance				string			`json:"amount" bson:"amount"`
}

type Volume struct {
	Base				float64			`json:"base" bson:"base"`
	Quote				float64			`json:"quote" bson:"quote"`
}

type Chart struct {
	Time				string			`json:"time" bson:"time"`
	Open				string			`json:"open" bson:"open"`
	High				string			`json:"high" bson:"high"`
	Low					string			`json:"low" bson:"low"`
	Close				string			`json:"close" bson:"close"`
	Volume				Volume			`json:"volume" bson:"volume"`
}

type Orderbook struct {
	Asks				[]Tick			`json:"asks" bson:"asks"`
	Bids				[]Tick			`json:"bids" bson:"bids"`
}

type MarketLast struct {
	Price				float64			`json:"price" bson:"price"`
	Chart				Chart			`json:"chart" bson:"chart"`
}

type Market struct {
	Address				string			`json:"address" bson:"address"`
	Base				string			`json:"symbol" bson:"symbol"`
	Quote				string			`json:"name" bson:"name"`
	Price				string			`json:"price" bson:"price"`
	Volume				Volume			`json:"volume" bson:"volume"`
	Tick				string			`json:"tick" bson:"tick"`
	Fee					string			`json:"fee" bson:"fee"`
	Threshold			string			`json:"threshold" bson:"threshold"`
	Orderbook			Orderbook		`json:"orderbook" bson:"orderbook"`
	Last				MarketLast		`json:"last" bson:"last"`
}