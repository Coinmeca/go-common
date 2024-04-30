package repository

import (
	"context"

	common "github.com/coinmeca/go-common/chain"
	logger "github.com/coinmeca/go-common/logger"
	"github.com/coinmeca/go-common/model"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
	"strings"
	"time"
)

func SetMarketPrice(ctx context.Context, mp model.MarketPriceRow) {
	// add 24h
	deleteQuery := `delete from %s.market_price_24h where address='%s' and datetime < now()::timestamptz at time zone 'utc' - '1day'::interval`
	_, err := POOL.Exec(ctx, fmt.Sprintf(deleteQuery, SCHEMA, mp.Address))
	if err != nil {
		logger.Error("SetMarketPrice", "err", err)
	}

	nowUTC := time.Now().UTC()
	query := `insert into %s.market_price_24h(datetime,price,address,updated_at) values($1,$2,$3,$4)
ON CONFLICT ON CONSTRAINT pk_market_price_24h DO UPDATE SET price=$2,updated_at=$4`
	query = fmt.Sprintf(query, SCHEMA)
	_, err = POOL.Exec(ctx, query, mp.Time, mp.Price, mp.Address, nowUTC)
	if err != nil {
		logger.Error("SetMarketPrice", "err", err)
	} else {
		logger.Debug("SetMarketPrice", "address", mp.Address, "price", mp.Price)
	}

	// update daily
	today := nowUTC.Format("2006-01-02")
	query = `select open,high,low from %s.market_price_daily where address='%s' and day = '%s' limit 1`
	query = fmt.Sprintf(query, SCHEMA, mp.Address, today)
	row, _ := Row(ctx, query)

	if row == nil {
		AddDailyMarketPrice(ctx, mp, today)
	} else {
		UpdateDailyMarketPrice(ctx, mp, row, today)
	}
}

func SetMarketVolume(ctx context.Context, mv model.MarketVolumeRow) {
	deleteQuery := `delete from %s.market_volume_24h where updated_at < now()::timestamptz at time zone 'utc' - '1day'::interval`
	_, err := POOL.Exec(ctx, fmt.Sprintf(deleteQuery, SCHEMA))
	if err != nil {
		logger.Error("SetMarketVolume", "err", err)
	}

	nowUTC := time.Now().UTC()
	query := `insert into %s.market_volume_24h(datetime,owner,event_type,price,amount,quantity,block_no,block_hash,tx_hash,tx_index,address) 
values($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)`
	query = fmt.Sprintf(query, SCHEMA)
	_, err = POOL.Exec(ctx, query, nowUTC, mv.Owner, mv.EventType, mv.Price, mv.Amount, mv.Quantity, mv.BlockNo, mv.BlockHash, mv.TxHash, mv.TxIndex, mv.Address)
	if err != nil {
		logger.Error("SetMarketVolume", "err", err)
	} else {
		logger.Debug("SetMarketVolume", "volumeRow", mv)
	}
	setOverviewVolume(ctx, mv.Quantity, strings.ToLower(mv.EventType))

	today := nowUTC.Format("2006-01-02")
	query = `select volume from %s.market_volume_daily where address='%s' and day='%s' and event_type='%s' limit 1`
	query = fmt.Sprintf(query, SCHEMA, mv.Address, today, mv.EventType)
	row, _ := Row(ctx, query)

	if row == nil {
		AddDailyMarketVolume(ctx, mv, today)
	} else {
		UpdateDailyMarketVolume(ctx, mv, row, today)
	}
}

