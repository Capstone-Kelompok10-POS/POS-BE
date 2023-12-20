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



func TestCreateMembership_Success(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	mockMembershipRepo.On("FindByPhoneNumber", mock.Anything).Return(nil, nil)
	// Set up expectations for the Create method
	request := web.MembershipCreateRequest{
		Name:        "Goldie",
		CashierID: 1,
		PhoneNumber: "62819174713815",
	}

	createdMembership := &domain.Membership{
		ID:          1,
		Name:        "Goldie",
		CashierID: 1,
		PhoneNumber: "62819174713815",
	}

	mockMembershipRepo.On("Create", mock.Anything).Return(createdMembership, nil)

	// Create a context for testing
	ctx := echo.New().NewContext(nil, nil)

	// Call the method you want to test
	resultMembership, err := membershipService.CreateMembership(ctx, request)

	// Assert that the Create method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultMembership)
	assert.Equal(t, createdMembership, resultMembership)
}

func TestCreateMembership_Failure(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	mockMembershipRepo.On("FindByPhoneNumber", mock.Anything).Return(nil, nil)
	// Set up expectations for the Create method
	request := web.MembershipCreateRequest{
		Name:        "Goldie",
		CashierID:   1,
		PhoneNumber: "62819174713815",
	}

	mockMembershipRepo.On("Create", mock.Anything).Return(nil, errors.New("error creating membership"))

	// Create a context for testing
	ctx := echo.New().NewContext(nil, nil)

	// Call the method you want to test
	resultMembership, err := membershipService.CreateMembership(ctx, request)

	// Assert that the Create method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultMembership)
	assert.EqualError(t, err, "error creating membership error creating membership")
}

func TestCreateMembership_ValidationError(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	request := web.MembershipCreateRequest{
		Name:        "Goldie",
	}

	// Call the method you want to test
	_ , err := membershipService.CreateMembership(nil, request)

	// Assert that the Create method was called with the expected parameters
	
	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")
	mockMembershipRepo.AssertExpectations(t)
}


func TestCreateMembership_PhoneNumberExists(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	// Set up expectations for the FindByPhoneNumber method
	request := web.MembershipCreateRequest{
		Name:        "Goldie",
		CashierID:   1,
		PhoneNumber: "62819174713815",
	}

	existingMembership := &domain.Membership{
		ID:          1,
		Name:        "ExistingMember",
		CashierID:   2,
		PhoneNumber: "62819174713815",
	}

	mockMembershipRepo.On("FindByPhoneNumber", request.PhoneNumber).Return(existingMembership, errors.New("phone_number already exists"))

	// Create a context for testing
	ctx := echo.New().NewContext(nil, nil)

	// Call the method you want to test
	resultMembership, err := membershipService.CreateMembership(ctx, request)

	// Assert that the FindByPhoneNumber method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultMembership)
	assert.EqualError(t, err, "phone_number already exists")  // Fix the error message here
}

func TestDeleteMembership_Success(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	mockMembershipRepo.On("FindById", mock.Anything).Return(&domain.Membership{}, nil)
	// Set up expectations for the Delete method when membership is successfully deleted
	mockMembershipRepo.On("Delete", 1).Return(nil)

	// Call the method you want to test
	err := membershipService.DeleteMembership(1)

	// Assert that the Delete method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
}


func TestDeleteMembership_MembershipNotFound(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	mockMembershipRepo.On("FindById", mock.Anything).Return((*domain.Membership)(nil), errors.New("membership not found"))

	// Call the method you want to test
	err := membershipService.DeleteMembership(1)

	// Assert that the Delete method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.EqualError(t, err, "membership not found")
}


func TestDeleteMembership_ErrorDeletingMembership(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	mockMembershipRepo.On("FindById", mock.Anything).Return(&domain.Membership{}, nil)
	// Set up expectations for the Delete method when membership is successfully deleted
	mockMembershipRepo.On("Delete", 1).Return(errors.New("error deleting membership"))

	// Call the method you want to test
	err := membershipService.DeleteMembership(1)

	// Assert that the Delete method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.EqualError(t, err, "error deleting membership: error deleting membership")
}


func TestFindAllMemberships_Success(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	// Set up expectations for the FindAll method when memberships are found
	expectedMemberships := []domain.Membership{
		{ID: 1, Name: "Gold", PhoneNumber: "628123456789"},
		{ID: 2, Name: "Silver", PhoneNumber: "6281234567890"},
	}
	mockMembershipRepo.On("FindAll").Return(expectedMemberships, len(expectedMemberships), nil)

	// Call the method you want to test
	resultMemberships, total, err := membershipService.FindAll()

	// Assert that the FindAll method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultMemberships)
	assert.Equal(t, expectedMemberships, resultMemberships)
	assert.Equal(t, len(expectedMemberships), total)
}


func TestFindAllMemberships_ErrorFindingMemberships(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	// Set up expectations for the FindAll method when an error occurs
	mockMembershipRepo.On("FindAll").Return(nil, 0, errors.New("error when get membership"))

	// Call the method you want to test
	resultMemberships, total, err := membershipService.FindAll()

	// Assert that the FindAll method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultMemberships)
	assert.EqualError(t, err, "error when get membership")
	assert.Equal(t, 0, total)
}

func TestFindMembershipById_Success(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	// Set up expectations for the FindById method when membership is found
	expectedMembership := &domain.Membership{
		ID:          1,
		Name:        "Goldie",
		CashierID:   1,
		PhoneNumber: "62819174713815",
	}

	mockMembershipRepo.On("FindById", 1).Return(expectedMembership, nil)

	// Call the method you want to test
	resultMembership, err := membershipService.FindById(1)

	// Assert that the FindById method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultMembership)
	assert.Equal(t, expectedMembership, resultMembership)
}

