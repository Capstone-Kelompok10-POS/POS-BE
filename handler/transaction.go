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

type TransactionHandler interface {
	CreateTransactionHandler(ctx echo.Context) error
	NotificationPayment(ctx echo.Context) error
	UpdateStatusTransactionPaymentHandler(ctx echo.Context) error
	GetTransactionHandler(ctx echo.Context) error
	GetTransactionsHandler(ctx echo.Context) error
	GetRecentTransactionsHandler(ctx echo.Context) error
	GetTransactionMonthlyHandler(ctx echo.Context) error
	GetTransactionYearlyHandler(ctx echo.Context) error
	GetTransactionDailyHandler(ctx echo.Context) error
	GetCashierTransactionsHandler(ctx echo.Context) error
	GetMembershipTransactionsHandler(ctx echo.Context) error
	FindPaginationTransaction(ctx echo.Context) error
}

type TransactionHandlerImpl struct {
	TransactionService services.TransactionService
}

func NewTransactionHandler(stockService services.TransactionService) TransactionHandler {
	return &TransactionHandlerImpl{TransactionService: stockService}
}

func (c *TransactionHandlerImpl) CreateTransactionHandler(ctx echo.Context) error {
	cashierId := middleware.ExtractTokenCashierId(ctx)
	transactionCreateRequest := web.TransactionCreateRequest{CashierID: uint(cashierId)}
	err :=ctx.Bind(&transactionCreateRequest)
	if err != nil{
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := c.TransactionService.CreateTransaction(transactionCreateRequest)
	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		case strings.Contains(err.Error(), "failed to decrease product stock"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("insufficient stock"))
		case strings.Contains(err.Error(), "failed to convert point"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("insufficient point membership"))
		default:
			logrus.Error(err.Error())
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to create transaction"))
		}
	}

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Success to create transaction", result))
}

func (c *TransactionHandlerImpl) NotificationPayment(ctx echo.Context) error {
	var notificationPayload map[string]interface{}
	err := ctx.Bind(&notificationPayload)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("failed to bind request"))
	}

	err = c.TransactionService.NotificationPayment(notificationPayload)
	if err != nil {
		if strings.Contains(err.Error(), "error when get order id") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("invoice not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to get notification payment"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Notification Update", nil))

}

func (c *TransactionHandlerImpl) GetTransactionHandler(ctx echo.Context) error {
	transactionID := ctx.Param("id")
	transactionIDInt, err := strconv.Atoi(transactionID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid param id"))
	}
	result, err := c.TransactionService.FindById(transactionIDInt)
	if err != nil {
		if strings.Contains(err.Error(), "transaction not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get transaction data error"))
	}
	response := res.TransactionDomainToTransactionResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Transaction", response))
}

func (c *TransactionHandlerImpl) GetTransactionsHandler(ctx echo.Context) error {
	transactions, totalTransactions, err := c.TransactionService.FindAllTransaction()
	if err != nil {
		if strings.Contains(err.Error(), "transaction not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get all transaction error"))
	}
	response := res.ConvertTransactionResponse(transactions) 

	return ctx.JSON(http.StatusOK, helpers.SuccessResponseWithTotal("Successfully Get Data Transaction", response, totalTransactions))

}

func (c *TransactionHandlerImpl) GetTransactionMonthlyHandler(ctx echo.Context) error {
	transactionsMonthly, err := c.TransactionService.FindByMonthly()
	if err != nil {
		if strings.Contains(err.Error(), "transaction not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction revenue monthly not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get all transaction monthly revenue error"))
	}
	response := res.ConvertTransactionMonthlyRevenueResponse(transactionsMonthly) 

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Transaction Revenue Monthly", response))

}

func (c *TransactionHandlerImpl) GetTransactionYearlyHandler(ctx echo.Context) error {
	transactionsYearly, err := c.TransactionService.FindByYearly()
	if err != nil {
		if strings.Contains(err.Error(), "transaction not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction revenue yearly not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get all transaction revenue yearly error"))
	}
	response := res.TransactionYearlyRevenueDomainToTransactionYearlyRevenueResponse(transactionsYearly) 

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Transaction Revenue Yearly", response))

}


func (c *TransactionHandlerImpl) GetTransactionDailyHandler(ctx echo.Context) error {
	transactionsDaily, err := c.TransactionService.FindByDaily()
	if err != nil {
		if strings.Contains(err.Error(), "error when get transaction daily") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction revenue daily not found"))
		}
		if strings.Contains(err.Error(), "transaction daily not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction revenue daily not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get transaction revenue Daily error"))
	}
	response := res.TransactionDailyDomainToTransactionDailyResponse(transactionsDaily) 

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Transaction Revenue Daily", response))

}

func (c *TransactionHandlerImpl) GetRecentTransactionsHandler(ctx echo.Context) error {
	transactions, err := c.TransactionService.FindRecentTransaction()
	if err != nil {
		if strings.Contains(err.Error(), "transaction not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get all transaction error"))
	}
	response := res.ConvertTransactionResponse(transactions) 

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Transaction", response))

}

func (c *TransactionHandlerImpl) GetCashierTransactionsHandler(ctx echo.Context) error {
	cashierId := middleware.ExtractTokenCashierId(ctx)
	transactions, err := c.TransactionService.FindByCashierIdTransaction(int(cashierId))
	if err != nil {
		if strings.Contains(err.Error(), "error when get transaction by cashier") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction by cashierId not found"))
		}
		if strings.Contains(err.Error(), "transaction by cashier not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction by cashierId not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get all transaction by cashier error"))
	}
	response := res.ConvertTransactionResponse(transactions) 

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Transaction By CashierID", response))

}

func (c *TransactionHandlerImpl) GetMembershipTransactionsHandler(ctx echo.Context) error {
	membershipID := ctx.Param("id")
	membershipIDInt, err := strconv.Atoi(membershipID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid param membershipId"))
	}
	transactions, err := c.TransactionService.FindByMembershipIdTransaction(membershipIDInt)
	if err != nil {
		if strings.Contains(err.Error(), "error when get transaction by membership") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction by membership Id not found"))
		}
		if strings.Contains(err.Error(), "transaction by membership not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction by membershipId not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get all transaction by membershipId error"))
	}
	response := res.ConvertTransactionResponse(transactions) 

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Transaction By MembershipID", response))

}

func (c *TransactionHandlerImpl) UpdateStatusTransactionPaymentHandler(ctx echo.Context) error {
	transactionPaymentUpdateRequest := web.TransactionPaymentUpdateRequest{}
	err :=ctx.Bind(&transactionPaymentUpdateRequest)
	if err != nil{
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}
	invoice := transactionPaymentUpdateRequest.Invoice
	result, err := c.TransactionService.ManualPayment(invoice)
	if err != nil {
		if strings.Contains(err.Error(), "transaction not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get transaction data error"))
	}
	response := res.TransactionDomainToTransactionResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Transaction", response))
}

func (c *TransactionHandlerImpl) FindPaginationTransaction(ctx echo.Context) error {
	orderBy := ctx.QueryParam("orderBy")
	QueryLimit := ctx.QueryParam("limit")
	QueryPage := ctx.QueryParam("page")

	response, meta, err := c.TransactionService.FindPaginationTransaction(orderBy, QueryLimit, QueryPage)
	if err != nil {

		if strings.Contains(err.Error(), "transaction is empty") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("transaction not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed get pagination transaction"))
	}

	tranactionResponse := res.ConvertTransactionResponse(response)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponseWithMeta("succesfully get data product", tranactionResponse, meta))

}