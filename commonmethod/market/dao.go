package market

type Recent struct {
	Time     string `json:"time" bson:"time"`
	Type     int    `json:"type" bson:"type"`
	Market   string `json:"market" bson:"market"`
	Owner    string `json:"owner" bson:"owner"`
	Sell     string `json:"sell" bson:"sell"`
	Price    string `json:"symbol" bson:"symbol"`
	Amount   string `json:"amount" bson:"amount"`
	Buy      string `json:"buy" bson:"buy"`
	Quantity string `json:"quantity" bson:"quantity"`
	TxHash   string `json:"txHash" bson:"txHash"`
}

type Tick struct {
	Price  string `json:"price" bson:"price"`
	Amount string `json:"amount" bson:"amount"`
}

type ChartInfo struct {
	Time   string `json:"time" bson:"time"`
	Open   string `json:"open" bson:"open"`
	High   string `json:"high" bson:"high"`
	Low    string `json:"low" bson:"low"`
	Close  string `json:"close" bson:"close"`
	Volume Volume `json:"volume" bson:"volume"`
}