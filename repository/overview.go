package repository

import (
	"context"
	cv 	"github.com/coinmeca/go-common/chain"
	"github.com/coinmeca/go-common/logger"
	"github.com/coinmeca/go-common/model"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
	"time"
)

func latestOverviewVolume(ctx context.Context, volumeType string) (decimal.Decimal, decimal.Decimal, error) {
	query := `select amount,%s from %s.overview_volume order by day desc limit 1`
	query = fmt.Sprintf(query, volumeType, SCHEMA)
	row, err := Row(ctx, query)
	if err != nil {
		return decimal.Zero, decimal.Zero, err
	}
	a, v := row[0].(pgtype.Numeric), row[1].(pgtype.Numeric)
	return decimal.NewFromBigInt(a.Int, a.Exp), decimal.NewFromBigInt(v.Int, v.Exp), nil
}

func setOverviewVolume(ctx context.Context, volume decimal.Decimal, volumeType string) {
	nowUTC := time.Now().UTC()
	volumeAmount, volumeEach, err := latestOverviewVolume(ctx, volumeType)
	if err != nil {
		logger.Error("setOverviewVolume", "err", err)
	}
	volumeAmount = volumeAmount.Add(volume)
	volumeEach = volumeEach.Add(volume)

	query := `insert into %s.overview_volume(day,amount,%s,updated_at) values($1,$2,$3,$4)
ON CONFLICT ON CONSTRAINT pk_overview_volume DO UPDATE SET amount=$2,%s=$3,updated_at=$4`
	query = fmt.Sprintf(query, SCHEMA, volumeType, volumeType)
	_, err = POOL.Exec(ctx, query, nowUTC.Format("2006-01-02"), volumeAmount, volumeEach, nowUTC)

	if err != nil {
		logger.Error("setOverviewVolume", "err", err, "query", query)
	} else {
		logger.Debug("setOverviewVolume", "volumeAmount", volumeAmount, "volumeType", volumeType, "volumeEach", volumeEach)
	}
}

func AddOverviewVolume(ctx context.Context, today string) {
	query := fmt.Sprintf(`insert into %s.overview_volume(day) values($1)`, SCHEMA)
	_, err := POOL.Exec(ctx, query, today)

	if err != nil {
		logger.Error("AddOverviewVolume", "err", err)
	} else {
		logger.Debug("AddOverviewVolume", "today", today)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// for API

func ContractInfo(ctx context.Context, id int) (*model.ContractApi, error) {
	query := `select app,market,vault,farm from public.contracts where chain_id=%d limit 1`
	row, err := Row(ctx, fmt.Sprintf(query, id))
	if err != nil {
		logger.Error("ContractInfo", "err", err)
		return nil, err
	}

	return &model.ContractApi{
		App:    row[0].(string),
		Market: row[1].(string),
		Vault:  row[2].(string),
		Farm:   row[3].(string),
	}, nil
}

func OverviewInfo(ctx context.Context) (*model.OverviewApi, error) {
	var info model.OverviewApi

	// volume
	query := `select sum(amount) from %s.overview_volume`
	query = fmt.Sprintf(query, SCHEMA)
	row, err := Row(ctx, query)
	if err != nil {
		logger.Error("OverviewInfo", "err",err)
		return nil, err
	} else if row != nil {
		v := row[0].(pgtype.Numeric)
		info.TotalVolume, _ = decimal.NewFromBigInt(v.Int, v.Exp).Float64()
	}

	// value (valueLocked)
	query = `select sum(valueLocked) from %s.vault_price_daily group by day order by day desc limit 1`
	query = fmt.Sprintf(query, SCHEMA)
	row, err = Row(ctx, query)
	if err != nil {
		logger.Error("OverviewInfo", "err", err)
		return nil, err
	} else if row != nil {
		v := row[0].(pgtype.Numeric)
		info.TotalValue, _ = decimal.NewFromBigInt(v.Int, v.Exp).Float64()
	}

	return &info, nil
}

func OverviewChartForValue(ctx context.Context) ([]model.BarChartApi, error) {
	query := `select day, sum(valueLocked) from %s.vault_price_daily group by day order by day limit 30`
	query = fmt.Sprintf(query, SCHEMA)
	rows, err := Rows(ctx, query)
	if err != nil {
		logger.Error("OverviewChartForValue", "err", err, "query", query)
		return nil, err
	}

	var chartData []model.BarChartApi
	for _, row := range rows {
		var d model.BarChartApi
		d.Datetime = row[0].(time.Time).Format(cv.DateFormat)
		a := row[1].(pgtype.Numeric)
		d.Value, _ = decimal.NewFromBigInt(a.Int, a.Exp).Float64()
		chartData = append(chartData, d)
	}
	return chartData, nil
}

func OverviewChartForVolume(ctx context.Context) ([]model.BarChartApi, error) {
	query := `select day, amount from %s.overview_volume order by day limit 30`
	query = fmt.Sprintf(query, SCHEMA)
	rows, err := Rows(ctx, query)
	if err != nil {
		logger.Error("OverviewChartForVolume", "err", err, "query", query)
		return nil, err
	}

	var chartData []model.BarChartApi
	for _, row := range rows {
		var d model.BarChartApi
		d.Datetime = row[0].(time.Time).Format(cv.DateFormat)
		a := row[1].(pgtype.Numeric)
		d.Value, _ = decimal.NewFromBigInt(a.Int, a.Exp).Float64()
		chartData = append(chartData, d)
	}
	return chartData, nil
}
