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
	Group.GET("", ProductHandler.FindPaginationProduct)
	Group.GET("/:id", ProductHandler.GetProductHandler)
	Group.GET("/all", ProductHandler.GetProductsHandler)
	Group.GET("/search/:name", ProductHandler.GetProductByNameHandler)
	Group.GET("/category/:productTypeID", ProductHandler.GetProductByCategoryHandler)
	Group.PUT("/:id", ProductHandler.UpdateProductHandler)
	Group.DELETE("/:id", ProductHandler.DeleteProductHandler)
}
