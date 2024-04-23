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
