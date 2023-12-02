package repository

import (
<<<<<<< Updated upstream
	"fmt"
	"qbills/models/domain"
	"qbills/utils/helpers/firebase"

	"github.com/labstack/echo/v4"
=======
	"qbills/models/domain"

>>>>>>> Stashed changes
	"gorm.io/gorm"
)

type MembershipCardRepository interface {
<<<<<<< Updated upstream
	UploadBarcodeToFirebase(ctx echo.Context, membership domain.Membership) (string, error)
	UpdateBarcode(id int, barcode string) error
	PrintMembershipCard(ctx echo.Context, id int) (*domain.Membership, error)
=======
	UpdateBarcode(id int, barcode string) (*domain.Membership, error)
>>>>>>> Stashed changes
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

<<<<<<< Updated upstream
func (repository *MembershipCardRepositoryImpl) UploadBarcodeToFirebase(ctx echo.Context, membership domain.Membership) (string, error) {
	barcode, err := firebase.GenerateBarcodeAndUploadToFirebase(ctx, membership.CodeMember.String())
	if err != nil {
		return "", fmt.Errorf("error upload %s", err.Error())
	}
	return barcode, nil
}

func (repository *MembershipCardRepositoryImpl) UpdateBarcode(id int, barcode string) error {
    result := repository.DB.Model(&domain.Membership{}).Where("id = ?", id).Update("Barcode", barcode)
    if result.Error != nil {
        return result.Error
    }
    return nil
}

func (repository *MembershipCardRepositoryImpl) PrintMembershipCard(ctx echo.Context, id int) (*domain.Membership, error) {
    membership, err := repository.FindById(id)
    if err != nil {
        return nil, err
    }

	barcode, err := repository.UploadBarcodeToFirebase(ctx, *membership)
	if err != nil {
        return nil, fmt.Errorf("error uploading barcode %s", err.Error())
    }

	membership.Barcode = barcode

    // Update hanya kolom barcode di database
    if err := repository.UpdateBarcode(int(membership.ID), barcode); err != nil {
        return nil, fmt.Errorf("error updating barcode in membership record %s", err.Error())
    }

    return membership, nil
}
=======


func (repository *MembershipCardRepositoryImpl) UpdateBarcode(id int, barcode string) (*domain.Membership, error) {
	membership := domain.Membership{}
    result := repository.DB.Model(&membership).Where("id = ?", id).Update("Barcode", barcode)
    if result.Error != nil {
        return nil, result.Error
    }
    return &membership, nil
}

>>>>>>> Stashed changes
