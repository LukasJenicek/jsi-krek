package main

import (
	"net/http"
	"log"
	"github.com/LukasJenicek/jsi-krek/services/api"
)

func main() {
	addr := ":9099"

	server := http.Server{
		Addr:    addr,
		Handler: api.Router(),
	}

	log.Printf("serving on port %s", addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
