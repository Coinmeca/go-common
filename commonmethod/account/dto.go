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
	TradeTypeMargin
	TradeTypeDeposit
	TradeTypeWithdraw
	TradeTypeStake
	TradeTypeUnstake
	TradeTypeHarvest
	TradeTypeClaim
	TradeTypeBuyTransfer
	TradeTypeSellTransfer
	TradeTypeLongTransfer
	TradeTypeShortTransfer
	TradeTypeStakeTransfer
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
		"Margin",
		"Deposit",
		"Withdraw",
		"Stake",
		"Unstake",
		"Harvest",
		"Claim",
		"BuyTransfer",
		"SellTransfer",
		"LongTransfer",
		"ShortTransfer",
		"StakeTransfer",
	}[t]
}

type Trade struct {
	Type    bool
	Address *string
	Amount  *primitive.Decimal128
	Value   *primitive.Decimal128
}
