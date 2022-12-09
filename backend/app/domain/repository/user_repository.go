package repository

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
)

type UserRepository interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	GetByID(ctx context.Context, id uint64) (entity.User, error)
}
