package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

const HealthResponse = "Hello"

func (a *Application) Routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(HealthResponse))
	})

	return r
}
