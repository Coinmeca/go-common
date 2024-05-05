package vault

import "time"

type Recent struct {
	Time				string			`json:"time" bson:"time"`
	Type				int				`json:"type" bson:"type"`
	User				string			`json:"user" bson:"user"`
	Token				string			`json:"token" bson:"token"`
	Symbol				string			`json:"symbol" bson:"symbol"`
	Volume				string			`json:"volume" bson:"volume"`
	Meca				string			`json:"meca" bson:"meca"`
	Share				string			`json:"share" bson:"share"`
	TxHash				string			`json:"txHash" bson:"txHash"`
	UpdateAt			time.Time		`json:"updateAt" bson:"updateAt"`
}

type VaultLast struct {
	Exchange			string			`json:"exchange" bson:"exchange"`
	Weight				string			`json:"weight" bson:"weight"`
	Tl					string			`json:"tl" bson:"tl"`
	Tvl					string			`json:"tvl" bson:"tvl"`
	Value				string			`json:"value" bson:"value"`
	Recent				Recent			`json:"recent" bson:"Recent"`
}

type Vault struct {
	Key					bool			`json:"key" bson:"key"`
	Address				string			`json:"address" bson:"address"`
	Symbol				string			`json:"symbol" bson:"symbol"`
	Name				string			`json:"name" bson:"name"`
	Decimals			int				`json:"decimals" bson:"decimals"`
	Exchange			string			`json:"exchange" bson:"exchange"`
	Rate				string			`json:"rate" bson:"rate"`
	Tl					string			`json:"tl" bson:"tl"`
	Tvl					string			`json:"tvl" bson:"tvl"`
	Weight				string			`json:"weight" bson:"weight"`
	Deposit				string			`json:"deposit" bson:"deposit"`
	Deposit24h			string			`json:"deposit24h" bson:"deposit24h"`
	Withdraw			string			`json:"withdraw" bson:"withdraw"`
	Withdraw24h			string			`json:"withdraw24h" bson:"withdraw24h"`
	PerToken			string			`json:"perToken" bson:"perToken"`
	TokenPer			string			`json:"tokenPer" bson:"tokenPer"`
	Mint				string			`json:"mint" bson:"mint"`
	Burn				string			`json:"burn" bson:"burn"`
	Value				string			`json:"value" bson:"value"`
	Recents				[]Recent		`json:"recents" bson:"recents"`
	Last				VaultLast		`json:"last" bson:"last"`
}