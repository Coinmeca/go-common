package account

import (
	"github.com/ethereum/go-ethereum/common"
)

type History struct {
	Asset				common.Address	`json:"asset" bson:"address`
	TotalBuy			float64			`json:"totalBuy" bson:"totalBuy"`
	TotalSell			float64			`json:"totalSell" bson:"totalSell"`
	TotalReturn			float64			`json:"totalReturn" bson:"totalReturn"`
	TotalReturnRate		float64			`json:"totalReturnRate" bson:"totalReturnRate"`
	AverageBuy			float64			`json:"averageBuy" bson:"averageBuy"`
	AverageSell			float64			`json:"averageSell" bson:"averageSell"`
	AverageReturn		float64			`json:"averageReturn" bson:"averageReturn"`
	AverageReturnRate	float64			`json:"averageReturnRate" bson:"averageReturnRate"`
	TotalLending		float64			`json:"totalLending" bson:"totalLending"`
	Using				float64			`json:"using" bson:"using"`
}