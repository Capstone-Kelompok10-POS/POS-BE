package services

import (
	"fmt"
	"qbills/models/domain"
	"qbills/models/web"
	"qbills/repository"
	req "qbills/utils/request"

	"gorm.io/gorm"
)

type MembershipPointService interface {
	UpdateMembershipPointService(request web.MembershipPointCreate) (*domain.MembershipPoint, error)
	FindAllMembershipPointByIdService(membershipId uint) ([]domain.MembershipPoint, error)
	FindByIdMembershipPointService(id uint) (*domain.MembershipPoint, error)
	FindIncreaseMembershipPointService() ([]domain.MembershipPoint, error)
	FindDecreaseMembershipPointService() ([]domain.MembershipPoint, error)
}

type MembershipPointServiceImpl struct {
	MembershipPointRepository repository.MembershipPointRepository
	MembershipRepository      repository.MembershipRepository
}

func NewMembershipPointService(repository repository.MembershipPointRepository, MembershipRepository repository.MembershipRepository) *MembershipPointServiceImpl {
	return &MembershipPointServiceImpl{
		MembershipPointRepository: repository,
		MembershipRepository:      MembershipRepository,
	}
}

func (service *MembershipPointServiceImpl) UpdateMembershipPointService(request web.MembershipPointCreate) (*domain.MembershipPoint, error) {
	req := req.MembershipPointCreateToMembershipPointDomain(request)
	var tx *gorm.DB
	membershipIdInt := int(req.MembershipID)

	membership, err := service.MembershipRepository.FindById(membershipIdInt)

	membership.TotalPoint += uint(req.Point)

	if membership.TotalPoint < 0 {
		return nil, fmt.Errorf("points decrease more than already point")
	}

	if err != nil {
		return nil, err
	}

	_, err = service.MembershipRepository.UpdatePointNoTx(membership, membershipIdInt)

	if err != nil {
		return nil, err
	}

	result, err := service.MembershipPointRepository.Create(tx, req)

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *MembershipPointServiceImpl) FindAllMembershipPointByIdService(membershipId uint) ([]domain.MembershipPoint, error) {

	result, err := service.MembershipPointRepository.FindAllByMembershipId(membershipId)
	if len(result) == 0 {
		return nil, fmt.Errorf("membership point not found")
	}
	if err != nil {
		return nil, err
	}
	
	return result, nil
}

func (service *MembershipPointServiceImpl) FindByIdMembershipPointService(id uint) (*domain.MembershipPoint, error) {

	result, err := service.MembershipPointRepository.FindById(id)
	if err != nil {
		return nil, fmt.Errorf("membership Point not found")
	}

	return result, nil

}

func (service *MembershipPointServiceImpl) FindIncreaseMembershipPointService() ([]domain.MembershipPoint, error) {

	result, err := service.MembershipPointRepository.FindIncreasePoint()

	if err != nil {
		return nil, err
	}

	return result, nil
}

func (service *MembershipPointServiceImpl) FindDecreaseMembershipPointService() ([]domain.MembershipPoint, error) {

	result, err := service.MembershipPointRepository.FindDecreasePoint()

	if err != nil {
		return nil, err
	}

	return result, nil
}
