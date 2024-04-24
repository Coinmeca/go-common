package market

type MarketInfo struct {
	Address 	string		`json:"address" bson:"address"`
	Base 		string		`json:"base" bson:"base"`
	Quote 		string		`json:"quote" bson:"quote"`
	Price  		float64		`json:"price" bson:"price"`
	Tick  		float64		`json:"tick" bson:"tick"`
	Fee		  	float64		`json:"fee" bson:"fee"`
	Threshold	string		`json:"threshold" bson:"threshold"`
	Lock		bool		`json:"lock" bson:"lock"`
	Orderbook	Orderbook	`json:"orderbook" bson:"orderbook"`
}

type Recent struct {
	Type    	int     `json:"type" bson:"type"`
	From    	string  `json:"from" bson:"from"`
	Address 	string  `json:"address" bson:"address"`
	Buy 		string  `json:"buy" bson:"buy"`
	Sell 		string  `json:"sell" bson:"sell"`
	Time    	string  `json:"time" bson:"time"`
	Price  		string  `json:"symbol" bson:"symbol"`
	Amount  	string  `json:"amount" bson:"amount"`
	Quantity    string  `json:"quantity" bson:"quantity"`
	TxHash  	string  `json:"txHash" bson:"txHash"`
}

type Tick struct {
	Price  		string  `json:"symbol" bson:"symbol"`
	Amount  	string  `json:"amount" bson:"amount"`
}