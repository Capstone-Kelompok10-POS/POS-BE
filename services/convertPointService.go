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

type ConvertPointService interface {
	CreateConvertPoint(ctx echo.Context, request web.ConvertPointRequest) (*domain.ConvertPoint, error)
	UpdateConvertPoint(ctx echo.Context, request web.ConvertPointRequest, id int) (*domain.ConvertPoint, error)
	FindById(ctx echo.Context, id int) (*domain.ConvertPoint, error)
	FindAll() ([]domain.ConvertPoint, error)
	DeleteConvertPoint(ctx echo.Context, id int) error
}

type ConvertPointServiceImpl struct {
	ConvertPointRepository repository.ConvertPointRepository
	Validate               *validator.Validate
}

func NewConvertPointService(ConvertPointRepository repository.ConvertPointRepository, validate *validator.Validate) *ConvertPointServiceImpl {
	return &ConvertPointServiceImpl{
		ConvertPointRepository: ConvertPointRepository,
		Validate:               validate,
	}
}

func (service *ConvertPointServiceImpl) CreateConvertPoint(ctx echo.Context, request web.ConvertPointRequest) (*domain.ConvertPoint, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	convertPoint := req.ConvertPointCreateRequestToConvertPointDomain(request)

	result, err := service.ConvertPointRepository.Create(convertPoint)
	if err != nil {
		return nil, fmt.Errorf("error when creating convertPoint: %s", err.Error())
	}

	return result, nil
}

func (service *ConvertPointServiceImpl) UpdateConvertPoint(ctx echo.Context, request web.ConvertPointRequest, id int) (*domain.ConvertPoint, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingConvertPoint, _ := service.ConvertPointRepository.FindById(id)
	if existingConvertPoint == nil {
		return nil, fmt.Errorf("convert point not found")
	}
	convertPoint := req.ConvertPointUpdateRequestToConvertPointDomain(request)

	result, err := service.ConvertPointRepository.Update(convertPoint, id)
	if err != nil {
		return nil, fmt.Errorf("error when update convertPoint: %s", err.Error())
	}
	return result, nil
}

func (service *ConvertPointServiceImpl) FindById(ctx echo.Context, id int) (*domain.ConvertPoint, error) {
	convertPoint, _ := service.ConvertPointRepository.FindById(id)
	if convertPoint == nil {
		return nil, fmt.Errorf("convert point not found")
	}

	return convertPoint, nil
}

func (service *ConvertPointServiceImpl) FindAll() ([]domain.ConvertPoint, error) {
	convertPoint, _ := service.ConvertPointRepository.FindAll()
	if convertPoint == nil {
		return nil, fmt.Errorf("convert point not found")
	}

	return convertPoint, nil
}

func (service *ConvertPointServiceImpl) DeleteConvertPoint(ctx echo.Context, id int) error {
	convertPoint, _ := service.ConvertPointRepository.FindById(id)
	if convertPoint == nil {
		return fmt.Errorf("convert point not found")
	}

	err := service.ConvertPointRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error when deleting convert point: %s", err)
	}

	return nil
}
