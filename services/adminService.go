package services

import (
	"fmt"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/repository"
	"qbills/utils/helpers"
	req "qbills/utils/request"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type AdminService interface {
	CreateAdmin(ctx echo.Context, request web.AdminCreateRequest) (*domain.Admin, error)
	LoginAdmin(ctx echo.Context, request web.AdminLoginRequest) (*domain.Admin, error)
	UpdateAdmin(ctx echo.Context, request web.AdminUpdateRequest, id int) (*domain.Admin, error)
	FindById(ctx echo.Context, id int) (*domain.Admin, error)
	FindAll(ctx echo.Context) ([]domain.Admin, error)
	FindByUsername(ctx echo.Context, name string) (*domain.Admin, error)
	DeleteAdmin(ctx echo.Context, id int) error
}

type AdminServiceImpl struct {
	AdminRepository repository.AdminRepository
	Validate        *validator.Validate
}

func NewAdminService(adminRepository repository.AdminRepository, validate *validator.Validate) *AdminServiceImpl {
	return &AdminServiceImpl{
		AdminRepository: adminRepository,
		Validate:        validate,
	}
}

func (service *AdminServiceImpl) CreateAdmin(ctx echo.Context, request web.AdminCreateRequest) (*domain.Admin, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingAdmin, _ := service.AdminRepository.FindByUsername(request.Username)
	if existingAdmin != nil {
		return nil, fmt.Errorf("username already exists")
	}

	admin := req.AdminCreateRequestToAdminDomain(request)

	admin.Password = helpers.HashPassword(admin.Password)
	result, err := service.AdminRepository.Create(admin)

	if err != nil {
		return nil, fmt.Errorf("error creating admin %s", err.Error())
	}

	return result, nil
}

func (service *AdminServiceImpl) LoginAdmin(ctx echo.Context, request web.AdminLoginRequest) (*domain.Admin, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingAdmin, err := service.AdminRepository.FindByUsername(request.Username)
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	admin := req.AdminLoginRequestToAdminDomain(request)

	err = helpers.ComparePassword(existingAdmin.Password, admin.Password)
	if err != nil {
		return nil, fmt.Errorf("invalid username or password")
	}

	return existingAdmin, nil

}

func (service *AdminServiceImpl) UpdateAdmin(ctx echo.Context, request web.AdminUpdateRequest, id int) (*domain.Admin, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingAdmin, _ := service.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	admin := req.AdminUpdateRequestToAdminDomain(request)
	if existingAdmin.Username != admin.Username {
		existingAdminUsername, _ := service.AdminRepository.FindByUsername(admin.Username)
		if existingAdminUsername != nil {
			return nil, fmt.Errorf("username already exists")
		}
	}
	
	admin.Password = helpers.HashPassword(admin.Password)
	result, err := service.AdminRepository.Update(admin, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating data admin: %s", err.Error())
	}

	return result, nil
}

func (service *AdminServiceImpl) FindById(ctx echo.Context, id int) (*domain.Admin, error) {
	existingAdmin, _ := service.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	return existingAdmin, nil
}

func (service *AdminServiceImpl) FindAll(ctx echo.Context) ([]domain.Admin, error) {
	admins, err := service.AdminRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("admins not found")
	}

	return admins, nil
}

func (service *AdminServiceImpl) FindByUsername(ctx echo.Context, name string) (*domain.Admin, error) {
	admin, _ := service.AdminRepository.FindByUsername(name)
	if admin == nil {
		return nil, fmt.Errorf("admin not found")
	}

	return admin, nil
}

func (service *AdminServiceImpl) DeleteAdmin(ctx echo.Context, id int) error {
	existingAdmin, _ := service.AdminRepository.FindById(id)
	if existingAdmin == nil {
		return fmt.Errorf("admin not found")
	}

	err := service.AdminRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting Admin: %s", err)
	}

	return nil
}
