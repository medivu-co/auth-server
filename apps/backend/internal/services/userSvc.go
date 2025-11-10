package services

import (
	"github.com/pkg/errors"
	"medivu.co/auth/crypt"
	"medivu.co/auth/internal/repositories"
	"medivu.co/auth/postgres/sqlc"
)

var (
	ErrInvalidCredentials = errors.New("invalid credentials")
)

type UserSvc interface {
	BasicAuthenticate(email, password string) (*sqlc.User, error)
}

type userSvc struct {
	userRepo repositories.UserRepo
}

func NewUserSvc(userRepo repositories.UserRepo) UserSvc {
	return &userSvc{
		userRepo: userRepo,
	}
}

func (s *userSvc) BasicAuthenticate(email, password string) (*sqlc.User, error) {
	user, err := s.userRepo.GetByEmail(email)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by email")
	}
	if !crypt.BcryptCompare(user.PasswordHash, password) {
		return nil, errors.Wrap(ErrInvalidCredentials, "invalid password")
	}
	return user, nil
}

