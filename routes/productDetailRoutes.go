package routes

import (
	"os"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"

	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductDetailRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	productDetailRepository := repository.NewProductDetailRepository(db)
	productDetailService := services.NewProductDetailService(productDetailRepository, validate)
	productDetailHandler := handler.NewProductDetailHandler(productDetailService)

	detailGroup := e.Group("api/v1/product/detail")

	detailGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	detailGroup.POST("", productDetailHandler.CreateProductDetailHandler)
	detailGroup.GET("", productDetailHandler.GetProductDetailsHandler)
	detailGroup.GET("/:id", productDetailHandler.GetProductDetailHandler)
	detailGroup.PUT("/:id", productDetailHandler.UpdateProductDetailHandler)
	detailGroup.DELETE("/:id", productDetailHandler.DeleteProductDetailHandler)
}
