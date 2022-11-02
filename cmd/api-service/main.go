package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"
)

type Application struct {
	Config Config
}

type Config struct {
	Port int
}

func main() {
	var cfg Config

	flag.IntVar(&cfg.Port, "port", 4000, "API Service port")
	flag.Parse()

	app := &Application{
		Config: cfg,
	}

	srv := &http.Server{
		Addr:         fmt.Sprintf(":%d", cfg.Port),
		Handler:      app.Routes(),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  15 * time.Second,
	}

	log.Printf("Starting API Service on port %d", cfg.Port)
	log.Fatal(srv.ListenAndServe())
}
