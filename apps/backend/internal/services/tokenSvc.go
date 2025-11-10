package services

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/pkg/errors"
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
	RefreshTokens(accessToken *AccessToken, refreshToken *RefreshToken) (newAccessToken *AccessToken, newRefreshToken *RefreshToken, err error)
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

	// Declare token variables
	accessToken = &AccessToken{}
	refreshToken = &RefreshToken{
		UserAgent: userAgent,
	}
	
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
	refreshToken.Token = uuid.New().String()
	refreshToken.ExpireAt = time.Now().Add(time.Duration(7*24*time.Hour)) // 7 days validity

	// Save refresh token to repository
	err = s.refreshTokenRepo.SaveRefreshToken(refreshToken.Token, user.ID, userAgent, refreshToken.ExpireAt)
	if err != nil {
		return nil, nil, errors.Wrap(err, "failed to save refresh token")
	}
	return accessToken, refreshToken, nil
}

func (s *tokenSvc) RefreshTokens(accessToken *AccessToken, refreshToken *RefreshToken) (newAccessToken *AccessToken, newRefreshToken *RefreshToken, err error) {
	// TODO: Implement token refresh logic
	return nil, nil, nil
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