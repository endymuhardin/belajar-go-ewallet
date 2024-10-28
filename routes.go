package main

import (
	"encoding/json"
	"net/http"
)

func (app *App) CreateRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("GET /hello/{name}", app.helloHandler)
	return router
}

func (app *App) helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	cust := app.store.GetCustomer(name)
	json.NewEncoder(w).Encode(cust)
}
