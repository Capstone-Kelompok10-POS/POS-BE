package services

import (
	"fmt"
	"qbills/models/domain"
	"qbills/repository"
<<<<<<< Updated upstream
	// "qbills/utils/helpers/firebase"

	// "github.com/go-playground/validator"
=======
	"qbills/utils/helpers/firebase"

>>>>>>> Stashed changes
	"github.com/labstack/echo/v4"
)

type MembershipCardService interface {
	PrintMembershipCard(ctx echo.Context, id int) (*domain.Membership, error)
<<<<<<< Updated upstream
	FindById(ctx echo.Context, id int) (*domain.Membership, error)
=======
	UploadBarcodeToFirebase(ctx echo.Context, membership domain.Membership) (string, error)
>>>>>>> Stashed changes
}

type MembershipCardServiceImpl struct {
	MembershipCardRepository repository.MembershipCardRepository
	// Validate                 *validator.Validate
}

func NewMembershipCardService(membershipCardRepository repository.MembershipCardRepository) *MembershipCardServiceImpl {
	return &MembershipCardServiceImpl{
		MembershipCardRepository: membershipCardRepository,
		// Validate:                 validate,
	}
}

<<<<<<< Updated upstream
func (service *MembershipCardServiceImpl) FindById(ctx echo.Context, id int) (*domain.Membership, error) {
	existingMembership, _ := service.MembershipCardRepository.PrintMembershipCard(ctx, id)
	if existingMembership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	return existingMembership, nil
}

func (service *MembershipCardServiceImpl) PrintMembershipCard(ctx echo.Context, id int) (*domain.Membership, error) {
	existingMembership, _ := service.MembershipCardRepository.FindById(id)
	if existingMembership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	result, err := service.MembershipCardRepository.PrintMembershipCard(ctx, id)
=======
func (service *MembershipCardServiceImpl) PrintMembershipCard(ctx echo.Context, id int) (*domain.Membership, error) {
	membership, _ := service.MembershipCardRepository.FindById(id)
	if membership == nil {
		return nil, fmt.Errorf("membership not found")
	}

	barcode, err := service.UploadBarcodeToFirebase(ctx, *membership)
	if err != nil {
        return nil, fmt.Errorf("error uploading barcode %s", err.Error())
    }

	membership.Barcode = barcode

    // Update hanya kolom barcode di database
	result, err := service.MembershipCardRepository.UpdateBarcode(int(membership.ID), membership.Barcode)
>>>>>>> Stashed changes
	if err != nil {
		return nil, fmt.Errorf("error creating membership card: %s", err.Error())
	}

	return result, nil
<<<<<<< Updated upstream
=======
}

func (repository *MembershipCardServiceImpl) UploadBarcodeToFirebase(ctx echo.Context, membership domain.Membership) (string, error) {
	barcode, err := firebase.GenerateBarcodeAndUploadToFirebase(ctx, membership.CodeMember.String())
	if err != nil {
		return "", fmt.Errorf("error upload %s", err.Error())
	}
	return barcode, nil
>>>>>>> Stashed changes
}