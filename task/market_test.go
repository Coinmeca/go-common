package task

import (
	"context"
	//cv "dex-server/internal/configs"
	cv "coinmeca-go_common/utils"
	//"dex-server/internal/model"
	"coinmeca-go_common/model"
	//repo "dex-server/repository"
	repo "coinmeca-go_common/repository"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/shopspring/decimal"
)

func TestRandomMarketVolume(t *testing.T) {
	ctx := context.Background()
	repo.InitDB(ctx, "arbitgo")
	defer repo.CloseDB()

	tokens, err := repo.MarketTokenInfo(CTX)
	if err != nil {
		panic(err)
	}

	// TODO: must increase BlockNo base
	row := model.MarketVolumeRow{
		TxHash:    "0x448ea94872267a10c944bbabf9fc505d2522ec1ad879aa4db08109fe7f7a7246",
		TxIndex:   5,
		Price:     decimal.NewFromInt(1300),
		Quantity:  decimal.NewFromInt(4),
		BlockNo:   35827978,
		Owner:     "0x5B38Da6a701c568545dCfcB03FcB875f56beddC4",
		BlockHash: "0xb1f96e9ca9043087ef3507ab247acb6f4b29da736bc268011f2f0d047751bc7f",
	}

	r := rand.New(rand.NewSource(99))
	for i := range [100]int32{} {
		seed := int32(r.Intn(100))
		price := int64(1000 + (seed * 5))

		event := "BUY"
		if !(seed%2 == 0 || seed%3 == 0) {
			event = "SELL"
		}

		row.Address = tokens[0].Address
		if i%3 != 0 {
			idx := r.Intn(6)
			row.Address = tokens[idx].Address
		}

		row.Price, row.EventType, row.Amount = decimal.NewFromInt(price), event, decimal.NewFromInt32(seed*100)
		row.TxIndex, row.BlockNo, row.Quantity = uint32(seed), row.BlockNo+uint64(i), decimal.NewFromInt32(int32(r.Intn(20)))
		repo.SetMarketVolume(ctx, row)
		//fmt.Println(row.Address, row.EventType, row.Price, row.Amount)
	}

	//b, s, v := GetMarketVolume(cv.ETHAddressMumbai)

	//b, s, v = GetMarketVolume(cv.USDTAddressMumbai)
}

func TestRandomMarketPrice(t *testing.T) {
	ctx := context.Background()
	repo.InitDB(ctx, "arbitgo")
	defer repo.CloseDB()

	tokens, err := repo.MarketTokenInfo(CTX)
	if err != nil {
		panic(err)
	}

	ts := time.Now().UTC()
	r := rand.New(rand.NewSource(99))
	for i := range [30]int32{} {
		seed := int32(r.Intn(100))
		price := int64(1000 + (seed * 10))

		minute := ((ts.Minute() + i) / 5) * 5
		hour := ts.Hour()
		if minute >= 60 {
			hour += 1
			minute -= 60
		}
		row := model.MarketPriceRow{
			Price: decimal.NewFromInt(price),
			Time:  fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", ts.Year(), ts.Month(), ts.Day(), hour, minute),
		}

		row.Address = tokens[0].Address // market address
		if i%3 != 0 {
			idx := r.Intn(6)
			row.Address = tokens[idx].Address
		}

		repo.SetMarketPrice(ctx, row)
		//fmt.Println(row.Address, row.Price, row.Time)
	}
}

func TestMarketDashboard(t *testing.T) {
	ctx := context.Background()
	repo.InitDB(ctx, "arbitgo")
	defer repo.CloseDB()

	address := cv.ETHxDAIMumbai
	repo.MarketDashboard(ctx, address)
}
