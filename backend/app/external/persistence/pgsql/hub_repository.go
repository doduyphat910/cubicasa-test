package pgsql

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"gorm.io/gorm"
)

type HubRepository struct {
	db gorm.DB
}

func NewHubRepository() HubRepository {
	db := gorm.DB{}
	return HubRepository{db: db}
}

func (repo HubRepository) Create(ctx context.Context, hub entity.Hub) (entity.Hub, error) {
	return entity.Hub{}, nil
}
