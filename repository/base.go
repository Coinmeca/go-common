package repository

import (
	"context"
	"fmt"

	"github.com/coinmeca/go-common/logger"
)

func Rows(ctx context.Context, query string) ([][]interface{}, error) {
	//fmt.Println(query)
	rows, err := POOL.Query(ctx, query)
	if err != nil {
		logger.Error("Rows", "err", err)
		return nil, err
	}
	defer rows.Close()

	if rows == nil {
		return nil, fmt.Errorf("no rows")
	}

	var result [][]interface{} = nil
	for rows.Next() {
		data, _ := rows.Values()
		if len(data) == 1 && data[0] == nil {
			continue
		}
		result = append(result, data)
	}

	return result, nil
}

func Row(ctx context.Context, query string) ([]interface{}, error) {
	// query must include limit 1
	row, err := POOL.Query(ctx, query)
	if err != nil {
		logger.Error("Row", "err", err)
		return nil, err
	}
	defer row.Close()

	if row == nil {
		return nil, fmt.Errorf("no row")
	}

	if row.Next() {
		data, _ := row.Values()
		return data, nil
	}
	return nil, fmt.Errorf("no data")
}

func Scan(ctx context.Context, query string) (interface{}, error) {
	// query must have one field
	var res interface{}
	err := POOL.QueryRow(ctx, query).Scan(&res)
	if err != nil {
		logger.Error("Scan", "err", err, "query", query)
		return nil, err
	}

	if res == nil {
		return nil, fmt.Errorf("no scan")
	}

	return res, nil
}

func Truncate(ctx context.Context, table string) {
	query := fmt.Sprintf(`truncate %s.%s`, SCHEMA, table)
	_, err := POOL.Exec(ctx, query)
	if err != nil {
		logger.Error("Truncate", "err", err, "query", query)
	}
}
