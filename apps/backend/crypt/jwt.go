package crypt

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/hex"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"medivu.co/auth/envs"
)

func NewJWTToken(claims jwt.MapClaims, ttl time.Duration) (signedToken string, expiresAt time.Time, err error) {
	privateKeyString := envs.JWTP256PrivateKeyHex()

	privateKeyBytes, err := hex.DecodeString(privateKeyString)
	if err != nil {
		return
	}

	now := time.Now()
	expiresAt = now.Add(ttl)
	token := jwt.New(jwt.SigningMethodES256)

	// Set standard headers and claims
	token.Header["kid"] = 1
	
	claims["iat"] = now.Unix()
	claims["exp"] = expiresAt.Unix()
	token.Claims = claims

	privateKey, err := ecdsa.ParseRawPrivateKey(elliptic.P256(),privateKeyBytes)
	if err != nil {
		return
	}
	signedToken, err = token.SignedString(privateKey)
	if err != nil {
		return
	}
	return
}

