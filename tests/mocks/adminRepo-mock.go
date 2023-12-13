package mocks

import (
	domain "qbills/models/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockAdminRepository is a mock of AdminRepository interface.
type MockAdminRepository struct {
        ctrl     *gomock.Controller
        recorder *MockAdminRepositoryMockRecorder
}

// MockAdminRepositoryMockRecorder is the mock recorder for MockAdminRepository.
type MockAdminRepositoryMockRecorder struct {
        mock *MockAdminRepository
}

// NewMockAdminRepository creates a new mock instance.
func NewMockAdminRepository(ctrl *gomock.Controller) *MockAdminRepository {
        mock := &MockAdminRepository{ctrl: ctrl}
        mock.recorder = &MockAdminRepositoryMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminRepository) EXPECT() *MockAdminRepositoryMockRecorder {
        return m.recorder
}

// Create mocks base method.
func (m *MockAdminRepository) Create(arg0 *domain.Admin) (*domain.Admin, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Create", arg0)
        ret0, _ := ret[0].(*domain.Admin)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockAdminRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockAdminRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockAdminRepository) Delete(arg0 int) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Delete", arg0)
        ret0, _ := ret[0].(error)
        return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockAdminRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockAdminRepository)(nil).Delete), arg0)
}

// FindAll mocks base method.
func (m *MockAdminRepository) FindAll() ([]domain.Admin, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "FindAll")
        ret0, _ := ret[0].([]domain.Admin)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockAdminRepositoryMockRecorder) FindAll() *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAdminRepository)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockAdminRepository) FindById(arg0 int) (*domain.Admin, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "FindById", arg0)
        ret0, _ := ret[0].(*domain.Admin)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockAdminRepositoryMockRecorder) FindById(arg0 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockAdminRepository)(nil).FindById), arg0)
}

// FindByUsername mocks base method.
func (m *MockAdminRepository) FindByUsername(arg0 string) (*domain.Admin, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "FindByUsername", arg0)
        ret0, _ := ret[0].(*domain.Admin)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// FindByUsername indicates an expected call of FindByUsername.
func (mr *MockAdminRepositoryMockRecorder) FindByUsername(arg0 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockAdminRepository)(nil).FindByUsername), arg0)
}

// Update mocks base method.
func (m *MockAdminRepository) Update(arg0 *domain.Admin, arg1 int) (*domain.Admin, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "Update", arg0, arg1)
        ret0, _ := ret[0].(*domain.Admin)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockAdminRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockAdminRepository)(nil).Update), arg0, arg1)
}