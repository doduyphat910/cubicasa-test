package pgsql

import (
	"context"
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
	if err != nil {
		return entity.User{}, err
	}
	return user, nil
}

func (repo *UserRepository) GetByID(ctx context.Context, id uint64) (entity.User, error) {
	var (
		user entity.User
		err  error
	)
	err = repo.db.Joins("Team").Preload("Team.Hub").
		Where(&entity.User{ID: id}).
		Find(&user).Error

	return user, err
}
