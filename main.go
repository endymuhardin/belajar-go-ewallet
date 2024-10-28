package main

import (
	"log"
	"net/http"

	"github.com/endymuhardin/belajar-go-ewallet/storage"
)

type App struct {
	store storage.Storage
}

func NewApp(store storage.Storage) *App {
	return &App{
		store: store,
	}
}

func main() {

	serverPort := ":8080"

	pgstore := storage.NewPostgresStorage(
		"localhost",
		"54321",
		"ewallet",
		"ewallet123",
		"ewalletdb",
	)

	app := NewApp(pgstore)

	server := http.Server{
		Addr:    serverPort,
		Handler: app.CreateRoutes(),
	}

	log.Println("Starting server at " + serverPort)
	server.ListenAndServe()
}
