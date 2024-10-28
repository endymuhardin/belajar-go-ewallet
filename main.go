package main

import (
	"log"
	"net/http"
)

func main() {
	

	serverPort := ":8080"

	server := http.Server{
		Addr:    serverPort,
		Handler: CreateRoutes(),
	}

	log.Println("Starting server at " + serverPort)
	server.ListenAndServe()
}


