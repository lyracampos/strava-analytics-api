package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/healthz", HealthZ).Methods(http.MethodGet)

	srv := &http.Server{
		Addr:    "0.0.0.0:8080",
		Handler: router,
	}
	if err := srv.ListenAndServe(); err != nil {
		log.Println(err)
	}
}
