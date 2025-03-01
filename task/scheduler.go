package task

import (
	"context"
	"strings"
	"sync"
	"time"

	"github.com/coinmeca/go-common/logger"
	"github.com/coinmeca/go-common/model"
	rep "github.com/coinmeca/go-common/repository"
	"github.com/shopspring/decimal"
)

func AddDailyVolume() {
	t := time.Now()
	today := time.Now().UTC().Format("2006-01-02")
	wg := sync.WaitGroup{}

	// vault
	tokens, err := rep.VaultInfo(CTX)
	if err == nil {
		vr := model.VaultVolumeRow{
			Amount: decimal.Zero, MecaQuantity: decimal.Zero,
		}
		for _, event := range []string{"DEPOSIT", "WITHDRAW"} {
			for _, t := range tokens {
				wg.Add(1)
				vr.Symbol, vr.Token, vr.EventType = t.Symbol, t.Address, event
				go func(c context.Context, v model.VaultVolumeRow, t string) {
					rep.AddDailyVaultVolume(c, v, t)
					defer wg.Done()
				}(CTX, vr, today)
			}
		}
	}

	// market
	vm := model.MarketVolumeRow{
		Amount: decimal.Zero,
	}
	for _, event := range []string{"SELL", "BUY"} {
		for _, address := range CAxBOOKS {
			wg.Add(1)
			vm.EventType, vm.Address = event, strings.ToLower(address.Hex())
			go func(c context.Context, v model.MarketVolumeRow, t string) {
				rep.AddDailyMarketVolume(c, v, t)
				defer wg.Done()
			}(CTX, vm, today)
		}
	}

	// overview
	wg.Add(1)
	go func(c context.Context, t string) {
		rep.AddOverviewVolume(c, t)
		defer wg.Done()
	}(CTX, today)

	wg.Wait()
	logger.Info("AddDailyVolume", "time", time.Since(t).Milliseconds())
}

func SetMinutelyMarketData() {
	t := time.Now()
	wg := sync.WaitGroup{}

	for _, address := range CAxBOOKS {
		wg.Add(1)
		mv := model.MarketVolumeRow{Address: strings.ToLower(address.Hex()), Quantity: decimal.Zero}
		go func(c context.Context, v model.MarketVolumeRow) {
			rep.SetMinutelyMarketVolume(c, v)
			defer wg.Done()
		}(CTX, mv)

		wg.Add(1)
		mp := model.MarketPriceRow{Address: strings.ToLower(address.Hex())}
		go func(c context.Context, v model.MarketPriceRow) {
			rep.SetMinutelyMarketPrice(c, v)
			defer wg.Done()
		}(CTX, mp)
	}

	wg.Wait()
	logger.Info("SetMinutelyMarketData", "time", time.Since(t).Milliseconds())
}
