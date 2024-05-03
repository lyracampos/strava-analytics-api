package entities

import (
	"context"

	"github.com/lyracampos/strava-analytics-api/infra"
)

// recuperar o bearer token do context - ok
// utiliza-lo para realizar o request na api - ok
func List(ctx context.Context, listParams infra.ListParams) ([]*infra.Activity, error) {
	return infra.ListActivities(ctx, listParams)
}
