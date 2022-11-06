package main

import (
	"net/http"
	"log"
)

type Server struct{}

func (*Server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	w.Write([]byte{})
}

func main() {
	handler := Server{}
	addr := ":9099"

	server := http.Server{
		Addr:    addr,
		Handler: &handler,
	}

	log.Printf("serving on port %s", addr)

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
