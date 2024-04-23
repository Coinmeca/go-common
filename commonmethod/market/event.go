package market

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type EventBuy struct {
	Owner  		common.Address
	Sell		common.Address
	Amount 		*big.Int
	Price  		*big.Int
	Buy  		common.Address
	Quantity	*big.Int
}

type EventSell struct {
	Owner  		common.Address
	Sell		common.Address
	Amount 		*big.Int
	Price  		*big.Int
	Buy  		common.Address
	Quantity	*big.Int
}

type EventAsk struct {
	Owner  		common.Address
	Sell		common.Address
	Amount 		*big.Int
	Price  		*big.Int
	Buy  		common.Address
}

type EventBid struct {
	Owner  		common.Address
	Sell		common.Address
	Amount 		*big.Int
	Price  		*big.Int
	Buy  		common.Address
}

type EventClaim struct {
	Owner  		common.Address
	Sell		common.Address
	Amount 		*big.Int
	Price  		*big.Int
	Buy  		common.Address
	Quantity	*big.Int
}

type EventLong struct {
	Owner  		common.Address
	Pay			common.Address
	Price  		*big.Int
	Size 		*big.Int
	Leverage	*big.Int
	Threshold	*big.Int
	Item  		common.Address
	Quantity	*big.Int
}

type EventShort struct {
	Owner  		common.Address
	Pay			common.Address
	Price  		*big.Int
	Size 		*big.Int
	Leverage	*big.Int
	Threshold	*big.Int
}

type EventOpen struct {
	Category	*big.Int
	Owner  		common.Address
	Pay			common.Address
	Price  		*big.Int
	Size 		*big.Int
	Leverage	*big.Int
	Threshold	*big.Int
}

type EventClose struct {
	Category	*big.Int
	Owner  		common.Address
	Pay			common.Address
	Price  		*big.Int
	Size 		*big.Int
	Leverage	*big.Int
	Threshold	*big.Int
}

type EventLiquidation struct {
	Category	*big.Int
	Owner  		common.Address
	Pay			common.Address
	Price  		*big.Int
	Size 		*big.Int
	Leverage	*big.Int
	Threshold	*big.Int
}
