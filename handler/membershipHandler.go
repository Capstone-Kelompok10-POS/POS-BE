package handler

import (
	"fmt"
	"net/http"
	"qbills/models/web"
	"qbills/services"
	"qbills/utils/helpers"
	"qbills/utils/helpers/middleware"
	res "qbills/utils/response"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
	"github.com/sirupsen/logrus"
)

type MembershipHandler interface {
	RegisterMembershipHandler(ctx echo.Context) error
	UpdateMembershipHandler(ctx echo.Context) error 
	GetMembershipHandler(ctx echo.Context) error
	GetMembershipsHandler(ctx echo.Context) error
	GetMembershipByNameHandler(ctx echo.Context) error
	GetMembershipByPhoneNumber(ctx echo.Context) error
	DeleteMembershipHandler(ctx echo.Context) error
}

type MembershipHandlerImpl struct {
	MembershipService services.MembershipService
}

func NewMembershipHandler(membershipService services.MembershipService) MembershipHandler {
	return &MembershipHandlerImpl{MembershipService: membershipService}
}

func (c *MembershipHandlerImpl) RegisterMembershipHandler(ctx echo.Context) error {
	cashierId := middleware.ExtractTokenCashierId(ctx)
	if cashierId == 0.0 {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid token cashier"))
	}
	membershipCreateRequest := web.MembershipCreateRequest{CashierID: uint(cashierId)}
	err := ctx.Bind(&membershipCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid input"))
	}

	result, err := c.MembershipService.CreateMembership(ctx, membershipCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "phone_number already exist") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("phone number already exist"))
		}
		if strings.Contains(err.Error(), "number") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("phone number is not valid must contain only number value"))
		}
		logrus.Error(err.Error())
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign up error"))
	}

	response := res.MembershipDomainToMembershipResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("successfully created account membership", response))
}

func (c *MembershipHandlerImpl) GetMembershipHandler(ctx echo.Context) error {
	membershipId := ctx.Param("id")
	membershipIdInt, err := strconv.Atoi(membershipId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	result, err := c.MembershipService.FindById(ctx, membershipIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "membership not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("membership not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get membership data error"))
	}
	response := res.MembershipDomainToMembershipResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully get data membership", response))
}

func (c MembershipHandlerImpl) GetMembershipsHandler(ctx echo.Context) error {
	memberships, totalMemberships, err := c.MembershipService.FindAll(ctx)
	if err != nil {
		if strings.Contains(err.Error(), "memberships not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("memberships not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get memberships data error"))
	}

	response := res.ConvertMembershipResponse(memberships)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponseWithTotal("successfully get all data memberships", response, totalMemberships))
}

func (c MembershipHandlerImpl) GetMembershipByNameHandler(ctx echo.Context) error{
	membershipName := ctx.Param("name")

	result, err := c.MembershipService.FindByName(ctx, membershipName)
	if err != nil {
		if strings.Contains(err.Error(), "membership not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("membership not found"))
		}
		return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Get membership data by name error"))
	}
	response := res.MembershipDomainToMembershipResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succesfully get membership data by name", response))
}

func (c MembershipHandlerImpl) GetMembershipByPhoneNumber(ctx echo.Context) error{
	membershipPhoneNumber := ctx.Param("phoneNumber")

	result, err := c.MembershipService.FindByPhoneNumber(ctx, membershipPhoneNumber)
	if err != nil {
		if strings.Contains(err.Error(), "membership not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("membership not found"))
		}
		return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Get membership data by phone number error"))
	}
	response := res.MembershipDomainToMembershipResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("succesfully get membership data by phone number", response))
}

func (c MembershipHandlerImpl) UpdateMembershipHandler(ctx echo.Context) error {
	membershipId := ctx.Param("id")
	membershipIdInt, err := strconv.Atoi(membershipId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	membershipUpdateRequest := web.MembershipUpdateRequest{}
	err = ctx.Bind(&membershipUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	_, err = c.MembershipService.UpdateMembership(ctx, membershipUpdateRequest, membershipIdInt)
	fmt.Print(err)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}
		if strings.Contains(err.Error(), "membership not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("membership not found"))
		}
		if strings.Contains(err.Error(), "phone_number already exist") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("phone number already exist"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("update membership error"))
	}
	results, err := c.MembershipService.FindById(ctx, membershipIdInt)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid client input"))
	}

	response := res.MembershipDomainToMembershipResponse(results)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully updated data membership", response))
}

func (c MembershipHandlerImpl) DeleteMembershipHandler(ctx echo.Context) error {
	membershipId := ctx.Param("id")
	membershipIdInt, err := strconv.Atoi(membershipId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	err = c.MembershipService.DeleteMembership(ctx, membershipIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "membership not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("membership not found"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("delete data membership error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("successfully delete data membership", nil))
}