func SetMinutelyMarketPrice(ctx context.Context, mv model.MarketPriceRow) {
	deleteQuery := `delete from %s.market_price_minutely where updated_at < now()::timestamptz at time zone 'utc' - '1day'::interval`
	_, err := POOL.Exec(ctx, fmt.Sprintf(deleteQuery, SCHEMA))
	if err != nil {
		logger.Error("SetMinutelyMarketPrice", "err", err)
	}

	t := time.Now().UTC()
	query := `select price from %s.market_price_24h where address='%s' order by datetime desc limit 1`
	query = fmt.Sprintf(query, SCHEMA, mv.Address)
	row, _ := Row(ctx, query)

	pd := decimal.Zero
	if row != nil {
		p := row[0].(pgtype.Numeric)
		pd = decimal.NewFromBigInt(p.Int, p.Exp)
	}

	upsertQuery := `insert into %s.market_price_minutely(datetime,price,address,updated_at) values($1,$2,$3,$4)
ON CONFLICT ON CONSTRAINT pk_market_price_minutely DO UPDATE SET price=$2,updated_at=$4`
	upsertQuery = fmt.Sprintf(upsertQuery, SCHEMA)

	dateTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	_, err = POOL.Exec(ctx, upsertQuery, dateTime, pd, mv.Address, t)
	if err != nil {
		logger.Error("SetMinutelyMarketPrice", "err", err)
	} else {
		logger.Debug("SetMinutelyMarketPrice", "address", mv.Address)
	}
}

func SetMinutelyMarketVolume(ctx context.Context, mv model.MarketVolumeRow) {
	deleteQuery := `delete from %s.market_volume_minutely where updated_at < now()::timestamptz at time zone 'utc' - '1day'::interval`
	_, err := POOL.Exec(ctx, fmt.Sprintf(deleteQuery, SCHEMA))
	if err != nil {
		logger.Error("SetMinutelyMarketVolume", "err", err)
	}

	t := time.Now().UTC()
	m := t.Add(-1 * time.Minute)
	condTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", m.Year(), m.Month(), m.Day(), m.Hour(), m.Minute())
	query := `select sum(quantity) from %s.market_volume_24h where address='%s' and datetime>='%s' group by address`
	query = fmt.Sprintf(query, SCHEMA, mv.Address, condTime)
	row, _ := Row(ctx, query)

	vd := decimal.Zero
	if row != nil {
		v := row[0].(pgtype.Numeric)
		vd = decimal.NewFromBigInt(v.Int, v.Exp).Add(mv.Quantity)
	}

	upsertQuery := `insert into %s.market_volume_minutely(datetime,volume,address,updated_at) values($1,$2,$3,$4)
ON CONFLICT ON CONSTRAINT pk_market_volume_minutely DO UPDATE SET volume=$2,updated_at=$4`
	upsertQuery = fmt.Sprintf(upsertQuery, SCHEMA)

	dateTime := fmt.Sprintf("%d-%02d-%02d %02d:%02d:00", t.Year(), t.Month(), t.Day(), t.Hour(), t.Minute())
	_, err = POOL.Exec(ctx, upsertQuery, dateTime, vd, mv.Address, t)
	if err != nil {
		logger.Error("SetMinutelyMarketVolume", "err", err)
	} else {
		logger.Debug("SetMinutelyMarketVolume", "address", mv.Address, "quantity", mv.Quantity)
	}
}

func SetMarketToken(ctx context.Context, m model.MarketTokenRow) {
	query := `insert into %s.market_tokens(symbol,name,decimals,base,market,nft,quote,updated_at)
values($1,$2,$3,$4,$5,$6,$7,$8) ON CONFLICT ON CONSTRAINT pk_market_tokens
DO UPDATE SET name=$2,decimals=$3,base=$4,market=$5,nft=$6,quote=$7,updated_at=$8`
	query = fmt.Sprintf(query, SCHEMA)

	now := time.Now().UTC()
	_, err := POOL.Exec(ctx, query, m.Symbol, m.Name, m.Decimals, m.Base, m.Orderbook, m.NFT, m.Quote, now)

	if err != nil {
		logger.Error("SetMarketToken", "err", err)
	} else {
		logger.Debug("SetMarketToken", "marketTokenRow", m)
	}
}

