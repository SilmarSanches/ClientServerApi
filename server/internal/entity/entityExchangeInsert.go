package entity

import (
	"errors"
	"strconv"
	"time"
)

type ExchangeRate struct {
	Code       string    `json:"code"`
	CodeIn     string    `json:"codein"`
	Name       string    `json:"name"`
	High       float64   `json:"high"`
	Low        float64   `json:"low"`
	VarBid     float64   `json:"varBid"`
	PctChange  float64   `json:"pctChange"`
	Bid        float64   `json:"bid"`
	Ask        float64   `json:"ask"`
	Timestamp  time.Time `json:"timestamp"`
	CreateDate time.Time `json:"create_date"`
}

func NewExchangeInsert(data map[string]interface{}) (*ExchangeRate, error) {
	usdb, ok := data["USDBRL"].(map[string]interface{})
	if !ok {
		return nil, errors.New("formato inválido no retorno da API")
	}

	timestamp, err := parseTimestamp(usdb["timestamp"])
	if err != nil {
		return nil, err
	}

	createDate, err := time.Parse("2006-01-02 15:04:05", usdb["create_date"].(string))
	if err != nil {
		return nil, err
	}

	exchange := ExchangeRate{
		Code:       usdb["code"].(string),
		CodeIn:     usdb["codein"].(string),
		Name:       usdb["name"].(string),
		High:       parseFloat(usdb["high"]),
		Low:        parseFloat(usdb["low"]),
		VarBid:     parseFloat(usdb["varBid"]),
		PctChange:  parseFloat(usdb["pctChange"]),
		Bid:        parseFloat(usdb["bid"]),
		Ask:        parseFloat(usdb["ask"]),
		Timestamp:  timestamp,
		CreateDate: createDate,
	}

	return &exchange, nil
}

func parseFloat(value interface{}) float64 {
	switch v := value.(type) {
	case string:
		if f, err := strconv.ParseFloat(v, 64); err == nil {
			return f
		}
	case float64:
		return v
	}
	return 0.0
}

func parseTimestamp(value interface{}) (time.Time, error) {
	if v, ok := value.(string); ok {
		timestamp, err := strconv.ParseInt(v, 10, 64)
		if err != nil {
			return time.Time{}, err
		}
		return time.Unix(timestamp, 0), nil
	}
	return time.Time{}, errors.New("formato de timestamp inválido")
}
