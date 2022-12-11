package repository

import (
	"context"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/utils"
)

type TeamRepository interface {
	Create(ctx context.Context, team entity.Team) (entity.Team, error)
	Search(ctx context.Context, geoLocation entity.GeoLocation, paging utils.Paging) ([]entity.Team, error)
	GetByID(ctx context.Context, id uint64) (entity.Team, error)
}
