package market

type Recent struct {
	Type    	int     `json:"type" bson:"type"`
	Owner    	string  `json:"owner" bson:"owner"`
	Market		string	`json:"market" bson:"market"`
	Time    	string  `json:"time" bson:"time"`
	Buy 		string  `json:"buy" bson:"buy"`
	Price  		string  `json:"symbol" bson:"symbol"`
	Amount  	string  `json:"amount" bson:"amount"`
	Sell 		string  `json:"sell" bson:"sell"`
	Quantity    string  `json:"quantity" bson:"quantity"`
	TxHash  	string  `json:"txHash" bson:"txHash"`
}

type Tick struct {
	Price  		string  `json:"price" bson:"price"`
	Amount  	string  `json:"amount" bson:"amount"`
}

type ChartInfo struct {
	// From	string	`json:"from" bson:"from"`
	// Buy		string	`json:"buy" bson:"buy"`
	// Sell	string	`json:"sell" bson:"sell"`
	Time	string	`json:"time" bson:"time"`
	Open	string	`json:"open" bson:"open"`
	High	string	`json:"high" bson:"high"`
	Low		string	`json:"low" bson:"low"`
	Close	string	`json:"close" bson:"close"`
	Volume	Volume	`json:"volume" bson:"volume"`
}