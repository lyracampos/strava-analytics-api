package infrastructure

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"github.com/lyracampos/strava-analytics-api/internal/domain"
	"github.com/lyracampos/strava-analytics-api/internal/domain/contracts"
)

const listActivitiesURL = "https://www.strava.com/api/v3/athlete/activities?after=%d&before=%d"

var _ contracts.StravaGateway = (*stravaHTTP)(nil)

type stravaHTTP struct{}

func NewStravaHTTP() *stravaHTTP {
	return &stravaHTTP{}
}

func (g *stravaHTTP) ListActivities(ctx context.Context, input contracts.ListInput) (contracts.ListOutput, error) {
	stravaUrl := fmt.Sprintf(
		listActivitiesURL,
		input.After.Unix(),
		input.Before.Unix(),
	)
	req, err := http.NewRequest(http.MethodGet, stravaUrl, nil)
	if err != nil {
		return contracts.ListOutput{}, fmt.Errorf("erro ao tentar criar Request: %w", err)
	}
	token := ctx.Value("token")
	if token == nil {
		token = "f93b98c6fd851414f9315f83a1f5550e0f598670"
	}
	bearerToken := fmt.Sprintf("Bearer %s", token)
	req.Header.Add(
		"Authorization",
		bearerToken,
	)

	client := &http.Client{}
	res, err := client.Do(req)
	if err != nil {
		return contracts.ListOutput{}, fmt.Errorf("erro ao fazer o Request: %w", err)
	}

	if res.StatusCode == http.StatusUnauthorized {
		return contracts.ListOutput{}, domain.ErrUserNotAuthorized
	}

	if res.StatusCode != http.StatusOK {
		return contracts.ListOutput{}, fmt.Errorf("erro ao fazer o Request: %w", err)
	}

	resJson, err := io.ReadAll(res.Body)
	if err != nil {
		return contracts.ListOutput{}, fmt.Errorf("erro ao carregar o response: %w", err)
	}

	var stravaActivitiesList []*StravaActivity
	err = json.Unmarshal(resJson, &stravaActivitiesList)
	if err != nil {
		return contracts.ListOutput{}, fmt.Errorf("erro ao interpretar o response: %w", err)
	}
	activitiesList := activitiesConverter(stravaActivitiesList)
	return contracts.ListOutput{
		Activities: activitiesList,
		Total:      len(activitiesList),
	}, nil
}
