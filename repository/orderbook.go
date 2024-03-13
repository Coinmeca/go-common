package repository

import (
	"coinmeca-go_common/logger"
	"coinmeca-go_common/model"
	"context"
	"fmt"
	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
	"time"
)

func SetMarketOrderbook(ctx context.Context, o model.MarketOrderbook) {
	query := `insert into %s.market_orderbook(order_type,order_index,price,balance,address,updated_at)
values($1,$2,$3,$4,$5,$6) ON CONFLICT ON CONSTRAINT pk_market_orderbook
DO UPDATE SET price=$3,balance=$4,address=$5,updated_at=$6`
	query = fmt.Sprintf(query, SCHEMA)

	_, err := POOL.Exec(ctx, query, o.OrderType, o.OrderIndex, o.Price, o.Balance, o.Address, time.Now().UTC())
	if err != nil {
		logger.Error("SetMarketOrderbook", "err", err, "query", query)
	} else {
		logger.Debug("SetMarketOrderbook", "query", query)
	}
}

////////////////////////////////////////////////////////////////////////////////////////////////////////////////////////
// for API

func MarketOrderbook(ctx context.Context, address string) (*model.OrderbookPriceApi, error) {
	query := `select order_type, price, balance from %s.market_orderbook where address='%s' and order_index <= 30 order by order_index`
	query = fmt.Sprintf(query, SCHEMA, address)
	rows, err := Rows(ctx, query)
	if err != nil {
		logger.Error("MarketOrderbook", "err", err, "query", query)
		return nil, err
	}

	var asks []model.OrderbookPriceItem
	var bids []model.OrderbookPriceItem
	for _, r := range rows {
		orderType, p, b := r[0].(string), r[1].(pgtype.Numeric), r[2].(pgtype.Numeric)
		price, _ := decimal.NewFromBigInt(p.Int, p.Exp).Float64()
		balance, _ := decimal.NewFromBigInt(b.Int, b.Exp).Float64()
		if orderType == "ASK" {
			asks = append(asks, model.OrderbookPriceItem{Price: price, Balance: balance})
		} else if orderType == "BID" {
			bids = append(bids, model.OrderbookPriceItem{Price: price, Balance: balance})
		}
	}
	return &model.OrderbookPriceApi{Asks: asks, Bids: bids}, nil
}
