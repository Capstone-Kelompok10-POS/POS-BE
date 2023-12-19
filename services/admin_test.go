package services

import (
	"errors"
	"qbills/mocks"
	"qbills/models/domain"
	"qbills/models/web"
	"testing"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

func TestCreateAdmin_Success(t *testing.T) {
	mockRepository := new(mocks.AdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)

	mockRepository.On("FindByUsername", mock.Anything).Return(nil, nil)
	mockPasswordHandler.On("HashPassword", mock.Anything).Return("hashedPassword", nil)
	mockRepository.On("Create", mock.Anything).Return(&domain.Admin{}, nil)

	ctx := echo.New().NewContext(nil, nil)
	// Mock data
	_, err := adminService.CreateAdmin(ctx, web.AdminCreateRequest{
			SuperAdminID: 1,
			FullName:     "Acek",
			Username:     "acekasik",
			Password:     "chawni123",
	})

	// Assert the result
	assert.NoError(t, err)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestCreateAdmin_ExistingAdmin(t *testing.T) {
	mockRepository := new(mocks.AdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)

	// Simulasi admin dengan username yang sudah ada
	existingAdmin := &domain.Admin{ID: 1, Username: "acekasik"}
	mockRepository.On("FindByUsername", "acekasik").Return(existingAdmin, nil)

	ctx := echo.New().NewContext(nil, nil)
	// Mock data
	_, err := adminService.CreateAdmin(ctx, web.AdminCreateRequest{
		SuperAdminID: 1,
		FullName:     "Acek",
		Username:     "acekasik",
		Password:     "chawni123",
	})

	// Assert the result
	assert.Error(t, err)
	assert.EqualError(t, err, "username already exists")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestCreateAdmin_ValidationError(t *testing.T) {
	mockRepository := new(mocks.AdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)

	// Set up validation error
	mockValidator.RegisterValidation("customValidation", func(fl validator.FieldLevel) bool {
		return false
	})

	ctx := echo.New().NewContext(nil, nil)

	// Mock data
	_, err := adminService.CreateAdmin(ctx, web.AdminCreateRequest{
		SuperAdminID: 1,
		FullName:     "Acek",
		Username:     "acekasik",
		Password:     "123",
	})

	// Assert the validation error
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestCreateAdmin_Failure(t *testing.T) {
	mockRepository := new(mocks.AdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)

	// Set up expected calls
	mockRepository.On("FindByUsername", mock.Anything).Return(nil, nil)
	mockPasswordHandler.On("HashPassword", mock.Anything).Return("hashedPassword", nil)
	mockRepository.On("Create", mock.Anything).Return(nil, errors.New("failed to create admin"))

	ctx := echo.New().NewContext(nil, nil)
	// Mock data
	admin, err := adminService.CreateAdmin(ctx, web.AdminCreateRequest{
		SuperAdminID: 1,
		FullName:     "Acek",
		Username:     "acekasik",
		Password:     "chawni123",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, admin)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

// func TestLoginAdmin_Success(t *testing.T) {
// 	mockRepository := new(mocks.AdminRepository)
// 	mockValidator := validator.New()
// 	mockPasswordHandler := new(mocks.PasswordHandler)

// 	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)
// 	adminRequest := web.AdminLoginRequest{
// 		Username: "vilanatasya",
// 		Password: "vila12345",
// 	}

// 	mockRepository.On("FindByUsername", adminRequest.Username).Return(, nil)
	
// 	mockPasswordHandler.On("ComparePassword", "hashedPassword", adminRequest.Password).Return(nil)


// 	ctx := echo.New().NewContext(nil, nil)
// 	// Mock data
// 	_, err := adminService.CreateAdmin(ctx, web.AdminCreateRequest{
// 			SuperAdminID: 1,
// 			FullName:     "Acek",
// 			Username:     "acekasik",
// 			Password:     "chawni123",
// 	})

// 	// Assert the result
// 	assert.NoError(t, err)

// 	// Assert that the expected calls were made
// 	mockRepository.AssertExpectations(t)
// 	mockPasswordHandler.AssertExpectations(t)
// }
