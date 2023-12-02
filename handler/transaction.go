package handler

import (
	"fmt"
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

type TransactionHandler interface {
	CreateTransactionHandler(ctx echo.Context) error
	GetTransactionHandler(ctx echo.Context) error
}

type TransactionHandlerImpl struct {
	TransactionService services.TransactionService
}

func NewTransactionHandler(stockService services.TransactionService) TransactionHandler {
	return &TransactionHandlerImpl{TransactionService: stockService}
}

func (c *TransactionHandlerImpl) CreateTransactionHandler(ctx echo.Context) error {
	cashier := middleware.ExtractTokenCashierId(ctx)
	transactionCreateRequest := web.TransactionCreateRequest{CashierID: uint(cashier)}
	err :=ctx.Bind(&transactionCreateRequest)
	if err != nil{
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}
	result, err := c.TransactionService.CreateTransaction(ctx, transactionCreateRequest)
	fmt.Print(err)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		case strings.Contains(err.Error(), "failed to decrease product stock"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("insufficient stock"))
		default:
			logrus.Error(err.Error())
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to create transaction"))
		}
	}

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Success to create transaction", result))
}

func (c *TransactionHandlerImpl) GetTransactionHandler(ctx echo.Context) error {
	transactionID := ctx.Param("id")
	transactionIDInt, err := strconv.Atoi(transactionID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid param id"))
	}
	result, err := c.TransactionService.FindById(ctx, transactionIDInt)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("admin not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get admin data error"))
	}
	response := res.TransactionDomainToTransactionResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Transaction", response))

}