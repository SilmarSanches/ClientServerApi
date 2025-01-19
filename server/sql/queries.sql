-- name: InsertExchangeRate :exec
INSERT INTO exchange_rate (code, codein, name, high, low, varBid, pctChange, bid, ask, timestamp, create_date)
VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?);

-- name: GetExchangeRate :one
SELECT * FROM exchange_rate WHERE id = ?;

-- name: ListExchangeRates :many
SELECT * FROM exchange_rate ORDER BY create_date DESC;