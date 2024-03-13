package vault

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type EventWithdraw struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
	Meca   *big.Int
}

type EventDeposit struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
	Meca   *big.Int
}
