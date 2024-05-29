package usecases

import (
	"context"
	"fmt"

	"github.com/lyracampos/strava-analytics-api/internal/domain/contracts"
)

type ListActiviesUseCase struct {
	stravaGateway contracts.StravaGateway
}

func NewListActiviesUseCase(stravaGateway contracts.StravaGateway) *ListActiviesUseCase {
	return &ListActiviesUseCase{
		stravaGateway: stravaGateway,
	}
}

func (u *ListActiviesUseCase) Execute(ctx context.Context, input contracts.ListInput) (contracts.ListOutput, error) {
	output, err := u.stravaGateway.ListActivities(ctx, input)
	if err != nil {
		return contracts.ListOutput{}, fmt.Errorf("falhou ao listar atividades na API do strava: %w", err)
	}

	return output, nil
}
