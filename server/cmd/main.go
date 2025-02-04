package main

import (
	"database/sql"
	"log"
	"net/http"
	"os"

	_ "github.com/mattn/go-sqlite3"
	"github.com/silmarsanches/clientserverapi/server/config"
	"github.com/silmarsanches/clientserverapi/server/internal/infra/db"
	"github.com/silmarsanches/clientserverapi/server/internal/web/controllers"
	"github.com/silmarsanches/clientserverapi/server/internal/web/routes"
	"github.com/silmarsanches/clientserverapi/server/internal/web/server"
)

func main() {
	workingDir, err := os.Getwd()
    if err != nil {
        log.Fatalf("Erro ao obter o diretório de trabalho atual: %v", err)
    }
    log.Printf("Diretório de trabalho atual: %s", workingDir)

	appConfig, err := config.LoadConfig(workingDir)
	if err != nil {
		log.Fatalf("Erro ao carregar o arquivo de configuração: %v", err)
	}

	dbConnection, err := sql.Open("sqlite3", appConfig.Database)
	if err != nil {
		log.Fatalf("Erro ao conectar ao banco de dados: %v", err)
	}
	defer func(dbConnection *sql.DB) {
		err := dbConnection.Close()
		if err != nil {
			log.Printf("Erro ao fechar a conexão com o banco de dados: %v", err)
		}
	}(dbConnection)

	queries := db.New(dbConnection)

	usecase := InitializeExchangeInsertUseCase(queries, appConfig)
	controller := controllers.NewExchangeController(usecase)
	exchangeRoutes := routes.ExchangeRoutes(controller)

	srv := server.NewServer(exchangeRoutes)

	log.Println("Servidor iniciado na porta 8080")
	if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("Erro ao iniciar o servidor: %v", err)
	}
}
