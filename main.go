package main

import (
	"log"
	"net/http"
)

func main() {
	router := http.NewServeMux()
	router.HandleFunc("/hello/{name}", helloHandler)

	serverPort := ":8080"

	server := http.Server{
		Addr:    serverPort,
		Handler: router,
	}

	log.Println("Starting server at " + serverPort)
	server.ListenAndServe()
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PathValue("name")
	w.Write([]byte("Halo " + name))
}
