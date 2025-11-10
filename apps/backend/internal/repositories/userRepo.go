package repositories

import (
	"context"

	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"medivu.co/auth/postgres/sqlc"
)

type UserRepo interface {
	GetByID(id int32) (*sqlc.User, error)
	GetByEmail(email string) (*sqlc.User, error)
}

type userRepoWithDB struct {
	db *pgx.Conn
}

func NewUserRepo(db *pgx.Conn) UserRepo {
	return &userRepoWithDB{
		db: db,
	}
}

func (r *userRepoWithDB) GetByID(id int32) (*sqlc.User, error) {
	ctx := context.Background()
	query := sqlc.New(r.db)
	user, err := query.GetUserByID(ctx, id)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by id from db")
	}
	return &user, nil
}

func (r *userRepoWithDB) GetByEmail(email string) (*sqlc.User, error) {
	ctx := context.Background()
	query := sqlc.New(r.db)
	user, err := query.GetUserByEmail(ctx, email)
	if err != nil {
		return nil, errors.Wrap(err, "failed to get user by email from db")
	}
	return &user, nil
}