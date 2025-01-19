// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/silmarsanches/clientserverapi/server/config"
	"github.com/silmarsanches/clientserverapi/server/internal/infra/db"
	"github.com/silmarsanches/clientserverapi/server/internal/infra/services"
	"github.com/silmarsanches/clientserverapi/server/internal/usecase"
)

import (
	_ "github.com/mattn/go-sqlite3"
)

// Injectors from wire.go:

func InitializeExchangeInsertUseCase(queries *db.Queries, appConfig *config.Config) *usecase.ExchangeInsertUseCase {
	httpExternalServiceExchange := services.NewHttpExternalServiceExchange(appConfig)
	exchangeInsertUseCase := usecase.NewExchangeInsertUseCase(queries, appConfig, httpExternalServiceExchange)
	return exchangeInsertUseCase
}

// wire.go:

var setUseCaseExchangeDependency = wire.NewSet(services.NewHttpExternalServiceExchange, wire.Bind(new(services.ExternalServiceExchangeInterface), new(*services.HttpExternalServiceExchange)))
