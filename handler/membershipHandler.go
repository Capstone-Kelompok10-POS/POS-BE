package handler

import (
	"net/http"
	"qbills/models/web"
	"qbills/services"
	"qbills/utils/helpers"
	res "qbills/utils/response"
	"strings"

	"github.com/labstack/echo/v4"
)

type MembershipHandler interface {
	RegisterMembershipHandler(ctx echo.Context) error
}

type MembershipHandlerImpl struct {
	MembershipService services.MembershipService
}

func NewMembershipHandler(membershipService services.MembershipService) MembershipHandler {
	return &MembershipHandlerImpl{MembershipService: membershipService}
}

func (c *MembershipHandlerImpl) RegisterMembershipHandler(ctx echo.Context) error {
	membershipCreateRequest := web.MembershipCreateRequest{}
	err := ctx.Bind(&membershipCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid input"))
	}

	result, err := c.MembershipService.CreateMembership(ctx, membershipCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation error") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("invalid validation"))
		}

		if strings.Contains(err.Error(), "name already exist") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("name already exist"))
		}

		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("sign up error"))
	}

	response := res.MembershipDomainToMembershipResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("successfully created account membership", response))
}
