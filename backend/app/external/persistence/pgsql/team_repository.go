package pgsql

import (
	"context"
	"fmt"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/utils"
	"gorm.io/gorm"
)

type TeamRepository struct {
	db *gorm.DB
}

func NewTeamRepository() *TeamRepository {
	return &TeamRepository{db: utils.GetDB()}
}

func (repo *TeamRepository) Create(ctx context.Context, team entity.Team) (entity.Team, error) {
	err := repo.db.Model(&team).Create(map[string]interface{}{
		"HubID":       team.HubID,
		"GeoLocation": fmt.Sprintf("(%v, %v)", team.GeoLocation.Long, team.GeoLocation.Lat),
	}).Error

	return team, err
}

func (repo *TeamRepository) Search(ctx context.Context, geoLocation entity.GeoLocation, paging utils.Paging) ([]entity.Team, error) {
	var teams []entity.Team
	err := repo.db.Order(fmt.Sprintf("geo_location <-> point(%v, %v)", geoLocation.Long, geoLocation.Lat)).
		Scopes(utils.Paginate(paging)).
		Preload("Hub").
		Find(&teams).Error
	return teams, err
}

func (repo *TeamRepository) GetByID(ctx context.Context, id uint64) (entity.Team, error) {
	var (
		team entity.Team
		err  error
	)
	err = repo.db.Take(&team, id).Error
	return team, err
}
