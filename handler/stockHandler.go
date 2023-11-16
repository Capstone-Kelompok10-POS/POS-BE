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

type StockHandler interface {
	CreateIncreaseStockHandler(ctx echo.Context) error
	CreateDecreaseStockHandler(ctx echo.Context) error
	FindAllStockHandler(ctx echo.Context) error
	FindByIdStockHandler(ctx echo.Context) error
	FindIncreaseStockHandler(ctx echo.Context) error
	FindDecreaseStockHandler(ctx echo.Context) error
}

type StockHandlerImpl struct {
	StockService services.StockService
}

func NewStockHandler(stockService services.StockService) StockHandler {
	return &StockHandlerImpl{StockService: stockService}
}

func (c *StockHandlerImpl) CreateIncreaseStockHandler(ctx echo.Context) error {
	request := new(web.StockCreateRequest)

	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := c.StockService.CreateIncreaseStockService(ctx, *request)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		default:
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to increase stock"))
		}
	}

	response := res.StockDomainToStockResponse(result)

	responseCustom := res.StockResponseToStockResponseCreate(response)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("success increase stock", responseCustom))
}

func (c *StockHandlerImpl) CreateDecreaseStockHandler(ctx echo.Context) error {
	request := new(web.StockCreateRequest)

	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := c.StockService.CreateDecreaseStockService(ctx, *request)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		case strings.Contains(err.Error(), "reduction amount is more than the stock amount"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("reduction amount is more than the stock amount"))
		default:
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to increase stock"))
		}
	}

	response := res.StockDomainToStockResponse(result)

	responseCustom := res.StockResponseToStockResponseCreate(response)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("success increase stock", responseCustom))
}

func (c *StockHandlerImpl) FindAllStockHandler(ctx echo.Context) error {
	result, err := c.StockService.FindAllStockService(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "update stock not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("stock not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("stock data error"))
	}

	response := res.ConvertStockResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get all update stock", response))

}

func (c *StockHandlerImpl) FindByIdStockHandler(ctx echo.Context) error {
	stockID := ctx.Param("id")
	stockIDInt, err := strconv.Atoi(stockID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid param id"))
	}

	stock, err := c.StockService.FindByIdStockService(ctx, uint(stockIDInt))
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMessage := "Get stock data error"

		if strings.Contains(err.Error(), "stocks not found") {
			statusCode = http.StatusNotFound
			errorMessage = "stock not found"
		}

		return ctx.JSON(statusCode, helpers.ErrorResponse(errorMessage))
	}

	response := res.StockDomainToStockResponse(stock)

	responseCustom := res.StockResponseToStockResponseCustom(response)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Success get stock by id", responseCustom))
}

func (c *StockHandlerImpl) FindIncreaseStockHandler(ctx echo.Context) error {
	result, err := c.StockService.FindIncreaseStockService(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "increase stock not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("increase stock not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("increase stock data error"))
	}

	response := res.ConvertStockResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get all Increase stock", response))

}

func (c *StockHandlerImpl) FindDecreaseStockHandler(ctx echo.Context) error {
	result, err := c.StockService.FindDecreaseStockService(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "decrease stock not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("decrease stock not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("decrease stock data error"))
	}

	response := res.ConvertStockResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get all decrease stock", response))

}
