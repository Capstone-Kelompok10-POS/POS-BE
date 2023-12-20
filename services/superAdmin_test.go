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
)

func TestLoginSuperAdmin_Success(t *testing.T) {
	mockRepository := new(mocks.SuperAdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	superAdminService := NewSuperAdminService(mockRepository, mockValidator, mockPasswordHandler)
	superAdminRequest := web.SuperAdminLoginRequest{
		Username: "vilanatasya",
		Password: "vila12345",
	}

	mockRepository.On("FindByUsername", superAdminRequest.Username).Return(&domain.SuperAdmin{ID: 1, Password: "hashedPassword"}, nil)

	mockPasswordHandler.On("ComparePassword", "hashedPassword", superAdminRequest.Password).Return(nil)

	ctx := echo.New().NewContext(nil, nil)
	result, err := superAdminService.LoginSuperAdmin(ctx, superAdminRequest)
	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestLoginSuperAdmin_Failure(t *testing.T) {
	mockRepository := new(mocks.SuperAdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	superAdminService := NewSuperAdminService(mockRepository, mockValidator, mockPasswordHandler)
	superAdminRequest := web.SuperAdminLoginRequest{
		Username: "vilanatasya",
		Password: "vila12345",
	}

	mockRepository.On("FindByUsername", superAdminRequest.Username).Return(&domain.SuperAdmin{ID: 1, Password: "hashedPassword"}, nil)

	mockPasswordHandler.On("ComparePassword", "hashedPassword", superAdminRequest.Password).Return(errors.New("invalid username or password"))

	ctx := echo.New().NewContext(nil, nil)
	result, err := superAdminService.LoginSuperAdmin(ctx, superAdminRequest)
	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestLoginSuperAdmin_ValidationError(t *testing.T) {
	mockRepository := new(mocks.SuperAdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	superAdminService := NewSuperAdminService(mockRepository, mockValidator, mockPasswordHandler)
	superAdminRequest := web.SuperAdminLoginRequest{
		Username: "vilanatasya",
		Password: "",
	}

	ctx := echo.New().NewContext(nil, nil)
	_ , err := superAdminService.LoginSuperAdmin(ctx, superAdminRequest)
	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "Validation error")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}


func TestLoginSuperAdmin_UsernameOrPasswordError(t *testing.T) {
	mockRepository := new(mocks.SuperAdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	superAdminService := NewSuperAdminService(mockRepository, mockValidator, mockPasswordHandler)
	superAdminRequest := web.SuperAdminLoginRequest{
		Username: "vilanatasya",
		Password: "vila12345",
	}

	mockRepository.On("FindByUsername", superAdminRequest.Username).Return(&domain.SuperAdmin{ID: 1, Password: "hashedPassword"}, errors.New("invalid username or password"))

	ctx := echo.New().NewContext(nil, nil)
	_ , err := superAdminService.LoginSuperAdmin(ctx, superAdminRequest)
	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid username or password")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestSuperAdminFindById_Success(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	superAdminRepoMock := new(mocks.SuperAdminRepository)

	// Create an instance of the service with the mock
	superAdminService := NewSuperAdminService(superAdminRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindById method
	expectedSuperAdmin := &domain.SuperAdmin{
		ID:             1,
		Username:       "john.doe",
		Password:       "hashedPassword",
	}

	superAdminRepoMock.On("FindById", 1).Return(expectedSuperAdmin, nil)

	// Call the method you want to test
	resultAdmin, err := superAdminService.FindById(nil, 1)

	// Assert that the FindById method was called with the expected parameters
	superAdminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultAdmin)
	assert.Equal(t, expectedSuperAdmin, resultAdmin)
}

func TestSuperAdminFindById_AdminNotFound(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	superAdminRepoMock := new(mocks.SuperAdminRepository)

	// Create an instance of the service with the mock
	superAdminService := NewSuperAdminService(superAdminRepoMock, mockValidator, mockPasswordHandler)

	superAdminRepoMock.On("FindById", 1).Return(nil, nil)

	// Call the method you want to test
	resultAdmin, err := superAdminService.FindById(nil, 1)

	// Assert that the FindById method was called with the expected parameters
	superAdminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultAdmin)
	assert.EqualError(t, err, "SuperAdmin not found")
}

func TestSuperAdminFindAll_Success(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	superAdminRepoMock := new(mocks.SuperAdminRepository)

	// Create an instance of the service with the mock
	superAdminService := NewSuperAdminService(superAdminRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindAll method
	expectedSuperAdmins := []domain.SuperAdmin{
		{
			ID:           1,
			Username:     "johndoe",
			Password:     "hashedPassword",
		},
		{
			ID:           2,
			Username:     "janedoe",
			Password:     "hashedPassword",
		},
	}

	superAdminRepoMock.On("FindAll").Return(expectedSuperAdmins, nil)

	// Call the method you want to test
	resultSuperAdmins, err := superAdminService.FindAll(nil)

	// Assert that the FindAll method was called with the expected parameters
	superAdminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultSuperAdmins)
	assert.Equal(t, expectedSuperAdmins, resultSuperAdmins)
}

func TestSuperAdminFindAll_NoSuperAdminsFound(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	superAdminRepoMock := new(mocks.SuperAdminRepository)

	// Create an instance of the service with the mock
	superAdminService := NewSuperAdminService(superAdminRepoMock, mockValidator, mockPasswordHandler)


	superAdminRepoMock.On("FindAll").Return(nil, errors.New("SuperAdmins not found"))

	// Call the method you want to test
	resultSuperAdmins, err := superAdminService.FindAll(nil)

	// Assert that the FindAll method was called with the expected parameters
	superAdminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultSuperAdmins)
	assert.EqualError(t, err, "SuperAdmins not found")
}