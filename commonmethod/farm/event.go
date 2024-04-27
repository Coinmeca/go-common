package farm

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type EventStake struct {
	User   common.Address
	Amount *big.Int
}

type EventUnstake struct {
	User   common.Address
	Amount *big.Int
}

type EventHarvest struct {
	User   common.Address
	Amount *big.Int
}

type EventClaim struct {
	User   common.Address
	Amount *big.Int
}

type EventMain struct {
	Farm  common.Address
	Token common.Address
	Id    *big.Int
}

type EventDerive struct {
	Main     common.Address
	Token    common.Address
	Id       *big.Int
	Derive   common.Address
	RToken   common.Address
	DeriveId *big.Int
}