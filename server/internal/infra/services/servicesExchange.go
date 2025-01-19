package services

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/silmarsanches/clientserverapi/server/config"
)

type ExternalServiceExchangeInterface interface {
	GetExchangeRate(ctx context.Context) (map[string]interface{}, error)
}

type HttpExternalServiceExchange struct {
	BaseUrl    string
	HttpClient *http.Client
	appConfig  config.Config
}

func NewHttpExternalServiceExchange(appConfig *config.Config) *HttpExternalServiceExchange {
	return &HttpExternalServiceExchange{
		BaseUrl: appConfig.URLDolar,
		HttpClient: &http.Client{
			Timeout: 5 * time.Second,
		},
		appConfig: *appConfig,
	}
}

func (h *HttpExternalServiceExchange) GetExchangeRate(ctx context.Context) (map[string]interface{}, error) {
	ctx, cancel := context.WithTimeout(ctx, 200*time.Millisecond)
	defer cancel()

	req, err := http.NewRequestWithContext(ctx, http.MethodGet, h.BaseUrl, nil)
	if err != nil {
		return nil, err
	}

	res, err := h.HttpClient.Do(req)
	if err != nil {
		if errors.Is(err, context.DeadlineExceeded) {
			log.Printf("Timeout de 200ms excedido ao consultar o serviço externo")
		}
		return nil, fmt.Errorf("Timeout de 200ms excedido ao consultar o serviço externo: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			log.Printf("Erro ao fechar o corpo da resposta: %v", err)
		}
	}(res.Body)

	if res.StatusCode != http.StatusOK {
		return nil, errors.New("erro ao consultar o serviço externo: " + res.Status)
	}

	var data map[string]interface{}
	err = json.NewDecoder(res.Body).Decode(&data)
	if err != nil {
		return nil, err
	}

	return data, nil
}
