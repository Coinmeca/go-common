package repository

import (
	"coinmeca-go_common/logger"
	"coinmeca-go_common/model"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
)

func SetCMCQuote(ctx context.Context, quotes []model.StableCoinQuote) {
	t := time.Now().UTC()
	day := t.Format("2006-01-02")

	for _, q := range quotes {
		query := `insert into public.stablecoin_quote_daily(day,price,slug,updated_at) values($1,$2,$3,$4)
ON CONFLICT ON CONSTRAINT pk_stablecoin_quote_daily DO UPDATE SET price=$2,updated_at=$4`
		_, err := POOL.Exec(ctx, query, day, q.Price, q.Symbol, t)
		if err != nil {
			logger.Error("SetCMCQuote", "err", err, "query", query)
		} else {
			logger.Debug("SetCMCQuote", "symbol", q.Symbol, "price", q.Price)
		}
		fmt.Printf("%+v\n", q)
	}
}

func latestCMCQuote(ctx context.Context, symbol string) (decimal.Decimal, error) {
	query := "select price from public.stablecoin_quote_daily where slug='%s' order by day desc limit 1"
	row, err := Row(ctx, fmt.Sprintf(query, symbol))
	if err != nil || row == nil {
		return decimal.NewFromInt32(1), fmt.Errorf("no quote %v", err)
	}
	p := row[0].(pgtype.Numeric)
	return decimal.NewFromBigInt(p.Int, p.Exp), nil
}

func SetDailyCurrencyRate(ctx context.Context, rate decimal.Decimal) {
	t := time.Now().UTC()
	day := t.Format("2006-01-02")
	query := fmt.Sprintf(`select high,low from sample.currency_rate_daily where day = '%s' limit 1`, day)
	row, _ := Row(ctx, query)

	if row == nil {
		insertQuery := `insert into sample.currency_rate_daily(day,rate,open,high,low) values($1,$2,$3,$4,$5)`
		_, err := POOL.Exec(ctx, insertQuery, day, rate, rate, rate, rate)
		if err != nil {
			logger.Error("SetDailyCurrencyRate", "err", err, "query", insertQuery)
		} else {
			logger.Debug("SetDailyCurrencyRate", "day", day, "rate", rate)
		}
	} else {
		h, l := row[0].(pgtype.Numeric), row[1].(pgtype.Numeric)
		highRate, lowRate, closeRate := decimal.NewFromBigInt(h.Int, h.Exp), decimal.NewFromBigInt(l.Int, l.Exp), rate
		if rate.GreaterThan(highRate) {
			highRate = rate
		} else if rate.LessThan(lowRate) {
			lowRate = rate
		}

		updateQuery := `update sample.currency_rate_daily set rate=$1,high=$2,low=$3,close=$4,updated_at=$5 where day = '%s'`
		updateQuery = fmt.Sprintf(updateQuery, day)
		_, err := POOL.Exec(ctx, updateQuery, rate, highRate, lowRate, closeRate, t.Format("2006-01-02 15:04:05"))
		if err != nil {
			logger.Error("SetDailyCurrencyRate", "err", err, "query", updateQuery)
		} else {
			logger.Debug("SetDailyCurrencyRate", "day", day, "rate", rate, "highRate", highRate, "lowRate", lowRate)
		}
	}
}

func Add24hCurrencyRate(ctx context.Context, rate decimal.Decimal) {
	deleteQuery := `delete from sample.currency_rate_24h where last_updated < now()::timestamptz at time zone 'utc' - '1day'::interval`
	_, err := POOL.Exec(ctx, deleteQuery)
	if err != nil {
		logger.Error("Add24hCurrencyRate", "err", err)
	}

	query := `insert into sample.currency_rate_24h(last_updated,rate) values($1,$2)`
	_, err = POOL.Exec(ctx, query, time.Now().UTC(), rate)
	if err != nil {
		logger.Error("Add24hCurrencyRate", "err", err, "query", query)
	} else {
		logger.Debug("Add24hCurrencyRate", "rate", rate)
	}
}

func LastCurrencyRate(ctx context.Context) (decimal.Decimal, error) {
	query := "select rate from sample.currency_rate_24h order by last_updated desc limit 1"
	rate, err := Scan(ctx, query)
	if err != nil {
		return decimal.Zero, fmt.Errorf("no rate %v", err)
	}
	return rate.(decimal.Decimal), nil
}
