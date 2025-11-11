package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
	"github.com/redis/go-redis/v9"
	"medivu.co/auth/crypt"
	"medivu.co/auth/envs"
	"medivu.co/auth/internal/repositories"
)

const (
	DefaultScope = "user:rw"
)

var (
	ErrUserNotFound = errors.New("user not found")
)

type TokenSvc interface {
	GenerateTokens(userID int32, userAgent string) (accessToken *AccessToken, refreshToken *RefreshToken, err error)
	RefreshTokens(refreshTokenString string) (newAccessToken *AccessToken, newRefreshToken *RefreshToken, err error)
}

type tokenSvc struct {
	refreshTokenRepo repositories.RefreshTokenRepo
	userRepo         repositories.UserRepo
}
func NewTokenSvc(refreshTokenRepo repositories.RefreshTokenRepo, userRepo repositories.UserRepo) TokenSvc {
	return &tokenSvc{
		refreshTokenRepo: refreshTokenRepo,
		userRepo:         userRepo,
	}
}

func (s *tokenSvc) GenerateTokens(userID int32, userAgent string) (accessToken *AccessToken, refreshToken *RefreshToken, err error) {
	user, err := s.userRepo.GetByID(userID)
	if err != nil {
		if errors.Cause(err) == pgx.ErrNoRows {
			return nil, nil, errors.Wrap(ErrUserNotFound, "user not found for token generation")
		}
		return nil, nil, errors.Wrap(err, "failed to get user by ID for token generation")
	}

	// Declare access token variable
	accessToken = &AccessToken{}
	
	// Generate refresh token
	accessToken.Token, accessToken.ExpireAt, err = crypt.NewJWTToken(
		jwt.MapClaims{
			"user_id": user.ID,
			"scope":   DefaultScope,
		},
		time.Duration(envs.JWTExpirationSec())*time.Second,
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate access token")
	}

	// Generate refresh token
	refreshToken = &RefreshToken{
		Token:     uuid.New().String(),
		UserAgent: userAgent,
		ExpireAt:  time.Now().Add(time.Duration(7*24*time.Hour)),
	}

	// Save refresh token to repository
	err = s.refreshTokenRepo.SaveRefreshToken(refreshToken.Token, user.ID, userAgent, refreshToken.ExpireAt)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to save refresh token")
	}
	return accessToken, refreshToken, nil
}

func (s *tokenSvc) RefreshTokens(refreshTokenString string) (newAccessToken *AccessToken, newRefreshToken *RefreshToken, err error) {
	
	tokenRecord, err := s.refreshTokenRepo.GetTokenRecord(refreshTokenString)
	if err != nil {
		if errors.Cause(err) == redis.Nil {
			return nil, nil, errors.Wrap(ErrUserNotFound, "refresh token not found")
		}
		return nil, nil, errors.Wrap(err, "failed to get refresh token record")
	}
	newAccessToken = &AccessToken{}

	newAccessToken.Token, newAccessToken.ExpireAt, err = crypt.NewJWTToken(
		jwt.MapClaims{
			"user_id": tokenRecord.UserID,
			"scope":   DefaultScope,
		},
		time.Duration(envs.JWTExpirationSec())*time.Second,
	)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to generate access token")
	}


	newRefreshToken = &RefreshToken{
		Token:     tokenRecord.Token,
		ExpireAt:  time.Now().Add(time.Duration(7*24*time.Hour)), // 7 days validity,
		UserAgent: tokenRecord.UserAgent,
	}
	// Update refresh token expiration in repository
	err = s.refreshTokenRepo.DeleteToken(refreshTokenString)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to delete old refresh token")
	}
	err = s.refreshTokenRepo.SaveRefreshToken(newRefreshToken.Token, tokenRecord.UserID, newRefreshToken.UserAgent, newRefreshToken.ExpireAt)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to save new refresh token")
	}
	return newAccessToken, newRefreshToken, nil
}

type AccessToken struct {
	Token    string
	ExpireAt time.Time
	Scope    string
}

type RefreshToken struct {
	Token     string
	ExpireAt  time.Time
	UserAgent string
}