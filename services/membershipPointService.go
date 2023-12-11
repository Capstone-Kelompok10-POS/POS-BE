package services

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/repository"
	"qbills/utils/helpers"
	req "qbills/utils/request"
)

type MembershipPointService interface {
	UpdateMembershipPointService(ctx echo.Context, request web.MembershipPointCreate) (*domain.MembershipPoint, error)
	FindAllMembershipPointService(ctx echo.Context) ([]domain.MembershipPoint, error)
	FindByIdMembershipPointService(ctx echo.Context, id uint) (*domain.MembershipPoint, error)
	FindIncreaseMembershipPointService(ctx echo.Context) ([]domain.MembershipPoint, error)
	FindDecreaseMembershipPointService(ctx echo.Context) ([]domain.MembershipPoint, error)
}

type MembershipPointServiceImpl struct {
	MembershipPointRepository repository.MembershipPointRepository
	MembershipRepository      repository.MembershipRepository
	validate                  *validator.Validate
}

func NewMembershipPointService(repository repository.MembershipPointRepository, MembershipRepository repository.MembershipRepository, validate *validator.Validate) *MembershipPointServiceImpl {
	return &MembershipPointServiceImpl{
		MembershipPointRepository: repository,
		MembershipRepository:      MembershipRepository,
		validate:                  validate,
	}
}

func (service *MembershipPointServiceImpl) UpdateMembershipPointService(ctx echo.Context, request web.MembershipPointCreate) (*domain.MembershipPoint, error) {
	err := service.validate.Struct(request)
	if err != nil {
		return nil, helpers.ValidationError(ctx, err)
	}

	req := req.MembershipPointCreateToMembershipPointDomain(request)

	membershipIdInt := int(req.MembershipID)

	membership, err := service.MembershipRepository.FindById(membershipIdInt)

	membership.TotalPoint += req.Point

	if membership.TotalPoint < 0 {
		return nil, fmt.Errorf("points decrease more than already point")
	}

	if err != nil {
		return nil, err
	}

	_, err = service.MembershipRepository.Update(membership, membershipIdInt)

	if err != nil {
		return nil, err
	}

	result, err := service.MembershipPointRepository.Create(req)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *MembershipPointServiceImpl) FindAllMembershipPointService(ctx echo.Context) ([]domain.MembershipPoint, error) {

	result, err := service.MembershipPointRepository.FindAll()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *MembershipPointServiceImpl) FindByIdMembershipPointService(ctx echo.Context, id uint) (*domain.MembershipPoint, error) {

	result, err := service.MembershipPointRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("Membership Point not found")
	}

	return result, nil

}

func (service *MembershipPointServiceImpl) FindIncreaseMembershipPointService(ctx echo.Context) ([]domain.MembershipPoint, error) {

	result, err := service.MembershipPointRepository.FindIncreasePoint()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *MembershipPointServiceImpl) FindDecreaseMembershipPointService(ctx echo.Context) ([]domain.MembershipPoint, error) {

	result, err := service.MembershipPointRepository.FindDecreasePoint()

	if err != nil {
		return nil, err
	}

	return result, nil
}
