package handler

import (
	"net/http"
	"qbills/models/web"
	"qbills/services"
	"qbills/utils/helpers"
	"qbills/utils/helpers/middleware"
	res "qbills/utils/response"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type SuperAdminHandler interface {
	LoginSuperAdminHandler(ctx echo.Context) error
	GetSuperAdminHandler(ctx echo.Context) error
	GetSuperAdminsHandler(ctx echo.Context) error
}

type SuperAdminHandlerImpl struct {
	SuperAdminService services.SuperAdminService
}

func NewSuperAdminHandler(superAdminService services.SuperAdminService) SuperAdminHandler {
	return &SuperAdminHandlerImpl{SuperAdminService: superAdminService}
}



func (c *SuperAdminHandlerImpl) LoginSuperAdminHandler(ctx echo.Context) error {
	superAdminLoginRequest := web.SuperAdminLoginRequest{}

	err := ctx.Bind(&superAdminLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response, err := c.SuperAdminService.LoginSuperAdmin(ctx, superAdminLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}

		if strings.Contains(err.Error(), "invalid email or password") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid email or password"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign in error"))
	}

	superAdminLoginResponse := res.SuperAdminDomainToSuperAdminLoginResponse(response)

	token, err := middleware.GenerateTokenSuperAdmin(response.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("generate jwt token error"))
	}

	superAdminLoginResponse.Token = token

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Succesfully Sign In", superAdminLoginResponse))
}

func (c *SuperAdminHandlerImpl) GetSuperAdminHandler(ctx echo.Context) error {
	superAdmintId := ctx.Param("id")
	superAdminIdInt, err := strconv.Atoi(superAdmintId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	result, err := c.SuperAdminService.FindById(ctx, superAdminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Super Admin not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Super Admin not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get Super Admin data error"))
	}
	response := res.SuperAdminDomainToSuperAdminResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Super Admin", response))
}

func (c SuperAdminHandlerImpl) GetSuperAdminsHandler(ctx echo.Context) error {
	result, err := c.SuperAdminService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "Super Admins not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Super Admins not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get Super Admins data error"))
	}

	response := res.ConvertSuperAdminResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Succesfully Get All data SuperAdmins", response))
}



