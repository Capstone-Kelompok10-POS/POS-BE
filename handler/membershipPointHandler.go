package handler

import (
	"net/http"
	"qbills/models/web"
	"qbills/services"
	"qbills/utils/helpers"
	res "qbills/utils/response"
	"strconv"
	"strings"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type MembershipPointHandler interface {
	UpdateMembershipPointHandler(ctx echo.Context) error
	FindAllMembershipPointHandler(ctx echo.Context) error
	FindByIdMembershipPointHandler(ctx echo.Context) error
	FindIncreaseMembershipPointHandler(ctx echo.Context) error
	FindDecreaseMembershipPointHandler(ctx echo.Context) error
}

type MembershipPointImpl struct {
	membershipPoint services.MembershipPointService
	validate                  *validator.Validate
}

func NewMembershipPointHandler(membershipPointService services.MembershipPointService, validate *validator.Validate) MembershipPointHandler {
	return &MembershipPointImpl{
		membershipPoint: membershipPointService,
		validate:                  validate,}
}

func (c *MembershipPointImpl) UpdateMembershipPointHandler(ctx echo.Context) error {
	err := c.validate.Struct(ctx)
	if err != nil {
		return helpers.ValidationError(ctx, err)
	}

	request := new(web.MembershipPointCreate)

	if err := ctx.Bind(request); err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	result, err := c.membershipPoint.UpdateMembershipPointService(*request)

	if err != nil {
		switch {
		case strings.Contains(err.Error(), "validation error"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		case strings.Contains(err.Error(), "stock decrease more than stock"):
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("point decrease more than avabile point"))
		default:
			return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("failed to update point"))
		}
	}

	response := res.MembershipDomainToMembershipCreateResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("success update membership point", response))
}

func (c *MembershipPointImpl) FindAllMembershipPointHandler(ctx echo.Context) error {
	membershipID := ctx.Param("id")
	membershipIDInt, err := strconv.Atoi(membershipID)
	if err != nil{
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}
	result, err := c.membershipPoint.FindAllMembershipPointByIdService(uint(membershipIDInt))

	if err != nil {
		if strings.Contains(err.Error(), "membership point not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("membership point record not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("membership point data record error"))
	}

	response := res.ConvertMembershipPointResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get all update memebrship point", response))

}

func (c *MembershipPointImpl) FindByIdMembershipPointHandler(ctx echo.Context) error {
	membershipPointID := ctx.Param("id")
	membershipPointIDInt, err := strconv.Atoi(membershipPointID)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid param id"))
	}

	stock, err := c.membershipPoint.FindByIdMembershipPointService(uint(membershipPointIDInt))
	if err != nil {
		statusCode := http.StatusInternalServerError
		errorMessage := "Get membership point data error"

		if strings.Contains(err.Error(), "membership point not found") {
			statusCode = http.StatusNotFound
			errorMessage = "membership point not found"
		}

		return ctx.JSON(statusCode, helpers.ErrorResponse(errorMessage))
	}

	response := res.MembershipDomainToMembershipPointResponse(stock)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Success get membership point by id", response))
}

func (c *MembershipPointImpl) FindIncreaseMembershipPointHandler(ctx echo.Context) error {
	result, err := c.membershipPoint.FindIncreaseMembershipPointService()

	if err != nil {
		if strings.Contains(err.Error(), "increase membership point not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("increase membership point not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("increase membership point data error"))
	}

	response := res.ConvertMembershipPointResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get all Increase membership point", response))

}

func (c *MembershipPointImpl) FindDecreaseMembershipPointHandler(ctx echo.Context) error {
	result, err := c.membershipPoint.FindDecreaseMembershipPointService()

	if err != nil {
		if strings.Contains(err.Error(), "decrease membership point not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("decrease membership point not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("decrease membership point data error"))
	}

	response := res.ConvertMembershipPointResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success get all decrease membership point", response))

}
