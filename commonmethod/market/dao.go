package market

type Recent struct {
	Type    	int     `json:"type" bson:"type"`
	From    	string  `json:"from" bson:"from"`
	Address 	string  `json:"address" bson:"address"`
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