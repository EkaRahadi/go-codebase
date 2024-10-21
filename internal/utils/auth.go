package utils

import (
	"github.com/EkaRahadi/go-codebase/internal/config"
	"golang.org/x/crypto/bcrypt"
)

type AuthUtil interface {
	HashAndSalt(password string) (string, error)
	ComparePassword(hashedPassword, inputPassword string) bool
}

type authUtilImpl struct{}

func NewAuthUtil(c *config.Config) AuthUtil {
	return &authUtilImpl{}
}

func (a authUtilImpl) HashAndSalt(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}

	return string(hash), nil
}

func (a authUtilImpl) ComparePassword(hashedPassword string, inputPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(inputPassword))
	return err == nil
}
