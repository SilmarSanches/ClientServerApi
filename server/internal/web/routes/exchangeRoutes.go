package routes

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/silmarsanches/clientserverapi/server/internal/web/controllers"
)

func ExchangeRoutes(controller *controllers.ExchangeController) http.Handler {
	r := chi.NewRouter()
	r.Get("/", controller.InsertExchange)
	return r
}
