package usecase

import (
	"context"
	"errors"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/repository"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/repository/mock"
	"github.com/doduyphat910/cubicasa-test/backend/app/usecase/dto"
	"github.com/doduyphat910/cubicasa-test/backend/app/utils"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestNewTeamUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mHubRepo := mock.NewMockHubRepository(mockCtrl)
	mTeamRepo := mock.NewMockTeamRepository(mockCtrl)

	type args struct {
		hubRepo  repository.HubRepository
		teamRepo repository.TeamRepository
	}
	type want struct {
		teamUC *TeamUseCase
	}
	type TestCase struct {
		name string
		args args
		want want
	}

	testCases := []TestCase{
		{
			name: "test new team usecase",
			args: args{hubRepo: mHubRepo, teamRepo: mTeamRepo},
			want: want{&TeamUseCase{hubRepo: mHubRepo, teamRepo: mTeamRepo}},
		},
	}
	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			got := NewTeamUseCase(testCases[i].args.teamRepo, testCases[i].args.hubRepo)
			if !reflect.DeepEqual(got, testCases[i].want.teamUC) {
				t.Errorf("TestNewTeamUseCase got:%v != want: %v", got, testCases[i].want.teamUC)
			}
		})
	}
}

func TestTeamUseCase_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mHubRepo := mock.NewMockHubRepository(mockCtrl)
	mTeamRepo := mock.NewMockTeamRepository(mockCtrl)
	uc := &TeamUseCase{hubRepo: mHubRepo, teamRepo: mTeamRepo}

	type args struct {
		ctx  context.Context
		team entity.Team
	}
	type want struct {
		team entity.Team
		err  error
	}
	type TestCase struct {
		name      string
		args      args
		want      want
		setupFunc func()
	}

	dummyErr := errors.New("dummy error")
	ctx := context.Background()
	teamEnt := entity.Team{HubID: 1}
	wantTeamEnt := entity.Team{ID: 1, HubID: 1}

	testCases := []TestCase{
		{
			name: "create team error",
			args: args{
				ctx:  ctx,
				team: teamEnt,
			},
			want: want{
				team: entity.Team{},
				err:  dummyErr,
			},
			setupFunc: func() {
				mHubRepo.EXPECT().GetByID(ctx, teamEnt.HubID).Return(entity.Hub{}, dummyErr)
			},
		},
		{
			name: "create team success",
			args: args{
				ctx:  ctx,
				team: teamEnt,
			},
			want: want{
				team: wantTeamEnt,
				err:  nil,
			},
			setupFunc: func() {
				mHubRepo.EXPECT().GetByID(ctx, teamEnt.HubID).Return(entity.Hub{ID: 1}, nil)
				mTeamRepo.EXPECT().Create(ctx, teamEnt).Return(wantTeamEnt, nil)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupFunc()
			gotTeam, gotErr := uc.Create(testCase.args.ctx, testCase.args.team)
			if gotErr != testCase.want.err {
				t.Errorf("TestTeamUseCase_Create error = %v, wantErr %v", gotErr, testCase.want.err)
				return
			}
			if diff := cmp.Diff(gotTeam, testCase.want.team); diff != "" {
				t.Errorf("TestTeamUseCase_Create gotTeam != wantTeam: %s", diff)
			}
		})
	}
}

func TestTeamUseCase_Search(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mTeamRepo := mock.NewMockTeamRepository(mockCtrl)
	uc := &TeamUseCase{teamRepo: mTeamRepo}

	type args struct {
		ctx    context.Context
		dtoReq dto.SearchTeamRequest
	}
	type want struct {
		teams []entity.Team
		err   error
	}
	type TestCase struct {
		name      string
		args      args
		want      want
		setupFunc func()
	}

	ctx := context.Background()
	wantTeamEnts := []entity.Team{
		{ID: 1, GeoLocation: entity.GeoLocation{Lat: 123.123, Long: 456.456}},
	}
	dtoReq := dto.SearchTeamRequest{
		Lat:  123.123,
		Long: 456.456,
		Paging: utils.Paging{
			Size:   1,
			Number: 1,
		},
	}

	testCases := []TestCase{
		{
			name: "create team error",
			args: args{
				ctx:    ctx,
				dtoReq: dtoReq,
			},
			want: want{
				teams: wantTeamEnts,
				err:   nil,
			},
			setupFunc: func() {
				mTeamRepo.EXPECT().Search(ctx, entity.GeoLocation{Lat: dtoReq.Lat, Long: dtoReq.Long}, dtoReq.Paging).
					Return(wantTeamEnts, nil)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupFunc()
			gotTeams, gotErr := uc.Search(testCase.args.ctx, testCase.args.dtoReq)
			if gotErr != testCase.want.err {
				t.Errorf("TestTeamUseCase_Create error = %v, wantErr %v", gotErr, testCase.want.err)
				return
			}
			if diff := cmp.Diff(gotTeams, testCase.want.teams); diff != "" {
				t.Errorf("TestTeamUseCase_Create gotTeam != wantTeam: %s", diff)
			}
		})
	}
}
