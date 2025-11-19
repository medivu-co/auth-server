package controllers

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"encoding/base64"
	"encoding/hex"
	"fmt"

	"github.com/gofiber/fiber/v2"
	"medivu.co/auth/envs"
)

type WellKnownCtrl struct {
}

func NewWellKnownCtrl() *WellKnownCtrl {
	return &WellKnownCtrl{
	}
}

func (ctrl *WellKnownCtrl) GetOpenIDConfig(c *fiber.Ctx) error {
	publicHostURL := envs.PublicHostURL()
	return c.JSON(fiber.Map{
		"issuer":                                publicHostURL,
		"authorization_endpoint":                publicHostURL + "/authorize",
		"token_endpoint":                        publicHostURL + "/token",
		"userinfo_endpoint":                     publicHostURL + "/userinfo",
		"jwks_uri":                             publicHostURL + "/.well-known/jwks.json",
		"id_token_signing_alg_values_supported": []string{"ES256"},
	})
}

func (ctrl *WellKnownCtrl) GetJWKS(c *fiber.Ctx) error {

	pubKeyHexString := envs.JWTP256PublicKeyHex()
	pubKeyBytes, err := hex.DecodeString(pubKeyHexString)
	if err != nil {
		return fiber.NewError(500, "failed to decode public key")
	}
	pubKey, err := ecdsa.ParseUncompressedPublicKey(elliptic.P256(), pubKeyBytes)
	if err != nil {
		fmt.Println(err)
		return fiber.NewError(500, "failed to parse public key")
	}

	xBase64SafeString := base64.RawURLEncoding.EncodeToString(pubKey.X.Bytes())
	yBase64SafeString := base64.RawURLEncoding.EncodeToString(pubKey.Y.Bytes())
	

	return c.JSON(fiber.Map{
		"keys": []fiber.Map{
			{
				"kty": "EC",
				"use": "sig",
				"crv": "P-256",
				"kid": "1",
				"x":  xBase64SafeString,
				"y":  yBase64SafeString,
			},
		},
	})


}

