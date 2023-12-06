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
	FindById(id int) (*domain.Membership, error)
<<<<<<< Updated upstream
	FindByName(name string) (*domain.Membership, error)
	FindAll() ([]domain.Membership, error)
	FindByPhoneNumber(phoneNumber string) (*domain.Membership, error)
=======
<<<<<<< Updated upstream
	FindByName(name string) (*domain.Membership, error)	
	FindAll() ([]domain.Membership, error)
	FindByTelephone(telephone string) (*domain.Membership, error)
=======
<<<<<<< Updated upstream
	FindByName(name string) (*domain.Membership, error)
	FindAll() ([]domain.Membership, error)
	FindByPhoneNumber(phoneNumber string) (*domain.Membership, error)
=======
<<<<<<< Updated upstream
	FindByName(name string) (*domain.Membership, error)	
	FindAll() ([]domain.Membership, error)
	FindByTelephone(telephone string) (*domain.Membership, error)
=======
	FindByName(name string) (*domain.Membership, error)
	FindAll() ([]domain.Membership, int,  error)
	FindByPhoneNumber(phoneNumber string) (*domain.Membership, error)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
>>>>>>> Stashed changes
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
		PhoneNumber: membership.PhoneNumber})
	if result.Error != nil {
		return nil, result.Error
	}

	return membership, nil
}

func (repository *MembershipRepositoryImpl) UpdatePoint(tx *gorm.DB, membership *domain.Membership) error {

	if err := tx.Model(&schema.Membership{}).Where("id = ?", membership.ID).Where("deleted_at IS NULL").Update("point", membership.Point).Error; err != nil {
		return err
	}

	return nil
}

func (repository *MembershipRepositoryImpl) FindById(id int) (*domain.Membership, error) {
	membership := domain.Membership{}

	result := repository.DB.First(&membership, id)
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

<<<<<<< Updated upstream
func (repository *MembershipRepositoryImpl) FindAll() ([]domain.Membership, error) {
	membership := []domain.Membership{}
<<<<<<< Updated upstream

	result := repository.DB.Find(&membership)
=======
	query := "SELECT * FROM memberships WHERE deleted_at IS NULL"
	result := repository.DB.Raw(query).Scan(&membership)
=======
func (repository *MembershipRepositoryImpl) FindAll() ([]domain.Membership, int, error) {
	memberships := []domain.Membership{}

	result := repository.DB.Where("deleted_at IS NULL").Find(&memberships)
>>>>>>> Stashed changes
>>>>>>> Stashed changes
	if result.Error != nil {
		return nil, 0, result.Error
	}
	totalMembership := len(memberships)
	return memberships, totalMembership, nil
}

func (repository *MembershipRepositoryImpl) Delete(id int) error {
	result := repository.DB.Delete(&schema.Membership{}, id)
	if result.Error != nil {
		return result.Error
	}
	return nil

}
