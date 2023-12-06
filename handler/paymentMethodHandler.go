package handler

import (
	"fmt"
	"net/http"
	"qbills/models/web"
	"qbills/services"
	"qbills/utils/helpers"
	res "qbills/utils/response"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/sirupsen/logrus"
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
		if strings.Contains(err.Error(), "number") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("Payment Type ID is not valid must contain only number value"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("create payment method error"))
	}

	response := res.PaymentMethodDomainToPaymentMethodResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("successfully created payment method", response))
}

func (c *PaymentMethodHandlerImpl) UpdatePaymentMethodHandler(ctx echo.Context) error {
	paymentMethodId := ctx.Param("id")
	paymentMethodIdInt, err := strconv.Atoi(paymentMethodId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid payment method ID"))
	}

	paymentUpdateRequest := web.PaymentMethodRequest{}
	if err := ctx.Bind(&paymentUpdateRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid request payload"))
	}

	// Update Payment Method
	if _, err := c.service.UpdatePaymentMethod(ctx, paymentUpdateRequest, paymentMethodIdInt); err != nil {
		fmt.Print(err)
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("validation failed"))
		} else if strings.Contains(err.Error(), "payment method not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment method not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to update payment method"))
	}

	// Find and return the updated payment method
	results, err := c.service.FindById(ctx, paymentMethodIdInt)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to retrieve updated data"))
	}

	response := res.PaymentMethodDomainToPaymentMethodResponse(results)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully updated payment method", response))
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
