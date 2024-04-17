package controller

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/lyracampos/strava-analytics-api/entities"
)

func Activities(w http.ResponseWriter, r *http.Request) {
	activities := entities.List()
	err := json.NewEncoder(w).Encode(activities)
	if err != nil {
		log.Printf("erro ao retornar dados")
	}
}
