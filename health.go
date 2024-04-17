package main

import (
	"log"
	"net/http"
)

func HealthZ(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("api v3 is running"))
	if err != nil {
		log.Printf("CheckStatus - write failed: %v", err)
	}
}
