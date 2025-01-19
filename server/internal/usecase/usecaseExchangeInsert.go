package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/silmarsanches/clientserverapi/server/config"
	"github.com/silmarsanches/clientserverapi/server/internal/entity"
	"github.com/silmarsanches/clientserverapi/server/internal/infra/db"
	"github.com/silmarsanches/clientserverapi/server/internal/infra/services"
)

type ExchangeInsertUseCase struct {
	DB        db.Queries
	External  services.ExternalServiceExchangeInterface
	appConfig config.Config
}

func NewExchangeInsertUseCase(db *db.Queries, appConfig *config.Config, external services.ExternalServiceExchangeInterface) *ExchangeInsertUseCase {
	return &ExchangeInsertUseCase{
		DB:        *db,
		External:  external,
		appConfig: *appConfig,
	}
}

func (e *ExchangeInsertUseCase) InsertExchange(ctx context.Context) (*entity.ExchangeRate, error) {
	data, err := e.External.GetExchangeRate(ctx)
	if err != nil {
		return nil, err
	}

	exchange, err := entity.NewExchangeInsert(data)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(ctx, 10*time.Millisecond)
	defer cancel()

	err = e.DB.InsertExchangeRate(ctx, db.InsertExchangeRateParams{
		Code:       exchange.Code,
		Codein:     exchange.CodeIn,
		Name:       exchange.Name,
		High:       exchange.High,
		Low:        exchange.Low,
		Varbid:     exchange.VarBid,
		Pctchange:  exchange.PctChange,
		Bid:        exchange.Bid,
		Ask:        exchange.Ask,
		Timestamp:  exchange.Timestamp.Unix(),
		CreateDate: exchange.CreateDate.Format("2006-01-02 15:04:05"),
	})
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Printf("Timeout de 10ms excedido ao inserir no banco de dados")
		}
		return nil, fmt.Errorf("Timeout de 10ms excedido ao inserir no banco de dados: %v", err)
	}
	return exchange, nil
}
