package farm

type FarmTokenInfo struct {
	Decimals int    `json:"decimals" bson:"decimals"`
	Address  string `json:"address" bson:"address"`
	Symbol   string `json:"symbol" bson:"symbol"`
	Name     string `json:"name" bson:"name"`
}
