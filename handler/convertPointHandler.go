package handler

import (
	"net/http"
	"qbills/models/web"
	"qbills/services"
	"qbills/utils/helpers"
	res "qbills/utils/response"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
	"github.com/sirupsen/logrus"
>>>>>>> Stashed changes
>>>>>>> Stashed changes
)

type ConvertPointHandler interface {
	CreateConvertPointHandler(ctx echo.Context) error
	UpdateConvertPointHandler(ctx echo.Context) error
	GetConvertPointHandler(ctx echo.Context) error
	GetAllConvertPointHandler(ctx echo.Context) error
	DeleteConvertPointHandler(ctx echo.Context) error
}

type ConvertPointHandlerImpl struct {
	ConvertPointService services.ConvertPointService
}

func NewConvertPointHandler(convertPointService services.ConvertPointService) ConvertPointHandler {
	return &ConvertPointHandlerImpl{ConvertPointService: convertPointService}
}

func (c *ConvertPointHandlerImpl) CreateConvertPointHandler(ctx echo.Context) error {
	convertPointCreateRequest := web.ConvertPointRequest{}
	err := ctx.Bind(&convertPointCreateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Client Inputed ConvertPoint"))
	}
	result, err := c.ConvertPointService.CreateConvertPoint(ctx, convertPointCreateRequest)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Validation"))
		}
<<<<<<< Updated upstream
=======
<<<<<<< Updated upstream
=======
		if strings.Contains(err.Error(), "numeric") {
			return ctx.JSON(http.StatusConflict, helpers.ErrorResponse("Point is not valid must contain only numeric value"))
		}
		logrus.Error(err.Error())
>>>>>>> Stashed changes
>>>>>>> Stashed changes
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Create ConvertPoint Error"))
	}

	response := res.ConvertPointDomainToConvertPointResponse(result)

	return ctx.JSON(http.StatusCreated, helpers.SuccessResponse("Succesfully create ConvertPoint", response))
}

func (c *ConvertPointHandlerImpl) UpdateConvertPointHandler(ctx echo.Context) error {
	convertPointId := ctx.Param("id")
	convertPointIdInt, err := strconv.Atoi(convertPointId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Invalid Param Id"))
	}

	convertPointUpdateRequest := web.ConvertPointRequest{}
	err = ctx.Bind(&convertPointUpdateRequest)
	if err != nil {
		return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid Client Input"))
	}
	_, err = c.ConvertPointService.UpdateConvertPoint(ctx, convertPointUpdateRequest, convertPointIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "validation failed") {
			return ctx.JSON(http.StatusBadRequest, helpers.ErrorResponse("Invalid validation"))
		}

		if strings.Contains(err.Error(), "ConvertPoint not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("ConvertPoint not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Update convertPoint error"))
	}
	results, err := c.ConvertPointService.FindById(ctx, convertPointIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Convert Point not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Convert Point not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get Convert Point Data Error"))
	}

	response := res.ConvertPointDomainToConvertPointResponse(results)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Succesfully Updated ConvertPoint Data", response))
}

func (c *ConvertPointHandlerImpl) GetConvertPointHandler(ctx echo.Context) error {
	convertPointId := ctx.Param("id")
	convertPointIdInt, err := strconv.Atoi(convertPointId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Invalid Param ID oo"))
	}
	result, err := c.ConvertPointService.FindById(ctx, convertPointIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "Convert Point not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("Convert Point not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get Convert Point Data Error"))
	}

	response := res.ConvertPointDomainToConvertPointResponse(result)

	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get Data ConvertPoint", response))
}

func (c *ConvertPointHandlerImpl) GetAllConvertPointHandler(ctx echo.Context) error {
	result, err := c.ConvertPointService.FindAll()
	if err != nil {
		if strings.Contains(err.Error(), "convertPoint not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("ConvertPoint not found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Get All ConvertPoint Data Error"))
	}
	response := res.ConvertCPointResponse(result)
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Get All Data ConvertPoint", response))
}

func (c *ConvertPointHandlerImpl) DeleteConvertPointHandler(ctx echo.Context) error {
	convertPointId := ctx.Param("id")
	convertPointIdInt, err := strconv.Atoi(convertPointId)
	if err != nil {
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Invalid Param ID"))
	}
	err = c.ConvertPointService.DeleteConvertPoint(ctx, convertPointIdInt)
	if err != nil {
		if strings.Contains(err.Error(), "convertPoint not found") {
			return ctx.JSON(http.StatusNotFound, helpers.ErrorResponse("ConvertPoint Not Found"))
		}
		return ctx.JSON(http.StatusInternalServerError, helpers.ErrorResponse("Delete ConvertPoint Data Error"))
	}
	return ctx.JSON(http.StatusOK, helpers.SuccessResponse("Successfully Delete Data ConvertPoint", nil))
}
