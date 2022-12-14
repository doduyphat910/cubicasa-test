// Code generated by MockGen. DO NOT EDIT.
// Source: team_repository.go

// Package mock is a generated GoMock package.
package mock

import (
	context "context"
	reflect "reflect"

	entity "github.com/doduyphat910/cubicasa-test/backend/app/domain/entity"
	utils "github.com/doduyphat910/cubicasa-test/backend/app/utils"
	gomock "github.com/golang/mock/gomock"
)

// MockTeamRepository is a mock of TeamRepository interface.
type MockTeamRepository struct {
	ctrl     *gomock.Controller
	recorder *MockTeamRepositoryMockRecorder
}

// MockTeamRepositoryMockRecorder is the mock recorder for MockTeamRepository.
type MockTeamRepositoryMockRecorder struct {
	mock *MockTeamRepository
}

// NewMockTeamRepository creates a new mock instance.
func NewMockTeamRepository(ctrl *gomock.Controller) *MockTeamRepository {
	mock := &MockTeamRepository{ctrl: ctrl}
	mock.recorder = &MockTeamRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockTeamRepository) EXPECT() *MockTeamRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockTeamRepository) Create(ctx context.Context, team entity.Team) (entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", ctx, team)
	ret0, _ := ret[0].(entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockTeamRepositoryMockRecorder) Create(ctx, team interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockTeamRepository)(nil).Create), ctx, team)
}

// GetByID mocks base method.
func (m *MockTeamRepository) GetByID(ctx context.Context, id uint64) (entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetByID", ctx, id)
	ret0, _ := ret[0].(entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetByID indicates an expected call of GetByID.
func (mr *MockTeamRepositoryMockRecorder) GetByID(ctx, id interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetByID", reflect.TypeOf((*MockTeamRepository)(nil).GetByID), ctx, id)
}

// Search mocks base method.
func (m *MockTeamRepository) Search(ctx context.Context, geoLocation entity.GeoLocation, paging utils.Paging) ([]entity.Team, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Search", ctx, geoLocation, paging)
	ret0, _ := ret[0].([]entity.Team)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Search indicates an expected call of Search.
func (mr *MockTeamRepositoryMockRecorder) Search(ctx, geoLocation, paging interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Search", reflect.TypeOf((*MockTeamRepository)(nil).Search), ctx, geoLocation, paging)
}
