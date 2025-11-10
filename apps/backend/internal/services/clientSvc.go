package services

import (
	"slices"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"medivu.co/auth/internal/repositories"
)

type ClientSvc interface {
	IsClientValid(clientID uuid.UUID, redirectURI string) (bool, error)
}

type clientSvc struct {
	clientRepo repositories.ClientRepo
}

func NewClientSvc(clientRepo repositories.ClientRepo) ClientSvc {
	return &clientSvc{
		clientRepo: clientRepo,
	}
}
func (s *clientSvc) IsClientValid(clientID uuid.UUID, redirectURI string) (bool, error) {
	client, err := s.clientRepo.FindByID(clientID)
	if err != nil {
		if errors.Cause(err) == pgx.ErrNoRows {
			return false, nil
		} else {
			return false, errors.Wrap(err, "failed to find client by ID")
		}
	}
	if !slices.Contains(client.AllowedRedirectUris, redirectURI) {
		return false, nil
	}
	return true, nil
}