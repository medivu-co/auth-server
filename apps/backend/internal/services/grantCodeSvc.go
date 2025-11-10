package services

import (
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"medivu.co/auth/crypt"
	"medivu.co/auth/internal/repositories"
	"medivu.co/auth/postgres/sqlc"
)

var (
	ErrInvalidGrantCode = errors.New("invalid grant code record data")
)

type GrantCodeSvc interface {
	GenerateGrantCode(user *sqlc.User, clientID uuid.UUID, redirectURI, scope, codeChallenge string) (grantCode string, err error)
	GetUserIDFromCode(code, codeVerifier, clientID, redirectURI string) (userID int32, err error)
}

type grantCode struct {
	grantCodeRepo repositories.GrantRecordRepo
}

func NewGrantCodeSvc(grantCodeRepo repositories.GrantRecordRepo) GrantCodeSvc {
	return &grantCode{
		grantCodeRepo: grantCodeRepo,
	}
}


func (s *grantCode) GenerateGrantCode(user *sqlc.User, clientID uuid.UUID, redirectURI, scope, codeChallenge string) (grantCode string, err error) {
	// Generate a grant code
	grantCode = uuid.New().String()
	err = s.grantCodeRepo.SaveCode(grantCode, codeChallenge, clientID, redirectURI, scope, user.ID)
	if err != nil {
		return "", errors.Wrap(err, "failed to save grant code")
	}
	return grantCode, nil
}

func (s *grantCode) GetUserIDFromCode(code, codeVerifier, clientID, redirectURI string) (userID int32, err error) {
	grantCodeRecord, err := s.grantCodeRepo.GetByCode(code)
	if err != nil {
		if errors.Cause(err) == redis.Nil {
			return 0, errors.Wrap(ErrInvalidGrantCode, "grant code not found")
		}
		return 0, errors.Wrap(err, "failed to get grant code record by code")
	}
	if grantCodeRecord.CodeChallenge != crypt.SHA256Hex([]byte(codeVerifier)) && grantCodeRecord.ClientID.String() != clientID && grantCodeRecord.RedirectURI != redirectURI {
		return 0, errors.Wrap(ErrInvalidGrantCode, "failed to validate grant code")
	}
	err = s.grantCodeRepo.DeleteByCode(code)
	if err != nil {
		return 0, errors.Wrap(err, "failed to delete grant code after use")
	}
	return grantCodeRecord.UserID, nil
}


