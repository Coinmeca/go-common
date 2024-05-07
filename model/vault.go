package model

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
)

type VaultAbiInfo struct {
	Key      bool
	Addr     common.Address
	Name     string
	Symbol   string
	Decimals uint8
	Treasury *big.Int
	Rate     *big.Int
	Weight   *big.Int
	Need     *big.Int
}

type VaultEventAbiInfo struct { // deposit, withdraw
	Address common.Address
	Token   common.Address
	Amount  *big.Int
	Meca    *big.Int // deposit +meca, withdraw -meca
}

type VaultVolumeRow struct {
	Event        string
	EventType    string
	Owner        string
	BlockHash    string
	BlockNo      uint64
	TxHash       string
	TxIndex      uint32
	Token        string
	Symbol       string
	Amount       decimal.Decimal // exchanged
	MecaQuantity decimal.Decimal
}

type VaultTokenRow struct {
	IsKey    bool
	Address  string
	Name     string
	Symbol   string
	Decimals uint8
}

type VaultPriceRow struct {
	Time     string
	Address  string
	Symbol   string
	Treasury decimal.Decimal // balance
	Price    decimal.Decimal // rate to exchange
	Weight   decimal.Decimal // token slot has meca
	Need     decimal.Decimal
	ValueLocked      decimal.Decimal // by usd
}

type VaultHistoryApi struct {
	EventType string  `json:"type"`
	Volume    float64 `json:"volume"`
	Meca      float64 `json:"meca"`
	Share     float64 `json:"share"`
}

type VaultDashboardApi struct {
	Exchange         float64 `json:"exchange"`
	ExchangeChange   float64 `json:"exchange_change"`
	ExchangeRate     float64 `json:"exchange_change_rate"`
	TotalLocked      float64 `json:"total_locked"`
	TotalChange      float64 `json:"total_change"`
	TotalValueLocked float64 `json:"total_value_locked"`
	ValueLockedChange        float64 `json:"valueLocked_change"`
	Weight           float64 `json:"weight"`
	WeightChange     float64 `json:"weight_change"`
	Deposit          float64 `json:"deposit"`
	Withdraw         float64 `json:"withdraw"`
	Earn             float64 `json:"earn"`
	Burn             float64 `json:"burn"`
	MecaPerToken     float64 `json:"meca_per_token"`
	TokenPerMeca     float64 `json:"token_per_meca"`
}

type VaultOverviewApi struct {
	IsKey          bool    `json:"key"`
	Name           string  `json:"name"`
	Decimals       int32   `json:"decimals"`
	Symbol         string  `json:"symbol"`
	Address        string  `json:"address"`
	Exchange       float64 `json:"exchange_rate"`
	ExchangeChange float64 `json:"exchange_rate_change"`
	TVL            float64 `json:"valueLocked"`
	TVLChange      float64 `json:"valueLocked_change"`
	Volume         float64 `json:"total_volume"`
	VolumeChange   float64 `json:"total_volume_change"`
}
