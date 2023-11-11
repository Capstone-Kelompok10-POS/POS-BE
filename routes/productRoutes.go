package routes

import (
	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
	"os"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
	"qbills/utils/helpers/middleware"
)

func ProductRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	productRepository := repository.NewProductRepository(db)
	productService := services.NewProductService(productRepository, validate)
	ProductHandler := handler.NewProductHandler(productService)

	Group := e.Group("api/v1/product")
	Group.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	Group.POST("", ProductHandler.CreateProductHandler, middleware.AuthMiddleware("Admin"))
	Group.GET("", ProductHandler.GetProductsHandler)
	Group.GET("/:id", ProductHandler.GetProductHandler)
	Group.GET("/search/:name", ProductHandler.GetProductByNameHandler)
	Group.PUT("/:id", ProductHandler.UpdateProductHandler, middleware.AuthMiddleware("Admin"))
	Group.DELETE("/:id", ProductHandler.DeleteProductHandler, middleware.AuthMiddleware("Admin"))
}