func TestFindMembershipById_NotFound(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	// Set up expectations for the FindById method when membership is not found
	mockMembershipRepo.On("FindById", 1).Return(nil, errors.New("membership not found"))

	// Call the method you want to test
	resultMembership, err := membershipService.FindById(1)

	// Assert that the FindById method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultMembership)
	assert.EqualError(t, err, "membership not found")
}

func TestFindMembershipByName_Success(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	// Set up expectations for the FindByName method when membership is found
	expectedMembership := &domain.Membership{
		ID:          1,
		Name:        "Goldie",
		CashierID:   1,
		PhoneNumber: "62819174713815",
	}

	mockMembershipRepo.On("FindByName", mock.Anything).Return(expectedMembership, nil)

	// Call the method you want to test
	resultMembership, err := membershipService.FindByName("goldie")

	// Assert that the FindByName method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultMembership)
	assert.Equal(t, expectedMembership, resultMembership)
}


func TestFindMembershipByName_EmptyName(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	// Set up expectations for the FindByName method with an empty string
	mockMembershipRepo.On("FindByName", "").Return(nil, errors.New("membership not found"))

	// Call the method you want to test with an empty name
	resultMembership, err := membershipService.FindByName("")

	// Assert that the FindByName method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultMembership)
	assert.EqualError(t, err, "membership not found")
}

func TestFindByPhoneNumber_Success(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	// Set up expectations for the FindByPhoneNumber method
	expectedMembership := &domain.Membership{
		ID:          1,
		Name:        "Goldie",
		CashierID:   1,
		PhoneNumber: "62819174713815",
	}

	mockMembershipRepo.On("FindByPhoneNumber", "62819174713815").Return(expectedMembership, nil)

	// Call the method you want to test with the phone number
	resultMembership, err := membershipService.FindByPhoneNumber("62819174713815")

	// Assert that the FindByPhoneNumber method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultMembership)
	assert.Equal(t, expectedMembership, resultMembership)
}

func TestFindByPhoneNumber_MembershipNotFound(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	// Set up expectations for the FindByPhoneNumber method when the membership is not found
	mockMembershipRepo.On("FindByPhoneNumber", "62819174713815").Return(nil, errors.New("membership not found"))

	// Call the method you want to test with the phone number
	resultMembership, err := membershipService.FindByPhoneNumber("62819174713815")

	// Assert that the FindByPhoneNumber method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultMembership)
	assert.EqualError(t, err, "membership not found")
}

func TestFindTopMember_Success(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	// Set up expectations for the FindTopMember method
	expectedMemberships := []domain.Membership{
		{
			ID:          1,
			Name:        "Goldie",
			CashierID:   1,
			TotalPoint: 5,
			PhoneNumber: "62819174713815",
		},
		{
			ID:          2,
			Name:        "Silver",
			CashierID:   2,
			TotalPoint: 10,
			PhoneNumber: "62819174713816",
		},
	}
	mockMembershipRepo.On("FindTopMember").Return(expectedMemberships, nil)

	// Call the method you want to test
	resultMemberships, err := membershipService.FindTopMember()

	// Assert that the FindTopMember method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultMemberships)
	assert.Equal(t, expectedMemberships, resultMemberships)
}

func TestFindTopMember_Error(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	// Set up expectations for the FindTopMember method when an error occurs
	mockMembershipRepo.On("FindTopMember").Return(nil, errors.New("error when get membership"))

	// Call the method you want to test
	resultMemberships, err := membershipService.FindTopMember()

	// Assert that the FindTopMember method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Nil(t, resultMemberships)
	assert.EqualError(t, err, "error when get membership")
}

func TestUpdateMembership_Success(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)
	mockMembershipRepo.On("FindById", mock.Anything).Return(&domain.Membership{}, nil)
	mockMembershipRepo.On("FindByPhoneNumber", mock.Anything).Return(nil, nil)

	// Set up expectations for the UpdateMembership method when the update is successful
	request := web.MembershipUpdateRequest{
		Name:        "UpdatedName",
		CashierID:   2,
		PhoneNumber: "62819174713816",
	}

	updatedMembership := &domain.Membership{
		ID:          1,
		Name:        "UpdatedName",
		CashierID:   2,
		PhoneNumber: "62819174713816",
	}

	mockMembershipRepo.On("Update", mock.Anything, 1).Return(updatedMembership, nil)

	// Create a context for testing
	ctx := echo.New().NewContext(nil, nil)

	// Call the method you want to test
	resultMembership, err := membershipService.UpdateMembership(ctx, request, 1)

	// Assert that the UpdateMembership method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.NoError(t, err)
	assert.NotNil(t, resultMembership)
	assert.Equal(t, updatedMembership, resultMembership)
}

func TestUpdateMembership_ValidationError(t *testing.T) {
	mockValidator := validator.New()
	mockMembershipRepo := new(mocks.MembershipRepository)

	// Create an instance of the service with the mock repository
	membershipService := NewMembershipService(mockMembershipRepo, mockValidator)

	// Create a context for testing
	ctx := echo.New().NewContext(nil, nil)

	// Call the method you want to test
	_ , err := membershipService.UpdateMembership(ctx, web.MembershipUpdateRequest{}, 1)

	// Assert that the FindById method was called with the expected parameters
	mockMembershipRepo.AssertExpectations(t)

	// Perform your assertions based on the test scenario
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "validation failed")

}

