package routes

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
)

func ProductDetailRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	productDetailRepository := repository.NewProductDetailRepository(db)
	productDetailService := services.NewProductDetailService(productDetailRepository, validate)
	productDetailHandler := handler.NewProductDetailHandler(productDetailService)

	Group := e.Group("api/v1/productDetail")

	Group.POST("", productDetailHandler.CreateProductDetailHandler)
	Group.GET("", productDetailHandler.GetProductDetailsHandler)
	Group.GET("/:id", productDetailHandler.GetProductDetailHandler)
	Group.PUT("/:id", productDetailHandler.UpdateProductDetailHandler)
	Group.DELETE("/:id", productDetailHandler.DeleteProductDetailHandler)
}
