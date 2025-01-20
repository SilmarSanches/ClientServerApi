//go:build wireinject
// +build wireinject

package main

import (
	"github.com/google/wire"
	"github.com/silmarsanches/clientserverapi/client/config"
	"github.com/silmarsanches/clientserverapi/client/internal/infra/services"
	"github.com/silmarsanches/clientserverapi/client/internal/usecase"
)

var setUsecaseGetDolarExchangeRate = wire.NewSet(
	services.NewHttpExternalServiceDolar,
	wire.Bind(new(services.ServiceDolarInterface), new(*services.HttpExternalServiceDolar)),
)

func InitializeGetDolarExchangeRate(appConfig *config.Config) *usecase.GetDolarExchangeRateUseCase {
	wire.Build(
		setUsecaseGetDolarExchangeRate,
		usecase.NewGetDolarExchangeRateUseCase,
	)
	return &usecase.GetDolarExchangeRateUseCase{}
}
