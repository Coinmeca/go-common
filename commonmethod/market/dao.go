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
	Amount				string			`json:"amount" bson:"amount"`
}

type ChartInfo struct {
	Time				string			`json:"time" bson:"time"`
	Open				string			`json:"open" bson:"open"`
	High				string			`json:"high" bson:"high"`
	Low					string			`json:"low" bson:"low"`
	Close				string			`json:"close" bson:"close"`
	Volume				Volume			`json:"volume" bson:"volume"`
}

type MarketOrderbook struct {
	Asks				[]Tick			`json:"asks" bson:"asks"`
	Bids				[]Tick			`json:"bids" bson:"bids"`
}

type MarketVolume struct {
	Base				float64			`json:"base" bson:"base"`
	Quote				float64			`json:"quote" bson:"quote"`
}

type MarketLast struct {
	Price				float64			`json:"price" bson:"price"`
	Chart				ChartInfo		`json:"chart" bson:"chart"`
}

type Vault struct {
	Address				string			`json:"address" bson:"address"`
	Base				string			`json:"symbol" bson:"symbol"`
	Quote				string			`json:"name" bson:"name"`
	Price				float64			`json:"price" bson:"price"`
	Volume				MarketVolume	`json:"volume" bson:"volume"`
	Tick				float64			`json:"tick" bson:"tick"`
	Fee					float64			`json:"fee" bson:"fee"`
	Threshold			float64			`json:"threshold" bson:"threshold"`
	Orderbook			float64			`json:"orderbook" bson:"orderbook"`
	Last				MarketLast		`json:"last" bson:"last"`
}