func AddDailyMarketVolume(ctx context.Context, v model.MarketVolumeRow, date string) {
	insertQuery := `insert into %s.market_volume_daily(day,volume,event_type,address) values($1,$2,$3,$4)`
	insertQuery = fmt.Sprintf(insertQuery, SCHEMA)

	_, err := POOL.Exec(ctx, insertQuery, date, v.Quantity, v.EventType, v.Address)
	if err != nil {
		//logger.Error.Printf("add daily volume error: %v (%v)\n", err, insertQuery)
		logger.Error("AddDailyMarketVolume", "err", err)
	} else {
		logger.Debug("AddDailyMarketVolume", "address", v.Address, "eventType", v.EventType, "quantity", v.Quantity)
	}
}

func UpdateDailyMarketVolume(ctx context.Context, mv model.MarketVolumeRow, row []interface{}, date string) {
	v := row[0].(pgtype.Numeric)
	vd := decimal.NewFromBigInt(v.Int, v.Exp)
	vq := vd.Add(mv.Quantity)

	updateQuery := `update %s.market_volume_daily set volume=$1,updated_at=$2 where address='%s' and day='%s' and event_type='%s'`
	updateQuery = fmt.Sprintf(updateQuery, SCHEMA, mv.Address, date, mv.EventType)
	_, err := POOL.Exec(ctx, updateQuery, vq, time.Now().UTC())
	if err != nil {
		logger.Error("UpdateDailyMarketVolume", "err", err)
	} else {
		logger.Debug("UpdateDailyMarketVolume", "date", date, "quantity", vq)
	}
}

func AddDailyMarketPrice(ctx context.Context, p model.MarketPriceRow, date string) {
	insertQuery := `insert into %s.market_price_daily(day,open,high,low,address) values($1,$2,$3,$4,$5)`
	insertQuery = fmt.Sprintf(insertQuery, SCHEMA)

	_, err := POOL.Exec(ctx, insertQuery, date, p.Price, p.Price, p.Price, p.Address)
	if err != nil {
		logger.Error("AddDailyMarketPrice", "err", err, "query", insertQuery)
	} else {
		logger.Debug("AddDailyMarketPrice", "date", date, "price", p.Price)
	}
}

func UpdateDailyMarketPrice(ctx context.Context, mp model.MarketPriceRow, row []interface{}, date string) {
	o, h, l := row[0].(pgtype.Numeric), row[1].(pgtype.Numeric), row[2].(pgtype.Numeric)
	op, hp, lp := decimal.NewFromBigInt(o.Int, o.Exp), decimal.NewFromBigInt(h.Int, h.Exp), decimal.NewFromBigInt(l.Int, l.Exp)
	if mp.Price.GreaterThan(hp) {
		hp = mp.Price
	} else if mp.Price.LessThan(lp) {
		lp = mp.Price
	}
	changed, changedRate := mp.Price.Sub(op), decimal.Zero
	if !op.IsZero() {
		changedRate = mp.Price.Div(op).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100))
	}

	updateQuery := `update %s.market_price_daily set high=$1,low=$2,close=$3,change=$4,change_rate=$5,price=$6,updated_at=$7 where address='%s' and day = '%s'`
	updateQuery = fmt.Sprintf(updateQuery, SCHEMA, mp.Address, date)
	_, err := POOL.Exec(ctx, updateQuery, hp, lp, mp.Price, changed, changedRate, mp.Price, time.Now().UTC())
	if err != nil {
		logger.Error("UpdateDailyMarketPrice", "err", err)
	} else {
		logger.Debug("UpdateDailyMarketPrice", "date", date, "changed", changed, "changedRate", changedRate)
	}
}

func LastBlockNoFromMarketVolume(ctx context.Context) (int64, error) {
	query := `select block_no from %s.market_volume_24h order by block_no desc limit 1`
	query = fmt.Sprintf(query, SCHEMA)
	row, err := Row(ctx, query)
	if err != nil {
		return 0, err
	}
	return row[0].(int64), nil
}

