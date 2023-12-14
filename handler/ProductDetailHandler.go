package handler

import (
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

type ProductDetailHandler interface {
	CreateProductDetailHandler(ctx echo.Context) error
	UpdateProductDetailHandler(ctx echo.Context) error
	GetProductDetailHandler(ctx echo.Context) error
	GetProductDetailsHandler(ctx echo.Context) error
	DeleteProductDetailHandler(ctx echo.Context) error
}

type ProductDetailHandlerImpl struct {
	ProductDetailService services.ProductDetailService
}

func NewProductDetailHandler(ProductDetailService services.ProductDetailService) ProductDetailHandler {
	return &ProductDetailHandlerImpl{ProductDetailService: ProductDetailService}
}

func (c *ProductDetailHandlerImpl) CreateProductDetailHandler(ctx echo.Context) error {
	ProductDetailRequest := new(web.ProductDetailCreate)

	if err := ctx.Bind(ProductDetailRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := c.ProductDetailService.CreateProductDetail(ctx, *ProductDetailRequest)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		case strings.Contains(err.Error(), "Cannot add or update a child row: a foreign key constraint fails"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		case strings.Contains(err.Error(), "numeric"):
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("price is not valid must contain only numeric value"))
		default:
			logrus.Error(err.Error())
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to create product Detail"))
		}
	}

	response := res.ProductDetailDomainToProductDetailCreateResponses(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("success create product Detail", response))
}

func (c *ProductDetailHandlerImpl) UpdateProductDetailHandler(ctx echo.Context) error {
	ProductDetailId := ctx.Param("id")

	ProductDetailIdInt, err := strconv.Atoi(ProductDetailId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid param id"))
	}

	// Define a pointer to the ProductTypeUpdate struct
	ProductDetailRequest := &web.ProductDetailCreate{}

	if err := ctx.Bind(ProductDetailRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	productTypeIdUint := uint(ProductDetailIdInt)

	_, err = c.ProductDetailService.UpdateProductDetail(ctx, *ProductDetailRequest, productTypeIdUint)

	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "product detail not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product detail not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("update product detail error"))
	}

	result, err := c.ProductDetailService.FindById(ctx, productTypeIdUint)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response := res.ProductDetailDomainToProductDetailResponses(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success update product detail", response))
}

func (c *ProductDetailHandlerImpl) GetProductDetailHandler(ctx echo.Context) error {
	ProductDetailID := ctx.Param("id")
	ProductDetailIDInt, err := strconv.Atoi(ProductDetailID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid param id"))
	}

	ProductDetail, err := c.ProductDetailService.FindById(ctx, uint(ProductDetailIDInt))
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMessage := "Get product detail data error"

		if strings.Contains(err.Error(), "product detail not found") {
			statusCode = http.StatusNotFound
			errorMessage = "Product detail not found"
		}
		logrus.Error(err.Error())
		return ctx.JSON(statusCode, helpers.ErrorResponse(errorMessage))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Success get product type", res.ProductDetailDomainToProductDetailResponses(ProductDetail)))
}

func (c *ProductDetailHandlerImpl) GetProductDetailsHandler(ctx echo.Context) error {
	result, err := c.ProductDetailService.FindAll(ctx)

	if err != nil {

		if strings.Contains(err.Error(), "record not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product detail not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get product detail data error"))
	}

	response := res.ConvertProductDetailResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get all data product detail", response))

}

func (c *ProductDetailHandlerImpl) DeleteProductDetailHandler(ctx echo.Context) error {
	ProductDetailId := ctx.Param("id")
	ProductDetailIdInt, err := strconv.Atoi(ProductDetailId)
	ProductDetailIdUint := uint(ProductDetailIdInt)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	err = c.ProductDetailService.DeleteProductDetail(ctx, ProductDetailIdUint)

	if err != nil {
		if strings.Contains(err.Error(), "product detail not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product detail not found"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete data product detail error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succesfully delete data product detail", nil))

}
