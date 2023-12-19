package password

import (
	"os"

	"golang.org/x/crypto/bcrypt"
)


type PasswordHandler interface {
	HashPassword(password string) string
	ComparePassword(hash, password string) error
}

type PasswordHandlerImpl struct {}

func NewPasswordHandler() PasswordHandler {
	return &PasswordHandlerImpl{}
}

func (p *PasswordHandlerImpl) HashPassword(password string) string {
	result, _ := bcrypt.GenerateFromPassword([]byte(password + os.Getenv("SALT")), bcrypt.DefaultCost)
	return string(result)
}

func (p *PasswordHandlerImpl) ComparePassword(hash, password string) error {
	passwordWithSalt := []byte(password + os.Getenv("SALT"))
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(passwordWithSalt))
	if err != nil {
		return err
	}

	return nil
}