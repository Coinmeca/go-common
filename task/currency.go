package task

import (
	"github.com/coinmeca/go-common/chain"
	"github.com/coinmeca/go-common/logger"
	"github.com/coinmeca/go-common/model"
	repo "github.com/coinmeca/go-common/repository"
	"github.com/coinmeca/go-common/utils"
	"encoding/json"
	"fmt"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

func FetchCurrencyRate() {
	t := time.Now()

	rate, err := USDCurrencyRate()
	if err != nil {
		rate, _ = repo.LastCurrencyRate(CTX)
	}

	repo.Add24hCurrencyRate(CTX, rate)
	repo.SetDailyCurrencyRate(CTX, rate)

	logger.Debug("FetchCurrencyRate", "time", time.Since(t).Milliseconds())
}

func FetchCMCQuote() {
	t := time.Now()

	quotes, err := CMCLatestQuote()
	if err != nil {
		logger.Error("FetchCMCQuote", "err", err)
		return
	}

	fmt.Printf("CMC quotes %v\n", len(quotes))
	repo.SetCMCQuote(CTX, quotes)

	logger.Debug("FetchCMCQuote", "time", time.Since(t).Milliseconds())
}

func USDCurrencyRate() (decimal.Decimal, error) {
	data, err := utils.ExternalData(chain.APICurrencyRate, "")
	if err != nil {
		logger.Error("USDCurrencyRate", "err", err)
		return decimal.Zero, err
	}

	var results []model.CurrencyRate
	if err = json.Unmarshal(data, &results); err != nil {
		logger.Error("USDCurrencyRate", "err", err)
		return decimal.Zero, err
	}
	if len(results) > 0 {
		res := results[0]
		if res.Unit == "USD" {
			logger.Debug("USDCurrencyRate", "res", res)
			return res.Rate, nil
		}
	}
	return decimal.Zero, fmt.Errorf("mismatch currency rate")
}

func CMCLatestQuote() ([]model.StableCoinQuote, error) {
	cmcIds := []string{"3408", "4943", "825"} // USDC, DAI, USDT
	uri := fmt.Sprintf("%s?id=%s", chain.CMCLatestQuote, strings.Join(cmcIds, ","))
	data, err := utils.ExternalData(uri, utils.GetConfig().CMCApiKey)

	if err != nil {
		return nil, err
	}

	var cmc model.CMCLatestQuote
	if err = json.Unmarshal(data, &cmc); err != nil {
		return nil, err
	}
	//fmt.Printf("%+v\n", string(data))
	var quotes [3]model.StableCoinQuote
	for i, tokenId := range cmcIds {
		coin := cmc.Data[tokenId]
		q, s := coin.Quote, coin.Symbol
		quotes[i] = q["USD"]
		quotes[i].Symbol = s
	}
	return quotes[:], nil
}
