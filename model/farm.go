package model

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/shopspring/decimal"
	"math/big"
)

type FarmAbiInfo struct {
	Farm        common.Address `json:"farm"`
	Name        string         `json:"name"`
	Stake       common.Address `json:"stake"`
	StakeSymbol string         `json:"stake_symbol"`
	StakeName   string         `json:"stake_name"`
	Earn        common.Address `json:"earn"`
	EarnSymbol  string         `json:"earn_symbol"`
	EarnName    string         `json:"earn_name"`
	Start       *big.Int       `json:"start"`
	Period      *big.Int       `json:"period"`
	Goal        *big.Int       `json:"goal"`
	Rewards     *big.Int       `json:"rewards"`
	Locked      *big.Int       `json:"locked"`
}

type FarmRow struct {
	Time        string
	Address     string
	Name        string
	Stake       string
	StakeSymbol string
	StakeName   string
	Earn        string
	EarnSymbol  string
	EarnName    string
	Start       uint64
	Period      uint64
	Goal        uint64
	Rewards     decimal.Decimal
	Locked      decimal.Decimal
}

type FarmListApi struct {
	Address string  `json:"address"`
	Rewards float64 `json:"rewards"`
	//RewardsRate float64 `json:"rewards_rate"`
	Locked float64 `json:"locked"`
	//LockedRate  float64 `json:"locked_rate"`
	Stake       string `json:"stake"`
	StakeSymbol string `json:"stake_symbol"`
	StakeName   string `json:"stake_name"`
	Earn        string `json:"earn"`
	EarnSymbol  string `json:"earn_symbol"`
	EarnName    string `json:"earn_name"`
	Name        string `json:"name"`
	Start       uint64 `json:"start"`
	Period      uint64 `json:"period"`
	Goal        uint64 `json:"goal"`
}

type FarmLockedChartApi struct {
	Datetime string  `json:"time"`
	Rewards  float64 `json:"rewards"`
	Locked   float64 `json:"locked"`
}
