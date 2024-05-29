package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"time"

	"github.com/lyracampos/strava-analytics-api/internal/domain"
	"github.com/lyracampos/strava-analytics-api/internal/domain/contracts"
	"github.com/lyracampos/strava-analytics-api/internal/usecases"
	"go.uber.org/zap"
)

type ActivitiesHandler struct {
	log                   *zap.SugaredLogger
	listActivitiesUsecase usecases.ListActiviesUseCase
}

func NewActivitiesHandler(log *zap.SugaredLogger, listActivitiesUsecase usecases.ListActiviesUseCase) *ActivitiesHandler {
	return &ActivitiesHandler{
		log:                   log,
		listActivitiesUsecase: listActivitiesUsecase,
	}
}

func (h *ActivitiesHandler) ListActivities(w http.ResponseWriter, r *http.Request) {
	h.log.Info("ListActivities - iniciou")

	input, err := readInput(r)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(err.Error()))
		if err != nil {
			h.log.Errorf("ListActivities - erro ao exibir o response erro: %v", err)
		}

		return
	}

	activities, err := h.listActivitiesUsecase.Execute(r.Context(), input)
	if err != nil {
		h.log.Errorf("ListActivities - erro ao executar o use case: %v", err)
		msgErr, statusCode := handlerError(err)
		w.WriteHeader(statusCode)
		_, err = w.Write([]byte(msgErr))
		if err != nil {
			h.log.Errorf("ListActivities - erro ao exibir o response erro: %v", err)
		}

		return
	}

	h.log.Info("ListActivities - finalizou com sucesso")

	err = json.NewEncoder(w).Encode(activities)
	if err != nil {
		h.log.Errorf("ListActivities - erro ao retornar dados: %v", err)
	}
}

func readInput(r *http.Request) (contracts.ListInput, error) {
	var after time.Time
	var err error
	if r.URL.Query().Get("start") != "" {
		after, err = time.Parse("2006-01-02", r.URL.Query().Get("start"))
		if err != nil {
			return contracts.ListInput{}, fmt.Errorf("data inicial inválida: %w", err)
		}
	}
	var before time.Time
	if r.URL.Query().Get("end") != "" {
		before, err = time.Parse("2006-01-02", r.URL.Query().Get("end"))
		if err != nil {
			return contracts.ListInput{}, fmt.Errorf("data final inválida: %w", err)
		}
	}
	return contracts.ListInput{
		Page:    1,
		PerPage: 100,
		After:   after,
		Before:  before,
	}, nil
}

func handlerError(err error) (string, int) {
	switch {
	case errors.Is(err, domain.ErrUserNotAuthorized):
		return err.Error(), http.StatusUnauthorized
	default:
		return err.Error(), http.StatusInternalServerError
	}
}
