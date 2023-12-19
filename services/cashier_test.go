package services_test

import (
	"errors"
	"qbills/mocks"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/services"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateCashier_Success(t *testing.T) {
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	cashierService := services.NewCashierService(mockRepository, mockValidator, mockPasswordHandler)

	mockRepository.On("FindByUsername", mock.Anything).Return(nil, nil)
	mockPasswordHandler.On("HashPassword", mock.Anything).Return("hashedPassword", nil)
	mockRepository.On("Create", mock.Anything).Return(&domain.Cashier{}, nil)

	ctx := echo.New().NewContext(nil, nil)
	// Mock data
	_, err := cashierService.CreateCashier(ctx, web.CashierCreateRequest{
		AdminID:  1,
		Fullname: "Halim",
		Username: "Halim",
		Password: "1223",
	})

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestCreateCashier_ExistingCashier(t *testing.T) {
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	cashierService := services.NewCashierService(mockRepository, mockValidator, mockPasswordHandler)

	// Simulate an existing cashier with the same username
	existingCashier := &domain.Cashier{ID: 1, Username: "johndoe"}
	mockRepository.On("FindByUsername", "johndoe").Return(existingCashier, nil)

	ctx := echo.New().NewContext(nil, nil)
	// Mock data
	_, err := cashierService.CreateCashier(ctx, web.CashierCreateRequest{
		AdminID:  1,
		Fullname: "halim",
		Username: "halim",
		Password: "1234",
	})

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "username already exists")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestCreateCashier_ValidationError(t *testing.T) {
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	cashierService := services.NewCashierService(mockRepository, mockValidator, mockPasswordHandler)

	// Set up validation error
	mockValidator.RegisterValidation("customValidation", func(fl validator.FieldLevel) bool {
		return false
	})

	ctx := echo.New().NewContext(nil, nil)

	// Mock data
	_, err := cashierService.CreateCashier(ctx, web.CashierCreateRequest{
		AdminID:  1,
		Fullname: "John Doe",
		Username: "johndoe",
		Password: "123", // Invalid password (less than 8 characters)
	})

	// Assert the validation error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestCreateCashier_Failure(t *testing.T) {
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	cashierService := services.NewCashierService(mockRepository, mockValidator, mockPasswordHandler)

	// Set up expected calls
	mockRepository.On("FindByUsername", mock.Anything).Return(nil, nil)
	mockPasswordHandler.On("HashPassword", mock.Anything).Return("hashedPassword", nil)
	mockRepository.On("Create", mock.Anything).Return(nil, errors.New("failed to create cashier"))

	ctx := echo.New().NewContext(nil, nil)
	// Mock data
	cashier, err := cashierService.CreateCashier(ctx, web.CashierCreateRequest{
		AdminID:  1,
		Fullname: "John Doe",
		Username: "johndoe",
		Password: "password123",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, cashier)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}
