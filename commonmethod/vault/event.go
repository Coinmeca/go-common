package vault

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
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

type EventListing struct {
	Owner    common.Address
	Token    common.Address
	Quantity *big.Int
	Proof    *big.Int
}

type EventDeposit struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
	Mint   *big.Int
}

type EventWithdraw struct {
	Owner  common.Address
	Token  common.Address
	Amount *big.Int
	Burn   *big.Int
}