func MarketTokenInfo(ctx context.Context) ([]*model.TokenInfo, error) {
	query := `select market, decimals from %s.market_tokens`
	query = fmt.Sprintf(query, SCHEMA)
	rows, err := Rows(ctx, query)
	if err != nil {
		// TODO: edit elog
		//logger.Error.Printf("market token error: %v (%v)\n", err, query)
		return nil, err
	}

	var tokens []*model.TokenInfo
	for _, r := range rows {
		t := model.TokenInfo{Address: r[0].(string), Decimals: r[1].(int32)}
		tokens = append(tokens, &t)
	}
	return tokens, nil
}

//////////////////////////////////////////////////////////////////////////////////
// for API

func MarketStats(ctx context.Context, params ...string) ([]model.MarketStatApi, error) {
	orderbook := params[0]
	today := time.Now().UTC().Format("2006-01-02")

	query := `select mt.market,nft,quote,mp.price,mp.change_rate,mv.volume,mt.decimals,mt.base,mt.symbol,mt.name from %s.market_tokens mt 
left join (select address,price,change_rate from %s.market_price_daily where day = '%s') mp on mp.address = mt.market 
left join (select address,sum(volume)as volume from %s.market_volume_daily where day = '%s' group by 1) mv on mv.address = mt.market where %s`

	condition := "1=1"
	if orderbook != "" {
		condition = fmt.Sprintf("market='%s'", orderbook)
	}
	query = fmt.Sprintf(query, SCHEMA, SCHEMA, today, SCHEMA, today, condition)
	rows, err := Rows(ctx, query)
	if err != nil {
		logger.Debug("MarketStats", "query", query)
		return nil, err
	}

	var marketStats []model.MarketStatApi
	for _, r := range rows {
		stat := model.MarketStatApi{
			Orderbook: r[0].(string), NFT: r[1].(string), Quote: r[2].(string),
			Decimals: r[6].(int32), Base: r[7].(string), Symbol: r[8].(string), Name: r[9].(string),
		}
		if r[3] != nil {
			p := r[3].(pgtype.Numeric)
			stat.Price, _ = decimal.NewFromBigInt(p.Int, p.Exp).Float64()
		}
		if r[4] != nil {
			c := r[4].(pgtype.Numeric)
			stat.Change, _ = decimal.NewFromBigInt(c.Int, c.Exp).Float64()
		}
		if r[5] != nil {
			v := r[5].(pgtype.Numeric)
			stat.Volume, _ = decimal.NewFromBigInt(v.Int, v.Exp).Float64()
		}
		marketStats = append(marketStats, stat)
	}
	return marketStats, nil
}

