package mock_repository

import (
	"qbills/models/domain"

	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

type MockMembershipRepository struct {
	ctrl *gomock.Controller 
	recorder *MockMembershipRepositoryMockRecorder 
}

type MockMembershipRepositoryMockRecorder struct {
	mock *MockMembershipRepository 
}

func NewMockMembershipRepository(ctrl *gomock.Controller) *MockMembershipRepository {
	mock := &MockMembershipRepository{ctrl: ctrl}
	mock.recorder = &MockMembershipRepositoryMockRecorder{mock}
	return mock
}

func (m *MockMembershipRepository) EXPECT() *MockMembershipRepositoryMockRecorder {
	return m.recorder
}

func (m *MockMembershipRepository) Create(arg0 *domain.Membership) (*domain.Membership, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*domain.Membership)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMembershipRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockMembershipRepository)(nil).Create), arg0)
}

func (m *MockMembershipRepository) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

func (mr *MockMembershipRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockMembershipRepository)(nil).Delete), arg0)
}

func (m *MockMembershipRepository) FindAll() ([]domain.Membership, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]domain.Membership)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

func (mr *MockMembershipRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockMembershipRepository)(nil).FindAll))
}

func (m *MockMembershipRepository) FindById(arg0 int) ([]domain.Membership, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].([]domain.Membership)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMembershipRepositoryMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockMembershipRepository)(nil).FindById), arg0)
}

func (m *MockMembershipRepository) FindByName(arg0 string) ([]domain.Membership, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByName", arg0)
	ret0, _ := ret[0].([]domain.Membership)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMembershipRepositoryMockRecorder) FindByName(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByName", reflect.TypeOf((*MockMembershipRepository)(nil).FindByName), arg0)
}

func (m *MockMembershipRepository) FindByPhoneNumber(arg0 string) ([]domain.Membership, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByPhoneNumber", arg0)
	ret0, _ := ret[0].([]domain.Membership)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMembershipRepositoryMockRecorder) FindByPhoneNumber(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByPhoneNumber", reflect.TypeOf((*MockMembershipRepository)(nil).FindByPhoneNumber), arg0)
}

func (m *MockMembershipRepository) Update(arg0 *domain.Membership, arg1 int) ([]domain.Membership, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].([]domain.Membership)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

func (mr *MockMembershipRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockMembershipRepository)(nil).Update), arg0, arg1)
}