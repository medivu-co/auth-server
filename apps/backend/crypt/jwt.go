package crypt

import (
	"time"

	"github.com/golang-jwt/jwt/v5"
	"medivu.co/auth/envs"
)

func NewJWTToken(claims jwt.MapClaims, ttl time.Duration) (signedToken string, expiresAt time.Time, err error) {
	secret := []byte(envs.JWTSecretKey())

	expiresAt = time.Now().Add(ttl)
	claims["exp"] = expiresAt.Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err = token.SignedString(secret)
	if err != nil {
		return
	}
	return
}

func ParseJWTToken(signedToken string) (claims jwt.MapClaims, err error) {
	secret := []byte(envs.JWTSecretKey())
	
	token, err := jwt.Parse(signedToken, func(token *jwt.Token) (any, error) {
		return secret, nil
	})
	if err != nil {
		return
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		err = jwt.ErrTokenInvalidClaims
		return
	}
	return claims, nil
}