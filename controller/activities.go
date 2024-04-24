package controller

import (
	"encoding/json"
	"log"
	"net/http"
	"time"

	"github.com/lyracampos/strava-analytics-api/entities"
)

func Activities(w http.ResponseWriter, r *http.Request) {
	var after time.Time
	var err error
	if r.URL.Query().Get("start") != "" {
		after, err = time.Parse("2006-01-02", r.URL.Query().Get("start"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte("data inicial inválida"))
			if err != nil {
				log.Printf("erro ao mostrar o erro: %v", err)
			}
			return
		}
	}
	var before time.Time
	if r.URL.Query().Get("end") != "" {
		before, err = time.Parse("2006-01-02", r.URL.Query().Get("end"))
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte("data final inválida"))
			if err != nil {
				log.Printf("erro ao mostrar o erro: %v", err)
			}
			return
		}
	}
	listParams := entities.ListParams{
		Page:    1,
		PerPage: 100,
		After:   after,
		Before:  before,
	}
	activities, err := entities.List(listParams)
	if err != nil {
		log.Printf("erro ao listar atividades")
	}
	err = json.NewEncoder(w).Encode(activities)
	if err != nil {
		log.Printf("erro ao retornar dados")
	}

}
