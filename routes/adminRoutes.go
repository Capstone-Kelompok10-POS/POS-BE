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

func AdminRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	adminRepository := repository.NewAdminRepository(db)
	adminService := services.NewAdminService(adminRepository, validate)
	AdminHandler := handler.NewAdminHandler(adminService)

	adminGroup := e.Group("api/v1/admin")

	adminGroup.POST("/register", AdminHandler.RegisterAdminHandler)
	adminGroup.POST("/login", AdminHandler.LoginAdminHandler)

	adminGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	adminGroup.GET("/:id", AdminHandler.GetAdminHandler, middleware.AuthMiddleware("Admin"))
	adminGroup.GET("", AdminHandler.GetAdminsHandler, middleware.AuthMiddleware("Admin"))
	adminGroup.GET("/name/:name", AdminHandler.GetAdminByNameHandler, middleware.AuthMiddleware("Admin"))
	adminGroup.PUT("/:id", AdminHandler.UpdateAdminHandler, middleware.AuthMiddleware("Admin"))
	adminGroup.DELETE("/:id", AdminHandler.DeleteAdminHandler, middleware.AuthMiddleware("Admin"))
}