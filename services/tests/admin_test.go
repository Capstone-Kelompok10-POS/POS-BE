package tests

// import (
// 	"qbills/mocks"
// 	"qbills/models/domain"
// 	"qbills/models/web"
// 	"qbills/services"
// 	"testing"

// 	"github.com/golang/mock/gomock"
// 	"github.com/labstack/echo/v4"
// 	"github.com/stretchr/testify/assert"
// )

// func TestCreateAdmin(t *testing.T) {
// 	ctrl := gomock.NewController(t)
// 	defer ctrl.Finish()

// 	// Mock AdminRepository
// 	mockAdminRepo := mocks.NewAdminRepository(ctrl)

// 	// Mock PasswordHandler
// 	mockPasswordHandler := mocks.NewPasswordHandler(ctrl)

// 	// Mock AdminService
// 	adminService := services.NewAdminServiceImpl(mockAdminRepo, mockPasswordHandler)

// 	// Mock echo.Context
// 	mockContext := new(echo.Echo).NewContext(nil, nil)

// 	// Mock data
// 	mockRequest := web.AdminCreateRequest{
// 		SuperAdminID: 1,
// 		FullName:     "Anggi",
// 		Username:     "anggiafrika",
// 		Password:     "123anggikebali",
// 	}

// 	// Expectations: Call to HashPassword on PasswordHandler will occur
// 	mockPasswordHandler.EXPECT().HashPassword(gomock.Any()).Return("hashedPassword", nil)

// 	// Expectations: Call to FindByUsername on AdminRepository will return nil
// 	mockAdminRepo.EXPECT().FindByUsername(gomock.Any(), gomock.Any()).Return(nil)

// 	// Expectations: Call to Create on AdminRepository will occur
// 	mockAdminRepo.EXPECT().Create(gomock.Any()).Return(&domain.Admin{}, nil)

// 	// Call the function under test
// 	admin, err := adminService.CreateAdmin(mockContext, mockRequest)

// 	// Check the result
// 	assert.NoError(t, err)
// 	assert.NotNil(t, admin)

// 	// Assert that the expected calls were made
// 	mockPasswordHandler.EXPECT().HashPassword(gomock.Any()).Times(1)
// 	mockAdminRepo.EXPECT().FindByUsername(gomock.Any(), gomock.Any()).Times(1)
// 	mockAdminRepo.EXPECT().Create(gomock.Any()).Times(1)
// }