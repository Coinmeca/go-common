package vault

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type EventWithdraw struct {
	Owner  common.Address `json:"_owner"`
	Token  common.Address `json:"_token"`
	Amount *big.Int       `json:"_amount"`
	Meca   *big.Int       `json:"_meca"`
}
