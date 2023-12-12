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

func MemberShipPointRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	membershipPointRepository := repository.NewMembershipPointRepository(db)
	membershipRepository := repository.NewMembershipRepository(db)
	membershipPointService := services.NewMembershipPointService(membershipPointRepository, membershipRepository)
	membershipPointHandler := handler.NewMembershipPointHandler(membershipPointService, validate)

	StockGroup := e.Group("/api/v1/membership/point")

	StockGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	StockGroup.POST("", membershipPointHandler.UpdateMembershipPointHandler)
	StockGroup.GET("s/:id", membershipPointHandler.FindAllMembershipPointHandler)
	StockGroup.GET("/:id", membershipPointHandler.FindByIdMembershipPointHandler)
	StockGroup.GET("/get/increase", membershipPointHandler.FindIncreaseMembershipPointHandler)
	StockGroup.GET("/get/decrease", membershipPointHandler.FindDecreaseMembershipPointHandler)

}
