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

	Group := e.Group("api/v1/productType")

	Group.POST("", ProductTypeHandler.CreateProductTypeHandler)
	Group.GET("", ProductTypeHandler.GetProductTypesHandler)
	Group.GET("/:id", ProductTypeHandler.GetProductTypeHandler)
	Group.GET("/search/:name", ProductTypeHandler.GetProductTypeByName)
	Group.PUT("/:id", ProductTypeHandler.UpdateProductTypeHandler)
	Group.DELETE("/:id", ProductTypeHandler.DeleteProductTypeHandler)
}
