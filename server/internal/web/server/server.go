package server

import (
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/silmarsanches/clientserverapi/server/internal/web/middlewares"
)

func NewServer(exchangeRoutes http.Handler) *http.Server {
	r := chi.NewRouter()

	r.Use(middleware.Recoverer)
	r.Use(middlewares.MiddlewareLog)

	r.Mount("/cotacao", exchangeRoutes)

	return &http.Server{
		Addr:    ":8080",
		Handler: r,
	}
}
