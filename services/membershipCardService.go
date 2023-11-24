package services

import (
	"fmt"
	"qbills/models/domain"
	"qbills/repository"

	// "github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type MembershipCardService interface {
	PrintMembershipCard(ctx echo.Context, id int) (*domain.Membership, error)
	FindById(ctx echo.Context, id int) (*domain.Membership, error)
}

type MembershipCardServiceImpl struct {
	MembershipCardRepository repository.MembershipCardRepository
	// Validate                 *validator.Validate
}

func NewMembershipCardService(membershipCardRepository repository.MembershipCardRepository) *MembershipCardServiceImpl {
	return &MembershipCardServiceImpl{
		MembershipCardRepository: membershipCardRepository,
		// Validate:                 validate,
	}
}

func (service *MembershipCardServiceImpl) FindById(ctx echo.Context, id int) (*domain.Membership, error) {
	existingMembership, _ := service.MembershipCardRepository.PrintMembershipCard(id)
	if existingMembership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	return existingMembership, nil
}

func (service *MembershipCardServiceImpl) PrintMembershipCard(ctx echo.Context, id int) (*domain.Membership, error) {
	existingMembership, _ := service.MembershipCardRepository.FindById(id)
	if existingMembership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	result, err := service.MembershipCardRepository.PrintMembershipCard(id)
	if err != nil {
		return nil, fmt.Errorf("error creating membership card: %s", err.Error())
	}

	return result, nil
}
