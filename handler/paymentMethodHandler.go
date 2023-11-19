package handler

import (
	"github.com/labstack/echo/v4"
	"net/http"
	"qbills/models/web"
	"qbills/services"
	"qbills/utils/helpers"
	res "qbills/utils/response"
	"strconv"
	"strings"
)

type PaymentMethodHandler interface {
	CreatePaymentMethodHandler(ctx echo.Context) error
	UpdatePaymentMethodHandler(ctx echo.Context) error
	GetPaymentMethodHandler(ctx echo.Context) error
	GetPaymentMethodsHandler(ctx echo.Context) error
	GetPaymentMethodByNameHandler(ctx echo.Context) error
	DeletePaymentMethodHandler(ctx echo.Context) error
}

type PaymentMethodHandlerImpl struct {
	service services.PaymentMethodService
}

func NewPaymentMethodHandler(service services.PaymentMethodService) PaymentMethodHandler {
	return &PaymentMethodHandlerImpl{service: service}
}

func (c *PaymentMethodHandlerImpl) CreatePaymentMethodHandler(ctx echo.Context) error {
	paymentMethodCreateRequest := web.PaymentMethodRequest{}
	err := ctx.Bind(&paymentMethodCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid input"))
	}

	result, err := c.service.CreatePaymentMethod(ctx, paymentMethodCreateRequest)

	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("create payment method error"))
	}

	response := res.PaymentMethodDomainToPaymentMethodResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("successfully created payment method", response))
}

func (c *PaymentMethodHandlerImpl) UpdatePaymentMethodHandler(ctx echo.Context) error {
	paymentMethodId := ctx.Param("id")
	paymentMethodIdInt, err := strconv.Atoi(paymentMethodId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	paymentUpdateRequest := web.PaymentMethodRequest{}
	err = ctx.Bind(&paymentUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	_, err = c.service.UpdatePaymentMethod(ctx, paymentUpdateRequest, paymentMethodIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "payment method not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment method not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("update payment method error"))
	}
	results, err := c.service.FindById(ctx, paymentMethodIdInt)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response := res.PaymentMethodDomainToPaymentMethodResponse(results)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully updated data payment method", response))
}

func (c *PaymentMethodHandlerImpl) GetPaymentMethodHandler(ctx echo.Context) error {
	paymentMethodId := ctx.Param("id")
	paymentMethodIdInt, err := strconv.Atoi(paymentMethodId)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	result, err := c.service.FindById(ctx, paymentMethodIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "payment method not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment method not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get payment method data error"))
	}
	response := res.PaymentMethodDomainToPaymentMethodResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully get data payment method", response))
}

func (c *PaymentMethodHandlerImpl) GetPaymentMethodsHandler(ctx echo.Context) error {
	result, err := c.service.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "payment method not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment method not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get payment method data error"))
	}

	response := res.ConvertPaymentMethodResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully get all data payment method", response))
}

func (c *PaymentMethodHandlerImpl) GetPaymentMethodByNameHandler(ctx echo.Context) error {
	paymentMethodName := ctx.Param("name")

	result, err := c.service.FindByName(ctx, paymentMethodName)
	if err != nil {
		if strings.Contains(err.Error(), "payment method not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment method not found"))
		}
		return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Get payment method data by name error"))
	}
	response := res.PaymentMethodDomainToPaymentMethodResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succesfully get payment method data by name", response))
}

func (c *PaymentMethodHandlerImpl) DeletePaymentMethodHandler(ctx echo.Context) error {
	paymentMethodId := ctx.Param("id")
	paymentMethodIdInt, err := strconv.Atoi(paymentMethodId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	err = c.service.DeletePaymentMethod(ctx, paymentMethodIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "payment method not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment method not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete data payment method error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully delete data payment method", nil))
}
