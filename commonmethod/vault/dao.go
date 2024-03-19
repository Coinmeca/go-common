package vault

type Recent struct {
	Type    int     `json:"type" bson:"type"`
	Symbol  string  `json:"symbol" bson:"symbol"`
	Address string  `json:"address" bson:"address"`
	Time    string  `json:"time" bson:"time"`
	Volume  string  `json:"volume" bson:"volume"`
	Meca    string  `json:"meca" bson:"meca"`
	Share   float64 `json:"share" bson:"share"`
	TxHash  string  `json:"txHash" bson:"txHash"`
}
