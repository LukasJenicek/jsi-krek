package api

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

func Router() chi.Router {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
		w.Write([]byte{})
	})

	return r
}
