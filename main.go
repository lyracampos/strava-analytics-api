package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	"github.com/gorilla/mux"
	"github.com/lyracampos/strava-analytics-api/config"
)

const defaultConfigFilePath = "./config/config.yaml"

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	var configFilePath string

	flag.StringVar(&configFilePath, "c", defaultConfigFilePath, "File path with app configs file.")
	flag.Parse()

	config, err := config.NewConfig(configFilePath)
	if err != nil {
		return fmt.Errorf("failed to build config: %w", err)
	}

	router := mux.NewRouter()
	router.HandleFunc("/healthz", HealthZ).Methods(http.MethodGet)

	address := fmt.Sprintf("%s:%d", config.API.Host, config.API.Port)

	server := &http.Server{
		Addr:         address,
		Handler:      router,
		WriteTimeout: time.Second * time.Duration(config.API.WriteTimeout),
		ReadTimeout:  time.Second * time.Duration(config.API.ReadTimeout),
		IdleTimeout:  time.Second * time.Duration(config.API.IdleTimeout),
	}

	go func() {
		log.Printf("running API HTTP server at: %s", address)

		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// We'll accept graceful shutdowns when quit via SIGINT (Ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (Ctrl+/) will not be caught.
	signal.Notify(c, os.Interrupt)

	// Block until a signal is received.
	sig := <-c
	log.Println("Got signal:", sig)

	// gracefully shutdown the server, waiting max 30 seconds for current operations to complete
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// Doesn't block if no connections, but will otherwise wait
	// until the timeout deadline.
	if err := server.Shutdown(ctx); err != nil {
		log.Println("Error shutting down server: %w", err)
	}

	log.Println("shutting down")
	return nil
}
