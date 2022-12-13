package pgsql

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/utils"
	"gorm.io/gorm"
)

type HubRepository struct {
	db *gorm.DB
}

func NewHubRepository() *HubRepository {
	return &HubRepository{db: utils.GetDB()}
}

func (repo *HubRepository) Create(ctx context.Context, hub entity.Hub) (entity.Hub, error) {
	err := repo.db.Create(&hub).Error
	return hub, err
}

func (repo *HubRepository) GetByID(ctx context.Context, id uint64) (entity.Hub, error) {
	var (
		hub entity.Hub
		err error
	)
	err = repo.db.Take(&hub, id).Error
	return hub, err
}
