package repository

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
)

type TeamRepository interface {
	Create(ctx context.Context, team entity.Team) (entity.Team, error)
	Search(ctx context.Context) (entity.Team, error)
}
