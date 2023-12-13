package mocks

import (
	domain "qbills/models/domain"
	web "qbills/models/web"
	reflect "reflect"

	gomock "github.com/golang/mock/gomock"
	echo "github.com/labstack/echo/v4"
)

// MockAdminService is a mock of AdminService interface.
type MockAdminService struct {
        ctrl     *gomock.Controller
        recorder *MockAdminServiceMockRecorder
}

// MockAdminServiceMockRecorder is the mock recorder for MockAdminService.
type MockAdminServiceMockRecorder struct {
        mock *MockAdminService
}

// NewMockAdminService creates a new mock instance.
func NewMockAdminService(ctrl *gomock.Controller) *MockAdminService {
        mock := &MockAdminService{ctrl: ctrl}
        mock.recorder = &MockAdminServiceMockRecorder{mock}
        return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockAdminService) EXPECT() *MockAdminServiceMockRecorder {
        return m.recorder
}

// CreateAdmin mocks base method.
func (m *MockAdminService) CreateAdmin(arg0 echo.Context, arg1 web.AdminCreateRequest) (*domain.Admin, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "CreateAdmin", arg0, arg1)
        ret0, _ := ret[0].(*domain.Admin)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// CreateAdmin indicates an expected call of CreateAdmin.
func (mr *MockAdminServiceMockRecorder) CreateAdmin(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "CreateAdmin", reflect.TypeOf((*MockAdminService)(nil).CreateAdmin), arg0, arg1)
}

// DeleteAdmin mocks base method.
func (m *MockAdminService) DeleteAdmin(arg0 echo.Context, arg1 int) error {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "DeleteAdmin", arg0, arg1)
        ret0, _ := ret[0].(error)
        return ret0
}

// DeleteAdmin indicates an expected call of DeleteAdmin.
func (mr *MockAdminServiceMockRecorder) DeleteAdmin(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "DeleteAdmin", reflect.TypeOf((*MockAdminService)(nil).DeleteAdmin), arg0, arg1)
}

// FindAll mocks base method.
func (m *MockAdminService) FindAll(arg0 echo.Context) ([]domain.Admin, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "FindAll", arg0)
        ret0, _ := ret[0].([]domain.Admin)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// FindAll indicates an expected call of FindAll.
func (mr *MockAdminServiceMockRecorder) FindAll(arg0 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindAll", reflect.TypeOf((*MockAdminService)(nil).FindAll), arg0)
}

// FindById mocks base method.
func (m *MockAdminService) FindById(arg0 echo.Context, arg1 int) (*domain.Admin, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "FindById", arg0, arg1)
        ret0, _ := ret[0].(*domain.Admin)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// FindById indicates an expected call of FindById.
func (mr *MockAdminServiceMockRecorder) FindById(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindById", reflect.TypeOf((*MockAdminService)(nil).FindById), arg0, arg1)
}

// FindByUsername mocks base method.
func (m *MockAdminService) FindByUsername(arg0 echo.Context, arg1 string) (*domain.Admin, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "FindByUsername", arg0, arg1)
        ret0, _ := ret[0].(*domain.Admin)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// FindByUsername indicates an expected call of FindByUsername.
func (mr *MockAdminServiceMockRecorder) FindByUsername(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "FindByUsername", reflect.TypeOf((*MockAdminService)(nil).FindByUsername), arg0, arg1)
}

// LoginAdmin mocks base method.
func (m *MockAdminService) LoginAdmin(arg0 echo.Context, arg1 web.AdminLoginRequest) (*domain.Admin, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "LoginAdmin", arg0, arg1)
        ret0, _ := ret[0].(*domain.Admin)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// LoginAdmin indicates an expected call of LoginAdmin.
func (mr *MockAdminServiceMockRecorder) LoginAdmin(arg0, arg1 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "LoginAdmin", reflect.TypeOf((*MockAdminService)(nil).LoginAdmin), arg0, arg1)
}

// UpdateAdmin mocks base method.
func (m *MockAdminService) UpdateAdmin(arg0 echo.Context, arg1 web.AdminUpdateRequest, arg2 int) (*domain.Admin, error) {
        m.ctrl.T.Helper()
        ret := m.ctrl.Call(m, "UpdateAdmin", arg0, arg1, arg2)
        ret0, _ := ret[0].(*domain.Admin)
        ret1, _ := ret[1].(error)
        return ret0, ret1
}

// UpdateAdmin indicates an expected call of UpdateAdmin.
func (mr *MockAdminServiceMockRecorder) UpdateAdmin(arg0, arg1, arg2 interface{}) *gomock.Call {
        mr.mock.ctrl.T.Helper()
        return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "UpdateAdmin", reflect.TypeOf((*MockAdminService)(nil).UpdateAdmin), arg0, arg1, arg2)
}