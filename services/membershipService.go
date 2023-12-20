package services

import (
	"fmt"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/repository"
	"qbills/utils/helpers"
	"strings"

	req "qbills/utils/request"

	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
)

type MembershipService interface {
	CreateMembership(ctx echo.Context, request web.MembershipCreateRequest) (*domain.Membership, error)
	UpdateMembership(ctx echo.Context, request web.MembershipUpdateRequest, id int) (*domain.Membership, error)
	FindById(id int) (*domain.Membership, error)
	FindByName(name string) (*domain.Membership, error)
	FindByPhoneNumber(phoneNumber string) (*domain.Membership, error)
	FindAll() ([]domain.Membership,int, error)
	FindTopMember() ([]domain.Membership, error)
	DeleteMembership(id int) error
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

	existingMembership, _ := service.MembershipRepository.FindByPhoneNumber(request.PhoneNumber)
	if existingMembership != nil {
		return nil, fmt.Errorf("phone_number already exists")
	}

	membership := req.MembershipCreateRequestToMembershipDomain(request)

	result, err := service.MembershipRepository.Create(membership)
	if err != nil {
		return nil, fmt.Errorf("error creating membership %s", err.Error())
	}

	return result, nil
}


func (service *MembershipServiceImpl) FindById(id int) (*domain.Membership, error) {
	existingMembership, _ := service.MembershipRepository.FindById(id)
	if existingMembership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	return existingMembership, nil
}

func (service *MembershipServiceImpl) FindAll() ([]domain.Membership, int, error) {
	memberships, totalMembership, err := service.MembershipRepository.FindAll()
	if err != nil {
		return nil, 0, fmt.Errorf("error when get membership")
	}

	return memberships, totalMembership, nil
}

func (service *MembershipServiceImpl) FindTopMember() ([]domain.Membership, error) {
	memberships, err := service.MembershipRepository.FindTopMember()
	if err != nil {
		return nil, fmt.Errorf("error when get membership")
	}

	return memberships, nil
}

func (service *MembershipServiceImpl) FindByName(name string) (*domain.Membership, error) {
	name = strings.ToLower(name)
	membership, err := service.MembershipRepository.FindByName(name)
	if err != nil {
		return nil, fmt.Errorf("membership not found")
	}
	if membership.Name == "" {
		return nil, fmt.Errorf("membership not found")
	}

	return membership, nil
}

func (service *MembershipServiceImpl) FindByPhoneNumber(phoneNumber string) (*domain.Membership, error) {
	membership, _ := service.MembershipRepository.FindByPhoneNumber(phoneNumber)
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
	if existingMembership.PhoneNumber != membership.PhoneNumber {
		existingMembershipPhoneNumber, _ := service.MembershipRepository.FindByPhoneNumber(membership.PhoneNumber)
		if existingMembershipPhoneNumber != nil {
			return nil, fmt.Errorf("phone_number already exist")
		}
	}
	result, err := service.MembershipRepository.Update(membership, id)
	if err != nil {
		return nil, fmt.Errorf("error when updating data membership: %s", err.Error())
	}

	return result, nil
}

func (service *MembershipServiceImpl) DeleteMembership(id int) error {
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
