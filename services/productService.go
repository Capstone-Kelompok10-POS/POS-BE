package services

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/repository"
	"qbills/utils/helpers"
	req "qbills/utils/request"
)

type ProductService interface {
	CreateProductService(ctx echo.Context, request web.ProductCreateRequest) (*domain.Product, error)
	UpdateProductService(ctx echo.Context, request web.ProductUpdateRequest, id uint) (*domain.Product, error)
	FindByIdProductService(ctx echo.Context, id uint) (*domain.Product, error)
	FindByNameProductService(ctx echo.Context, name string) ([]*domain.Product, error)
	FindAllProductService(ctx echo.Context) ([]domain.Product, error)
	DeleteProductService(ctx echo.Context, id uint) error
}

type ProductServiceImpl struct {
	ProductRepository repository.ProductRepository
	validate          *validator.Validate
}

func NewProductService(productRepository repository.ProductRepository, validate *validator.Validate) *ProductServiceImpl {
	return &ProductServiceImpl{
		ProductRepository: productRepository,
		validate:          validate,
	}
}

func (service *ProductServiceImpl) CreateProductService(ctx echo.Context, request web.ProductCreateRequest) (*domain.Product, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	product := req.ProductCreateRequestToProductDomain(request)

	result, err := service.ProductRepository.Create(product)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *ProductServiceImpl) UpdateProductService(ctx echo.Context, request web.ProductUpdateRequest, id uint) (*domain.Product, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	exitingAdmin, err := service.ProductRepository.FindById(id)

	if exitingAdmin == nil {
		return nil, fmt.Errorf("product not found")
	}

	product := req.ProductUpdateRequestToProductDomain(request)

	result, err := service.ProductRepository.Update(product, id)

	if err != nil {
		return nil, fmt.Errorf("error when updating data product: %s", err.Error())
	}

	return result, nil

}

func (service *ProductServiceImpl) FindByIdProductService(ctx echo.Context, id uint) (*domain.Product, error) {
	result, err := service.ProductRepository.FindById(id)

	if err != nil {
		return nil, fmt.Errorf("product not found")
	}

	return result, nil
}

func (service *ProductServiceImpl) FindAllProductService(ctx echo.Context) ([]domain.Product, error) {
	product, err := service.ProductRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return product, nil
}

func (service *ProductServiceImpl) FindByNameProductService(ctx echo.Context, name string) ([]*domain.Product, error) {
	products, err := service.ProductRepository.FindByName(name)

	if err != nil {
		return nil, fmt.Errorf("failed to find products with the name %s: %s", name, err.Error())
	}

	return products, nil
}

func (service *ProductServiceImpl) DeleteProductService(ctx echo.Context, id uint) error {
	exitingProduct, _ := service.ProductRepository.FindById(id)
	if exitingProduct == nil {
		return fmt.Errorf("product not found")
	}

	err := service.ProductRepository.Delete(id)

	if err != nil {
		return err
	}

	return nil
}
