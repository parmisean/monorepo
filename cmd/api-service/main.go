// API Service
package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type application struct {
	Config config
}

type config struct {
	Port int
}

func main() {
	var cfg config

	flag.IntVar(&cfg.Port, "port", 4000, "API Service port")
	flag.Parse()

	app := &application{
		Config: cfg,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Starting API Service on port %d", cfg.Port)
	log.Fatal(srv.ListenAndServe())
}
