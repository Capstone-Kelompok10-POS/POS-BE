package services

import (
	"fmt"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/repository"
	"qbills/utils/helpers"
	req "qbills/utils/request"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type MembershipService interface {
	CreateMembership(ctx echo.Context, request web.MembershipCreateRequest) (*domain.Membership, error)
	UpdateMembership(ctx echo.Context, request web.MembershipUpdateRequest, id int) (*domain.Membership, error)
	FindById(ctx echo.Context, id int) (*domain.Membership, error)
}

type MembershipServiceImpl struct {
	MembershipRepository repository.MembershipRepository
	Validate *validator.Validate
}

func NewMembershipService(membershipRepository repository.MembershipRepository, validate *validator.Validate) *MembershipServiceImpl {
	return &MembershipServiceImpl{
		MembershipRepository: membershipRepository,
		Validate: validate,
	}
}

func (service *MembershipServiceImpl) CreateMembership(ctx echo.Context, request web.MembershipCreateRequest) (*domain.Membership, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingMembership, _:= service.MembershipRepository.FindByName(request.Name)
	if existingMembership != nil {
		return nil, fmt.Errorf("name already exist")
	}
	membership := req.MembershipCreateRequestToMembershipDomain(request)

	result, err := service.MembershipRepository.Create(membership)

	if err != nil {
		return nil, fmt.Errorf("error creating membership %s", err.Error())
	}

	return result, nil
}