package farm

type Recent struct {
	Type    	int     `json:"type" bson:"type"`
	Owner    	string  `json:"owner" bson:"owner"`
	Farm		string	`json:"farm" bson:"farm"`
	Amount		string	`json:"amount bson:"amount"`
	TxHash  	string  `json:"txHash" bson:"txHash"`
}
