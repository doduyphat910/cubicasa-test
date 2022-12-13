package pgsql

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/aggregate"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/utils"
	"gorm.io/gorm"
)

type UserRepository struct {
	db *gorm.DB
}

func NewUserRepository() *UserRepository {
	return &UserRepository{db: utils.GetDB()}
}

func (repo *UserRepository) Create(ctx context.Context, user entity.User) (entity.User, error) {
	err := repo.db.Create(&user).Error
	return user, err
}

func (repo *UserRepository) GetByID(ctx context.Context, id uint64) (aggregate.UserAggregate, error) {
	var (
		userAggregate aggregate.UserAggregate
		err           error
	)

	err = repo.db.Model(&entity.User{}).
		Select("users.id, users.team_id, users.type, users.created_at, users.updated_at, " +
			"teams.id, teams.hub_id, teams.geo_location, teams.created_at, teams.updated_at, " +
			"hubs.id, hubs.name, hubs.created_at, hubs.updated_at").
		Joins("join teams on team_id = users.team_id").
		Joins("join hubs on teams.hub_id = hubs.id").
		Where(&entity.User{ID: id}).
		Scan(&userAggregate).Error

	return userAggregate, err
}
