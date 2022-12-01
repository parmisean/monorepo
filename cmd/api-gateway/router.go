package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const healthResponse = "available"

func (a *application) routes() *chi.Mux {
	r := chi.NewRouter()

	r.Get("/health", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		_, err := w.Write([]byte(healthResponse))
		if err != nil {
			log.Printf("Error: %v", err)
		}
	})

	return r
}
