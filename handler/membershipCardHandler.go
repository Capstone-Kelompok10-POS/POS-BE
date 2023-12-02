package handler

import (
	"net/http"
	"qbills/services"
	"qbills/utils/helpers"
	res "qbills/utils/response"
<<<<<<< Updated upstream
=======

>>>>>>> Stashed changes
	// res "qbills/utils/response"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
<<<<<<< Updated upstream
=======
	"github.com/sirupsen/logrus"
>>>>>>> Stashed changes
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
<<<<<<< Updated upstream
=======
		logrus.Error(err.Error())
>>>>>>> Stashed changes
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("get membership data error"))
	}

	response := res.MembershipCardDomainToMembershipCardResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("success print membership card", response))
}