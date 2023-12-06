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

type PaymentTypeService interface {
	CreatePaymentType(ctx echo.Context, request web.PaymentTypeRequest) (*domain.PaymentType, error)
	UpdatePaymentType(ctx echo.Context, request web.PaymentTypeRequest, id int) (*domain.PaymentType, error)
	FindById(ctx echo.Context, id int) (*domain.PaymentType, error)
	FindByName(ctx echo.Context, name string) (*domain.PaymentType, error)
	FindAll(ctx echo.Context) ([]domain.PaymentType, error)
	DeletePaymentType(ctx echo.Context, id int) error
}

type PaymentTypeServiceImpl struct {
	PaymentTypeRepository repository.PaymentTypeRepository
	Validate              *validator.Validate
}

func NewPaymentTypeService(repository repository.PaymentTypeRepository, validate *validator.Validate) *PaymentTypeServiceImpl {
	return &PaymentTypeServiceImpl{
		PaymentTypeRepository: repository,
		Validate:              validate,
	}
}

func (service *PaymentTypeServiceImpl) CreatePaymentType(ctx echo.Context, request web.PaymentTypeRequest) (*domain.PaymentType, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	paymentType := req.PaymentTypeRequestToPaymentTypeDomain(request)

	ExistingpaymentType, _ := service.PaymentTypeRepository.FindByName(request.TypeName)
	if ExistingpaymentType != nil {
		return nil, fmt.Errorf("payment type is already exists")
	}

	result, err := service.PaymentTypeRepository.Create(paymentType)

	if err != nil {
		return nil, fmt.Errorf("error creating payment type %s", err.Error())
	}

	return result, nil
}

func (service *PaymentTypeServiceImpl) UpdatePaymentType(ctx echo.Context, request web.PaymentTypeRequest, id int) (*domain.PaymentType, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingPaymentType, _ := service.PaymentTypeRepository.FindById(id)
	if existingPaymentType == nil {
		return nil, fmt.Errorf("payment type not found")
	}

	paymentType := req.PaymentTypeRequestToPaymentTypeDomain(request)
	if existingPaymentType.TypeName != paymentType.TypeName {
		existingPaymentType, _ := service.PaymentTypeRepository.FindByName(paymentType.TypeName)
		if existingPaymentType != nil {
			return nil, fmt.Errorf("payment type name is already exists")
		}
	}
	fmt.Println(paymentType)
	result, err := service.PaymentTypeRepository.Update(paymentType, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating data payment type: %s", err.Error())
	}

	return result, nil
}

func (service *PaymentTypeServiceImpl) FindById(ctx echo.Context, id int) (*domain.PaymentType, error) {
	existingPaymentType, _ := service.PaymentTypeRepository.FindById(id)
	if existingPaymentType == nil {
		return nil, fmt.Errorf("payment type not found")
	}

	return existingPaymentType, nil
}

func (service *PaymentTypeServiceImpl) FindAll(ctx echo.Context) ([]domain.PaymentType, error) {
	paymentType, err := service.PaymentTypeRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("payment type not found")
	}

	return paymentType, nil
}

func (service *PaymentTypeServiceImpl) FindByName(ctx echo.Context, name string) (*domain.PaymentType, error) {
	paymentType, _ := service.PaymentTypeRepository.FindByName(name)
	if paymentType == nil {
		return nil, fmt.Errorf("payment type not found")
	}

	return paymentType, nil
}

func (service *PaymentTypeServiceImpl) DeletePaymentType(ctx echo.Context, id int) error {
	existingPaymentType, _ := service.PaymentTypeRepository.FindById(id)
	if existingPaymentType == nil {
		return fmt.Errorf("payment type not found")
	}

	err := service.PaymentTypeRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting payment type: %s", err)
	}

	return nil
}
