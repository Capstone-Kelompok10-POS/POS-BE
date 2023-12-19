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

func ConvertPointRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	convertPointRepository := repository.NewConvertPointRepository(db)
	convertPointService := services.NewConvertPointService(convertPointRepository, validate)
	ConvertPointHandler := handler.NewConvertPointHandler(convertPointService)

	convertPointGroup := e.Group("api/v1/convert/point")

	convertPointGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	convertPointGroup.POST("", ConvertPointHandler.CreateConvertPointHandler, middleware.AuthMiddleware("Admin"))
	convertPointGroup.GET("/:id", ConvertPointHandler.GetConvertPointHandler)
	convertPointGroup.GET("s", ConvertPointHandler.GetAllConvertPointHandler)
	convertPointGroup.PUT("/:id", ConvertPointHandler.UpdateConvertPointHandler, middleware.AuthMiddleware("Admin"))
	convertPointGroup.DELETE("/:id", ConvertPointHandler.DeleteConvertPointHandler, middleware.AuthMiddleware("Admin"))
}