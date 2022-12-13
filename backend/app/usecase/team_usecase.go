package usecase

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/repository"
	"github.com/doduyphat910/cubicasa-test/backend/app/usecase/dto"
)

type TeamUseCaser interface {
	Create(ctx context.Context, team entity.Team) (entity.Team, error)
	Search(ctx context.Context, dtoReq dto.SearchTeamRequest) ([]entity.Team, error)
}

type TeamUseCase struct {
	hubRepo  repository.HubRepository
	teamRepo repository.TeamRepository
}

func NewTeamUseCase(teamRepo repository.TeamRepository, hubRepo repository.HubRepository) *TeamUseCase {
	return &TeamUseCase{
		hubRepo:  hubRepo,
		teamRepo: teamRepo,
	}
}

func (uc *TeamUseCase) Create(ctx context.Context, team entity.Team) (entity.Team, error) {
	_, err := uc.hubRepo.GetByID(ctx, team.HubID)
	if err != nil {
		return entity.Team{}, err
	}
	teamEnt, err := uc.teamRepo.Create(ctx, team)
	return teamEnt, err
}

func (uc *TeamUseCase) Search(ctx context.Context, dtoReq dto.SearchTeamRequest) ([]entity.Team, error) {
	teams, err := uc.teamRepo.Search(ctx, entity.GeoLocation{Lat: dtoReq.Lat, Long: dtoReq.Long}, dtoReq.Paging)
	return teams, err
}
