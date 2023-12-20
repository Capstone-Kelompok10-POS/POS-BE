package services

import (
	"errors"
	"fmt"
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

func TestLoginAdmin_Success(t *testing.T) {
	mockRepository := new(mocks.AdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)
	adminRequest := web.AdminLoginRequest{
		Username: "vilanatasya",
		Password: "vila12345",
	}

	mockRepository.On("FindByUsername", adminRequest.Username).Return(&domain.Admin{ID:1, Password: "hashedPassword"}, nil)
	
	mockPasswordHandler.On("ComparePassword", "hashedPassword", adminRequest.Password).Return(nil)

	ctx := echo.New().NewContext(nil, nil)
	result, err := adminService.LoginAdmin(ctx, adminRequest)
	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}



func TestLoginAdmin_Failure(t *testing.T) {
	mockRepository := new(mocks.AdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)
	adminRequest := web.AdminLoginRequest{
		Username: "vilanatasya",
		Password: "vila12345",
	}

	mockRepository.On("FindByUsername", adminRequest.Username).Return(&domain.Admin{ID:1, Password: "hashedPassword"}, nil)
	
	mockPasswordHandler.On("ComparePassword", "hashedPassword", adminRequest.Password).Return(errors.New("invalid username or password"))

	ctx := echo.New().NewContext(nil, nil)
	result, err := adminService.LoginAdmin(ctx, adminRequest)
	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}


func TestLoginAdmin_UsernameOrPasswordError(t *testing.T) {
	mockRepository := new(mocks.AdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)
	adminRequest := web.AdminLoginRequest{
		Username: "vilanatasya",
		Password: "vila12345",
	}

	mockRepository.On("FindByUsername", adminRequest.Username).Return(nil, errors.New("invalid username or password"))
	ctx := echo.New().NewContext(nil, nil)
	_ , err := adminService.LoginAdmin(ctx, adminRequest)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid username or password")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestLoginAdmin_ValidationError(t *testing.T) {
	mockRepository := new(mocks.AdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)
	adminRequest := web.AdminLoginRequest{
		Username: "vilanatasya",
		Password: "",
	}

	ctx := echo.New().NewContext(nil, nil)
	_ , err := adminService.LoginAdmin(ctx, adminRequest)
	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestUpdateAdmin_Success(t *testing.T) {
	// Create a new instance of the mock repository, validator, and password handler
	mockRepository := new(mocks.AdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	// Create an instance of the admin service with the mocks
	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)

	adminID := 1
	adminId := uint(adminID)
	adminUpdateRequest := web.AdminUpdateRequest{
		FullName: "Acek",
		Username: "acekasik",
		Password: "chawni123",
	}

	// Set up the expected return values from the mock repository
	mockRepository.On("FindById", adminID).Return(&domain.Admin{}, nil)
	mockRepository.On("FindByUsername", mock.Anything).Return(nil, nil)
	mockPasswordHandler.On("HashPassword", mock.Anything).Return("hashedPassword", nil)

	// Set up the expected call for the Update method
	mockRepository.On("Update", mock.Anything, adminID).Return(&domain.Admin{ID: adminId, FullName: adminUpdateRequest.FullName}, nil)

	updatedAdmin, err := adminService.UpdateAdmin(nil, adminUpdateRequest, adminID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, updatedAdmin)
	assert.Equal(t, adminUpdateRequest.FullName, updatedAdmin.FullName)

	// Assert that the expected calls were made to the mock repository
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}


func TestUpdateAdmin_ValidationError(t *testing.T) {
	// Create a new instance of the mock repository, validator, and password handler
	mockRepository := new(mocks.AdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	// Create an instance of the admin service with the mocks
	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)

	adminID := 1
	adminUpdateRequest := web.AdminUpdateRequest{
		FullName: "Acek",
		Username: "acekasik",
		Password: "123",
	}

	_ , err := adminService.UpdateAdmin(nil, adminUpdateRequest, adminID)


	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")

	// Assert that the expected calls were made to the mock repository
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestUpdateAdmin_AdminNotFound(t *testing.T) {
	// Create a new instance of the mock repository, validator, and password handler
	mockRepository := new(mocks.AdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	// Create an instance of the admin service with the mocks
	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)

	adminID := 1
	adminUpdateRequest := web.AdminUpdateRequest{
		FullName: "Acek",
		Username: "acekasik",
		Password: "chawni123",
	}

	// Set up the mock repository to return AdminNotFound error
	mockRepository.On("FindById", adminID).Return(nil, errors.New("Admin not found"))

	// Call the UpdateAdmin method
	updatedAdmin, err := adminService.UpdateAdmin(nil, adminUpdateRequest, adminID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, updatedAdmin)

	// Assert that the expected calls were made to the mock repository
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestUpdateAdmin_Failure(t *testing.T) {
	// Create a new instance of the mock repository, validator, and password handler
	mockRepository := new(mocks.AdminRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	// Create an instance of the admin service with the mocks
	adminService := NewAdminService(mockRepository, mockValidator, mockPasswordHandler)

	adminID := 1
	adminId := uint(adminID)
	adminUpdateRequest := web.AdminUpdateRequest{
		FullName: "Acek",
		Username: "acekasik",
		Password: "chawni123",
	}

	// Set up the expected return values from the mock repository
	mockRepository.On("FindById", adminID).Return(&domain.Admin{}, nil)
	mockRepository.On("FindByUsername", mock.Anything).Return(nil, nil)
	mockPasswordHandler.On("HashPassword", mock.Anything).Return("hashedPassword", nil)

	// Set up the expected call for the Update method
	mockRepository.On("Update", mock.Anything, adminID).Return(&domain.Admin{ID: adminId, FullName: adminUpdateRequest.FullName}, errors.New("failed to update admin"))

	updatedAdmin, err := adminService.UpdateAdmin(nil, adminUpdateRequest, adminID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, updatedAdmin)

	// Assert that the expected calls were made to the mock repository
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}



func TestAdminFindById_Success(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	adminRepoMock := new(mocks.AdminRepository)

	// Create an instance of the service with the mock
	adminService := NewAdminService(adminRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindById method
	expectedAdmin := &domain.Admin{
		ID:             1,
		SuperAdminID:   1,
		FullName:       "John Doe",
		Username:       "john.doe",
		Password:       "hashedPassword",
	}

	adminRepoMock.On("FindById", 1).Return(expectedAdmin, nil)

	// Call the method you want to test
	resultAdmin, err := adminService.FindById(nil, 1)

	// Assert that the FindById method was called with the expected parameters
	adminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultAdmin)
	assert.Equal(t, expectedAdmin, resultAdmin)
}


func TestAdminFindById_AdminNotFound(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	adminRepoMock := new(mocks.AdminRepository)

	// Create an instance of the service with the mock
	adminService := NewAdminService(adminRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindById method
	adminRepoMock.On("FindById", 1).Return(nil, nil) // Simulating admin not found

	// Call the method you want to test
	resultAdmin, err := adminService.FindById(nil, 1)

	// Assert that the FindById method was called with the expected parameters
	adminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultAdmin)
	assert.EqualError(t, err, "admin not found")
}


func TestAdminFindAll_Success(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	adminRepoMock := new(mocks.AdminRepository)

	// Create an instance of the service with the mock
	adminService := NewAdminService(adminRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindAll method
	expectedAdmins := []domain.Admin{
		{
			ID:           1,
			SuperAdminID: 1,
			FullName:     "John Doe",
			Username:     "john.doe",
			Password:     "hashedPassword",
		},
		{
			ID:           2,
			SuperAdminID: 1,
			FullName:     "Jane Doe",
			Username:     "jane.doe",
			Password:     "hashedPassword",
		},
	}

	adminRepoMock.On("FindAll").Return(expectedAdmins, nil)

	// Call the method you want to test
	resultAdmins, err := adminService.FindAll(nil)

	// Assert that the FindAll method was called with the expected parameters
	adminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultAdmins)
	assert.Equal(t, expectedAdmins, resultAdmins)
}

func TestAdminFindAll_NoAdminsFound(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	adminRepoMock := new(mocks.AdminRepository)

	// Create an instance of the service with the mock
	adminService := NewAdminService(adminRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindAll method when no admins are found
	adminRepoMock.On("FindAll").Return(nil, errors.New("admins not found"))

	// Call the method you want to test
	resultAdmins, err := adminService.FindAll(nil)

	// Assert that the FindAll method was called with the expected parameters
	adminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultAdmins)
	assert.EqualError(t, err, "admins not found")
}


func TestAdminFindByUsername_Success(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	adminRepoMock := new(mocks.AdminRepository)

	// Create an instance of the service with the mock
	adminService := NewAdminService(adminRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindByUsername method
	expectedAdmin := &domain.Admin{
		ID:           1,
		SuperAdminID: 1,
		FullName:     "John Doe",
		Username:     "johndoe",
		Password:     "hashedPassword",
	}

	adminRepoMock.On("FindByUsername", "johndoe").Return(expectedAdmin, nil)

	// Call the method you want to test
	resultAdmin, err := adminService.FindByUsername(nil, "johndoe")

	// Assert that the FindByUsername method was called with the expected parameters
	adminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultAdmin)
	assert.Equal(t, expectedAdmin, resultAdmin)
}

func TestAdminFindByUsername_AdminNotFound(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	adminRepoMock := new(mocks.AdminRepository)

	// Create an instance of the service with the mock
	adminService := NewAdminService(adminRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindByUsername method when admin is not found
	adminRepoMock.On("FindByUsername", "nonexistentUser").Return(nil, fmt.Errorf("admin not found"))

	// Call the method you want to test
	resultAdmin, err := adminService.FindByUsername(nil, "nonexistentUser")

	// Assert that the FindByUsername method was called with the expected parameters
	adminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultAdmin)
	assert.EqualError(t, err, "admin not found")
}

func TestDeleteAdmin_Success(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	adminRepoMock := new(mocks.AdminRepository)

	// Create an instance of the service with the mock
	adminService := NewAdminService(adminRepoMock, mockValidator, mockPasswordHandler)

	adminRepoMock.On("FindById", mock.Anything).Return(&domain.Admin{}, nil)
	// Set up expectations for the DeleteAdmin method when admin is successfully deleted
	adminRepoMock.On("Delete", 1).Return(nil)

	// Call the method you want to test
	err := adminService.DeleteAdmin(nil, 1)

	// Assert that the DeleteAdmin method was called with the expected parameters
	adminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
}

func TestDeleteAdmin_AdminNotFound(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	adminRepoMock := new(mocks.AdminRepository)

	// Create an instance of the service with the mock
	adminService := NewAdminService(adminRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindById method when admin is not found
	adminRepoMock.On("FindById", mock.Anything).Return((*domain.Admin)(nil), fmt.Errorf("admin not found"))

	// Call the method you want to test
	err := adminService.DeleteAdmin(nil, 1)

	// Assert that the FindById method was called with the expected parameters
	adminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.EqualError(t, err, "admin not found")
}

func TestDeleteAdmin_ErrorDeletingAdmin(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	adminRepoMock := new(mocks.AdminRepository)

	// Create an instance of the service with the mock
	adminService := NewAdminService(adminRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindById method when admin is found
	adminRepoMock.On("FindById", mock.Anything).Return(&domain.Admin{}, nil)

	// Set up expectations for the DeleteAdmin method when there is an error
	adminRepoMock.On("Delete", 1).Return(fmt.Errorf("error deleting admin"))

	// Call the method you want to test
	err := adminService.DeleteAdmin(nil, 1)

	// Assert that the FindById and DeleteAdmin methods were called with the expected parameters
	adminRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.EqualError(t, err, "error deleting Admin: error deleting admin")
}
