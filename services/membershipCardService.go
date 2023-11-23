package services

import (
	"fmt"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/repository"
	"qbills/utils/helpers"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type MembershipCardService interface {
	Print(ctx echo.Context, request web.MembershipCardPrintRequest, id int) (*domain.Membership, error)
	FindById(ctx echo.Context, id int) (*domain.Membership, error)
}

type MembershipCardServiceImpl struct {
	MembershipCardRepository repository.MembershipCardRepository
	Validate                 *validator.Validate
}

func NewMembershipCardService(membershipCardRepository repository.MembershipCardRepository, validate *validator.Validate) *MembershipCardServiceImpl {
	return &MembershipCardServiceImpl{
		MembershipCardRepository: membershipCardRepository,
		Validate:                 validate,
	}
}

func (service *MembershipCardServiceImpl) FindById(ctx echo.Context, id int) (*domain.Membership, error) {
	existingMembership, _ := service.MembershipCardRepository.FindById(id)
	if existingMembership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	return existingMembership, nil
}

func (service *MembershipCardServiceImpl) Print(ctx echo.Context, request web.MembershipCardPrintRequest, id int) (*domain.Membership, error){
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingMembership, _ := service.MembershipCardRepository.FindById(id)
	if existingMembership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	if err != nil {
		return nil, fmt.Errorf("error creating membership card %s", err.Error())
	}

	return existingMembership, nil
}
