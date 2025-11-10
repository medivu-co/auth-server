package repositories

import (
	"context"
	"strconv"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
)

type GrantRecordRepo interface{
	SaveCode(code string, codeChallenge string, clientID uuid.UUID, redirectURI string, scope string, userID int32) error
	GetByCode(code string) (*GrantRecord, error)
	DeleteByCode(code string) error
}

type GrantRecord struct {
	Code          string
	CodeChallenge string
	ClientID      uuid.UUID
	RedirectURI   string
	UserID        int32
}

type grantRecordRepoWithRDB struct {
	rdb *redis.Client
}

func NewGrantRecordRepo(rdb *redis.Client) GrantRecordRepo {
	return &grantRecordRepoWithRDB{
		rdb: rdb,
	}
}
func (r *grantRecordRepoWithRDB) SaveCode(code string, codeChallenge string, clientID uuid.UUID, redirectURI string, scope string, userID int32) error {
	// RDB Set Hash with Expiration
	err := r.rdb.HSet(context.Background(), "grant_record:"+code, map[string]interface{}{
		"code_challenge": codeChallenge,
		"client_id":      clientID.String(),
		"redirect_uri":   redirectURI,
		"scope":          scope,
		"user_id":       userID,
	}).Err()
	if err != nil {
		return errors.Wrap(err, "failed to save grant record in redis")
	}
	err = r.rdb.Expire(context.Background(), "grant_record:"+code, time.Minute*3).Err()
	if err != nil {
		return errors.Wrap(err, "failed to set expiration for grant record in redis")	
	}
	return nil
}

func (r *grantRecordRepoWithRDB) GetByCode(code string) (*GrantRecord, error) {
	// TODO: Implement the logic to find the grant record by code in Redis
	result, err := r.rdb.HGetAll(context.Background(), "grant_record:"+code).Result()
	if err != nil {
		return nil, errors.Wrap(err, "failed to get grant record from redis")
	}
	if len(result) == 0 {
		return nil, errors.Wrap(redis.Nil, "grant record not found")
	}
	userID, err := strconv.Atoi(result["user_id"])
	if err != nil {
		return nil, errors.Wrap(err, "failed to convert user_id to int32")
	}
	clientID, err := uuid.Parse(result["client_id"])
	if err != nil {
		return nil, errors.Wrap(err, "failed to parse client_id to uuid")
	}	
	return &GrantRecord{
		Code:          code,
		CodeChallenge: result["code_challenge"],
		ClientID:      clientID,
		RedirectURI:   result["redirect_uri"],
		UserID:       int32(userID),
	}, nil
}
func (r *grantRecordRepoWithRDB) DeleteByCode(code string) error {
	err := r.rdb.Del(context.Background(), "grant_record:"+code).Err()
	if err != nil {
		return errors.Wrap(err, "failed to delete grant record from redis")
	}
	return nil
}
	
