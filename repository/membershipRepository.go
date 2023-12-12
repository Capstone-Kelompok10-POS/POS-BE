package repository

import (
	"qbills/models/domain"
	"qbills/models/schema"

	req "qbills/utils/request"
	res "qbills/utils/response"

	"gorm.io/gorm"
)

type MembershipRepository interface {
	Create(membership *domain.Membership) (*domain.Membership, error)
	Update(membership *domain.Membership, id int) (*domain.Membership, error)
	UpdatePoint(tx *gorm.DB, membership *domain.Membership) error
	UpdatePointNoTx(membership *domain.Membership, id int) (*domain.Membership, error)
	FindById(id int) (*domain.Membership, error)
	FindByName(name string) (*domain.Membership, error)
	FindAll() ([]domain.Membership, int, error)
	FindTopMember() ([]domain.Membership, error)
	FindByPhoneNumber(phoneNumber string) (*domain.Membership, error)
	Delete(id int) error
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
		Name:        membership.Name,
		PhoneNumber: membership.PhoneNumber,
		TotalPoint:  membership.TotalPoint,
	})
	if result.Error != nil {
		return nil, result.Error
	}

	return membership, nil
}

func (repository *MembershipRepositoryImpl) UpdatePointNoTx(membership *domain.Membership, id int) (*domain.Membership, error) {
	result := repository.DB.Table("memberships").Where("id = ?", id).Update("total_point", membership.TotalPoint)
	if result.Error != nil {
		return nil, result.Error
	}

	return membership, nil
}

func (repository *MembershipRepositoryImpl) UpdatePoint(tx *gorm.DB, membership *domain.Membership) error {
  
	if err := tx.Model(&schema.Membership{}).Where("id = ?", membership.ID).Where("deleted_at IS NULL").Update("total_point", membership.TotalPoint).Error; err != nil {
		return err
	}

	return nil
}

func (repository *MembershipRepositoryImpl) FindById(id int) (*domain.Membership, error) {
	membership := domain.Membership{}

	result := repository.DB.Where("deleted_at IS NULL").First(&membership, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &membership, nil
}

func (repository *MembershipRepositoryImpl) FindByPhoneNumber(phoneNumber string) (*domain.Membership, error) {
	membership := domain.Membership{}

	result := repository.DB.Where("phone_number = ?", phoneNumber).Where("deleted_at IS NULL").First(&membership)
	if result.Error != nil {
		return nil, result.Error
	}

	return &membership, nil
}

func (repository *MembershipRepositoryImpl) FindByName(name string) (*domain.Membership, error) {
	membership := domain.Membership{}
	result := repository.DB.Where("deleted_at IS NULL AND name LIKE ?", "%"+name+"%").Find(&membership)
	if result.Error != nil {
		return nil, result.Error
	}

	return &membership, nil
}

func (repository *MembershipRepositoryImpl) FindAll() ([]domain.Membership, int, error) {
	memberships := []domain.Membership{}

	query := "SELECT * FROM memberships WHERE deleted_at IS NULL"
	result := repository.DB.Raw(query).Scan(&memberships)
	if result.Error != nil {
		return nil, 0, result.Error
	}
	totalMembership := len(memberships)
	return memberships, totalMembership, nil
}

func (repository *MembershipRepositoryImpl) FindTopMember() ([]domain.Membership, error) {
	memberships := []domain.Membership{}

	query := "SELECT * FROM memberships WHERE deleted_at IS NULL ORDER BY point DESC LIMIT 3"
	result := repository.DB.Raw(query).Scan(&memberships)
	if result.Error != nil {
		return nil, result.Error
	}

	return memberships, nil
}

func (repository *MembershipRepositoryImpl) Delete(id int) error {
	result := repository.DB.Where("deleted_at IS NULL").Delete(&schema.Membership{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
