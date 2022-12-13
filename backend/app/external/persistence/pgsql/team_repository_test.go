package pgsql

import (
	"context"
	"fmt"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/utils"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"regexp"
	"testing"
)

func TestNewTeamRepository(t *testing.T) {
	want := &TeamRepository{}
	repo := NewTeamRepository()

	if diff := cmp.Diff(repo, want,
		cmpopts.IgnoreFields(TeamRepository{}, "db")); diff != "" {
		t.Error(diff)
	}
}

func TestTeamRepository_Create(t *testing.T) {
	db, mock, err := utils.OpenTestDBConnection()
	if err != nil {
		return
	}

	repo := &TeamRepository{
		db: db,
	}

	teamEnt := entity.Team{HubID: 1, GeoLocation: entity.GeoLocation{Long: 123, Lat: 456}}
	getLocationArg := fmt.Sprintf("(%v, %v)", teamEnt.GeoLocation.Long, teamEnt.GeoLocation.Lat)
	t.Run("create success", func(t *testing.T) {
		query := "INSERT INTO \"teams\" (\"geo_location\",\"hub_id\") VALUES ($1,$2) RETURNING \"id\""
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(getLocationArg, teamEnt.HubID).
			WillReturnRows(
				sqlmock.NewRows([]string{"id"}).AddRow(teamEnt.ID),
			)
		mock.ExpectCommit()

		got, err := repo.Create(context.Background(), teamEnt)
		if err != nil {
			t.Error(err)
			return
		}
		diff := cmp.Diff(got, teamEnt, cmpopts.IgnoreFields(entity.Team{}, "CreatedAt", "UpdatedAt"))
		if diff != "" {
			t.Error(diff)
		}
	})
}

func TestTeamRepository_Search(t *testing.T) {
	db, mock, err := utils.OpenTestDBConnection()
	if err != nil {
		return
	}

	repo := &TeamRepository{
		db: db,
	}

	paging := utils.Paging{Size: 2, Number: 1}
	geoLocation := entity.GeoLocation{Long: 123, Lat: 456}
	teams := []entity.Team{{ID: 1, HubID: 1, Hub: entity.Hub{ID: 1}}}
	hub := entity.Hub{ID: 1}

	t.Run("search", func(t *testing.T) {
		getTeamsQuery := fmt.Sprintf(
			"SELECT * FROM \"teams\" ORDER BY geo_location <-> point(%v, %v) LIMIT %v",
			geoLocation.Long, geoLocation.Lat, paging.Size,
		)
		mock.ExpectQuery(regexp.QuoteMeta(getTeamsQuery)).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "hub_id"}).AddRow(1, 1),
			)
		mock.ExpectQuery(regexp.QuoteMeta("SELECT * FROM \"hubs\" WHERE \"hubs\".\"id\" = $1")).
			WillReturnRows(
				sqlmock.NewRows([]string{"id"}).AddRow(hub.ID),
			)

		got, err := repo.Search(context.Background(), geoLocation, paging)
		if err != nil {
			t.Error(err)
			return
		}
		diff := cmp.Diff(got, teams)
		if diff != "" {
			t.Error(diff)
		}
	})
}

func TestTeamRepository_GetByID(t *testing.T) {
	db, mock, err := utils.OpenTestDBConnection()
	if err != nil {
		return
	}

	repo := &TeamRepository{
		db: db,
	}

	id := uint64(1)
	teamEnt := entity.Team{ID: 1, HubID: 1}
	t.Run("get by id", func(t *testing.T) {
		query := "SELECT * FROM \"teams\" WHERE \"teams\".\"id\" = $1 LIMIT 1"
		mock.ExpectQuery(regexp.QuoteMeta(query)).
			WithArgs(id).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "hub_id"}).AddRow(1, 1),
			)

		got, err := repo.GetByID(context.Background(), id)
		if err != nil {
			t.Error(err)
			return
		}
		diff := cmp.Diff(got, teamEnt)
		if diff != "" {
			t.Error(diff)
		}
	})
}
