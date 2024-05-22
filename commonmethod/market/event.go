package market

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

const (
	EventNameBid         = "Bid"
	EventNameAsk         = "Ask"
	EventNameBuy         = "Buy"
	EventNameSell        = "Sell"
	EventNameLong        = "Long"
	EventNameShort       = "Short"
	EventNameClaim       = "Claim"
	EventNameLiquidated  = "Liquidated"
	EventNameMargin      = "Margin"
	EventNameModify      = "Modify"
	EventNameOpen        = "Open"
	EventNameClose       = "Close"
	EventNameCancel      = "Cancel"
	EventNameFee         = "Fee"
	EventNameReward      = "Reward"
	EventNameThreshold   = "Threshold"
	EventNameCallLimit   = "CallLimit"
	EventNameLiquidation = "Liquidation"
)

type EventBuy struct {
	Owner    common.Address
	Sell     common.Address
	Amount   *big.Int
	Price    *big.Int
	Buy      common.Address
	Quantity *big.Int
}

type EventSell struct {
	Owner    common.Address
	Sell     common.Address
	Amount   *big.Int
	Price    *big.Int
	Buy      common.Address
	Quantity *big.Int
}

type EventBid struct {
	Owner    common.Address
	Sell     common.Address
	Amount   *big.Int
	Price    *big.Int
	Buy      common.Address
	Quantity *big.Int
}

type EventAsk struct {
	Owner    common.Address
	Sell     common.Address
	Amount   *big.Int
	Price    *big.Int
	Buy      common.Address
	Quantity *big.Int
}

type EventModify struct {
	Owner        common.Address
	BeforeAmount *big.Int
	AfterAmount  *big.Int
	BeforePrice  *big.Int
	AfterPrice   *big.Int
}

type EventClaim struct {
	Owner    common.Address
	Sell     common.Address
	Amount   *big.Int
	Price    *big.Int
	Buy      common.Address
	Quantity *big.Int
}

type EventLong struct {
	Owner     common.Address
	Pay       common.Address
	Price     *big.Int
	Size      *big.Int
	Leverage  *big.Int
	Threshold *big.Int
}

type EventShort struct {
	Owner     common.Address
	Pay       common.Address
	Price     *big.Int
	Size      *big.Int
	Leverage  *big.Int
	Threshold *big.Int
}

type EventOpen struct {
	Category  uint8
	Owner     common.Address
	Pay       common.Address
	Price     *big.Int
	Amount    *big.Int
	Leverage  *big.Int
	Threshold *big.Int
}

type EventClose struct {
	Category  uint8
	Owner     common.Address
	Pay       common.Address
	Price     *big.Int
	Amount    *big.Int
	Leverage  *big.Int
	Threshold *big.Int
}

type EventLiquidation struct {
	Category  uint8
	Owner     common.Address
	Pay       common.Address
	Price     *big.Int
	Size      *big.Int
	Leverage  *big.Int
	Threshold *big.Int
}

type EventMargin struct {
	Owner        common.Address
	BeforeAmount *big.Int
	AfterAmount  *big.Int
	BeforeMargin *big.Int
	AfterMargin  *big.Int
}

type EventCancel struct {
	Owner  common.Address
	Sell   common.Address
	Amount *big.Int
	Price  *big.Int
	Buy    common.Address
}

type EventFee struct {
	Old *big.Int
	New *big.Int
}

type EventYield struct {
	Old *big.Int
	New *big.Int
}

type EventReward struct {
	Old *big.Int
	New *big.Int
}

type EventThreshold struct {
	Old *big.Int
	New *big.Int
}

type EventCallLimit struct {
	Old *big.Int
	New *big.Int
}

type EventTransferOrder struct {
	From     common.Address
	To       common.Address
	Sell     common.Address
	Amount   *big.Int
	Buy      common.Address
	Quantity *big.Int
}

type EventTransferPosition struct {
	From     common.Address
	To       common.Address
	Category uint8
	State    uint8
	Pay      common.Address
	Amount   *big.Int
	Leverage *big.Int
}
