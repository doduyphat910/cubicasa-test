package repository

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/aggregate"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
)

//go:generate mockgen -destination=./mock/mock_$GOFILE -source=$GOFILE -package=mock

type UserRepository interface {
	Create(ctx context.Context, user entity.User) (entity.User, error)
	GetByID(ctx context.Context, id uint64) (aggregate.UserAggregate, error)
}
