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

func SetFarmData(ctx context.Context, f model.FarmRow) {
	deleteQuery := `delete from %s.farm_24h where updated_at < now()::timestamptz at time zone 'utc' - '1day'::interval`
	_, err := POOL.Exec(ctx, fmt.Sprintf(deleteQuery, SCHEMA))
	if err != nil {
		logger.Error("SetFarmData", "err", err)
	}

	rate := f.Rewards.Div(decimal.NewFromInt(1))
	if !f.Locked.IsZero() {
		rate = f.Rewards.Div(f.Locked)
	}

	apy, interval := rate, f.Goal-f.Start
	if interval > 0 {
		days := decimal.NewFromInt(int64(interval)).Div(decimal.NewFromInt(86400))
		apy = rate.Div(days)
	}

	nowUTC := time.Now().UTC()
	query := `insert into %s.farm_24h(datetime,address,name,stake,earn,rewards,locked,start,period,goal,apy,updated_at) 
values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12) ON CONFLICT ON CONSTRAINT pk_farm_24h DO UPDATE SET rewards=$6,locked=$7,apy=$11,updated_at=$12`
	query = fmt.Sprintf(query, SCHEMA)
	_, err = POOL.Exec(ctx, query, f.Time, f.Address, f.Name, f.Stake, f.Earn, f.Rewards, f.Locked, f.Start, f.Period, f.Goal, apy, nowUTC)
	if err != nil {
		logger.Error("SetFarmData", "err", err, "query", query)
	} else {
		logger.Debug("SetFarmData", "time", f.Time)
	}

	today := nowUTC.Format("2006-01-02")
	query = `select rewards_open,locked_open from %s.farm_daily where address='%s' and day='%s' limit 1`
	query = fmt.Sprintf(query, SCHEMA, f.Address, today)
	row, _ := Row(ctx, query)

	if row == nil {
		AddDailyFarm(ctx, f, today, apy)
	} else {
		or, ol := row[0].(pgtype.Numeric), row[1].(pgtype.Numeric)
		ord, old := decimal.NewFromBigInt(or.Int, or.Exp), decimal.NewFromBigInt(ol.Int, ol.Exp)

		// rewards rate
		orcr := decimal.Zero
		if !ord.IsZero() {
			orcr = f.Rewards.Div(ord).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100))
		}
		// locked rate
		olcr := decimal.Zero
		if !old.IsZero() {
			olcr = f.Locked.Div(old).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100))
		}

		updateQuery := `update %s.farm_daily set rewards=$1,rewards_rate=$2,locked=$3,locked_rate=$4,apy=$5,start=$6,goal=$7,updated_at=$8 where day='%s' and address='%s'`
		updateQuery = fmt.Sprintf(updateQuery, SCHEMA, today, f.Address)
		_, err = POOL.Exec(ctx, updateQuery, f.Rewards, orcr, f.Locked, olcr, apy, f.Start, f.Goal, nowUTC)
		if err != nil {
			logger.Error("SetFarmData", "err", err, "query", updateQuery)
		} else {
			//logger.Debug.Printf("update daily farm...%v, %v, %v\n", today, f.Rewards, f.Locked)
			logger.Debug("SetFarmData", "today", today, "reward", f.Rewards, "locked", f.Locked)
		}
	}
}

func AddDailyFarm(ctx context.Context, f model.FarmRow, today string, apy decimal.Decimal) {
	insertQuery := `insert into %s.farm_daily(day,rewards,rewards_open,locked,locked_open,address,name,stake,stake_symbol,stake_name,earn,earn_symbol,earn_name,start,period,goal,apy) 
values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13,$14,$15,$16,$17)`
	insertQuery = fmt.Sprintf(insertQuery, SCHEMA)

	_, err := POOL.Exec(ctx, insertQuery, today, f.Rewards, f.Rewards, f.Locked, f.Locked, f.Address, f.Name, f.Stake, f.StakeSymbol, f.StakeName, f.Earn, f.EarnSymbol, f.EarnName, f.Start, f.Period, f.Goal, apy)
	if err != nil {
		logger.Error("AddDailyFarm", "err", err, "query", insertQuery)
	} else {
		logger.Debug("AddDailyFarm", "address", f.Address, "reward", f.Rewards, "locked", f.Locked)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// for API

func FarmList(ctx context.Context) ([]model.FarmListApi, error) {
	query := `select stake,earn,name,period,rewards,locked,address,stake_symbol,stake_name,earn_symbol,earn_name,start,goal 
from %s.farm_daily where day >= '%s'`
	query = fmt.Sprintf(query, SCHEMA, time.Now().UTC().Format("2006-01-02"))
	rows, err := Rows(ctx, query)
	if err != nil {
		logger.Error("FarmList", "err", err, "query", query)
		return nil, err
	}

	var farmList []model.FarmListApi
	for _, row := range rows {
		r, l := row[4].(pgtype.Numeric), row[5].(pgtype.Numeric)
		rewards, _ := decimal.NewFromBigInt(r.Int, r.Exp).Float64()
		locked, _ := decimal.NewFromBigInt(l.Int, l.Exp).Float64()

		farm := model.FarmListApi{
			Rewards: rewards, Locked: locked,
			//RewardsRate: rewardsRate, LockedRate: lockedRate,
			Period: uint64(row[3].(int64)), Start: uint64(row[11].(int64)), Goal: uint64(row[12].(int64)),
			Stake: row[0].(string), StakeSymbol: row[7].(string), StakeName: row[8].(string),
			Earn: row[1].(string), EarnSymbol: row[9].(string), EarnName: row[10].(string),
			Name:    row[2].(string),
			Address: row[6].(string),
		}
		farmList = append(farmList, farm)
	}

	return farmList, nil
}

func FarmChartForLocked(ctx context.Context, address string) ([]model.FarmLockedChartApi, error) {
	query := `select day,rewards,locked,rewards_rate,locked_rate from %s.farm_daily where address='%s' order by day limit 30`
	query = fmt.Sprintf(query, SCHEMA, address)
	rows, err := Rows(ctx, query)
	if err != nil {
		logger.Error("FarmChartForLocked", "err", err, "query", query)
		return nil, err
	}

	var items []model.FarmLockedChartApi
	for _, row := range rows {
		r, l := row[1].(pgtype.Numeric), row[2].(pgtype.Numeric)
		rewards, _ := decimal.NewFromBigInt(r.Int, r.Exp).Float64()
		locked, _ := decimal.NewFromBigInt(l.Int, l.Exp).Float64()
		c := model.FarmLockedChartApi{
			Datetime: row[0].(time.Time).Format("2006-01-02"),
			Rewards:  rewards,
			Locked:   locked,
		}
		items = append(items, c)
	}
	return items, nil
}

func FarmChartForApy(ctx context.Context, address string) ([]model.BarChartApi, error) {
	query := `select day,apy from %s.farm_daily where address='%s' order by day limit 30`
	query = fmt.Sprintf(query, SCHEMA, address)
	rows, err := Rows(ctx, query)
	if err != nil {
		logger.Error("FarmChartForApy", "err", err, "query", query)
		return nil, err
	}

	var items []model.BarChartApi
	for _, row := range rows {
		r := row[1].(pgtype.Numeric)
		apy, _ := decimal.NewFromBigInt(r.Int, r.Exp).Float64()
		c := model.BarChartApi{
			Datetime: row[0].(time.Time).Format("2006-01-02"),
			Value:    apy,
		}
		items = append(items, c)
	}
	return items, nil
}
