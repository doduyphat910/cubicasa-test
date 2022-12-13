package usecase

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/aggregate"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/repository"
)

type UserUseCaser interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	GetByID(ctx context.Context, id uint64) (aggregate.UserAggregate, error)
}

type UserUseCase struct {
	teamRepo repository.TeamRepository
	userRepo repository.UserRepository
}

func NewUserUseCase(teamRepo repository.TeamRepository, userRepo repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		teamRepo: teamRepo,
		userRepo: userRepo,
	}
}

func (uc *UserUseCase) Create(ctx context.Context, user entity.User) (entity.User, error) {
	_, err := uc.teamRepo.GetByID(ctx, user.TeamID)
	if err != nil {
		return entity.User{}, err
	}
	userEnt, err := uc.userRepo.Create(ctx, user)
	return userEnt, err
}

func (uc *UserUseCase) GetByID(ctx context.Context, id uint64) (aggregate.UserAggregate, error) {
	user, err := uc.userRepo.GetByID(ctx, id)
	return user, err
}
