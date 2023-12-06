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

	"github.com/sirupsen/logrus"
)

type PaymentTypeHandler interface {
	CreatePaymentTypeHandler(ctx echo.Context) error
	UpdatePaymentTypeHandler(ctx echo.Context) error
	GetPaymentTypeHandler(ctx echo.Context) error
	GetPaymentTypesHandler(ctx echo.Context) error
	GetPaymentTypeByNameHandler(ctx echo.Context) error
	DeletePaymentTypeHandler(ctx echo.Context) error
}
type PaymentTypeHandlerImpl struct {
	service services.PaymentTypeService
}

func NewPaymentTypeHandler(service services.PaymentTypeService) PaymentTypeHandler {
	return &PaymentTypeHandlerImpl{service: service}
}

func (c *PaymentTypeHandlerImpl) CreatePaymentTypeHandler(ctx echo.Context) error {
	paymentTypeCreateRequest := web.PaymentTypeRequest{}
	err := ctx.Bind(&paymentTypeCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid input"))
	}

	result, err := c.service.CreatePaymentType(ctx, paymentTypeCreateRequest)

	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "alpha") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("payment method typename is not valid must contain only alphabetical characters"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("create payment type error"))
	}

	response := res.PaymentTypeDomainToPaymentTypeRespone(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("successfully created payment type", response))
}

func (c *PaymentTypeHandlerImpl) UpdatePaymentTypeHandler(ctx echo.Context) error {
	paymentTypeId := ctx.Param("id")
	paymentTypeIdInt, err := strconv.Atoi(paymentTypeId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	paymentTypeUpdateRequest := web.PaymentTypeRequest{}
	err = ctx.Bind(&paymentTypeUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	_, err = c.service.UpdatePaymentType(ctx, paymentTypeUpdateRequest, paymentTypeIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "membership not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment type not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("update payment type error"))
	}
	results, err := c.service.FindById(ctx, paymentTypeIdInt)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response := res.PaymentTypeDomainToPaymentTypeRespone(results)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully updated data payment type", response))
}

func (c *PaymentTypeHandlerImpl) GetPaymentTypeHandler(ctx echo.Context) error {
	paymentTypeId := ctx.Param("id")
	paymentTypeIdInt, err := strconv.Atoi(paymentTypeId)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	result, err := c.service.FindById(ctx, paymentTypeIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "payment type not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment type not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get payment type data error"))
	}
	response := res.PaymentTypeDomainToPaymentTypeRespone(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully get data payment type", response))
}

func (c *PaymentTypeHandlerImpl) GetPaymentTypesHandler(ctx echo.Context) error {
	result, err := c.service.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "payment type not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment type not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get payment type data error"))
	}

	response := res.ConvertPaymentTypeResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully get all data payment type", response))
}

func (c *PaymentTypeHandlerImpl) GetPaymentTypeByNameHandler(ctx echo.Context) error {
	paymentTypeName := ctx.Param("name")

	result, err := c.service.FindByName(ctx, paymentTypeName)
	if err != nil {
		if strings.Contains(err.Error(), "payment type not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment type not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Get payment type data by name error"))
	}
	response := res.PaymentTypeDomainToPaymentTypeRespone(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succesfully get payment type data by name", response))
}

func (c *PaymentTypeHandlerImpl) DeletePaymentTypeHandler(ctx echo.Context) error {
	paymentTypeId := ctx.Param("id")
	paymentTypeIdInt, err := strconv.Atoi(paymentTypeId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	err = c.service.DeletePaymentType(ctx, paymentTypeIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "payment type not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("payment type not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete data payment type error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully delete data payment type", nil))
}
