package vault

type VaultEventInfo struct {
	TotalDeposit  string `json:"totalDeposit" bson:"totalDeposit"`
	Deposit24h    string `json:"deposit24h" bson:"deposit24h"`
	TotalWithdraw string `json:"totalWithdraw" bson:"totalWithdraw"`
	Withdraw24h   string `json:"withdraw24h" bson:"withdraw24h"`
	Mint          string `json:"mint" bson:"mint"`
	Burn          string `json:"burn" bson:"burn"`
}
