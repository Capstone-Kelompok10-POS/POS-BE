package repository

import (
	"fmt"
	"qbills/models/domain"
	"time"

	"gorm.io/gorm"
)

type MembershipCardRepository interface {
	PrintMembershipCard(id int) (*domain.Membership, error)
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

func (repository *MembershipCardRepositoryImpl) PrintMembershipCard(id int) (*domain.Membership, error) {
	membership, err := repository.FindById(id)
	if err != nil {
		return nil, err
	}

	AvailableDate := time.Now().AddDate(1, 0, 0)

	fmt.Println("Name:", membership.Name)
	fmt.Println("Phone_Number:", membership.Phone_Number)
	fmt.Println("CodeMember:", membership.Code_Member)
	fmt.Println("Available until: ", AvailableDate)

	return membership, nil
}
