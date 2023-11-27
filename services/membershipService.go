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
	FindByName(ctx echo.Context, name string) (*domain.Membership, error)
	FindAll(ctx echo.Context) ([]domain.Membership, error)
	DeleteMembership(ctx echo.Context, id int) error
}

type MembershipServiceImpl struct {
	MembershipRepository repository.MembershipRepository
	Validate             *validator.Validate
}

func NewMembershipService(membershipRepository repository.MembershipRepository, validate *validator.Validate) *MembershipServiceImpl {
	return &MembershipServiceImpl{
		MembershipRepository: membershipRepository,
		Validate:             validate,
	}
}

func (service *MembershipServiceImpl) CreateMembership(ctx echo.Context, request web.MembershipCreateRequest) (*domain.Membership, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingMembership, _ := service.MembershipRepository.FindByPhoneNumber(request.Phone_Number)
	if existingMembership != nil {
		return nil, fmt.Errorf("phone_number already exist")
	}

	membership := req.MembershipCreateRequestToMembershipDomain(request)
	fmt.Println(membership)

	result, err := service.MembershipRepository.Create(membership)
	if err != nil {
		return nil, fmt.Errorf("error creating membership %s", err.Error())
	}

	fmt.Println(result.CodeMember)

	return result, nil
}


func (service *MembershipServiceImpl) FindById(ctx echo.Context, id int) (*domain.Membership, error) {
	existingMembership, _ := service.MembershipRepository.FindById(id)
	if existingMembership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	return existingMembership, nil
}

func (service *MembershipServiceImpl) FindAll(ctx echo.Context) ([]domain.Membership, error) {
	memberships, err := service.MembershipRepository.FindAll()
	if err != nil {
		return nil, fmt.Errorf("membership not found")
	}

	return memberships, nil
}

func (service *MembershipServiceImpl) FindByName(ctx echo.Context, name string) (*domain.Membership, error) {
	membership, _ := service.MembershipRepository.FindByName(name)
	if membership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	return membership, nil
}

func (service *MembershipServiceImpl) UpdateMembership(ctx echo.Context, request web.MembershipUpdateRequest, id int) (*domain.Membership, error) {
	err := service.Validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	existingMembership, _ := service.MembershipRepository.FindById(id)
	if existingMembership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	membership := req.MembershipUpdateRequestToMembershipDomain(request)
	result, err := service.MembershipRepository.Update(membership, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating data membership: %s", err.Error())
	}

	return result, nil
}

func (service *MembershipServiceImpl) DeleteMembership(ctx echo.Context, id int) error {
	existingMembership, _ := service.MembershipRepository.FindById(id)
	if existingMembership == nil {
		return fmt.Errorf("membership not found")
	}

	err := service.MembershipRepository.Delete(id)
	if err != nil {
		return fmt.Errorf("error deleting membership: %s", err)
	}

	return nil
}
