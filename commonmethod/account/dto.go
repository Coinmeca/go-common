package account

import "go.mongodb.org/mongo-driver/bson/primitive"

type TradeType int

const (
	TradeTypeOrder TradeType = iota
	TradeTypeBid
	TradeTypeAsk
	TradeTypeBuy
	TradeTypeSell
	TradeTypeLong
	TradeTypeShort
	TradeTypeOpen
	TradeTypeClose
	TradeTypeDeposit
	TradeTypeWithdraw
	TradeTypeStake
	TradeTypeUnstake
	TradeTypeHarvest
	TradeTypeClaim
)

func (t TradeType) String() string {
	return [...]string{
		"Order",
		"Bid",
		"Ask",
		"Buy",
		"Sell",
		"Long",
		"Short",
		"Open",
		"Close",
		"Deposit",
		"Withdraw",
		"Stake",
		"Unstake",
		"Harvest",
		"Claim",
	}[t]
}

type Trade struct {
	Type    bool
	Address *string
	Amount  *primitive.Decimal128
	Value   *primitive.Decimal128
}

type OpenPosition struct {
	Account   string               `json:"account" bson:"account"`
	ToAccount string               `json:"toAccount" bson:"toAccount"`
	ChainId   string               `json:"chainId" bson:"chainId"`
	Pay       string               `json:"pay" bson:"pay"`
	Leverage  primitive.Decimal128 `json:"leverage" bson:"leverage"`
	Item      string               `json:"item" bson:"item"`
	Size      primitive.Decimal128 `json:"size" bson:"size"`
	Pnl       primitive.Decimal128 `json:"pnl" bson:"pnl"`
	Category  int64                `json:"category" bson:"category"`
	State     int64                `json:"state" bson:"state"`
}
