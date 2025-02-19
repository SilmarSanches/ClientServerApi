// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"github.com/google/wire"
	"github.com/silmarsanches/clientserverapi/client/config"
	"github.com/silmarsanches/clientserverapi/client/internal/infra/services"
	"github.com/silmarsanches/clientserverapi/client/internal/usecase"
)

// Injectors from wire.go:

func InitializeGetDolarExchangeRate(appConfig *config.Config) *usecase.GetDolarExchangeRateUseCase {
	httpExternalServiceDolar := services.NewHttpExternalServiceDolar(appConfig)
	getDolarExchangeRateUseCase := usecase.NewGetDolarExchangeRateUseCase(appConfig, httpExternalServiceDolar)
	return getDolarExchangeRateUseCase
}

// wire.go:

var setUsecaseGetDolarExchangeRate = wire.NewSet(services.NewHttpExternalServiceDolar, wire.Bind(new(services.ServiceDolarInterface), new(*services.HttpExternalServiceDolar)))
