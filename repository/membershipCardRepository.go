package repository

import (
	"qbills/models/domain"

	"gorm.io/gorm"
)

type MembershipCardRepository interface {
	UpdateBarcode(id int, barcode string) (*domain.Membership, error)
	FindById(id int) (*domain.Membership, error)
}

type MembershipCardRepositoryImpl struct {
	DB *gorm.DB
}

func NewMembershipCardRepository(DB *gorm.DB) MembershipCardRepository {
	return &MembershipCardRepositoryImpl{DB: DB}
}

func (repository *MembershipCardRepositoryImpl) FindById(id int) (*domain.Membership, error) {
	membership := domain.Membership{}

	result := repository.DB.First(&membership, id)
	if result.Error != nil {
		return nil, result.Error
	}
	return &membership, nil
}


func (repository *MembershipCardRepositoryImpl) UpdateBarcode(id int, barcode string) (*domain.Membership, error) {
	membership := domain.Membership{}
    result := repository.DB.Model(&membership).Where("id = ?", id).Update("Barcode", barcode)
    if result.Error != nil {
        return nil, result.Error
    }
    return &membership, nil
}

