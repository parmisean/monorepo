package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const healthResponse = "Hello"

func (a *Application) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(healthResponse))
	})

	return r
}
