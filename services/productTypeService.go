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

type ProductTypeService interface {
	CreateProductType(ctx echo.Context, request web.ProductTypeCreate) (*domain.ProductType, error)
	UpdateProductType(ctx echo.Context, request web.ProductTypeUpdate, id uint) (*domain.ProductType, error)
	FindById(ctx echo.Context, id uint) (*domain.ProductType, error)
	FindAll(ctx echo.Context) ([]domain.ProductType, error)
	FindByName(ctx echo.Context, name string) (*domain.ProductType, error)
	DeleteProductType(ctx echo.Context, id uint) error
}

type ProductTypeImpl struct {
	ProductTypeRepository repository.ProductTypeRepository
	validate              *validator.Validate
}

func NewProductTypeService(productTypeRepository repository.ProductTypeRepository, validate *validator.Validate) *ProductTypeImpl {
	return &ProductTypeImpl{
		ProductTypeRepository: productTypeRepository,
		validate:              validate,
	}
}

func (service *ProductTypeImpl) CreateProductType(ctx echo.Context, request web.ProductTypeCreate) (*domain.ProductType, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	productType := req.ProductTypeCreateToProductTypeDomain(request)
	existingProductType, _ := service.ProductTypeRepository.FindByName(request.TypeName)
	if existingProductType != nil {
		return nil, fmt.Errorf("product type name already exists")
	}

	result, err := service.ProductTypeRepository.Create(productType)

	if err != nil {
		return nil, fmt.Errorf("error creating admin %s", err.Error())
	}

	return result, nil
}

func (service *ProductTypeImpl) UpdateProductType(ctx echo.Context, request web.ProductTypeUpdate, id uint) (*domain.ProductType, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingProductType, _ := service.ProductTypeRepository.FindById(id)
	if existingProductType == nil {
		return nil, fmt.Errorf("product type not found")
	}
	productType := req.ProductTypeUpdateToProductTypeDomain(request)

	if existingProductType.TypeName != productType.TypeName {
		existingProductType, _ := service.ProductTypeRepository.FindByName(productType.TypeName)
		if existingProductType != nil {
			return nil, fmt.Errorf("product type already exists")
		}
	}
	result, err := service.ProductTypeRepository.Update(productType, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating data product type: %s", err.Error())
	}

	return result, nil
}

func (service *ProductTypeImpl) FindById(ctx echo.Context, id uint) (*domain.ProductType, error) {
	existingProductType, _ := service.ProductTypeRepository.FindById(id)
	if existingProductType == nil {
		return nil, fmt.Errorf("product type not found")
	}

	return existingProductType, nil
}

func (service *ProductTypeImpl) FindAll(ctx echo.Context) ([]domain.ProductType, error) {
	productTypes, err := service.ProductTypeRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("product type not found")
	}

	return productTypes, nil
}

func (service *ProductTypeImpl) FindByName(ctx echo.Context, name string) (*domain.ProductType, error) {
	productType, _ := service.ProductTypeRepository.FindByName(name)
	if productType == nil {
		return nil, fmt.Errorf("product type not found")
	}

	return productType, nil
}

func (service *ProductTypeImpl) DeleteProductType(ctx echo.Context, id uint) error {
	exitingProductType, _ := service.ProductTypeRepository.FindById(id)
	if exitingProductType == nil {
		return fmt.Errorf("product type not found")
	}

	err := service.ProductTypeRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting Product Type: %s", err)
	}

	return nil
}
