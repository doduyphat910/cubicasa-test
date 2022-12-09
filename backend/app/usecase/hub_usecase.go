package usecase

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/repository"
)

type HubUseCaser interface {
	Create(ctx context.Context, hub entity.Hub) (entity.Hub, error)
}

type HubUseCase struct {
	hubRepo repository.HubRepository
}

func NewHubUseCase(hubRepo repository.HubRepository) HubUseCase {
	return HubUseCase{
		hubRepo: hubRepo,
	}
}

func (uc HubUseCase) Create(ctx context.Context, hub entity.Hub) (entity.Hub, error) {
	return entity.Hub{}, nil
}
