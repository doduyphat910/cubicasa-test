package pgsql

import (
	"context"
	"github.com/DATA-DOG/go-sqlmock"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/utils"
	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"regexp"
	"testing"
)

func TestNewHubRepository(t *testing.T) {
	want := &HubRepository{}
	repo := NewHubRepository()

	if diff := cmp.Diff(repo, want,
		cmpopts.IgnoreFields(HubRepository{}, "db")); diff != "" {
		t.Error(diff)
	}
}

func TestHubRepository_Create(t *testing.T) {
	db, mock, err := utils.OpenTestDBConnection()
	if err != nil {
		return
	}

	repo := &HubRepository{
		db: db,
	}

	hubEnt := entity.Hub{Name: "name"}
	t.Run("create success", func(t *testing.T) {
		insertHubQuery := "INSERT INTO \"hubs\" (\"name\",\"created_at\",\"updated_at\") VALUES ($1,$2,$3) RETURNING \"id\""
		mock.ExpectBegin()
		mock.ExpectQuery(regexp.QuoteMeta(insertHubQuery)).
			WithArgs(hubEnt.Name, utils.AnyTime{}, utils.AnyTime{}).
			WillReturnRows(
				sqlmock.NewRows([]string{"id"}).AddRow(hubEnt.ID),
			)
		mock.ExpectCommit()

		got, err := repo.Create(context.Background(), hubEnt)
		if err != nil {
			t.Error(err)
			return
		}
		diff := cmp.Diff(got, hubEnt, cmpopts.IgnoreFields(entity.Hub{}, "CreatedAt", "UpdatedAt"))
		if diff != "" {
			t.Error(diff)
		}
	})
}

func TestHubRepository_GetByID(t *testing.T) {
	db, mock, err := utils.OpenTestDBConnection()
	if err != nil {
		return
	}

	repo := &HubRepository{
		db: db,
	}

	id := uint64(1)
	hubEnt := entity.Hub{ID: 1, Name: "name"}
	t.Run("get by id", func(t *testing.T) {
		getHubQuery := "SELECT * FROM \"hubs\" WHERE \"hubs\".\"id\" = $1 LIMIT 1"
		mock.ExpectQuery(regexp.QuoteMeta(getHubQuery)).
			WithArgs(id).
			WillReturnRows(
				sqlmock.NewRows([]string{"id", "name"}).AddRow(1, "name"),
			)

		got, err := repo.GetByID(context.Background(), id)
		if err != nil {
			t.Error(err)
			return
		}
		diff := cmp.Diff(got, hubEnt)
		if diff != "" {
			t.Error(diff)
		}
	})
}
