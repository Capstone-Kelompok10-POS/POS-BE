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

func TestCreateCashier_Success(t *testing.T) {
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	cashierService := NewCashierService(mockRepository, mockValidator, mockPasswordHandler)

	mockRepository.On("FindByUsername", mock.Anything).Return(nil, nil)
	mockPasswordHandler.On("HashPassword", mock.Anything).Return("hashedPassword", nil)
	mockRepository.On("Create", mock.Anything).Return(&domain.Cashier{}, nil)

	ctx := echo.New().NewContext(nil, nil)
	// Mock data
	_, err := cashierService.CreateCashier(ctx, web.CashierCreateRequest{
			AdminID: 1,
			Fullname:     "Acek",
			Username:     "acekasik",
			Password:     "chawni123",
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

	cashierService := NewCashierService(mockRepository, mockValidator, mockPasswordHandler)

	// Simulasi cashier dengan username yang sudah ada
	existingCashier := &domain.Cashier{ID: 1, Username: "acekasik"}
	mockRepository.On("FindByUsername", "acekasik").Return(existingCashier, nil)

	ctx := echo.New().NewContext(nil, nil)
	// Mock data
	_, err := cashierService.CreateCashier(ctx, web.CashierCreateRequest{
		AdminID: 1,
		Fullname:     "Acek",
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

func TestCreateCashier_ValidationError(t *testing.T) {
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	cashierService := NewCashierService(mockRepository, mockValidator, mockPasswordHandler)


	ctx := echo.New().NewContext(nil, nil)

	// Mock data
	_, err := cashierService.CreateCashier(ctx, web.CashierCreateRequest{
		AdminID: 1,
		Fullname:     "Acek",
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

func TestCreateCashier_Failure(t *testing.T) {
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	cashierService := NewCashierService(mockRepository, mockValidator, mockPasswordHandler)

	// Set up expected calls
	mockRepository.On("FindByUsername", mock.Anything).Return(nil, nil)
	mockPasswordHandler.On("HashPassword", mock.Anything).Return("hashedPassword", nil)
	mockRepository.On("Create", mock.Anything).Return(nil, errors.New("failed to create cashier"))

	ctx := echo.New().NewContext(nil, nil)
	// Mock data
	cashier, err := cashierService.CreateCashier(ctx, web.CashierCreateRequest{
		AdminID: 1,
		Fullname:     "Acek",
		Username:     "acekasik",
		Password:     "chawni123",
	})

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, cashier)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestLoginCashier_Success(t *testing.T) {
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	cashierService := NewCashierService(mockRepository, mockValidator, mockPasswordHandler)
	cashierRequest := web.CashierLoginRequest{
		Username: "vilanatasya",
		Password: "vila12345",
	}

	mockRepository.On("FindByUsername", cashierRequest.Username).Return(&domain.Cashier{ID:1, Password: "hashedPassword"}, nil)
	
	mockPasswordHandler.On("ComparePassword", "hashedPassword", cashierRequest.Password).Return(nil)

	ctx := echo.New().NewContext(nil, nil)
	result, err := cashierService.LoginCashier(ctx, cashierRequest)
	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}



func TestLoginCashier_Failure(t *testing.T) {
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	cashierService := NewCashierService(mockRepository, mockValidator, mockPasswordHandler)
	cashierRequest := web.CashierLoginRequest{
		Username: "vilanatasya",
		Password: "vila12345",
	}

	mockRepository.On("FindByUsername", cashierRequest.Username).Return(&domain.Cashier{ID:1, Password: "hashedPassword"}, nil)
	
	mockPasswordHandler.On("ComparePassword", "hashedPassword", cashierRequest.Password).Return(errors.New("invalid username or password"))

	ctx := echo.New().NewContext(nil, nil)
	result, err := cashierService.LoginCashier(ctx, cashierRequest)
	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, result)

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}


func TestLoginCashier_UsernameOrPasswordError(t *testing.T) {
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	cashierService := NewCashierService(mockRepository, mockValidator, mockPasswordHandler)
	cashierRequest := web.CashierLoginRequest{
		Username: "vilanatasya",
		Password: "vila12345",
	}

	mockRepository.On("FindByUsername", cashierRequest.Username).Return(nil, errors.New("invalid username or password"))
	ctx := echo.New().NewContext(nil, nil)
	_ , err := cashierService.LoginCashier(ctx, cashierRequest)

	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "invalid username or password")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestLoginCashier_ValidationError(t *testing.T) {
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	cashierService := NewCashierService(mockRepository, mockValidator, mockPasswordHandler)
	cashierRequest := web.CashierLoginRequest{
		Username: "vilanatasya",
		Password: "",
	}

	ctx := echo.New().NewContext(nil, nil)
	_ , err := cashierService.LoginCashier(ctx, cashierRequest)
	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")

	// Assert that the expected calls were made
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestUpdateCashier_Success(t *testing.T) {
	// Create a new instance of the mock repository, validator, and password handler
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	// Create an instance of the cashier service with the mocks
	cashierService := NewCashierService(mockRepository, mockValidator, mockPasswordHandler)

	cashierID := 1
	cashierId := uint(cashierID)
	cashierUpdateRequest := web.CashierUpdateRequest{
		Fullname: "Acek",
		Username: "acekasik",
		Password: "chawni123",
	}

	// Set up the expected return values from the mock repository
	mockRepository.On("FindById", cashierID).Return(&domain.Cashier{}, nil)
	mockRepository.On("FindByUsername", mock.Anything).Return(nil, nil)
	mockPasswordHandler.On("HashPassword", mock.Anything).Return("hashedPassword", nil)

	// Set up the expected call for the Update method
	mockRepository.On("Update", mock.Anything, cashierID).Return(&domain.Cashier{ID: cashierId, Fullname: cashierUpdateRequest.Fullname}, nil)

	updatedCashier, err := cashierService.UpdateCashier(nil, cashierUpdateRequest, cashierID)

	// Assert the result
	assert.NoError(t, err)
	assert.NotNil(t, updatedCashier)
	assert.Equal(t, cashierUpdateRequest.Fullname, updatedCashier.Fullname)

	// Assert that the expected calls were made to the mock repository
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}


func TestUpdateCashier_ValidationError(t *testing.T) {
	// Create a new instance of the mock repository, validator, and password handler
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	// Create an instance of the cashier service with the mocks
	cashierService := NewCashierService(mockRepository, mockValidator, mockPasswordHandler)

	cashierID := 1
	cashierUpdateRequest := web.CashierUpdateRequest{
		Fullname: "Acek",
		Username: "acekasik",
		Password: "123",
	}

	_ , err := cashierService.UpdateCashier(nil, cashierUpdateRequest, cashierID)


	// Assert the result
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")

	// Assert that the expected calls were made to the mock repository
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestUpdateCashier_CashierNotFound(t *testing.T) {
	// Create a new instance of the mock repository, validator, and password handler
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	// Create an instance of the cashier service with the mocks
	cashierService := NewCashierService(mockRepository, mockValidator, mockPasswordHandler)

	cashierID := 1
	cashierUpdateRequest := web.CashierUpdateRequest{
		Fullname: "Acek",
		Username: "acekasik",
		Password: "chawni123",
	}

	// Set up the mock repository to return CashierNotFound error
	mockRepository.On("FindById", cashierID).Return(nil, errors.New("Cashier not found"))

	// Call the UpdateCashier method
	updatedCashier, err := cashierService.UpdateCashier(nil, cashierUpdateRequest, cashierID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, updatedCashier)

	// Assert that the expected calls were made to the mock repository
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}

func TestUpdateCashier_Failure(t *testing.T) {
	// Create a new instance of the mock repository, validator, and password handler
	mockRepository := new(mocks.CashierRepository)
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)

	// Create an instance of the cashier service with the mocks
	cashierService := NewCashierService(mockRepository, mockValidator, mockPasswordHandler)

	cashierID := 1
	cashierId := uint(cashierID)
	cashierUpdateRequest := web.CashierUpdateRequest{
		Fullname: "Acek",
		Username: "acekasik",
		Password: "chawni123",
	}

	// Set up the expected return values from the mock repository
	mockRepository.On("FindById", cashierID).Return(&domain.Cashier{}, nil)
	mockRepository.On("FindByUsername", mock.Anything).Return(nil, nil)
	mockPasswordHandler.On("HashPassword", mock.Anything).Return("hashedPassword", nil)

	// Set up the expected call for the Update method
	mockRepository.On("Update", mock.Anything, cashierID).Return(&domain.Cashier{ID: cashierId, Fullname: cashierUpdateRequest.Fullname}, errors.New("failed to update cashier"))

	updatedCashier, err := cashierService.UpdateCashier(nil, cashierUpdateRequest, cashierID)

	// Assert the result
	assert.Error(t, err)
	assert.Nil(t, updatedCashier)

	// Assert that the expected calls were made to the mock repository
	mockRepository.AssertExpectations(t)
	mockPasswordHandler.AssertExpectations(t)
}



func TestCashierFindById_Success(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	cashierRepoMock := new(mocks.CashierRepository)

	// Create an instance of the service with the mock
	cashierService := NewCashierService(cashierRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindById method
	expectedCashier := &domain.Cashier{
		ID:             1,
		AdminID:   1,
		Fullname:       "John Doe",
		Username:       "john.doe",
		Password:       "hashedPassword",
	}

	cashierRepoMock.On("FindById", 1).Return(expectedCashier, nil)

	// Call the method you want to test
	resultCashier, err := cashierService.FindById(nil, 1)

	// Assert that the FindById method was called with the expected parameters
	cashierRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultCashier)
	assert.Equal(t, expectedCashier, resultCashier)
}


func TestCashierFindById_CashierNotFound(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	cashierRepoMock := new(mocks.CashierRepository)

	// Create an instance of the service with the mock
	cashierService := NewCashierService(cashierRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindById method
	cashierRepoMock.On("FindById", 1).Return(nil, nil) // Simulating cashier not found

	// Call the method you want to test
	resultCashier, err := cashierService.FindById(nil, 1)

	// Assert that the FindById method was called with the expected parameters
	cashierRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultCashier)
	assert.EqualError(t, err, "cashier not found")
}


func TestCashierFindAll_Success(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	cashierRepoMock := new(mocks.CashierRepository)

	// Create an instance of the service with the mock
	cashierService := NewCashierService(cashierRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindAll method
	expectedCashiers := []domain.Cashier{
		{
			ID:           1,
			AdminID: 1,
			Fullname:     "John Doe",
			Username:     "john.doe",
			Password:     "hashedPassword",
		},
		{
			ID:           2,
			AdminID: 1,
			Fullname:     "Jane Doe",
			Username:     "jane.doe",
			Password:     "hashedPassword",
		},
	}

	cashierRepoMock.On("FindAll").Return(expectedCashiers, len(expectedCashiers), nil)

	// Call the method you want to test
	resultCashiers, totalCashier, err := cashierService.FindAll(nil)

	// Assert that the FindAll method was called with the expected parameters
	cashierRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultCashiers)
	assert.Equal(t, expectedCashiers, resultCashiers)
	assert.Equal(t, len(expectedCashiers), totalCashier)
}

func TestCashierFindAll_NoCashiersFound(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	cashierRepoMock := new(mocks.CashierRepository)

	// Create an instance of the service with the mock
	cashierService := NewCashierService(cashierRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindAll method when no cashiers are found
	cashierRepoMock.On("FindAll").Return(nil, 0, errors.New("cashier not found"))

	// Call the method you want to test
	resultCashiers, totalCashier, err := cashierService.FindAll(nil)

	// Assert that the FindAll method was called with the expected parameters
	cashierRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultCashiers, totalCashier)
	assert.EqualError(t, err, "cashier not found")
	assert.Equal(t, 0, totalCashier)
}


func TestCashierFindByUsername_Success(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	cashierRepoMock := new(mocks.CashierRepository)

	// Create an instance of the service with the mock
	cashierService := NewCashierService(cashierRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindByUsername method
	expectedCashier := &domain.Cashier{
		ID:           1,
		AdminID: 1,
		Fullname:     "John Doe",
		Username:     "johndoe",
		Password:     "hashedPassword",
	}

	cashierRepoMock.On("FindByUsername", "johndoe").Return(expectedCashier, nil)

	// Call the method you want to test
	resultCashier, err := cashierService.FindByUsername(nil, "johndoe")

	// Assert that the FindByUsername method was called with the expected parameters
	cashierRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultCashier)
	assert.Equal(t, expectedCashier, resultCashier)
}

func TestCashierFindByUsername_CashierNotFound(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	cashierRepoMock := new(mocks.CashierRepository)

	// Create an instance of the service with the mock
	cashierService := NewCashierService(cashierRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindByUsername method when cashier is not found
	cashierRepoMock.On("FindByUsername", "nonexistentUser").Return(nil, fmt.Errorf("cashier not found"))

	// Call the method you want to test
	resultCashier, err := cashierService.FindByUsername(nil, "nonexistentUser")

	// Assert that the FindByUsername method was called with the expected parameters
	cashierRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultCashier)
	assert.EqualError(t, err, "cashier not found")
}

func TestDeleteCashier_Success(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	cashierRepoMock := new(mocks.CashierRepository)

	// Create an instance of the service with the mock
	cashierService := NewCashierService(cashierRepoMock, mockValidator, mockPasswordHandler)

	cashierRepoMock.On("FindById", mock.Anything).Return(&domain.Cashier{}, nil)
	// Set up expectations for the DeleteCashier method when cashier is successfully deleted
	cashierRepoMock.On("Delete", 1).Return(nil)

	// Call the method you want to test
	err := cashierService.DeleteCashier(nil, 1)

	// Assert that the DeleteCashier method was called with the expected parameters
	cashierRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
}

func TestDeleteCashier_CashierNotFound(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	cashierRepoMock := new(mocks.CashierRepository)

	// Create an instance of the service with the mock
	cashierService := NewCashierService(cashierRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindById method when cashier is not found
	cashierRepoMock.On("FindById", mock.Anything).Return((*domain.Cashier)(nil), fmt.Errorf("cashier not found"))

	// Call the method you want to test
	err := cashierService.DeleteCashier(nil, 1)

	// Assert that the FindById method was called with the expected parameters
	cashierRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.EqualError(t, err, "cashier not found")
}

func TestDeleteCashier_ErrorDeletingCashier(t *testing.T) {
	// Create your mock
	mockValidator := validator.New()
	mockPasswordHandler := new(mocks.PasswordHandler)
	cashierRepoMock := new(mocks.CashierRepository)

	// Create an instance of the service with the mock
	cashierService := NewCashierService(cashierRepoMock, mockValidator, mockPasswordHandler)

	// Set up expectations for the FindById method when cashier is found
	cashierRepoMock.On("FindById", mock.Anything).Return(&domain.Cashier{}, nil)

	// Set up expectations for the DeleteCashier method when there is an error
	cashierRepoMock.On("Delete", 1).Return(fmt.Errorf("error deleting cashier"))

	// Call the method you want to test
	err := cashierService.DeleteCashier(nil, 1)

	// Assert that the FindById and DeleteCashier methods were called with the expected parameters
	cashierRepoMock.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.EqualError(t, err, "error deleting cashier: error deleting cashier")
}
