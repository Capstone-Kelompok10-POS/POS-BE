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

type AdminHandler interface {
	RegisterAdminHandler(ctx echo.Context) error
	LoginAdminHandler(ctx echo.Context) error
	UpdateAdminHandler(ctx echo.Context) error
	GetAdminHandler(ctx echo.Context) error
	GetAdminsHandler(ctx echo.Context) error
	GetAdminByNameHandler(ctx echo.Context) error
	DeleteAdminHandler(ctx echo.Context) error
}

type AdminHandlerImpl struct {
	AdminService services.AdminService
}

func NewAdminHandler(adminService services.AdminService) AdminHandler {
	return &AdminHandlerImpl{AdminService: adminService}
}

func (c *AdminHandlerImpl) RegisterAdminHandler(ctx echo.Context) error {
	adminCreateRequest := web.AdminCreateRequest{}
	err := ctx.Bind(&adminCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := c.AdminService.CreateAdmin(ctx, adminCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))

		}

		if strings.Contains(err.Error(), "username already exist") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("username already exist"))

		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign up error"))
	}

	response := res.AdminDomainToAdminResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("succesfully created account admin", response))
}

func (c *AdminHandlerImpl) LoginAdminHandler(ctx echo.Context) error {
	adminLoginRequest := web.AdminLoginRequest{}

	err := ctx.Bind(&adminLoginRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response, err := c.AdminService.LoginAdmin(ctx, adminLoginRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}

		if strings.Contains(err.Error(), "invalid username or password") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid username or password"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign in error"))
	}

	adminLoginResponse := res.AdminDomainToAdminLoginResponse(response)

	token, err := middleware.GenerateTokenAdmin(response.ID)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("generate jwt token error"))
	}

	adminLoginResponse.Token = token
	adminLoginResponse.Role = "Admin"
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Succesfully Sign In", adminLoginResponse))
}

func (c *AdminHandlerImpl) GetAdminHandler(ctx echo.Context) error {
	admintId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(admintId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	result, err := c.AdminService.FindById(ctx, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("admin not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get admin data error"))
	}
	response := res.AdminDomainToAdminResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data Admin", response))
}

func (c AdminHandlerImpl) GetAdminsHandler(ctx echo.Context) error {
	result, err := c.AdminService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "admins not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("admins not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get admins data error"))
	}

	response := res.ConvertAdminResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Succesfully Get All data Admins", response))
}

func (c AdminHandlerImpl) GetAdminByNameHandler(ctx echo.Context) error {
	adminName := ctx.Param("name")

	result, err := c.AdminService.FindByName(ctx, adminName)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("admin not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get admin data by name error"))
	}
	response := res.AdminDomainToAdminResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Succesfully get admin data by name", response))
}

func (c AdminHandlerImpl) UpdateAdminHandler(ctx echo.Context) error {
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	adminUpdateRequest := web.AdminUpdateRequest{}
	err = ctx.Bind(&adminUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	_, err = c.AdminService.UpdateAdmin(ctx, adminUpdateRequest, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("admin not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("update admin error"))
	}
	results, err := c.AdminService.FindById(ctx, adminIdInt)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response := res.AdminDomainToAdminResponse(results)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully updated data admin", response))
}

func (c AdminHandlerImpl) DeleteAdminHandler(ctx echo.Context) error {
	adminId := ctx.Param("id")
	adminIdInt, err := strconv.Atoi(adminId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	err = c.AdminService.DeleteAdmin(ctx, adminIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "admin not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("admin not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete data admin error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succesfully delete data admin", nil))
}