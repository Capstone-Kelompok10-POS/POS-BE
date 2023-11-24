package handler

import (
	"net/http"
	"qbills/services"
	"qbills/utils/helpers"
	// res "qbills/utils/response"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type MembershipCardHandler interface {
	PrintMembershipCardHandler(ctx echo.Context) error
}

type MembershipCardHandlerImpl struct {
	MembershipCardService services.MembershipCardService
}

func NewMembershipCardHandler(membershipCardService services.MembershipCardService) MembershipCardHandler {
	return &MembershipCardHandlerImpl{MembershipCardService: membershipCardService}
}

func (c *MembershipCardHandlerImpl) PrintMembershipCardHandler(ctx echo.Context) error {
	membershipCardId := ctx.Param("id")
	membershipIdInt, err := strconv.Atoi(membershipCardId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("invalid param id"))
	}

	result, err := c.MembershipCardService.PrintMembershipCard(ctx, membershipIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "membership not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("membership not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("get membership data error"))
	}

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success print membership card", result))
}