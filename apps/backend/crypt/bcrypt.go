package crypt

import (
	"github.com/pkg/errors"
	"golang.org/x/crypto/bcrypt"
)

func BcryptHash(password []byte) ([]byte, error) {
	hashed, err := bcrypt.GenerateFromPassword(password, 12)
	if err != nil {
		return nil, errors.Wrap(err, "failed to generate bcrypt hash")
	}
	return hashed, nil
}

func BcryptCompare(hashedPassword, password string) (bool) {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password)) == nil
}

