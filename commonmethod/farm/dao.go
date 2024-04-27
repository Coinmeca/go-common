package farm

type Recent struct {
	Type    	int     `json:"type" bson:"type"`
	Farm		string	`json:"farm" bson:"farm"`
	Owner    	string  `json:"owner" bson:"owner"`
	Amount		string	`json:"amount" bson:"amount"`
	TxHash  	string  `json:"txHash" bson:"txHash"`
}
