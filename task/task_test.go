package task

import (
	"context"
	//"dex-server/internal/model"
	"github.com/coinmeca/go-common/model"
	//repo "dex-server/repository"
	"fmt"
	"testing"
	"time"

	rep "github.com/coinmeca/go-common/repository"
	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

const (
	tokenSymbol  = "USDT"
	tokenAddress = "0x85e136cc3c5e3084c2d98d2857598c450e370843"
)

func TestVaultTask(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	rep.InitDB(ctx, "test")
	defer rep.CloseDB()

	ts := time.Now().UTC()
	t.Run("Dashboard", func(t *testing.T) {
		rep.Truncate(ctx, "vault_price_24h")
		rep.Truncate(ctx, "vault_price_daily")
		rep.Truncate(ctx, "vault_volume_24h")
		rep.Truncate(ctx, "vault_volume_daily")

		vpr := model.VaultPriceRow{
			Symbol:  tokenSymbol,
			Address: tokenAddress,
			Time:    fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", ts.Year(), ts.Month(), ts.Day(), ts.Hour(), 30),
		}

		vpr.Price = decimal.NewFromInt(2000)
		vpr.Treasury = decimal.NewFromInt(7500000)
		vpr.Weight = decimal.NewFromInt(400000)
		vpr.ValueLocked = vpr.Treasury.Mul(decimal.NewFromFloat32(0.9998))
		rep.SetVaultPrice(ctx, vpr)

		vpr.Time = fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", ts.Year(), ts.Month(), ts.Day(), ts.Hour(), 31)
		vpr.Price = decimal.NewFromInt(2100)
		vpr.Treasury = decimal.NewFromInt(7600000)
		vpr.Weight = decimal.NewFromInt(420000)
		vpr.ValueLocked = vpr.Treasury.Mul(decimal.NewFromFloat32(0.9997))
		rep.SetVaultPrice(ctx, vpr)

		vvr := model.VaultVolumeRow{
			Symbol:    tokenSymbol,
			Token:     tokenAddress,
			TxHash:    "0x6f3ed208e7467fae4ef1b31b7fa5fabc2518654fac41551b499400ec81db28d3",
			TxIndex:   5,
			BlockNo:   35827978,
			Owner:     "0xfe00fa244d69bfd8b1c108321c4713993f9ebb7c",
			BlockHash: "0x701849bdad60a9379f4c7820f80f64a4d393503eb6d23fd2fb0b3167125eced7",
		}
		vvr.EventType = "DEPOSIT"
		vvr.Amount = decimal.NewFromInt(5000)
		vvr.MecaQuantity = decimal.NewFromInt(770000)
		rep.SetVaultVolume(ctx, vvr)
		vvr.Amount = decimal.NewFromInt(4500)
		vvr.MecaQuantity = decimal.NewFromInt(690000)
		rep.SetVaultVolume(ctx, vvr)

		vvr.EventType = "WITHDRAW"
		vvr.Amount = decimal.NewFromInt(2000)
		vvr.MecaQuantity = decimal.NewFromInt(560000)
		rep.SetVaultVolume(ctx, vvr)
		vvr.Amount = decimal.NewFromInt(2500)
		vvr.MecaQuantity = decimal.NewFromInt(600000)
		rep.SetVaultVolume(ctx, vvr)

		d, err := rep.VaultDashboard(ctx, tokenAddress)
		assert.NoError(err)
		//fmt.Printf("%+v\n", d)
		assert.Equal(float64(2100), d.Exchange, "Exchange")
		assert.Equal(float64(100), d.ExchangeChange, "ExchangeChange")
		assert.Equal(float64(5), d.ExchangeRate, "ExchangeChangeRate")
		assert.Equal(float64(7600000), d.TotalLocked, "TotalLocked")
		assert.Equal(float64(100000), d.TotalChange, "TotalChange")
		assert.Equal(float64(420000), d.Weight, "Weight")
		assert.Equal(float64(20000), d.WeightChange, "WeightChange")
		assert.Equal(float64(9500), d.Deposit, "Deposit")
		assert.Equal(float64(4500), d.Withdraw, "Withdraw")
		assert.Equal(float64(1460000), d.Earn, "Earn")
		assert.Equal(float64(1160000), d.Burn, "Burn")
		assert.Equal(float64(2100), d.MecaPerToken, "MecaPerToken")
		assert.Equal(1809, int(d.TokenPerMeca*100), "TokenPerMeca")
	})
}

func TestMarketTask(t *testing.T) {
	assert := assert.New(t)

	ctx := context.Background()
	rep.InitDB(ctx, "test")
	defer rep.CloseDB()

	ts := time.Now().UTC()
	t.Run("Dashboard", func(t *testing.T) {
		rep.Truncate(ctx, "market_price_24h")
		rep.Truncate(ctx, "market_price_daily")
		rep.Truncate(ctx, "market_volume_24h")
		rep.Truncate(ctx, "market_volume_daily")

		mpr := model.MarketPriceRow{
			Address: tokenAddress,
			Time:    fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", ts.Year(), ts.Month(), ts.Day(), ts.Hour(), 40),
		}
		mpr.Price = decimal.NewFromInt(1800)
		rep.SetMarketPrice(ctx, mpr)

		mpr.Time = fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", ts.Year(), ts.Month(), ts.Day(), ts.Hour(), 42)
		mpr.Price = decimal.NewFromInt(2200)
		rep.SetMarketPrice(ctx, mpr)

		mvr := model.MarketVolumeRow{
			Address:   tokenAddress,
			TxHash:    "0x448ea94872267a10c944bbabf9fc505d2522ec1ad879aa4db08109fe7f7a7246",
			TxIndex:   5,
			BlockNo:   35827978,
			Owner:     "0xfe00fa244d69bfd8b1c108321c4713993f9ebb7c",
			BlockHash: "0xb1f96e9ca9043087ef3507ab247acb6f4b29da736bc268011f2f0d047751bc7f",
		}
		mvr.EventType = "BUY"
		mvr.Price = decimal.NewFromInt(1300)
		mvr.Quantity = decimal.NewFromInt(4)
		mvr.Amount = mvr.Price.Mul(mvr.Quantity)
		rep.SetMarketVolume(ctx, mvr)

		mvr.EventType = "BUY"
		mvr.Price = decimal.NewFromInt(1500)
		mvr.Quantity = decimal.NewFromInt(6)
		mvr.Amount = mvr.Price.Mul(mvr.Quantity)
		rep.SetMarketVolume(ctx, mvr)

		mvr.EventType = "SELL"
		mvr.Price = decimal.NewFromInt(1400)
		mvr.Quantity = decimal.NewFromInt(2)
		mvr.Amount = mvr.Price.Mul(mvr.Quantity)
		rep.SetMarketVolume(ctx, mvr)

		mvr.EventType = "SELL"
		mvr.Price = decimal.NewFromInt(1600)
		mvr.Quantity = decimal.NewFromInt(3)
		mvr.Amount = mvr.Price.Mul(mvr.Quantity)
		rep.SetMarketVolume(ctx, mvr)

		d, err := rep.MarketDashboard(ctx, tokenAddress)
		assert.NoError(err)
		//fmt.Printf("%+v\n", d)
		assert.Equal(float64(1600), d.Price, "Price")
		assert.Equal(float64(10), d.VolumeQuote, "VolumeQuote")
		assert.Equal(float64(5), d.VolumeBase, "VolumeBase")
		assert.Equal(float64(1600), d.High, "High")
		assert.Equal(float64(1300), d.Low, "Low")
		assert.Equal(2307, int(d.ChangeRate*100), "ChangeRate")
		assert.Equal(float64(300), d.Change, "Change")
	})

	t.Run("Chart", func(t *testing.T) {
		d, err := rep.MarketChartData(ctx, tokenAddress, "")
		assert.NoError(err)
		//fmt.Printf("%+v\n", d)
		assert.Equal(float64(1800), d.Price[0].Open, "Price.Open")
		assert.Equal(float64(2200), d.Price[0].Close, "Price.Close")
		assert.Equal(float64(15), d.Volume[0].Value, "Volume.Value")
		assert.Equal("BUY", d.Volume[0].OrderType, "Volume.Type")
	})
}
