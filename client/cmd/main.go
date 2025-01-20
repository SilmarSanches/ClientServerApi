package main

import (
	"github.com/silmarsanches/clientserverapi/client/config"
	"log"
)

func main() {
	appConfig, err := config.LoadConfig("./client/cmd")
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo de configuração: %v", err)
	}

	useCase := InitializeGetDolarExchangeRate(appConfig)
	result, err := useCase.GetDolarExchangeRate()
	if err != nil {
		log.Fatalf("Erro ao obter a cotação do dólar: %v", err)
	}
	log.Printf("Cotação do dólar: %v", result)
}
