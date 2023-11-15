package repository

import (
	"qbills/models/domain"

	req "qbills/utils/request"
	res "qbills/utils/response"

	"gorm.io/gorm"
)

type MembershipRepository interface {
	Create(membership *domain.Membership) (*domain.Membership, error)
	Update(membership *domain.Membership, id int) (*domain.Membership, error)
	FindById(id int) (*domain.Membership, error)
	FindByName(name string) (*domain.Membership, error)	
	// FindAll() ([]domain.Membership, error)
	// Delete(id int) error
}

type MembershipRepositoryImpl struct {
	DB *gorm.DB
}

func NewMembershipRepository(DB *gorm.DB) MembershipRepository {
	return &MembershipRepositoryImpl{DB: DB}
}

func (repository *MembershipRepositoryImpl) Create(membership *domain.Membership) (*domain.Membership, error) {
	membershipDB := req.MembershipDomainintoMembershipSchema(*membership)
	result := repository.DB.Create(&membershipDB)
	if result.Error != nil {
		return nil, result.Error
	}
	results := res.MembershipSchemaToMembershipDomain(membershipDB)

	return results, nil
}

func (repository *MembershipRepositoryImpl) Update(membership *domain.Membership, id int) (*domain.Membership, error) {
	result := repository.DB.Table("memberships").Where("id = ?", id).Updates(domain.Membership{
		Name: membership.Name,
		Telephone: membership.Telephone})
	if result.Error != nil {
		return nil, result.Error
	}

	return membership, nil
}

func (repository *MembershipRepositoryImpl) FindById(id int) (*domain.Membership, error) {
	membership := domain.Membership{}

	result := repository.DB.First(&membership, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &membership, nil
}

func (repository *MembershipRepositoryImpl) FindByName(name string) (*domain.Membership, error) {
	membership := domain.Membership{}

	result := repository.DB.Where("name = ?", name).First(&membership)
	if result.Error != nil {
		return nil, result.Error
	}

	return &membership, nil
}
