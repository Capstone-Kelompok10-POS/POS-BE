package repository

import (
	"qbills/models/domain"

	"gorm.io/gorm"
)

type MembershipRepository interface {
	Create(membership *domain.Membership) (*domain.Membership, error)
	Update(membership *domain.Membership, id int) (*domain.Membership, error)
	FindById(id int) (*domain.Membership, error)
	FindByName(name string) (*domain.Membership, error)	
	FindAll() ([]domain.Membership, error)
	Delete(id int) error
}

type MembershipRepositoryImpl struct {
	DB *gorm.DB
}

// func NewMembershipRepository(DB *gorm.DB) MembershipRepository {
// 	return &MembershipRepositoryImpl{DB: DB}
// }

// func (repository *MembershipRepositoryImpl) Create(membership *domain.Membership) (*domain.Membership, error) {
// 	// membershipDB := req.Membership
// }
