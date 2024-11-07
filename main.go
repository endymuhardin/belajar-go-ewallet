package main

import (
	"log"
	"net/http"

	"github.com/endymuhardin/belajar-go-ewallet/storage"
	"github.com/endymuhardin/belajar-go-ewallet/types"
)

type App struct {
	service types.WalletService
}

func NewApp(svc types.WalletService) *App {
	return &App{
		service: svc,
	}
}

func main() {

	serverPort := ":8080"

	pgstore := storage.NewPostgresStorage(
		"localhost",
		"5432",
		"ewallet",
		"ewallet123",
		"ewallet-db",
	)

	app := NewApp(pgstore)

	server := http.Server{
		Addr:    serverPort,
		Handler: createRoutes(app),
	}

	log.Println("Starting server at " + serverPort)
	server.ListenAndServe()
}
