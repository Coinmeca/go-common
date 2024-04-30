package task

import (
	"context"
	//"dex-server/internal/model"
	"github.com/coinmeca/go-common/model"
	//repo "dex-server/repository"
	repo "github.com/coinmeca/go-common/repository"
	"fmt"
	"github.com/shopspring/decimal"
	"math/big"
	"math/rand"
	"testing"
	"time"
)

func TestRandIndex(t *testing.T) {
	symbols := []string{"DAI", "USDT", "USDC", "MATIC", "ETH"}
	r := rand.New(rand.NewSource(99))
	for range [50]int32{} {
		idx := r.Intn(5)
		symbol := symbols[idx]
		fmt.Println(symbol)
	}
}

// stats
func TestRandomVaultPrice(t *testing.T) {
	ctx := context.Background()
	repo.InitDB(ctx, "arbitgo")
	defer repo.CloseDB()

	tokens, err := repo.MarketTokenInfo(CTX)
	if err != nil {
		panic(err)
	}

	ts := time.Now().UTC().Add(time.Hour * -5)
	r := rand.New(rand.NewSource(99))
	for i := range [100]int32{} {
		seed := int64(r.Intn(100))
		factor := int64(r.Intn(10))

		price := decimal.NewFromBigInt(big.NewInt(400+factor), -3)
		treasury := decimal.NewFromInt(750000 + (seed * 1000))
		weight := decimal.NewFromInt(400000 + (seed * 1000))

		idx := r.Intn(5)
		symbol := tokens[idx].Symbol
		switch symbol {
		case "ETH":
			price = decimal.NewFromInt(1400 + (seed * 10))
			treasury = decimal.NewFromInt(300 + seed)
		case "MATIC":
			price = decimal.NewFromInt(seed)
			treasury = decimal.NewFromInt(200000 + (seed * 1000))
		}

		minute := ((ts.Minute() + i) / 5) * 5
		hour := ts.Hour()
		if minute >= 60 {
			minute = minute % 5
			hour += 1
		}
		row := model.VaultPriceRow{
			Symbol:   symbol,
			Address:  tokens[idx].Address,
			Treasury: treasury,
			Price:    price,
			Weight:   weight,
			Time:     fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", ts.Year(), ts.Month(), ts.Day(), hour, minute),
		}
		repo.SetVaultPrice(ctx, row)
		//fmt.Println(row.Symbol, row.Price, row.Weight, row.Treasury)
	}
}

// volume
func TestRandomVaultVolume(t *testing.T) {
	ctx := context.Background()
	repo.InitDB(ctx, "arbitgo")
	defer repo.CloseDB()

	tokens, err := repo.MarketTokenInfo(CTX)
	if err != nil {
		panic(err)
	}

	row := model.VaultVolumeRow{
		TxHash:    "0x6f3ed208e7467fae4ef1b31b7fa5fabc2518654fac41551b499400ec81db28d3",
		TxIndex:   5,
		BlockNo:   35827978,
		Owner:     "0x2a27a94bf61e7bfc5cf5774019471ab2022a4c9a",
		BlockHash: "0x701849bdad60a9379f4c7820f80f64a4d393503eb6d23fd2fb0b3167125eced7",
	}

	r := rand.New(rand.NewSource(99))
	for i := range [100]int32{} {
		seed := int64(r.Intn(100))
		event := "DEPOSIT"
		if !(seed%2 == 0 || seed%3 == 0) {
			event = "WITHDRAW"
		}
		factor := int64(r.Intn(1000))
		idx := r.Intn(5)
		row.Symbol = tokens[idx].Symbol
		if event == "DEPOSIT" {
			switch row.Symbol {
			case "MATIC":
				row.Amount = decimal.NewFromInt(4000 + (seed * 30))
				row.MecaQuantity = decimal.NewFromInt(200000 + (seed * factor))
			case "USDC":
				row.Amount = decimal.NewFromInt(1200 + (seed * 30))
				row.MecaQuantity = decimal.NewFromInt(770000 + (seed * factor))
			case "USDT":
				row.Amount = decimal.NewFromInt(5000 + (seed * 30))
				row.MecaQuantity = decimal.NewFromInt(770000 + (seed * factor))
			case "DAI":
				row.Amount = decimal.NewFromInt(5500 + (seed * 30))
				row.MecaQuantity = decimal.NewFromInt(770000 + (seed * factor))
			case "ETH":
				row.Amount = decimal.NewFromInt(int64(r.Intn(10)))
				row.MecaQuantity = decimal.NewFromInt(300 + seed)
			}
		} else { // WITHDRAW
			switch row.Symbol {
			case "MATIC":
				row.Amount = decimal.NewFromInt(170 + seed)
				row.MecaQuantity = decimal.NewFromInt(200000 + (seed * factor))
			case "USDC":
				row.Amount = decimal.NewFromInt(2800 + (seed * 30))
				row.MecaQuantity = decimal.NewFromInt(770000 + (seed * factor))
			case "USDT":
				row.Amount = decimal.NewFromInt(2500 + (seed * 30))
				row.MecaQuantity = decimal.NewFromInt(770000 + (seed * factor))
			case "DAI":
				row.Amount = decimal.NewFromInt(14000 + (seed * 100))
				row.MecaQuantity = decimal.NewFromInt(770000 + (seed * factor))
			case "ETH":
				row.Amount = decimal.NewFromBigInt(big.NewInt(seed), -3)
				row.MecaQuantity = decimal.NewFromInt(300 + seed)
			}
		}

		row.EventType, row.TxIndex, row.BlockNo = event, uint32(seed), row.BlockNo+uint64(i)
		row.Token = tokens[idx].Address
		repo.SetVaultVolume(ctx, row)
		//fmt.Println(row.Symbol, row.EventType, row.Amount, row.MecaQuantity)
	}
}
