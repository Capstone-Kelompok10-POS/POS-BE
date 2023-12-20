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

func TestCreateConvertPoint_Success(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the handler with the mock service
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)

	mockConvertPointRepo.On("Create", mock.Anything, mock.Anything).Return(&domain.ConvertPoint{}, nil)

	// Create a request context (you can customize it based on your implementation)
	ctx := echo.New().NewContext(nil, nil)

	// Call the method you want to test
	_ , err := convertPointService.CreateConvertPoint(ctx, web.ConvertPointRequest{
		Point: 5,
		ValuePoint: 500,
	})

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)

	mockConvertPointRepo.AssertExpectations(t)
}


func TestCreateConvertPoint_ValidationError(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the handler with the mock service
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)

	// Create a request context (you can customize it based on your implementation)
	ctx := echo.New().NewContext(nil, nil)

	// Call the method you want to test
	_ , err := convertPointService.CreateConvertPoint(ctx, web.ConvertPointRequest{
		Point: 5,
	})

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")

	mockConvertPointRepo.AssertExpectations(t)
}

func TestCreateConvertPoint_Failure(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the handler with the mock service
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)

	mockConvertPointRepo.On("Create", mock.Anything, mock.Anything).Return(nil, errors.New("error when creating convertPoint"))

	// Create a request context (you can customize it based on your implementation)
	ctx := echo.New().NewContext(nil, nil)

	// Call the method you want to test
	convertPoint , err := convertPointService.CreateConvertPoint(ctx, web.ConvertPointRequest{
		Point: 5,
		ValuePoint: 500,
	})

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, convertPoint)

	mockConvertPointRepo.AssertExpectations(t)
}

func TestDeleteConvertPoint_Success(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the handler with the mock service
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)
	mockConvertPointRepo.On("FindById", mock.Anything).Return(&domain.ConvertPoint{}, nil)
	mockConvertPointRepo.On("Delete", 1).Return(nil)

	// Call the method you want to test
	err := convertPointService.DeleteConvertPoint(nil, 1)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)

	mockConvertPointRepo.AssertExpectations(t)
}

func TestDeleteConvertPoint_ConvertPointNotFound(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the handler with the mock service
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)
	mockConvertPointRepo.On("FindById", mock.Anything).Return(nil, errors.New("convert point not found"))

	// Call the method you want to test
	err := convertPointService.DeleteConvertPoint(nil, 1)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.EqualError(t, err, "convert point not found")

	mockConvertPointRepo.AssertExpectations(t)
}

func TestDeleteConvertPoint_ErrorDeletingConvertPoint(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the handler with the mock service
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)
	mockConvertPointRepo.On("FindById", mock.Anything).Return(&domain.ConvertPoint{}, nil)
	mockConvertPointRepo.On("Delete", 1).Return(errors.New("error when deleting convert point"))

	// Call the method you want to test
	err := convertPointService.DeleteConvertPoint(nil, 1)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.EqualError(t, err, "error when deleting convert point: error when deleting convert point")

	mockConvertPointRepo.AssertExpectations(t)
}

func TestConvertPointFindAll_Success(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the handler with the mock service
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)

	expectedConverPoint := []domain.ConvertPoint{
		{
			ID:           1,
			Point: 5,
			ValuePoint: 5000,
		},
		{
			ID:           2,
			Point: 10,
			ValuePoint: 10000,
		},
	}


	mockConvertPointRepo.On("FindAll").Return(expectedConverPoint, nil)

	resultConvertPoint, err := convertPointService.FindAll()

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, expectedConverPoint)
	assert.Equal(t, expectedConverPoint, resultConvertPoint)

	mockConvertPointRepo.AssertExpectations(t)
}

func TestConvertPointFindAll_NoConvertPointFound(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the handler with the mock service
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)

	mockConvertPointRepo.On("FindAll").Return(nil, errors.New("convert point not found"))

	resultConvertPoint, err := convertPointService.FindAll()

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultConvertPoint)
	assert.EqualError(t, err,"convert point not found")

	mockConvertPointRepo.AssertExpectations(t)
}

