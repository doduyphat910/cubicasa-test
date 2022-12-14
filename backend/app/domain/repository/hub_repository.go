package repository

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
)

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

type HubRepository interface {
	Create(ctx context.Context, hub entity.Hub) (entity.Hub, error)
	GetByID(ctx context.Context, id uint64) (entity.Hub, error)
}