func MarketChartData(ctx context.Context, address, period string) (*model.CommonChartApi, error) {
	query := `with volume as (select day,sum(volume) as vol from %s.market_volume_daily where address='%s' group by 1 order by 1)
select mp.day,open,high,low,close,coalesce(vol, 0),case when open>=close then 'SELL' else 'BUY' end as ev from %s.market_price_daily as mp
left join volume v on mp.day = v.day where address='%s' order by day limit 30`
	query = fmt.Sprintf(query, SCHEMA, address, SCHEMA, address)
	if period == "1min" {
		query = `with low_high as (
select day,open,high,low from %s.market_price_daily where address='%s' order by 1 desc limit 2
), vol_min as (
select address,datetime,volume from %s.market_volume_minutely where address='%s' and datetime >= now()::timestamptz at time zone 'utc' - '1day'::interval
)
select v.datetime,p.price,coalesce(volume, 0)as volume,open,high,low,case when lh.open>=p.price then 'BUY' else 'SELL' end as ev from vol_min as v 
left join %s.market_price_minutely as p on v.datetime = p.datetime and v.address = p.address 
left join low_high lh on date_trunc('day', v.datetime) = lh.day where v.address='%s' order by 1`
		query = fmt.Sprintf(query, SCHEMA, address, SCHEMA, address, SCHEMA, address)
	} else if period == "3min" || period == "5min" {
		query = `with low_high as (
select day,open,high,low from %s.market_price_daily where address='%s' order by 1 desc limit 2
), min_vol as (
select volume, datetime, floor(date_part('minute',datetime)::numeric/%d)*%d as m from %s.market_volume_minutely where address='%s' and datetime >= now()::timestamptz at time zone 'utc' - '1day'::interval
), sum_vol as (
select sum(volume) as volume, to_char(datetime,'yyyymmddhh24')||':'|| case when m < 10 then '0'||cast(m as text) else cast(m as text)end as d from min_vol group by 2
), dates_vol as (
select volume, '%s' as address, to_timestamp(to_date(substring(d,1,8),'YYYYMMDD')||' '||split_part(substring(d,9,12),':',1)||':'||split_part(substring(d,9,12),':',2),'YYYY-MM-DD HH24:MI')::timestamp as datetime from sum_vol
)
select v.datetime,coalesce(p.price,0)as price,coalesce(volume, 0)as volume,open,high,low,case when lh.open>=p.price then 'BUY' else 'SELL' end as ev from dates_vol as v 
left join %s.market_price_minutely as p on v.datetime = p.datetime and v.address = p.address 
left join low_high lh on date_trunc('day', v.datetime) = lh.day 
where v.address='%s' order by 1`
		min := 3
		if period == "5min" {
			min = 5
		}
		query = fmt.Sprintf(query, SCHEMA, address, min, min, SCHEMA, address, address, SCHEMA, address)
	}

	rows, err := Rows(ctx, query)
	if err != nil {
		return nil, err
	}

	var resultChart model.CommonChartApi
	for _, r := range rows {
		o, h, l, c := r[1].(pgtype.Numeric), r[2].(pgtype.Numeric), r[3].(pgtype.Numeric), r[4].(pgtype.Numeric)
		d, v, t := r[0].(time.Time), r[5].(pgtype.Numeric), r[6].(string)
		ds := d.Format("2006-01-02")
		if period != "" {
			ds = d.Format(common.DateFormat)
		}

		of, _ := decimal.NewFromBigInt(o.Int, o.Exp).Float64()
		hf, _ := decimal.NewFromBigInt(h.Int, h.Exp).Float64()
		lf, _ := decimal.NewFromBigInt(l.Int, l.Exp).Float64()
		cf, _ := decimal.NewFromBigInt(c.Int, c.Exp).Float64()
		price := model.PriceChartItem{Datetime: ds, Open: of, High: hf, Low: lf, Close: cf}
		resultChart.Price = append(resultChart.Price, price)

		vf, _ := decimal.NewFromBigInt(v.Int, v.Exp).Float64()
		volume := model.VolumeChartItem{Datetime: ds, Value: vf, OrderType: t}
		resultChart.Volume = append(resultChart.Volume, volume)
	}
	return &resultChart, nil
}

