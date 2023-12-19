package routes

import (
	"os"
	"qbills/handler"
	"qbills/repository"
	"qbills/services"
	"qbills/utils/helpers/middleware"
	"qbills/utils/helpers/password"

	"github.com/go-playground/validator"
	echoJwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func SuperAdminRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate, password password.PasswordHandler) {
	superAdminRepository := repository.NewSuperAdminRepository(db)
	superAdminService := services.NewSuperAdminService(superAdminRepository, validate, password)
	SuperAdminHandler := handler.NewSuperAdminHandler(superAdminService)

	superAdminGroup := e.Group("api/v1/super-admin")

	superAdminGroup.POST("/login", SuperAdminHandler.LoginSuperAdminHandler)

	superAdminGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	superAdminGroup.GET("/:id", SuperAdminHandler.GetSuperAdminHandler, middleware.AuthMiddleware("SuperAdmin"))
	superAdminGroup.GET("s", SuperAdminHandler.GetSuperAdminsHandler, middleware.AuthMiddleware("SuperAdmin"))
}