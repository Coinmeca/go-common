package market

type Recent struct {
	Time				string			`json:"time" bson:"time"`
	Type				int				`json:"type" bson:"type"`
	Market				string			`json:"market" bson:"market"`
	User				string			`json:"user" bson:"user"`
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

type Liquidity struct {
	Base				string			`json:"base" bson:"base"`
	Quote				string			`json:"quote" bson:"quote"`
}

type Volume struct {
	Base				string			`json:"base" bson:"base"`
	Quote				string			`json:"quote" bson:"quote"`
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

type Last struct {
	Price				string			`json:"price" bson:"price"`
	High				string			`json:"high" bson:"high"`
	Low					string			`json:"low" bson:"low"`
	Volume				Volume			`json:"volume" bson:"volume"`
	Chart				Chart			`json:"chart" bson:"chart"`
	Recent				Recent			`json:"recent" bson:"recent"`
}

type Market struct {
	Address				string			`json:"address" bson:"address"`
	Name				string			`json:"name" bson:"name"`
	Symbol				string			`json:"symbol" bson:"symbol"`
	Base				string			`json:"base" bson:"base"`
	Quote				string			`json:"quote" bson:"quote"`
	Price				string			`json:"price" bson:"price"`
	Liquidity			Liquidity		`json:"liquidity" bson:"liquidity"`
	Volume				Volume			`json:"volume" bson:"volume"`
	Tick				string			`json:"tick" bson:"tick"`
	Fee					string			`json:"fee" bson:"fee"`
	Lock				bool			`json:"lock" bson:"lock"`
	Threshold			string			`json:"threshold" bson:"threshold"`
	Orderbook			Orderbook		`json:"orderbook" bson:"orderbook"`
	Recents				[]Recent		`json:"recents" bson:"recents"`
	Last				Last			`json:"last" bson:"last"`
}