func TestConvertPointFindById_Success(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the handler with the mock service
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)

	expectedConverPoint := &domain.ConvertPoint{
			ID:           1,
			Point: 5,
			ValuePoint: 5000,
	}


	mockConvertPointRepo.On("FindById", 1).Return(expectedConverPoint, nil)
	ctx := echo.New().NewContext(nil, nil)
	resultConvertPoint, err := convertPointService.FindById(ctx, 1)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, expectedConverPoint)
	assert.Equal(t, expectedConverPoint, resultConvertPoint)

	mockConvertPointRepo.AssertExpectations(t)
}

func TestConvertPointFindById_ConvertPointNotFound(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the handler with the mock service
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)

	mockConvertPointRepo.On("FindById", 1).Return(nil, errors.New("convert point not found"))
	ctx := echo.New().NewContext(nil, nil)
	resultConvertPoint, err := convertPointService.FindById(ctx, 1)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultConvertPoint)
	assert.EqualError(t, err, "convert point not found")

	mockConvertPointRepo.AssertExpectations(t)
}


func TestUpdateConvertPoint_Success(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the service with the mock repository
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)

	// Set up expectations for the FindById method
	existingConvertPoint := &domain.ConvertPoint{
		ID:        1,
		Point:     100,
		ValuePoint: 500,
	}

	mockConvertPointRepo.On("FindById", 1).Return(existingConvertPoint, nil)

	// Set up expectations for the Update method
	updatedConvertPoint := &domain.ConvertPoint{
		ID:        1,
		Point:     150,
		ValuePoint: 750,
	}

	mockConvertPointRepo.On("Update", mock.Anything, 1).Return(updatedConvertPoint, nil)

	// Create a context for testing
	ctx := echo.New().NewContext(nil, nil)

	// Call the method you want to test
	resultConvertPoint, err := convertPointService.UpdateConvertPoint(ctx, web.ConvertPointRequest{
		Point:      150,
		ValuePoint: 750,
	}, 1)

	// Assert that the FindById and Update methods were called with the expected parameters
	mockConvertPointRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultConvertPoint)
	assert.Equal(t, updatedConvertPoint, resultConvertPoint)
}

func TestUpdateConvertPoint_ValidationError(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the service with the mock repository
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)

	convertPointID := 1
	updatedConvertPoint := web.ConvertPointRequest{
		Point:     150,
	}
	
	// Call the method you want to test
	_ , err := convertPointService.UpdateConvertPoint(nil, updatedConvertPoint, convertPointID)

	
	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")


	mockConvertPointRepo.AssertExpectations(t)
}

func TestUpdateConvertPoint_ConvertPointNotFound(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the service with the mock repository
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)

	// Set up expectations for the FindById method
	convertPointUpdateRequest := web.ConvertPointRequest{
		Point:     100,
		ValuePoint: 500,
	}

	mockConvertPointRepo.On("FindById", 1).Return(nil,errors.New("convert point not found"))


	// Call the method you want to test
	resultConvertPoint, err := convertPointService.UpdateConvertPoint(nil, convertPointUpdateRequest, 1)

	// Assert that the FindById and Update methods were called with the expected parameters
	mockConvertPointRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultConvertPoint)
}


func TestUpdateConvertPoint_Failure(t *testing.T) {
	mockValidator := validator.New()
	mockConvertPointRepo := new(mocks.ConvertPointRepository)

	// Create an instance of the service with the mock repository
	convertPointService := NewConvertPointService(mockConvertPointRepo, mockValidator)

	// Set up expectations for the FindById method
	convertPointUpdateRequest := web.ConvertPointRequest{
		Point:     100,
		ValuePoint: 500,
	}

	mockConvertPointRepo.On("FindById", 1).Return(&domain.ConvertPoint{}, nil)

	// Set up expectations for the Update method
	updatedConvertPoint := &domain.ConvertPoint{
		ID:        1,
		Point:     150,
		ValuePoint: 750,
	}

	mockConvertPointRepo.On("Update", mock.Anything, 1).Return(updatedConvertPoint, errors.New("error when update convertPoint"))


	// Call the method you want to test
	resultConvertPoint, err := convertPointService.UpdateConvertPoint(nil, convertPointUpdateRequest, 1)

	// Assert that the FindById and Update methods were called with the expected parameters
	mockConvertPointRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultConvertPoint)
}