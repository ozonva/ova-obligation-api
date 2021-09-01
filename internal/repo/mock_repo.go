// Code generated by MockGen. DO NOT EDIT.
// Source: repo.go

// Package repo is a generated GoMock package.
package repo

import (
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	entity "github.com/ozonva/ova-obligation-api/internal/entity"
)

// MockRepo is a mock of Repo interface.
type MockRepo struct {
	ctrl     *gomock.Controller
	recorder *MockRepoMockRecorder
}

// MockRepoMockRecorder is the mock recorder for MockRepo.
type MockRepoMockRecorder struct {
	mock *MockRepo
}

// NewMockRepo creates a new mock instance.
func NewMockRepo(ctrl *gomock.Controller) *MockRepo {
	mock := &MockRepo{ctrl: ctrl}
	mock.recorder = &MockRepoMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockRepo) EXPECT() *MockRepoMockRecorder {
	return m.recorder
}

// AddEntities mocks base method.
func (m *MockRepo) AddEntities(entities []*entity.Obligation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEntities", entities)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEntities indicates an expected call of AddEntities.
func (mr *MockRepoMockRecorder) AddEntities(entities interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEntities", reflect.TypeOf((*MockRepo)(nil).AddEntities), entities)
}

// AddEntity mocks base method.
func (m *MockRepo) AddEntity(entity *entity.Obligation) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "AddEntity", entity)
	ret0, _ := ret[0].(error)
	return ret0
}

// AddEntity indicates an expected call of AddEntity.
func (mr *MockRepoMockRecorder) AddEntity(entity interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "AddEntity", reflect.TypeOf((*MockRepo)(nil).AddEntity), entity)
}

// DescribeEntity mocks base method.
func (m *MockRepo) DescribeEntity(entityId uint64) (*entity.Obligation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "DescribeEntity", entityId)
	ret0, _ := ret[0].(*entity.Obligation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// DescribeEntity indicates an expected call of DescribeEntity.
func (mr *MockRepoMockRecorder) DescribeEntity(entityId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DescribeEntity", reflect.TypeOf((*MockRepo)(nil).DescribeEntity), entityId)
}

// ListEntities mocks base method.
func (m *MockRepo) ListEntities(limit, offset uint64) ([]entity.Obligation, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "ListEntities", limit, offset)
	ret0, _ := ret[0].([]entity.Obligation)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// ListEntities indicates an expected call of ListEntities.
func (mr *MockRepoMockRecorder) ListEntities(limit, offset interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "ListEntities", reflect.TypeOf((*MockRepo)(nil).ListEntities), limit, offset)
}

// RemoveEntity mocks base method.
func (m *MockRepo) RemoveEntity(entityId uint64) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "RemoveEntity", entityId)
	ret0, _ := ret[0].(error)
	return ret0
}

// RemoveEntity indicates an expected call of RemoveEntity.
func (mr *MockRepoMockRecorder) RemoveEntity(entityId interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "RemoveEntity", reflect.TypeOf((*MockRepo)(nil).RemoveEntity), entityId)
}
