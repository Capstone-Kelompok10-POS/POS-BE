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

	superAdminGroup := e.Group("api/v1/super-admin")

	superAdminGroup.POST("/login", SuperAdminHandler.LoginSuperAdminHandler)

	superAdminGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	superAdminGroup.GET("/:id", SuperAdminHandler.GetSuperAdminHandler, middleware.AuthMiddleware("SuperAdmin"))
	superAdminGroup.GET("s", SuperAdminHandler.GetSuperAdminsHandler, middleware.AuthMiddleware("SuperAdmin"))
}
