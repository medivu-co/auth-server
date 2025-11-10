package repositories

import (
	"context"
	"strconv"
	"time"

	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

// 2 weeks to expire
const refreshTokenExpiration = 14 * 24 * 60 * 60 * time.Second

type RefreshTokenRecord struct {
	Token    string
	UserID   int32
	UserAgent string
	LoginedAt time.Time
}

type RefreshTokenRepo interface {
	SaveRefreshToken(token string, userID int32, userAgent string, expiresAt time.Time) error
	GetTokenRecord(token string) (*RefreshTokenRecord, error)
	DeleteToken(token string) error
}

type refreshTokenRepoWithRDB struct {
	rdb *redis.Client
}

func NewRefreshTokenRepo(rdb *redis.Client) RefreshTokenRepo {
	return &refreshTokenRepoWithRDB{
		rdb: rdb,
	}
}

func (r *refreshTokenRepoWithRDB) SaveRefreshToken(token string, userID int32, userAgent string, expiresAt time.Time) error {
	// RDB Set Hash with Expiration
	err := r.rdb.HSet(context.Background(), "refresh_token:"+token, map[string]interface{}{
		"user_id":    userID,
		"user_agent": userAgent,
		"logined_at": time.Now().Unix(),
	}).Err()
	if err != nil {
		return errors.Wrap(err, "failed to save refresh token in redis")
	}
	err = r.rdb.ExpireAt(context.Background(), "refresh_token:"+token, expiresAt).Err()
	if err != nil {
		return errors.Wrap(err, "failed to set expiration for refresh token in redis")
	}
	return nil
}

func (r *refreshTokenRepoWithRDB) GetTokenRecord(token string) (*RefreshTokenRecord, error) {
	result, err := r.rdb.HGetAll(context.Background(), "refresh_token:"+token).Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get refresh token from redis")
	}
	if len(result) == 0 {
		return nil, errors.Wrap(redis.Nil, "refresh token not found")
	}
	userID, err := strconv.Atoi(result["user_id"])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse user_id from redis")
	}
	loginedAtUnix, err := strconv.ParseInt(result["logined_at"], 10, 64)
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse logined_at from redis")
	}
	return &RefreshTokenRecord{
		Token:     token,
		UserID:    int32(userID),
		UserAgent: result["user_agent"],
		LoginedAt: time.Unix(loginedAtUnix, 0),
	}, nil
}

func (r *refreshTokenRepoWithRDB) DeleteToken(token string) error {
	err := r.rdb.Del(context.Background(), "refresh_token:"+token).Err()
	if err != nil {
		return errors.Wrap(err, "failed to delete refresh token from redis")
	}
	return nil
}