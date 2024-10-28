package main

import "net/http"

func CreateRoutes() *http.ServeMux {
	router := http.NewServeMux()
	router.HandleFunc("/hello/{name}", helloHandler)
	return router
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Write([]byte("Halo " + name))
}
