package services

import (
	"encoding/json"
	"errors"
	"github.com/silmarsanches/clientserverapi/client/config"
	"io"
	"log"
	"net/http"
)

type ServiceDolarInterface interface {
	GetDolarExchangeRate() (map[string]interface{}, error)
}

type HttpExternalServiceDolar struct {
	appConfig  config.Config
	HttpClient *http.Client
}

func NewHttpExternalServiceDolar(appConfig *config.Config) *HttpExternalServiceDolar {
	return &HttpExternalServiceDolar{
		appConfig:  *appConfig,
		HttpClient: &http.Client{},
	}
}

func (s *HttpExternalServiceDolar) GetDolarExchangeRate() (map[string]interface{}, error) {
	req, err := http.NewRequest(http.MethodGet, s.appConfig.UrlApi, nil)
	if err != nil {
		return nil, err
	}

	res, err := s.HttpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Erro ao fechar o corpo da resposta: %v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("Erro ao obter cotação do dolar: " + res.Status)
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
