package market

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

type ChartInfo struct {
	From	string	`json:"from" bson:"from"`
	Sell	string	`json:"sell" bson:"sell"`
	Buy		string	`json:"buy" bson:"buy"`
	Time	string	`json:"time" bson:"time"`
	Open	string	`json:"open" bson:"open"`
	High	string	`json:"high" bson:"high"`
	Low		string	`json:"low" bson:"low"`
	Close	string	`json:"close" bson:"close"`
}