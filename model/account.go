package model

import (
	"github.com/jackc/pgtype"
	"github.com/shopspring/decimal"
)

type AccountRow struct {
	Address   string
	Balance   decimal.Decimal
	PickPrice pgtype.Numeric
}
