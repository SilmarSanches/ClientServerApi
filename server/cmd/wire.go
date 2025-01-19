//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/silmarsanches/clientserverapi/server/config"
	"github.com/silmarsanches/clientserverapi/server/internal/infra/db"
	"github.com/silmarsanches/clientserverapi/server/internal/infra/services"
	"github.com/silmarsanches/clientserverapi/server/internal/usecase"
)

var setUseCaseExchangeDependency = wire.NewSet(
	services.NewHttpExternalServiceExchange,
	wire.Bind(new(services.ExternalServiceExchangeInterface), new(*services.HttpExternalServiceExchange)),
)

func InitializeExchangeInsertUseCase(queries *db.Queries, appConfig *config.Config) *usecase.ExchangeInsertUseCase {
	wire.Build(
		setUseCaseExchangeDependency,
		usecase.NewExchangeInsertUseCase,
	)
	return &usecase.ExchangeInsertUseCase{}
}
