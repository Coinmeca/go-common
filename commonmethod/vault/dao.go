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
	Share				float64			`json:"share" bson:"share"`
	TxHash				string			`json:"txHash" bson:"txHash"`
	UpdateAt			time.Time		`json:"updateAt" bson:"updateAt"`
}

type VaultLast struct {
	Exchange			float64			`json:"exchange" bson:"exchange"`
	Tl					float64			`json:"tl" bson:"tl"`
	Tvl					float64			`json:"tvl" bson:"tvl"`
	Weight				float64			`json:"weight" bson:"weight"`
	Recent				Recent			`json:"recent" bson:"Recent"`
}

type Vault struct {
	Key					bool			`json:"key" bson:"key"`
	Address				string			`json:"address" bson:"address"`
	Symbol				string			`json:"symbol" bson:"symbol"`
	Name				string			`json:"name" bson:"name"`
	Decimals			int				`json:"decimals" bson:"decimals"`
	Exchange			float64			`json:"exchange" bson:"exchange"`
	Rate				float64			`json:"rate" bson:"rate"`
	Tl					float64			`json:"tl" bson:"tl"`
	Tvl					float64			`json:"tvl" bson:"tvl"`
	Weight				float64			`json:"weight" bson:"weight"`
	Deposit				float64			`json:"deposit" bson:"deposit"`
	Deposit24h			float64			`json:"deposit24h" bson:"deposit24h"`
	Withdraw			float64			`json:"withdraw" bson:"withdraw"`
	Withdraw24h			float64			`json:"withdraw24h" bson:"withdraw24h"`
	PerToken			float64			`json:"perToken" bson:"perToken"`
	TokenPer			float64			`json:"tokenPer" bson:"tokenPer"`
	Mint				float64			`json:"mint" bson:"mint"`
	Burn				float64			`json:"burn" bson:"burn"`
}
