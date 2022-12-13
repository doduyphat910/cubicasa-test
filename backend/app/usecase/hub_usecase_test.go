package usecase

import (
	"context"
	"errors"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/repository"
	"github.com/doduyphat910/cubicasa-test/backend/app/domain/repository/mock"
	"github.com/golang/mock/gomock"
	"github.com/google/go-cmp/cmp"
	"reflect"
	"testing"
)

func TestNewHubUseCase(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()
	mHubRepo := mock.NewMockHubRepository(mockCtrl)

	type args struct {
		hubRepo repository.HubRepository
	}
	type want struct {
		hubUC *HubUseCase
	}
	type TestCase struct {
		name string
		args args
		want want
	}

	testCases := []TestCase{
		{
			name: "test new hub usecase",
			args: args{hubRepo: mHubRepo},
			want: want{&HubUseCase{hubRepo: mHubRepo}},
		},
	}
	for i := range testCases {
		t.Run(testCases[i].name, func(t *testing.T) {
			got := NewHubUseCase(testCases[i].args.hubRepo)
			if !reflect.DeepEqual(got, testCases[i].want.hubUC) {
				t.Errorf("TestNewHubUseCase got:%v != want: %v", got, testCases[i].want.hubUC)
			}
		})
	}
}

func TestHubUseCase_Create(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	mHubRepo := mock.NewMockHubRepository(mockCtrl)
	uc := &HubUseCase{hubRepo: mHubRepo}

	type args struct {
		ctx context.Context
		hub entity.Hub
	}
	type want struct {
		hub entity.Hub
		err error
	}
	type TestCase struct {
		name      string
		args      args
		want      want
		setupFunc func()
	}

	err := errors.New("create error")
	ctx := context.Background()
	hubEnt := entity.Hub{ID: 1, Name: "test"}
	testCases := []TestCase{
		{
			name: "create hub error",
			args: args{
				ctx: ctx,
				hub: entity.Hub{},
			},
			want: want{
				hub: entity.Hub{},
				err: err,
			},
			setupFunc: func() {
				mHubRepo.EXPECT().Create(ctx, entity.Hub{}).Return(entity.Hub{}, err)
			},
		},
		{
			name: "create hub success",
			args: args{
				ctx: ctx,
				hub: entity.Hub{
					Name: "test",
				},
			},
			want: want{
				hub: hubEnt,
				err: nil,
			},
			setupFunc: func() {
				mHubRepo.EXPECT().Create(ctx, entity.Hub{Name: "test"}).Return(hubEnt, nil)
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			testCase.setupFunc()
			gotHub, gotErr := uc.Create(testCase.args.ctx, testCase.args.hub)
			if gotErr != testCase.want.err {
				t.Errorf("TestHubUseCase_Create error = %v, wantErr %v", gotErr, testCase.want.err)
				return
			}
			if diff := cmp.Diff(gotHub, testCase.want.hub); diff != "" {
				t.Errorf("TestHubUseCase_Create gotHub != wantHub: %s", diff)
			}
		})
	}
}
