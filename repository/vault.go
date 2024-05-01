package repository

import (
	"github.com/coinmeca/go-common/logger"
	"github.com/coinmeca/go-common/model"
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

func SetVaultPrice(ctx context.Context, v model.VaultPriceRow) {
	// add 24h
	deleteQuery := `delete from %s.vault_price_24h where symbol='%s' and datetime < now()::timestamptz at time zone 'utc' - '1day'::interval`
	_, err := POOL.Exec(ctx, fmt.Sprintf(deleteQuery, SCHEMA, v.Symbol))
	if err != nil {
		logger.Error("SetVaultPrice", "err", err)
	}

	nowUTC := time.Now().UTC()
	query := `insert into %s.vault_price_24h(datetime,symbol,treasury,price,weight,need,address,updated_at) values($1,$2,$3,$4,$5,$6,$7,$8)
ON CONFLICT ON CONSTRAINT pk_vault_price_24h DO UPDATE SET treasury=$3,price=$4,weight=$5,need=$6,updated_at=$8`
	query = fmt.Sprintf(query, SCHEMA)
	_, err = POOL.Exec(ctx, query, v.Time, v.Symbol, v.Treasury, v.Price, v.Weight, v.Need, v.Address, nowUTC)
	if err != nil {
		logger.Error("SetVaultPrice", "err", err, "query", query)
	} else {
		logger.Debug("SetVaultPrice", "symbol", v.Symbol, "treasury", v.Treasury)
	}

	// update daily
	today := nowUTC.Format("2006-01-02")
	query = `select treasury_open,weight_open,price_open,price_high,price_low,tvl_open from %s.vault_price_daily where symbol='%s' and day = '%s' limit 1`
	query = fmt.Sprintf(query, SCHEMA, v.Symbol, today)
	row, _ := Row(ctx, query)

	if row == nil {
		AddDailyVaultPrice(ctx, v, today)
	} else {
		UpdateDailyVaultPrice(ctx, v, row, today)
	}
}

func SetVaultTokenInfo(ctx context.Context, t model.VaultTokenRow) {
	query := `insert into %s.vault_tokens(is_key,address,symbol,name,decimals,updated_at)
values($1,$2,$3,$4,$5,$6) ON CONFLICT ON CONSTRAINT pk_vault_tokens
DO UPDATE SET is_key=$1,name=$4,decimals=$5,updated_at=$6`
	query = fmt.Sprintf(query, SCHEMA)

	now := time.Now().UTC()
	_, err := POOL.Exec(ctx, query, t.IsKey, t.Address, t.Symbol, t.Name, t.Decimals, now)
	if err != nil {
		logger.Error("SetVaultTokenInfo", "err", err)
	} else {
		logger.Debug("SetVaultTokenInfo", "vaultTokenRow", t)
	}
}

func VaultTokenInfo(ctx context.Context) ([]*model.TokenInfo, error) {
	query := `select address, symbol, decimals from %s.vault_tokens`
	query = fmt.Sprintf(query, SCHEMA)
	rows, err := Rows(ctx, query)
	if err != nil {
		logger.Error("VaultTokenInfo", "err", err)
		return nil, err
	}

	var tokens []*model.TokenInfo
	for _, r := range rows {
		t := model.TokenInfo{Address: r[0].(string), Symbol: r[1].(string), Decimals: r[2].(int32)}
		tokens = append(tokens, &t)
	}
	return tokens, nil
}

func SetVaultVolume(ctx context.Context, v model.VaultVolumeRow) {
	deleteQuery := `delete from %s.vault_volume_24h where updated_at < now()::timestamptz at time zone 'utc' - '1day'::interval`
	_, err := POOL.Exec(ctx, fmt.Sprintf(deleteQuery, SCHEMA))
	if err != nil {
		logger.Error("SetVaultVolume", "err", err)
	}

	nowUTC := time.Now().UTC()
	query := `insert into %s.vault_volume_24h(datetime,symbol,owner,token,event_type,amount,meca_quantity,block_no,block_hash,tx_hash,tx_index) 
values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	query = fmt.Sprintf(query, SCHEMA)
	_, err = POOL.Exec(ctx, query, nowUTC, v.Symbol, v.Owner, v.Token, v.EventType, v.Amount, v.MecaQuantity, v.BlockNo, v.BlockHash, v.TxHash, v.TxIndex)
	if err != nil {
		logger.Error("SetVaultVolume", "err", err)
	} else {
		logger.Debug("SetVaultVolume", "vaultVolumeRow", v)
	}
	setOverviewVolume(ctx, v.Amount, strings.ToLower(v.EventType))

	today := nowUTC.Format("2006-01-02")
	query = `select amount, meca_quantity from %s.vault_volume_daily where symbol='%s' and day='%s' and event_type='%s' limit 1`
	query = fmt.Sprintf(query, SCHEMA, v.Symbol, today, v.EventType)
	row, _ := Row(ctx, query)

	if row == nil {
		AddDailyVaultVolume(ctx, v, today)
	} else {
		UpdateDailyVaultVolume(ctx, v, row, today)
	}
}

func AddDailyVaultPrice(ctx context.Context, v model.VaultPriceRow, date string) {
	insertQuery := `insert into %s.vault_price_daily(symbol,day,treasury,treasury_open,price,price_open,price_high,price_low,weight,weight_open,tvl,tvl_open,address) 
values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11,$12,$13)`
	insertQuery = fmt.Sprintf(insertQuery, SCHEMA)

	_, err := POOL.Exec(ctx, insertQuery, v.Symbol, date, v.Treasury, v.Treasury, v.Price, v.Price, v.Price, v.Price, v.Weight, v.Weight, v.Tvl, v.Tvl, v.Address)
	if err != nil {
		logger.Error("AddDailyVaultPrice", "err", err, "query", insertQuery)
	} else {
		logger.Debug("AddDailyVaultPrice", "date", date, "symbol", v.Symbol)
	}
}

func UpdateDailyVaultPrice(ctx context.Context, v model.VaultPriceRow, row []interface{}, date string) {
	t, w, p := row[0].(pgtype.Numeric), row[1].(pgtype.Numeric), row[5].(pgtype.Numeric)
	o, h, l := row[2].(pgtype.Numeric), row[3].(pgtype.Numeric), row[4].(pgtype.Numeric)
	tr, wt, tv := decimal.NewFromBigInt(t.Int, t.Exp), decimal.NewFromBigInt(w.Int, w.Exp), decimal.NewFromBigInt(p.Int, p.Exp)
	op, hp, lp := decimal.NewFromBigInt(o.Int, o.Exp), decimal.NewFromBigInt(h.Int, h.Exp), decimal.NewFromBigInt(l.Int, l.Exp)

	if v.Price.GreaterThan(hp) {
		hp = v.Price
	} else if v.Price.LessThan(lp) {
		lp = v.Price
	}
	// treasury change
	trc, trcr := v.Treasury.Sub(tr), decimal.Zero
	if !tr.IsZero() {
		trcr = v.Treasury.Div(tr).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100))
	}
	// rate -> price change
	rpc, rpcr := v.Price.Sub(op), decimal.Zero
	if !op.IsZero() {
		rpcr = v.Price.Div(op).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100))
	}
	// weight change
	wtc := v.Weight.Sub(wt)
	// tvl change
	tvc, tvcr := v.Tvl.Sub(tv), decimal.Zero
	if !tv.IsZero() {
		tvcr = v.Tvl.Div(tv).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100))
	}

	updateQuery := `update %s.vault_price_daily set treasury=$1,treasury_change=$2,price=$3,price_change=$4,price_rate=$5,price_high=$6,price_low=$7,
weight=$8,weight_change=$9,tvl=$10,tvl_change=$11,treasury_rate=$12,tvl_rate=$13,updated_at=$14 where symbol='%s' and day = '%s'`
	updateQuery = fmt.Sprintf(updateQuery, SCHEMA, v.Symbol, date)
	_, err := POOL.Exec(ctx, updateQuery, v.Treasury, trc, v.Price, rpc, rpcr, hp, lp, v.Weight, wtc, v.Tvl, tvc, trcr, tvcr, time.Now().UTC())
	if err != nil {
		logger.Error("UpdateDailyVaultPrice", "err", err)
	} else {
		logger.Debug("UpdateDailyVaultPrice", "symbol", v.Symbol, "date", date, "trc", trc, "rpc", rpc, "wtc", wtc, "tvc", tvc)
	}
}

func AddDailyVaultVolume(ctx context.Context, v model.VaultVolumeRow, date string) {
	insertQuery := `insert into %s.vault_volume_daily(symbol,day,amount,meca_quantity,event_type,address) values($1,$2,$3,$4,$5,$6)`
	insertQuery = fmt.Sprintf(insertQuery, SCHEMA)

	_, err := POOL.Exec(ctx, insertQuery, v.Symbol, date, v.Amount, v.MecaQuantity, v.EventType, v.Token)
	if err != nil {
		logger.Error("AddDailyVaultVolume", "err", err)
	} else {
		logger.Debug("AddDailyVaultVolume", "symbol", v.Symbol, "eventType", v.EventType, "amount", v.Amount)
	}
}

func UpdateDailyVaultVolume(ctx context.Context, v model.VaultVolumeRow, row []interface{}, date string) {
	a, q := row[0].(pgtype.Numeric), row[1].(pgtype.Numeric)
	ad, qd := decimal.NewFromBigInt(a.Int, a.Exp), decimal.NewFromBigInt(q.Int, q.Exp)
	amount, quantity := ad.Add(v.Amount), qd.Add(v.MecaQuantity)

	updateQuery := `update %s.vault_volume_daily set amount=$1,meca_quantity=$2,updated_at=$3 where symbol='%s' and day='%s' and event_type='%s'`
	updateQuery = fmt.Sprintf(updateQuery, SCHEMA, v.Symbol, date, v.EventType)
	_, err := POOL.Exec(ctx, updateQuery, amount, quantity, time.Now().UTC())
	if err != nil {
		logger.Error("UpdateDailyVaultVolume", "err", err)
	} else {
		logger.Debug("UpdateDailyVaultVolume", "date", date, "amount", amount)
	}
}

func LastBlockNoFromVaultVolume(ctx context.Context) (int64, error) {
	query := `select block_no from %s.vault_volume_24h order by block_no desc limit 1`
	query = fmt.Sprintf(query, SCHEMA)
	row, err := Row(ctx, query)
	if err != nil {
		return 0, err
	}
	return row[0].(int64), nil
}

func latestWeight(ctx context.Context, address string) (decimal.Decimal, error) {
	query := `select weight from %s.vault_price_daily where address='%s' order by day desc limit 1`
	query = fmt.Sprintf(query, SCHEMA, address)
	row, err := Row(ctx, query)
	if err != nil {
		return decimal.Decimal{}, err
	}
	w := row[0].(pgtype.Numeric)
	return decimal.NewFromBigInt(w.Int, w.Exp), nil
}

func getUsdtAddress(ctx context.Context) (string, error) {
	query := `select address from %s.vault_tokens where symbol='USDT' limit 1`
	query = fmt.Sprintf(query, SCHEMA)
	row, err := Row(ctx, query)
	if err != nil {
		logger.Error("getUsdtAddress", "err", err, "query", query)
		return "", err
	}

	return row[0].(string), nil
}

func LatestUsdPrice(ctx context.Context, symbol string) (decimal.Decimal, error) {
	usdPrice := decimal.NewFromInt32(1)

	if symbol == "ETH" || symbol == "MATIC" {
		address, err := getUsdtAddress(ctx)
		if err != nil {
			return usdPrice, err
		}
		query := `select price from %s.market_price_daily where address='%s' order by day desc limit 1`
		query = fmt.Sprintf(query, SCHEMA, address)
		row, err := Row(ctx, query)
		if err != nil {
			return usdPrice, err
		} else if row != nil {
			p := row[0].(pgtype.Numeric)
			basePrice := decimal.NewFromBigInt(p.Int, p.Exp)
			cmcPrice, _ := latestCMCQuote(ctx, "USDT")
			usdPrice = basePrice.Mul(cmcPrice)
		}
	} else {
		usdPrice, _ = latestCMCQuote(ctx, symbol)
	}

	return usdPrice, nil
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// for API

func VaultTokenList(ctx context.Context) ([]model.VaultOverviewApi, error) {
	query := `select is_key,name,vt.symbol,decimals,vt.address,price,price_rate,treasury,treasury_rate,tvl,tvl_rate
from %s.vault_tokens vt left join %s.vault_price_daily vp on vt.address = vp.address where vp.day='%s'`
	query = fmt.Sprintf(query, SCHEMA, SCHEMA, time.Now().UTC().Format("2006-01-02"))
	rows, err := Rows(ctx, query)
	if err != nil {
		//logger.Error.Printf("vault token error: %v (%v)\n", err, query)
		return nil, err
	}

	var overviewList []model.VaultOverviewApi
	for _, r := range rows {
		item := model.VaultOverviewApi{
			IsKey:    r[0].(bool),
			Symbol:   r[2].(string),
			Name:     r[1].(string),
			Decimals: r[3].(int32),
			Address:  r[4].(string),
		}

		p, pc := r[5].(pgtype.Numeric), r[6].(pgtype.Numeric)
		item.Exchange, _ = decimal.NewFromBigInt(p.Int, p.Exp).Float64()
		item.ExchangeChange, _ = decimal.NewFromBigInt(pc.Int, pc.Exp).Float64()

		t, to := r[7].(pgtype.Numeric), r[8].(pgtype.Numeric)
		item.Volume, _ = decimal.NewFromBigInt(t.Int, t.Exp).Float64()
		item.VolumeChange, _ = decimal.NewFromBigInt(to.Int, to.Exp).Float64()

		v, vc := r[9].(pgtype.Numeric), r[10].(pgtype.Numeric)
		item.TVL, _ = decimal.NewFromBigInt(v.Int, v.Exp).Float64()
		item.TVLChange, _ = decimal.NewFromBigInt(vc.Int, vc.Exp).Float64()

		overviewList = append(overviewList, item)
	}
	return overviewList, nil
}

func VaultHistory(ctx context.Context, address string) ([]model.VaultHistoryApi, error) {
	weight, e := latestWeight(ctx, address)
	if e != nil {
		weight = decimal.NewFromInt32(1)
	}

	query := `select event_type,amount,meca_quantity from %s.vault_volume_24h where token='%s' order by datetime`
	query = fmt.Sprintf(query, SCHEMA, address)
	rows, err := Rows(ctx, query)
	if err != nil {
		return nil, err
	}

	var historyList []model.VaultHistoryApi
	for _, p := range rows {
		event, a, q := p[0].(string), p[1].(pgtype.Numeric), p[2].(pgtype.Numeric)
		af, _ := decimal.NewFromBigInt(a.Int, a.Exp).Float64()
		qd := decimal.NewFromBigInt(q.Int, q.Exp)
		share, _ := qd.Div(weight).Float64()
		qf, _ := qd.Float64()

		item := model.VaultHistoryApi{EventType: event, Volume: af, Meca: qf, Share: share}
		historyList = append(historyList, item)
	}
	return historyList, err
}

func VaultDashboard(ctx context.Context, address string) (*model.VaultDashboardApi, error) {
	today := time.Now().UTC().Format("2006-01-02")
	var deposit, earn, withdraw, burn float64

	// deposit
	query := `select sum(amount),sum(meca_quantity) from %s.vault_volume_24h 
where token='%s' and date_trunc('day',datetime)='%s' and event_type='DEPOSIT' group by event_type`
	query = fmt.Sprintf(query, SCHEMA, address, today)
	row, err := Row(ctx, query)
	if err != nil {
		logger.Error("VaultDashboard", "err", err, "query", query)
	} else if row != nil {
		if row[0] != nil && row[1] != nil {
			d, e := row[0].(pgtype.Numeric), row[1].(pgtype.Numeric)
			deposit, _ = decimal.NewFromBigInt(d.Int, d.Exp).Float64()
			earn, _ = decimal.NewFromBigInt(e.Int, e.Exp).Float64()
		}
	}

	// withdraw
	query = `select sum(amount),sum(meca_quantity) from %s.vault_volume_24h 
where token='%s' and date_trunc('day',datetime)='%s' and event_type='WITHDRAW' group by event_type`
	query = fmt.Sprintf(query, SCHEMA, address, today)
	row, err = Row(ctx, query)
	if err != nil {
		logger.Error("VaultDashboard", "err", err, "query", query)
	} else if row != nil {
		if row[0] != nil && row[1] != nil {
			w, b := row[0].(pgtype.Numeric), row[1].(pgtype.Numeric)
			withdraw, _ = decimal.NewFromBigInt(w.Int, w.Exp).Float64()
			burn, _ = decimal.NewFromBigInt(b.Int, b.Exp).Float64()
		}
	}

	query = `select price,price_change,price_rate,treasury,treasury_change,weight,weight_change from %s.vault_price_daily 
where address='%s' and "day"='%s' limit 1`
	query = fmt.Sprintf(query, SCHEMA, address, today)
	row, err = Row(ctx, query)
	if err != nil {
		logger.Error("VaultDashboard", "err", err, "query", query)
	}

	var price, priceChange, priceRate, treasury, treasuryChange, weight, weightChange float64
	if row != nil {
		if row[0] != nil && row[1] != nil && row[2] != nil {
			p, pc, pr := row[0].(pgtype.Numeric), row[1].(pgtype.Numeric), row[2].(pgtype.Numeric)
			price, _ = decimal.NewFromBigInt(p.Int, p.Exp).Float64()
			priceChange, _ = decimal.NewFromBigInt(pc.Int, pc.Exp).Float64()
			priceRate, _ = decimal.NewFromBigInt(pr.Int, pr.Exp).Float64()
		}
		if row[3] != nil && row[4] != nil {
			to, tc := row[3].(pgtype.Numeric), row[4].(pgtype.Numeric)
			treasury, _ = decimal.NewFromBigInt(to.Int, to.Exp).Float64()
			treasuryChange, _ = decimal.NewFromBigInt(tc.Int, tc.Exp).Float64()
		}
		if row[5] != nil && row[6] != nil {
			wo, wc := row[5].(pgtype.Numeric), row[6].(pgtype.Numeric)
			weight, _ = decimal.NewFromBigInt(wo.Int, wo.Exp).Float64()
			weightChange, _ = decimal.NewFromBigInt(wc.Int, wc.Exp).Float64()
		}
	}

	dashboard := model.VaultDashboardApi{
		Exchange: price, ExchangeChange: priceChange, ExchangeRate: priceRate,
		TotalLocked: treasury, TotalChange: treasuryChange,
		Weight: weight, WeightChange: weightChange,
		Deposit: deposit, Earn: earn,
		Withdraw: withdraw, Burn: burn,
		MecaPerToken: price,
	}

	if weight != 0 {
		dashboard.TokenPerMeca = treasury / weight
	}
	return &dashboard, nil
}

func VaultChartDataForRate(ctx context.Context, address string) (*model.CommonChartApi, error) {
	query := `select day,price_open,price_high,price_low,price from %s.vault_price_daily where address='%s' order by day limit 30`
	query = fmt.Sprintf(query, SCHEMA, address)
	rows, _ := Rows(ctx, query)

	var priceList []model.PriceChartItem
	for _, p := range rows {
		d, o, h, l, c := p[0].(time.Time), p[1].(pgtype.Numeric), p[2].(pgtype.Numeric), p[3].(pgtype.Numeric), p[4].(pgtype.Numeric)
		of, _ := decimal.NewFromBigInt(o.Int, o.Exp).Float64()
		hf, _ := decimal.NewFromBigInt(h.Int, h.Exp).Float64()
		lf, _ := decimal.NewFromBigInt(l.Int, l.Exp).Float64()
		cf, _ := decimal.NewFromBigInt(c.Int, c.Exp).Float64()
		item := model.PriceChartItem{Datetime: d.Format("2006-01-02"), Open: of, High: hf, Low: lf, Close: cf}
		priceList = append(priceList, item)
	}

	query = `select day, sum(amount), first_value(min(event_type)) over (order by min(amount)) as event_type
	from %s.vault_volume_daily where address='%s' group by 1 order by 1 limit 30`
	query = fmt.Sprintf(query, SCHEMA, address)
	rows, _ = Rows(ctx, query)

	var volumeList []model.VolumeChartItem
	for _, r := range rows {
		d, v, t := r[0].(time.Time), r[1].(pgtype.Numeric), r[2].(string)
		vf, _ := decimal.NewFromBigInt(v.Int, v.Exp).Float64()
		item := model.VolumeChartItem{Datetime: d.Format("2006-01-02"), Value: vf, OrderType: t}
		volumeList = append(volumeList, item)
	}
	return &model.CommonChartApi{Price: priceList, Volume: volumeList}, nil
}

func VaultChartDataForVolume(ctx context.Context, address string) ([]model.BarChartApi, error) {
	query := `select day,treasury from %s.vault_price_daily where address='%s' order by day limit 30`
	query = fmt.Sprintf(query, SCHEMA, address)
	rows, err := Rows(ctx, query)
	if err != nil {
		return nil, err
	}

	var volumeList []model.BarChartApi
	for _, p := range rows {
		d, t := p[0].(time.Time), p[1].(pgtype.Numeric)
		tf, _ := decimal.NewFromBigInt(t.Int, t.Exp).Float64()
		item := model.BarChartApi{Datetime: d.Format("2006-01-02"), Value: tf}
		volumeList = append(volumeList, item)
	}
	return volumeList, nil
}

func VaultChartDataForValue(ctx context.Context, address string) ([]model.BarChartApi, error) {
	query := `select day,tvl from %s.vault_price_daily where address='%s' order by day limit 30`
	query = fmt.Sprintf(query, SCHEMA, address)
	rows, err := Rows(ctx, query)
	if err != nil {
		return nil, err
	}

	var valueList []model.BarChartApi
	for _, p := range rows {
		d, t := p[0].(time.Time), p[1].(pgtype.Numeric)
		tf, _ := decimal.NewFromBigInt(t.Int, t.Exp).Float64()
		item := model.BarChartApi{Datetime: d.Format("2006-01-02"), Value: tf}
		valueList = append(valueList, item)
	}
	return valueList, nil
}
