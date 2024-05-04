package farm

type Recent struct {
	Time				string			`json:"time" bson:"time"`
	Type				int				`json:"type" bson:"type"`
	Farm				string			`json:"farm" bson:"farm"`
	User				string			`json:"owner" bson:"owner"`
	Amount				string			`json:"amount" bson:"amount"`
	Share				string			`json:"share" bson:"share"`
	TxHash				string			`json:"txHash" bson:"txHash"`
}


type Last struct {
	Tvl					string			`json:"tvl" bson:"tvl"`
	Staking				string			`json:"staking" bson:"staking"`
	Interest			string			`json:"interest" bson:"interest"`
	Staking24h			string			`json:"staking24h" bson:"staking24h"`
	Unstaking24h		string			`json:"unstaking24h" bson:"unstaking24h"`
	Interest24h			string			`json:"interest24h" bson:"interest24h"`
}

type Farm struct {
	Address				string			`json:"address" bson:"address"`
	Id					string			`json:"id" bson:"id"`
	Name				string			`json:"name" bson:"name"`
	Master				string			`json:"master" bson:"master"`
	Stake				string			`json:"stake" bson:"stake"`
	Earn				string			`json:"earn" bson:"earn"`
	Tvl					string			`json:"tvl" bson:"tvl"`
	Staking				string			`json:"staking" bson:"staking"`
	interest			string			`json:"interest" bson:"interest"`
	Staking24h			string			`json:"staking24h" bson:"staking24h"`
	Unstaking24h		string			`json:"unstaking24h" bson:"unstaking24h"`
	Interest24h			string			`json:"interest24h" bson:"interest24h"`
	Recents				Recent			`json:"recents" bson:"recents"`
	Last				Last			`json:"last" bson:"bson"`
}