package contracts

import (
	"context"
	"time"

	"github.com/lyracampos/strava-analytics-api/internal/domain/entities"
)

type ListInput struct {
	Before  time.Time
	After   time.Time
	Page    int
	PerPage int
}

type ListOutput struct {
	Activities []*entities.Activity
	Total      int
}

type StravaGateway interface {
	ListActivities(ctx context.Context, input ListInput) (ListOutput, error)
}
