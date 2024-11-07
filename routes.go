package main

import (
	"net/http"

	"github.com/endymuhardin/belajar-go-ewallet/handler"
)

func createRoutes(app *App) *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /hello/{name}", handler.HelloHandler(app.service))
	return router
}