func MarketDashboard(ctx context.Context, address string) (*model.MarketDashboardApi, error) {
	query := `select event_type,price,quantity from %s.market_volume_24h 
where address='%s' and date_trunc('day',datetime)='%s' order by datetime`
	query = fmt.Sprintf(query, SCHEMA, address, time.Now().UTC().Format("2006-01-02"))
	rows, err := Rows(ctx, query)
	if err != nil {
		//logger.Error.Printf("market volume error: %v (%v)\n", err, query)
		return nil, err
	}

	var d model.MarketDashboardApi
	var bt, qt, o, h, l, c decimal.Decimal // base (sell), quote (buy)
	if len(rows) == 0 {
		query = `select price from %s.market_price_daily where address='%s' order by 1 limit 1`
		query = fmt.Sprintf(query, SCHEMA, address)
		row, e := Row(ctx, query)
		if e != nil {
			return &d, e
		}
		p := row[0].(pgtype.Numeric)
		d.Price, _ = decimal.NewFromBigInt(p.Int, p.Exp).Float64()
		d.Low = d.Price
		d.High = d.Price
	} else {
		for i, ts := range rows {
			p, q := ts[1].(pgtype.Numeric), ts[2].(pgtype.Numeric)
			pd, qd := decimal.NewFromBigInt(p.Int, p.Exp), decimal.NewFromBigInt(q.Int, q.Exp)

			if i == 0 {
				l = pd
				o = pd // open
			} else {
				c = pd // close
			}

			if ts[0].(string) == "BUY" {
				qt = qt.Add(qd) // quote
			} else {
				bt = bt.Add(qd) // quote
			}

			if pd.LessThan(l) {
				l = pd // low
			} else if pd.GreaterThan(h) {
				h = pd // high
			}
		}

		d.Price, _ = c.Float64()
		d.VolumeQuote, _ = qt.Float64()
		d.VolumeBase, _ = bt.Float64()
		d.Low, _ = l.Float64()
		d.High, _ = h.Float64()
		d.Change, _ = c.Sub(o).Float64()
		if !o.IsZero() {
			d.ChangeRate, _ = c.Div(o).Sub(decimal.NewFromInt(1)).Mul(decimal.NewFromInt(100)).Float64()
		}
	}
	//fmt.Printf("o %v, h %v, l %v, c %v, q %v, b %v\n", o, h, l, c, qt, bt)
	return &d, nil
}

func MarketHistory(ctx context.Context, address string) ([]model.MarketHistoryApi, error) {
	query := `select event_type,quantity,price,datetime from %s.market_volume_24h where address='%s' order by datetime`
	query = fmt.Sprintf(query, SCHEMA, address)
	rows, err := Rows(ctx, query)
	if err != nil {
		return nil, err
	}

	var historyList []model.MarketHistoryApi
	for _, r := range rows {
		event, q, p, ts := r[0].(string), r[1].(pgtype.Numeric), r[2].(pgtype.Numeric), r[3].(time.Time)
		qf, _ := decimal.NewFromBigInt(q.Int, q.Exp).Float64()
		pf, _ := decimal.NewFromBigInt(p.Int, p.Exp).Float64()
		item := model.MarketHistoryApi{Type: event, Quantity: qf, Time: ts.Format(common.DateFormat), Price: pf}
		historyList = append(historyList, item)
	}
	return historyList, nil
}

func MarketTokenList(ctx context.Context) ([]model.MarketTokenApi, error) {
	query := `select symbol,name,decimals,base,market,nft,quote from %s.market_tokens`
	query = fmt.Sprintf(query, SCHEMA)
	tokens, err := Rows(ctx, query)
	if err != nil {
		return nil, err
	}

	var marketTokens []model.MarketTokenApi
	for _, t := range tokens {
		symbol, name, dec, base := t[0].(string), t[1].(string), t[2].(int32), t[3].(string)
		orderbook, nft, quote := t[4].(string), t[5].(string), t[6].(string)
		token := model.MarketTokenApi{
			Symbol: symbol, Name: name, Decimals: dec, Base: base, Orderbook: orderbook, NFT: nft, Quote: quote,
		}

		marketTokens = append(marketTokens, token)
	}
	return marketTokens, nil
}

