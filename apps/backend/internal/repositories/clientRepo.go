package repositories

import (
	"context"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"medivu.co/auth/postgres/sqlc"
)

type ClientRepo interface {
	FindByID(clientID uuid.UUID) (*sqlc.Client, error)
}

type clientRepoWithDB struct {
	db *pgx.Conn
}
func NewClientRepo(db *pgx.Conn) ClientRepo {
	return &clientRepoWithDB{
		db: db,
	}
}
func (r *clientRepoWithDB) FindByID(clientID uuid.UUID) (*sqlc.Client, error) {
	query := sqlc.New(r.db)
	client, err := query.GetClientByID(context.Background(), clientID)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get client by ID from db")
	}
	return &client, nil
}


