package pgsql

import (
	"context"
	"errors"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/aggregate"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/utils"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"gorm.io/gorm"
	"regexp"
	"testing"
)

func TestNewUserRepository(t *testing.T) {
	want := &UserRepository{}
	repo := NewUserRepository()

	if diff := cmp.Diff(repo, want,
		cmpopts.IgnoreFields(UserRepository{}, "db")); diff != "" {
		t.Error(diff)
	}
}

func TestNewUserRepository_Create(t *testing.T) {
	db, mock, err := utils.OpenTestDBConnection()
	if err != nil {
		return
	}

	repo := &UserRepository{
		db: db,
	}
	teamEnt := entity.Team{ID: 1}
	createErr := errors.New("create error")

	t.Run("create success", func(t *testing.T) {
		userEnt := entity.User{TeamID: teamEnt.ID, Type: "test"}
		insertUserQuery := "INSERT INTO \"users\" (\"team_id\",\"type\",\"created_at\",\"updated_at\") VALUES ($1,$2,$3,$4) RETURNING \"id\""
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(insertUserQuery)).
			WithArgs(userEnt.TeamID, userEnt.Type, utils.AnyTime{}, utils.AnyTime{}).
			WillReturnRows(sqlmock.NewRows([]string{"id", "team_id", "type"}).AddRow(1, 1, "test"))
		mock.ExpectCommit()

		got, err := repo.Create(context.Background(), userEnt)
		if err != nil {
			t.Error(err)
			return
		}

		userEnt.ID = 1
		diff := cmp.Diff(got, userEnt, cmpopts.IgnoreFields(entity.User{}, "CreatedAt", "UpdatedAt"))
		if diff != "" {
			t.Error(diff)
		}
	})

	t.Run("create fail", func(t *testing.T) {
		userEnt := entity.User{TeamID: teamEnt.ID, Type: "test"}
		insertUserQuery := "INSERT INTO \"users\" (\"team_id\",\"type\",\"created_at\",\"updated_at\") VALUES ($1,$2,$3,$4) RETURNING \"id\""
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(insertUserQuery)).
			WithArgs(userEnt.TeamID, userEnt.Type, utils.AnyTime{}, utils.AnyTime{}).
			WillReturnError(createErr)
		mock.ExpectRollback()

		_, gotErr := repo.Create(context.Background(), userEnt)
		if gotErr != createErr {
			t.Error("expect error")
			return
		}
	})
}

func TestUserRepository_GetByID(t *testing.T) {
	db, mock, err := utils.OpenTestDBConnection()
	if err != nil {
		return
	}

	repo := &UserRepository{db: db}
	id := uint64(1)
	userAggregate := aggregate.UserAggregate{User: entity.User{ID: 1, TeamID: 1}}

	t.Run("get by id", func(t *testing.T) {
		getUserQuery := "SELECT users.id, users.team_id, users.type, users.created_at, users.updated_at, " +
			"teams.id, teams.hub_id, teams.geo_location, teams.created_at, teams.updated_at, " +
			"hubs.id, hubs.name, hubs.created_at, hubs.updated_at " +
			"FROM \"users\" join teams on team_id = users.team_id join hubs on teams.hub_id = hubs.id " +
			"WHERE \"users\".\"id\" = $1"
		mock.ExpectQuery(regexp.QuoteMeta(getUserQuery)).
			WithArgs(id).
			WillReturnRows(sqlmock.NewRows([]string{"id", "team_id"}).AddRow(1, 1))

		got, err := repo.GetByID(context.Background(), id)
		if err != nil {
			t.Error(err)
			return
		}
		diff := cmp.Diff(got, userAggregate)
		if diff != "" {
			t.Error(diff)
		}
	})
	t.Run("get by id", func(t *testing.T) {
		getUserQuery := "SELECT users.id, users.team_id, users.type, users.created_at, users.updated_at, " +
			"teams.id, teams.hub_id, teams.geo_location, teams.created_at, teams.updated_at, " +
			"hubs.id, hubs.name, hubs.created_at, hubs.updated_at " +
			"FROM \"users\" join teams on team_id = users.team_id join hubs on teams.hub_id = hubs.id " +
			"WHERE \"users\".\"id\" = $1"
		mock.ExpectQuery(regexp.QuoteMeta(getUserQuery)).
			WithArgs(id).
			WillReturnError(gorm.ErrRecordNotFound)

		_, err := repo.GetByID(context.Background(), id)
		if err != gorm.ErrRecordNotFound {
			t.Errorf("err:%v != wantErr: %v", err, gorm.ErrRecordNotFound)
			return
		}
	})
}
