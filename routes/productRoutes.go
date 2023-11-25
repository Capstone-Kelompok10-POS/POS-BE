package routes

import (
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
)

func ProductRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	productRepository := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepository, validate)
	ProductHandler := handler.NewProductHandler(productService)

	Group := e.Group("api/v1/product")

	Group.POST("", ProductHandler.CreateProductHandler)
	Group.GET("", ProductHandler.GetProductsHandler)
	Group.GET("/:id", ProductHandler.GetProductHandler)
	Group.GET("/pagination", ProductHandler.FindPaginationProduct)
	Group.GET("/search/:name", ProductHandler.GetProductByNameHandler)
	Group.PUT("/:id", ProductHandler.UpdateProductHandler)
	Group.DELETE("/:id", ProductHandler.DeleteProductHandler)
}
