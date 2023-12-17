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

func MembershipRoutes(e *echo.Echo, db *gorm.DB, validate *validator.Validate) {
	membershipRepository := repository.NewMembershipRepository(db)
	membershipService := services.NewMembershipService(membershipRepository, validate)
	MembershipHandler := handler.NewMembershipHandler(membershipService)

	membershipGroup := e.Group("api/v1/membership")

	membershipGroup.Use(echoJwt.JWT([]byte(os.Getenv("SECRET_KEY"))))

	membershipGroup.POST("/register", MembershipHandler.RegisterMembershipHandler, middleware.AuthMiddleware("Cashier"))
	membershipGroup.GET("/:id", MembershipHandler.GetMembershipHandler)
	membershipGroup.GET("s", MembershipHandler.GetMembershipsHandler)
	membershipGroup.GET("s/top", MembershipHandler.GetTopMembershipsHandler)
	membershipGroup.GET("/name/:name", MembershipHandler.GetMembershipByNameHandler)
	membershipGroup.GET("/phone-number/:phoneNumber", MembershipHandler.GetMembershipByPhoneNumber)
	membershipGroup.PUT("/:id", MembershipHandler.UpdateMembershipHandler, middleware.AuthMiddleware("Admin"))
	membershipGroup.DELETE("/:id", MembershipHandler.DeleteMembershipHandler, middleware.AuthMiddleware("Admin"))

}