func LastMarketChart(ctx context.Context, address, period string) (*model.CommonChartApi, error) {
	query := `with volume as (select day,sum(volume) as vol from %s.market_volume_daily where address='%s' group by 1 order by 1)
select mp.day,open,high,low,close,coalesce(vol, 0),case when open>=close then 'SELL' else 'BUY' end as ev from %s.market_price_daily as mp
left join volume v on mp.day = v.day where address='%s' order by day desc limit 1`
	query = fmt.Sprintf(query, SCHEMA, address, SCHEMA, address)
	if period == "1min" {
		query = `with low_high as (
select day,open,high,low from %s.market_price_daily where address='%s' order by 1 desc limit 1
), vol_min as (
select datetime,volume,address from %s.market_volume_minutely where address='%s' and datetime >= now()::timestamptz at time zone 'utc' - '1day'::interval
)
select v.datetime,p.price,coalesce(volume, 0)as volume,open,high,low,case when lh.open>=p.price then 'BUY' else 'SELL' end as ev from vol_min as v 
left join %s.market_price_minutely as p on v.datetime = p.datetime and v.address = p.address 
left join low_high lh on date_trunc('day', v.datetime) = lh.day where v.address='%s' order by 1 desc limit 1`
		query = fmt.Sprintf(query, SCHEMA, address, SCHEMA, address, SCHEMA, address)
	} else if period == "3min" || period == "5min" {
		query = `with low_high as (
select day,open,high,low from %s.market_price_daily where address='%s' order by 1 desc limit 1
), min_vol as (
select volume, datetime, floor(date_part('minute',datetime)::numeric/%d)*%d as m from %s.market_volume_minutely where address='%s' and datetime >= now()::timestamptz at time zone 'utc' - '1day'::interval
), sum_vol as (
select sum(volume) as volume, to_char(datetime,'yyyymmddhh24')||':'|| case when m < 10 then '0'||cast(m as text) else cast(m as text)end as d from min_vol group by 2
), dates_vol as (
select volume, '%s' as address, to_timestamp(to_date(substring(d,1,8),'YYYYMMDD')||' '||split_part(substring(d,9,12),':',1)||':'||split_part(substring(d,9,12),':',2),'YYYY-MM-DD HH24:MI')::timestamp as datetime from sum_vol
)
select v.datetime,coalesce(p.price,0)as price,coalesce(volume, 0)as volume,open,high,low,case when lh.open>=p.price then 'BUY' else 'SELL' end as ev from dates_vol as v 
left join %s.market_price_minutely as p on v.datetime = p.datetime and v.address = p.address 
left join low_high lh on date_trunc('day', v.datetime) = lh.day 
where v.address='%s' order by 1 desc limit 1`
		min := 3
		if period == "5min" {
			min = 5
		}
		query = fmt.Sprintf(query, SCHEMA, address, min, min, SCHEMA, address, address, SCHEMA, address)
	}

	r, err := Row(ctx, query)
	if err != nil {
		return nil, err
	}

	lastResult := model.CommonChartApi{Price: make([]model.PriceChartItem, 1), Volume: make([]model.VolumeChartItem, 1)}
	o, h, l, c := r[1].(pgtype.Numeric), r[2].(pgtype.Numeric), r[3].(pgtype.Numeric), r[4].(pgtype.Numeric)
	d, v, t := r[0].(time.Time), r[5].(pgtype.Numeric), r[6].(string)
	ds := d.Format("2006-01-02")
	if period != "" {
		ds = d.Format(common.DateFormat)
	}

	of, _ := decimal.NewFromBigInt(o.Int, o.Exp).Float64()
	hf, _ := decimal.NewFromBigInt(h.Int, h.Exp).Float64()
	lf, _ := decimal.NewFromBigInt(l.Int, l.Exp).Float64()
	cf, _ := decimal.NewFromBigInt(c.Int, c.Exp).Float64()
	lastResult.Price[0] = model.PriceChartItem{Datetime: ds, Open: of, High: hf, Low: lf, Close: cf}

	vf, _ := decimal.NewFromBigInt(v.Int, v.Exp).Float64()
	lastResult.Volume[0] = model.VolumeChartItem{Datetime: ds, Value: vf, OrderType: t}

	return &lastResult, nil
}
