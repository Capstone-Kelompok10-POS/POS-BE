package handler

import (
<<<<<<< Updated upstream
	"github.com/labstack/echo/v4"
=======
>>>>>>> Stashed changes
	"net/http"
	"qbills/models/web"
	"qbills/services"
	"qbills/utils/helpers"
	res "qbills/utils/response"
	"strconv"
	"strings"
<<<<<<< Updated upstream
=======

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
>>>>>>> Stashed changes
)

type ProductTypeHandler interface {
	CreateProductTypeHandler(ctx echo.Context) error
	UpdateProductTypeHandler(ctx echo.Context) error
	GetProductTypeHandler(ctx echo.Context) error
	GetProductTypesHandler(ctx echo.Context) error
	GetProductTypeByName(ctx echo.Context) error
	DeleteProductTypeHandler(ctx echo.Context) error
}

type ProductTypeHandlerImpl struct {
	ProductTypeService services.ProductTypeService
}

func NewProductTypeHandler(productTypeService services.ProductTypeService) ProductTypeHandler {
	return &ProductTypeHandlerImpl{ProductTypeService: productTypeService}
}

func (c *ProductTypeHandlerImpl) CreateProductTypeHandler(ctx echo.Context) error {
	productTypeRequest := new(web.ProductTypeCreate)

	if err := ctx.Bind(productTypeRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := c.ProductTypeService.CreateProductType(ctx, *productTypeRequest)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
<<<<<<< Updated upstream
		default:
=======
		case strings.Contains(err.Error(), "TypeName alpha"): 
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("type name is not valid must contain only alphabetical characters"))
		case strings.Contains(err.Error(), "TypeDescription alpha"): 
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("type description is not valid must contain only alphabetical characters"))		
		default:
			logrus.Error(err.Error())
>>>>>>> Stashed changes
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to create product type"))
		}
	}

	response := res.ProductTypeDomainToProductTypeResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("success create product type", response))
}

func (c *ProductTypeHandlerImpl) UpdateProductTypeHandler(ctx echo.Context) error {
	productTypeId := ctx.Param("id")

	productTypeIdInt, err := strconv.Atoi(productTypeId)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid param id"))
	}

	// Define a pointer to the ProductTypeUpdate struct
	productTypeRequest := &web.ProductTypeUpdate{}

	if err := ctx.Bind(productTypeRequest); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	productTypeIdUint := uint(productTypeIdInt)

	_, err = c.ProductTypeService.UpdateProductType(ctx, *productTypeRequest, productTypeIdUint)

	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "product type not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product type not found"))
		}
<<<<<<< Updated upstream
=======
		logrus.Error(err.Error())
>>>>>>> Stashed changes
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("update product type error"))
	}

	result, err := c.ProductTypeService.FindById(ctx, productTypeIdUint)

	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response := res.ProductTypeDomainToProductTypeResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success update product type", response))
}

func (c *ProductTypeHandlerImpl) GetProductTypeHandler(ctx echo.Context) error {
	productTypeID := ctx.Param("id")
	productTypeIDInt, err := strconv.Atoi(productTypeID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid param id"))
	}

	productType, err := c.ProductTypeService.FindById(ctx, uint(productTypeIDInt))
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMessage := "Get product type data error"

		if strings.Contains(err.Error(), "product type not found") {
			statusCode = http.StatusNotFound
			errorMessage = "Product type not found"
		}
<<<<<<< Updated upstream

=======
		logrus.Error(err.Error())
>>>>>>> Stashed changes
		return ctx.JSON(statusCode, helpers.ErrorResponse(errorMessage))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Success get product type", res.ProductTypeDomainToProductTypeResponse(productType)))
}

func (c *ProductTypeHandlerImpl) GetProductTypesHandler(ctx echo.Context) error {
	result, err := c.ProductTypeService.FindAll(ctx)

	if err != nil {
		if strings.Contains(err.Error(), "product types not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product types not found"))
		}
<<<<<<< Updated upstream

=======
		logrus.Error(err.Error())
>>>>>>> Stashed changes
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get product types data error"))
	}

	response := res.ConvertProductTypeResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get all data product type", response))

}

func (c *ProductTypeHandlerImpl) GetProductTypeByName(ctx echo.Context) error {
	productTypeName := ctx.Param("name")

	result, err := c.ProductTypeService.FindByName(ctx, productTypeName)

	if err != nil {
		if strings.Contains(err.Error(), "product type not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product type not found"))
		}
<<<<<<< Updated upstream
=======
		logrus.Error(err.Error())
>>>>>>> Stashed changes
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get product type data by name error"))
	}

	response := res.ProductTypeDomainToProductTypeResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succsess get product type by name", response))
}

func (c *ProductTypeHandlerImpl) DeleteProductTypeHandler(ctx echo.Context) error {
	productTypeId := ctx.Param("id")
	productTypeIdInt, err := strconv.Atoi(productTypeId)
	productTypeIdUint := uint(productTypeIdInt)

	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	err = c.ProductTypeService.DeleteProductType(ctx, productTypeIdUint)

	if err != nil {
		if strings.Contains(err.Error(), "product type not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("product type not found"))
		}
<<<<<<< Updated upstream

=======
		logrus.Error(err.Error())
>>>>>>> Stashed changes
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete data product type error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succesfully delete data product type", nil))

}
