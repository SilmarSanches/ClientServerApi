// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0
// source: queries.sql

package db

import (
	"context"
)

const getExchangeRate = `-- name: GetExchangeRate :one
SELECT id, code, codein, name, high, low, varbid, pctchange, bid, ask, timestamp, create_date FROM exchange_rate WHERE id = ?
`

func (q *Queries) GetExchangeRate(ctx context.Context, id int64) (ExchangeRate, error) {
	row := q.db.QueryRowContext(ctx, getExchangeRate, id)
	var i ExchangeRate
	err := row.Scan(
		&i.ID,
		&i.Code,
		&i.Codein,
		&i.Name,
		&i.High,
		&i.Low,
		&i.Varbid,
		&i.Pctchange,
		&i.Bid,
		&i.Ask,
		&i.Timestamp,
		&i.CreateDate,
	)
	return i, err
}

const insertExchangeRate = `-- name: InsertExchangeRate :exec
INSERT INTO exchange_rate (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)
`

type InsertExchangeRateParams struct {
	Code       string
	Codein     string
	Name       string
	High       float64
	Low        float64
	Varbid     float64
	Pctchange  float64
	Bid        float64
	Ask        float64
	Timestamp  int64
	CreateDate string
}

func (q *Queries) InsertExchangeRate(ctx context.Context, arg InsertExchangeRateParams) error {
	_, err := q.db.ExecContext(ctx, insertExchangeRate,
		arg.Code,
		arg.Codein,
		arg.Name,
		arg.High,
		arg.Low,
		arg.Varbid,
		arg.Pctchange,
		arg.Bid,
		arg.Ask,
		arg.Timestamp,
		arg.CreateDate,
	)
	return err
}

const listExchangeRates = `-- name: ListExchangeRates :many
SELECT id, code, codein, name, high, low, varbid, pctchange, bid, ask, timestamp, create_date FROM exchange_rate ORDER BY create_date DESC
`

func (q *Queries) ListExchangeRates(ctx context.Context) ([]ExchangeRate, error) {
	rows, err := q.db.QueryContext(ctx, listExchangeRates)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []ExchangeRate
	for rows.Next() {
		var i ExchangeRate
		if err := rows.Scan(
			&i.ID,
			&i.Code,
			&i.Codein,
			&i.Name,
			&i.High,
			&i.Low,
			&i.Varbid,
			&i.Pctchange,
			&i.Bid,
			&i.Ask,
			&i.Timestamp,
			&i.CreateDate,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}
