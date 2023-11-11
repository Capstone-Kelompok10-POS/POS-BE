package routes

import (
	"os"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
	"qbills/utils/helpers/middleware"

	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func ProductTypeRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	productTypeRepository := repository.NewProductTypeRepository(db)
	productTypeService := services.NewProductTypeService(productTypeRepository, validate)
	ProductTypeHandler := handler.NewProductTypeHandler(productTypeService)


	adminGroup := e.Group("api/v1/product-type")

	adminGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	adminGroup.POST("", ProductTypeHandler.CreateProductTypeHandler, middleware.AuthMiddleware("Admin"))
	adminGroup.GET("", ProductTypeHandler.GetProductTypesHandler, middleware.AuthMiddleware("Admin"))
	adminGroup.GET("/:id", ProductTypeHandler.GetProductTypeHandler, middleware.AuthMiddleware("Admin"))
	adminGroup.PUT("/:id", ProductTypeHandler.UpdateProductTypeHandler, middleware.AuthMiddleware("Admin"))
	adminGroup.DELETE("/:id", ProductTypeHandler.DeleteProductTypeHandler, middleware.AuthMiddleware("Admin"))
}
