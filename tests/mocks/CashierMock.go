package mock_repository

import (
	domain "qbills/models/domain"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
)

// MockCashierRepository is a mock of CashierRepository interface.
type MockCashierRepository struct {
	ctrl     *gomock.Controller
	recorder *MockCashierRepositoryMockRecorder
}

// MockCashierRepositoryMockRecorder is the mock recorder for MockCashierRepository.
type MockCashierRepositoryMockRecorder struct {
	mock *MockCashierRepository
}

// NewMockCashierRepository creates a new mock instance.
func NewMockCashierRepository(ctrl *gomock.Controller) *MockCashierRepository {
	mock := &MockCashierRepository{ctrl: ctrl}
	mock.recorder = &MockCashierRepositoryMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockCashierRepository) EXPECT() *MockCashierRepositoryMockRecorder {
	return m.recorder
}

// Create mocks base method.
func (m *MockCashierRepository) Create(arg0 *domain.Cashier) (*domain.Cashier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Create", arg0)
	ret0, _ := ret[0].(*domain.Cashier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Create indicates an expected call of Create.
func (mr *MockCashierRepositoryMockRecorder) Create(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Create", reflect.TypeOf((*MockCashierRepository)(nil).Create), arg0)
}

// Delete mocks base method.
func (m *MockCashierRepository) Delete(arg0 int) error {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Delete", arg0)
	ret0, _ := ret[0].(error)
	return ret0
}

// Delete indicates an expected call of Delete.
func (mr *MockCashierRepositoryMockRecorder) Delete(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Delete", reflect.TypeOf((*MockCashierRepository)(nil).Delete), arg0)
}

// FindAll mocks base method.
func (m *MockCashierRepository) FindAll() ([]domain.Cashier, int, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindAll")
	ret0, _ := ret[0].([]domain.Cashier)
	ret1, _ := ret[1].(int)
	ret2, _ := ret[2].(error)
	return ret0, ret1, ret2
}

// FindAll indicates an expected call of FindAll.
func (mr *MockCashierRepositoryMockRecorder) FindAll() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockCashierRepository)(nil).FindAll))
}

// FindById mocks base method.
func (m *MockCashierRepository) FindById(arg0 int) (*domain.Cashier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindById", arg0)
	ret0, _ := ret[0].(*domain.Cashier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockCashierRepositoryMockRecorder) FindById(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockCashierRepository)(nil).FindById), arg0)
}

// FindByUsername mocks base method.
func (m *MockCashierRepository) FindByUsername(arg0 string) (*domain.Cashier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "FindByUsername", arg0)
	ret0, _ := ret[0].(*domain.Cashier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// FindByUsername indicates an expected call of FindByUsername.
func (mr *MockCashierRepositoryMockRecorder) FindByUsername(arg0 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockCashierRepository)(nil).FindByUsername), arg0)
}

// Update mocks base method.
func (m *MockCashierRepository) Update(arg0 *domain.Cashier, arg1 int) (*domain.Cashier, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "Update", arg0, arg1)
	ret0, _ := ret[0].(*domain.Cashier)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// Update indicates an expected call of Update.
func (mr *MockCashierRepositoryMockRecorder) Update(arg0, arg1 interface{}) *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "Update", reflect.TypeOf((*MockCashierRepository)(nil).Update), arg0, arg1)
}
