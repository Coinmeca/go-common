package farm

type Recent struct {
	Type    	int     `json:"type" bson:"type"`
	Time		string	`json:"time" bson:"time"`
	Farm		string	`json:"farm" bson:"farm"`
	User    	string  `json:"owner" bson:"owner"`
	Amount		string	`json:"amount" bson:"amount"`
	Share		string	`json:"share" bson:"share"`
	TxHash  	string  `json:"txHash" bson:"txHash"`
}
