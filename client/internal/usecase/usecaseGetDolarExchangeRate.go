package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/silmarsanches/clientserverapi/client/config"
	"github.com/silmarsanches/clientserverapi/client/internal/infra/services"
	"log"
	"os"
)

type GetDolarExchangeRateUseCase struct {
	External  services.ServiceDolarInterface
	appConfig config.Config
}

type DolarExchangeRate struct {
	Bid float64 `json:"bid"`
}

func NewGetDolarExchangeRateUseCase(appConfig *config.Config, external services.ServiceDolarInterface) *GetDolarExchangeRateUseCase {
	return &GetDolarExchangeRateUseCase{
		External:  external,
		appConfig: *appConfig,
	}
}

func (d *GetDolarExchangeRateUseCase) GetDolarExchangeRate() (float64, error) {
	data, err := d.External.GetDolarExchangeRate()
	if err != nil {
		return 0, err
	}

	var exchangeRate DolarExchangeRate
	jsonData, err := json.Marshal(data)
	if err != nil {
		return 0, err
	}

	err = json.Unmarshal(jsonData, &exchangeRate)
	if err != nil {
		return 0, err
	}

	err = appendToFile(exchangeRate.Bid)
	if err != nil {
		return 0, err
	}

	return exchangeRate.Bid, nil
}

func appendToFile(value float64) error {
	file, err := os.OpenFile("exchange_rate.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return err
	}
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			log.Printf("Erro ao fechar o arquivo: %v", err)
		}
	}(file)

	_, err = file.WriteString(fmt.Sprintf("Dólar: %.2f\n", value))
	if err != nil {
		return err
	}

	return nil
}
