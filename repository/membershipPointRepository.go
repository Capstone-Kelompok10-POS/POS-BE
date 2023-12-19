package repository

import (
	"qbills/models/domain"
	"qbills/utils/request"
	"qbills/utils/response"

	"gorm.io/gorm"
)

type MembershipPointRepository interface {
	Create(tx *gorm.DB, point *domain.MembershipPoint) (*domain.MembershipPoint, error)
	Update(point *domain.MembershipPoint, id int) (*domain.MembershipPoint, error)
	FindAllByMembershipId(membershipId uint) ([]domain.MembershipPoint, error)
	FindById(id uint) (*domain.MembershipPoint, error)
	FindIncreasePoint() ([]domain.MembershipPoint, error)
	FindDecreasePoint() ([]domain.MembershipPoint, error)
}

type MembershipPointRepositoryImpl struct {
	DB *gorm.DB
}

func NewMembershipPointRepository(DB *gorm.DB) MembershipPointRepository {
	return &MembershipPointRepositoryImpl{DB: DB}
}

func (repository *MembershipPointRepositoryImpl) Create(tx *gorm.DB, point *domain.MembershipPoint) (*domain.MembershipPoint, error) {

	req := request.MembershipPointDomainToMembershipPointSchema(point)

	result := tx.Create(&req)

	if result.Error != nil {
		return nil, result.Error
	}

	res := response.MembershipPointSchemaToMembershipPointDomain(req)

	return res, nil
}

func (repository *MembershipPointRepositoryImpl) Update(cashier *domain.MembershipPoint, id int) (*domain.MembershipPoint, error) {
	result := repository.DB.Table("membership_points").Where("id = ?", id).Updates(domain.MembershipPoint{
		ID:           cashier.ID,
		CreatedAt:    cashier.CreatedAt,
		MembershipID: cashier.MembershipID,
		Membership:   cashier.Membership,
		Point:        cashier.Point,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return cashier, nil
}

func (repository *MembershipPointRepositoryImpl) FindAllByMembershipId(membershipId uint) ([]domain.MembershipPoint, error) {
	point := []domain.MembershipPoint{}

	if err := repository.DB.Preload("Membership").Where("membership_id = ? ", membershipId).Find(&point).Error; err != nil {
		return nil, err
	}

	return point, nil
}

func (repository *MembershipPointRepositoryImpl) FindById(id uint) (*domain.MembershipPoint, error) {
	point := domain.MembershipPoint{}

	result := repository.DB.Preload("Membership").First(&point, id)

	if result.Error != nil {
		return nil, result.Error
	}

	return &point, nil
}

func (repository *MembershipPointRepositoryImpl) FindIncreasePoint() ([]domain.MembershipPoint, error) {
	point := []domain.MembershipPoint{}

	if err := repository.DB.Preload("Membership").Where("point > 0").Find(&point).Error; err != nil {
		return nil, err
	}

	return point, nil
}

func (repository *MembershipPointRepositoryImpl) FindDecreasePoint() ([]domain.MembershipPoint, error) {
	point := []domain.MembershipPoint{}

	if err := repository.DB.Preload("Membership").Where("point < 0").Find(&point).Error; err != nil {
		return nil, err
	}

	return point, nil
}
