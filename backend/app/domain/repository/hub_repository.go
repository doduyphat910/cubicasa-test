package repository

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
)

type HubRepository interface {
	Create(ctx context.Context, hub entity.Hub) (entity.Hub, error)
}
