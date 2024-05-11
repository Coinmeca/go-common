package vault

import (
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

const (
	EventNameKeyToken = "KeyToken"
	EventNameFee      = "Fee"
	EventNameReward   = "Reward"

	EventNameDeposit    = "Deposit"
	EventNameWithdraw   = "Withdraw"
	EventNameListing    = "Listing"
	EventNameDelisting  = "Delisting"
	EventNamePermission = "Permission"
)

type EventWithdraw struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
	Burn   *big.Int
}

type EventDeposit struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
	Mint   *big.Int
}
