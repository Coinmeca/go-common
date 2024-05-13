package farm

type FarmTokenInfo struct {
	Address  	string	`json:"address" bson:"address"`
	Name     	string	`json:"name" bson:"name"`
	Symbol   	string	`json:"symbol" bson:"symbol"`
	Decimals 	int64	`json:"decimals" bson:"decimals"`
}

type Open struct {
	ChainId		string	`json:"chainId" bson:"chainId"`
	Address		string	`json:"address" bson:"address"`	
	Token		string	`json:"token" bson:"token"`	
	Id			int64	`json:"id" bson:"id"`
}

type OpenDerive struct {
	ChainId		string	`json:"chainId" bson:"chainId"`
	MainAddress	string	`json:"mainAddress" bson:"mainAddress"`	
	MainToken	string	`json:"mainToken" bson:"mainToken"`	
	MainId		int64	`json:"mainId" bson:"mainId"`
	Address		string	`json:"address" bson:"address"`	
	Token		string	`json:"token" bson:"token"`	
	Id			int64	`json:"id" bson:"id"`
}