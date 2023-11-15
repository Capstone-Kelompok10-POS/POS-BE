package services

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/repository"
	"qbills/utils/helpers"
	req2 "qbills/utils/request"
)

type StockService interface {
	CreateIncreaseStockService(ctx echo.Context, request web.StockCreateRequest) (*domain.Stock, error)
	CreateDecreaseStockService(ctx echo.Context, request web.StockCreateRequest) (*domain.Stock, error)
	FindAllStockService(ctx echo.Context) ([]domain.Stock, error)
	FindByIdStockService(ctx echo.Context, id uint) (*domain.Stock, error)
	FindIncreaseStockService(ctx echo.Context) ([]domain.Stock, error)
	FindDecreaseStockService(ctx echo.Context) ([]domain.Stock, error)
}

type StockServiceImpl struct {
	StockRepository repository.StockRepository
	validate        *validator.Validate
}

func NewStockService(repository repository.StockRepository, validate *validator.Validate) *StockServiceImpl {
	return &StockServiceImpl{
		StockRepository: repository,
		validate:        validate,
	}
}

func (service *StockServiceImpl) CreateIncreaseStockService(ctx echo.Context, request web.StockCreateRequest) (*domain.Stock, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	req := req2.StockCreateRequestToStockDomain(request)

	result, err := service.StockRepository.Create(req)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *StockServiceImpl) CreateDecreaseStockService(ctx echo.Context, request web.StockCreateRequest) (*domain.Stock, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	req := req2.StockCreateRequestToStockDomain(request)

	req.Stock = -req.Stock

	result, err := service.StockRepository.Create(req)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *StockServiceImpl) FindAllStockService(ctx echo.Context) ([]domain.Stock, error) {

	result, err := service.StockRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *StockServiceImpl) FindByIdStockService(ctx echo.Context, id uint) (*domain.Stock, error) {

	result, err := service.StockRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("Product not found")
	}

	return result, nil

}

func (service *StockServiceImpl) FindIncreaseStockService(ctx echo.Context) ([]domain.Stock, error) {

	result, err := service.StockRepository.FindIncreaseStock()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *StockServiceImpl) FindDecreaseStockService(ctx echo.Context) ([]domain.Stock, error) {

	result, err := service.StockRepository.FindDecreaseStock()

	if err != nil {
		return nil, err
	}

	return result, nil
}
