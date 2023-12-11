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

type PaymentMethodService interface {
	CreatePaymentMethod(ctx echo.Context, request web.PaymentMethodRequest) (*domain.PaymentMethod, error)
	UpdatePaymentMethod(ctx echo.Context, request web.PaymentMethodRequest, id int) (*domain.PaymentMethod, error)
	FindById(ctx echo.Context, id int) (*domain.PaymentMethod, error)
	FindByName(ctx echo.Context, name string) (*domain.PaymentMethod, error)
	FindAll(ctx echo.Context) ([]domain.PaymentMethod, error)
	DeletePaymentMethod(ctx echo.Context, id int) error
}

type PaymentMethodServiceImpl struct {
	repository repository.PaymentMethodRepository
	validate   *validator.Validate
}

func NewPaymentMethodService(Repository repository.PaymentMethodRepository, validate *validator.Validate) *PaymentMethodServiceImpl {
	return &PaymentMethodServiceImpl{
		repository: Repository,
		validate:   validate,
	}
}

func (service *PaymentMethodServiceImpl) CreatePaymentMethod(ctx echo.Context, request web.PaymentMethodRequest) (*domain.PaymentMethod, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	paymentMethod := req.PaymentMethodRequestToPaymentMethodDomain(request)
	existingPaymentMethodName, _ := service.repository.FindByName(request.Name)
	if existingPaymentMethodName != nil {
		return nil, fmt.Errorf("payment method name already exists")
	}
	result, err := service.repository.Create(paymentMethod)

	if err != nil {
		return nil, fmt.Errorf("error creating payment method %s", err.Error())
	}

	return result, nil
}

func (service *PaymentMethodServiceImpl) UpdatePaymentMethod(ctx echo.Context, request web.PaymentMethodRequest, id int) (*domain.PaymentMethod, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingPaymentMethod, _ := service.repository.FindById(id)
	if existingPaymentMethod == nil {
		return nil, fmt.Errorf("payment method not found")
	}

	paymentMethod := req.PaymentMethodRequestToPaymentMethodDomain(request)
	if existingPaymentMethod.Name != paymentMethod.Name {
		existingPaymentMethodName, _ := service.repository.FindByName(paymentMethod.Name)
		if existingPaymentMethodName != nil {
			return nil, fmt.Errorf("payment method already exists")
		}
	}
	result, err := service.repository.Update(paymentMethod, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating data payment method: %s", err.Error())
	}

	return result, nil
}

func (service *PaymentMethodServiceImpl) FindById(ctx echo.Context, id int) (*domain.PaymentMethod, error) {
	existingPaymentMethod, _ := service.repository.FindById(id)
	if existingPaymentMethod == nil {
		return nil, fmt.Errorf("payment method not found")
	}

	return existingPaymentMethod, nil
}

func (service *PaymentMethodServiceImpl) FindAll(ctx echo.Context) ([]domain.PaymentMethod, error) {
	paymentMethod, err := service.repository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("payment method not found")
	}

	return paymentMethod, nil
}

func (service *PaymentMethodServiceImpl) FindByName(ctx echo.Context, name string) (*domain.PaymentMethod, error) {
	paymentMethod, _ := service.repository.FindByName(name)
	if paymentMethod == nil {
		return nil, fmt.Errorf("payment method not found")
	}

	return paymentMethod, nil
}

func (service *PaymentMethodServiceImpl) DeletePaymentMethod(ctx echo.Context, id int) error {
	existingPaymentMethod, _ := service.repository.FindById(id)
	if existingPaymentMethod == nil {
		return fmt.Errorf("payment method not found")
	}

	err := service.repository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting payment method: %s", err)
	}

	return nil
}
