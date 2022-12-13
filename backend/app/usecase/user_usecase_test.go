package usecase

import (
	"context"
	"errors"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/aggregate"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/repository"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestNewUserUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mTeamRepo := mock.NewMockTeamRepository(mockCtrl)
	mUserRepo := mock.NewMockUserRepository(mockCtrl)

	type args struct {
		teamRepo repository.TeamRepository
		userRepo repository.UserRepository
	}
	type want struct {
		userUC *UserUseCase
	}
	type TestCase struct {
		name string
		args args
		want want
	}

	testCases := []TestCase{
		{
			name: "test new user usecase",
			args: args{userRepo: mUserRepo, teamRepo: mTeamRepo},
			want: want{userUC: &UserUseCase{userRepo: mUserRepo, teamRepo: mTeamRepo}},
		},
	}
	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			got := NewUserUseCase(testCases[i].args.teamRepo, testCases[i].args.userRepo)
			if !reflect.DeepEqual(got, testCases[i].want.userUC) {
				t.Errorf("TestNewUserUseCase got:%v != want: %v", got, testCases[i].want.userUC)
			}
		})
	}
}

func TestUserUseCase_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mTeamRepo := mock.NewMockTeamRepository(mockCtrl)
	mUserRepo := mock.NewMockUserRepository(mockCtrl)
	uc := &UserUseCase{userRepo: mUserRepo, teamRepo: mTeamRepo}

	type args struct {
		ctx  context.Context
		user entity.User
	}
	type want struct {
		user entity.User
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
	userEnt := entity.User{TeamID: 1}
	wantUserEnt := entity.User{ID: 1, TeamID: 1}

	testCases := []TestCase{
		{
			name: "create user error",
			args: args{
				ctx:  ctx,
				user: userEnt,
			},
			want: want{
				user: entity.User{},
				err:  dummyErr,
			},
			setupFunc: func() {
				mTeamRepo.EXPECT().GetByID(ctx, userEnt.TeamID).Return(entity.Team{}, dummyErr)
			},
		},
		{
			name: "create user success",
			args: args{
				ctx:  ctx,
				user: userEnt,
			},
			want: want{
				user: wantUserEnt,
				err:  nil,
			},
			setupFunc: func() {
				mTeamRepo.EXPECT().GetByID(ctx, userEnt.TeamID).Return(entity.Team{ID: userEnt.TeamID}, nil)
				mUserRepo.EXPECT().Create(ctx, userEnt).Return(wantUserEnt, nil)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupFunc()
			gotUser, gotErr := uc.Create(testCase.args.ctx, testCase.args.user)
			if gotErr != testCase.want.err {
				t.Errorf("TestUserUseCase_Create error = %v, wantErr %v", gotErr, testCase.want.err)
				return
			}
			if diff := cmp.Diff(gotUser, testCase.want.user); diff != "" {
				t.Errorf("TestUserUseCase_Create gotUser != wantUser: %s", diff)
			}
		})
	}
}

func TestUserUseCase_GetByID(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mUserRepo := mock.NewMockUserRepository(mockCtrl)
	uc := &UserUseCase{userRepo: mUserRepo}

	type args struct {
		ctx context.Context
		id  uint64
	}
	type want struct {
		userAggregate aggregate.UserAggregate
		err           error
	}
	type TestCase struct {
		name      string
		args      args
		want      want
		setupFunc func()
	}

	ctx := context.Background()

	userAggregateWant := aggregate.UserAggregate{}
	id := uint64(1)

	testCases := []TestCase{
		{
			name: "create team error",
			args: args{
				ctx: ctx,
				id:  id,
			},
			want: want{
				userAggregate: userAggregateWant,
				err:           nil,
			},
			setupFunc: func() {
				mUserRepo.EXPECT().GetByID(ctx, id).Return(userAggregateWant, nil)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupFunc()
			gotUserAggregate, gotErr := uc.GetByID(testCase.args.ctx, testCase.args.id)
			if gotErr != testCase.want.err {
				t.Errorf("TestUserUseCase_GetByID error = %v, wantErr %v", gotErr, testCase.want.err)
				return
			}
			if diff := cmp.Diff(gotUserAggregate, testCase.want.userAggregate); diff != "" {
				t.Errorf("TestUserUseCase_GetByID gotUserAggregate != wantUserAggregate: %s", diff)
			}
		})
	}
}
