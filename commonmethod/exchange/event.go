package exchange

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	EventNameOpen       = "Open"
	EventNamePermission = "Permission"
	EventNameClose      = "Close"
)

type EventOpen struct {
	Market common.Address
	Base   common.Address
	Quote  common.Address
	Price  *big.Int
	Nft    common.Address
	App    common.Address
}
