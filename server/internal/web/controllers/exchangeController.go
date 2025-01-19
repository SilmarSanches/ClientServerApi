package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/silmarsanches/clientserverapi/server/internal/usecase"
	"github.com/silmarsanches/clientserverapi/server/internal/web/dtos/output"
)

type ExchangeController struct {
	Usecase *usecase.ExchangeInsertUseCase
}

func NewExchangeController(usecase *usecase.ExchangeInsertUseCase) *ExchangeController {
	return &ExchangeController{
		Usecase: usecase,
	}
}

func (c *ExchangeController) InsertExchange(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	exchange, err := c.Usecase.InsertExchange(ctx)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	exchangeResponse := output.ExchangeResponseDto{
		Bid: exchange.Bid,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err = json.NewEncoder(w).Encode(exchangeResponse)
	if err != nil {
		return
	}
}
