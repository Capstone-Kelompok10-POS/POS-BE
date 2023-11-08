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
)

type CashierHandler interface {
	LoginCashierHandler(ctx echo.Context) error
	UpdateCashierHandler(ctx echo.Context) error
}

type CashierHandlerImpl struct {
	CashierService services.CashierService
}

func NewCashierHandler(cashierService services.CashierService) CashierHandler {
	return &CashierHandlerImpl{CashierService: cashierService}
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

		if strings.Contains(err.Error(), "invalid email or password") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid email or password"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign in error"))
	}

	cashierLoginResponse := res.CashierDomainToCashierLoginResponse(response)

	token, err := middleware.GenerateTokenCashier(response.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("gnerate jwt token error"))
	}

	cashierLoginResponse.Token = token

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("success sign in", cashierLoginResponse))
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
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("update cashier error"))
	}
	results, err := c.CashierService.FindById(ctx, cashierIdInt)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response := res.CashierDomainToCashierResponse(results)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully updated data admin", response))
}