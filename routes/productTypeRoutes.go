package routes

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
)

func ProductTypeRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	productTypeRepository := repository.NewProductTypeRepository(db)
	productTypeService := services.NewProductTypeService(productTypeRepository, validate)
	ProductTypeHandler := handler.NewProductTypeHandler(productTypeService)

	e.POST("/productType", ProductTypeHandler.CreateProductTypeHandler)
	e.GET("/productTypes", ProductTypeHandler.GetProductTypesHandler)
	e.GET("/productType/:id", ProductTypeHandler.GetProductTypeHandler)
	e.PUT("/productType/:id", ProductTypeHandler.UpdateProductTypeHandler)
	e.DELETE("/productType/:id", ProductTypeHandler.DeleteProductTypeHandler)
}
