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

type ProductDetailService interface {
	CreateProductDetail(ctx echo.Context, request web.ProductDetailCreate) (*domain.ProductDetail, error)
	UpdateProductDetail(ctx echo.Context, request web.ProductDetailCreate, id uint) (*domain.ProductDetail, error)
	FindById(ctx echo.Context, id uint) (*domain.ProductDetail, error)
	FindAll(ctx echo.Context) ([]domain.ProductDetail, error)
	FindByProductId(ctx echo.Context, productId uint) ([]domain.ProductDetail, error)
	DeleteProductDetail(ctx echo.Context, id uint) error
}

type ProductDetailServiceImpl struct {
	ProductDetailRepository repository.ProductDetailRepository
	Validate                *validator.Validate
}

func NewProductDetailService(repository repository.ProductDetailRepository, validate *validator.Validate) *ProductDetailServiceImpl {
	return &ProductDetailServiceImpl{
		ProductDetailRepository: repository,
		Validate:                validate,
	}
}

func (service *ProductDetailServiceImpl) CreateProductDetail(ctx echo.Context, request web.ProductDetailCreate) (*domain.ProductDetail, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	ProductDetail := req.ProductDetailCreateToProductDomain(request)

	ProductDetail.TotalStock = 0

	result, err := service.ProductDetailRepository.Create(ProductDetail)

	if err != nil {
		return nil, fmt.Errorf("error creating product detail %s", err.Error())
	}

	return result, nil
}

func (service *ProductDetailServiceImpl) UpdateProductDetail(ctx echo.Context, request web.ProductDetailCreate, id uint) (*domain.ProductDetail, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingPaymentType, _ := service.ProductDetailRepository.FindById(id)
	if existingPaymentType == nil {
		return nil, fmt.Errorf("product detail not found")
	}

	productDetail := req.ProductDetailCreateToProductDomain(request)

	result, err := service.ProductDetailRepository.Update(productDetail, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating data product detail: %s", err.Error())
	}

	return result, nil
}

func (service *ProductDetailServiceImpl) FindById(ctx echo.Context, id uint) (*domain.ProductDetail, error) {
	existingProductDetail, _ := service.ProductDetailRepository.FindById(id)
	if existingProductDetail == nil {
		return nil, fmt.Errorf("product detail not found")
	}

	return existingProductDetail, nil
}

func (service *ProductDetailServiceImpl) FindAll(ctx echo.Context) ([]domain.ProductDetail, error) {
	productDetail, err := service.ProductDetailRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("product detail not found")
	}

	return productDetail, nil
}

func (service *ProductDetailServiceImpl) FindByProductId(ctx echo.Context, productId uint) ([]domain.ProductDetail, error) {
	productDetail, err := service.ProductDetailRepository.FindByProductId(productId)
	if err != nil {
		return nil, fmt.Errorf("product detail not found")
	}

	return productDetail, nil
}

func (service *ProductDetailServiceImpl) DeleteProductDetail(ctx echo.Context, id uint) error {
	existingProductDetail, _ := service.ProductDetailRepository.FindById(id)
	if existingProductDetail == nil {
		return fmt.Errorf("product detail not found")
	}

	err := service.ProductDetailRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting product detail: %s", err)
	}

	return nil
}
