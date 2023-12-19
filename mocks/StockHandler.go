// Code generated by mockery v0.0.0-dev. DO NOT EDIT.

package mocks

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// StockHandler is an autogenerated mock type for the StockHandler type
type StockHandler struct {
	mock.Mock
}

// FindAllStockHandler provides a mock function with given fields: ctx
func (_m *StockHandler) FindAllStockHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindByIdStockHandler provides a mock function with given fields: ctx
func (_m *StockHandler) FindByIdStockHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindDecreaseStockHandler provides a mock function with given fields: ctx
func (_m *StockHandler) FindDecreaseStockHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// FindIncreaseStockHandler provides a mock function with given fields: ctx
func (_m *StockHandler) FindIncreaseStockHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// UpdateStockHandler provides a mock function with given fields: ctx
func (_m *StockHandler) UpdateStockHandler(ctx echo.Context) error {
	ret := _m.Called(ctx)

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// NewStockHandler creates a new instance of StockHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewStockHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *StockHandler {
	mock := &StockHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}