package handler

import (
	"net/http"
	"qbills/models/web"
	"qbills/services"
	"qbills/utils/helpers"
	"qbills/utils/helpers/middleware"
	res "qbills/utils/response"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type CashierHandler interface {
	RegisterCashierHandler(ctx echo.Context) error
	LoginCashierHandler(ctx echo.Context) error
	UpdateCashierHandler(ctx echo.Context) error
	DeleteCashierHandler(ctx echo.Context) error
	GetCashierHandler(ctx echo.Context) error
	GetCashiersHandler(ctx echo.Context) error
	GetCashierByUsernameHandler(ctx echo.Context) error
}

type CashierHandlerImpl struct {
	CashierService services.CashierService
}

func NewCashierHandler(cashierService services.CashierService) CashierHandler {
	return &CashierHandlerImpl{CashierService: cashierService}
}

func (c *CashierHandlerImpl) RegisterCashierHandler(ctx echo.Context) error {
	cashierCreateRequest := web.CashierCreateRequest{}
	err := ctx.Bind(&cashierCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid input"))
	}

	result, err := c.CashierService.CreateCashier(ctx, cashierCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "username already exist") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("username already exist"))
		}
		if strings.Contains(err.Error(), "alphanum") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("username is not valid must contain only alphanumeric characters"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign up error"))
	}

	response := res.CashierDomainToCashierResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("successfully created account cashier", response))
}

func (c *CashierHandlerImpl) LoginCashierHandler(ctx echo.Context) error {
	cashierLoginRequest := web.CashierLoginRequest{}

	err := ctx.Bind(&cashierLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response, err := c.CashierService.LoginCashier(ctx, cashierLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}

		if strings.Contains(err.Error(), "invalid username or password") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid username or password"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign in error"))
	}

	cashierLoginResponse := res.CashierDomainToCashierLoginResponse(response)

	token, err := middleware.GenerateTokenCashier(response.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("gnerate jwt token error"))
	}

	cashierLoginResponse.Token = token

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success sign in", cashierLoginResponse))
}

func (c *CashierHandlerImpl) GetCashierHandler(ctx echo.Context) error {
	cashierId := ctx.Param("id")
	cashierIdInt, err := strconv.Atoi(cashierId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	result, err := c.CashierService.FindById(ctx, cashierIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "cashier not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("cashier not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get Cashier data error"))
	}
	response := res.CashierDomainToCashierResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully get data cashier", response))
}

func (c CashierHandlerImpl) GetCashierByUsernameHandler(ctx echo.Context) error {
	cashierName := ctx.Param("name")

	result, err := c.CashierService.FindByUsername(ctx, cashierName)
	if err != nil {
		if strings.Contains(err.Error(), "cashier not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("cashier not found"))
		}
    logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get cashier data by name error"))
	}
	response := res.CashierDomainToCashierResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully get cashier data by name", response))
}

func (c CashierHandlerImpl) GetCashiersHandler(ctx echo.Context) error {
	result, err := c.CashierService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "cashiers not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("cashiers not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get cashiers data error"))
	}

	response := res.ConvertCashierResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get All data cashiers", response))
}

func (c CashierHandlerImpl) UpdateCashierHandler(ctx echo.Context) error {
	cashierId := ctx.Param("id")
	cashierIdInt, err := strconv.Atoi(cashierId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	cashierUpdateRequest := web.CashierUpdateRequest{}
	err = ctx.Bind(&cashierUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	_, err = c.CashierService.UpdateCashier(ctx, cashierUpdateRequest, cashierIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "cashier not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("cashier not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("update cashier error"))
	}
	results, err := c.CashierService.FindById(ctx, cashierIdInt)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response := res.CashierDomainToCashierResponse(results)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully updated data admin", response))
}

func (c CashierHandlerImpl) DeleteCashierHandler(ctx echo.Context) error {
	cashierId := ctx.Param("id")
	cashierIdInt, err := strconv.Atoi(cashierId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	err = c.CashierService.DeleteCashier(ctx, cashierIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "cashier not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("cashier not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete data cashier error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully delete cashier", nil))
}
