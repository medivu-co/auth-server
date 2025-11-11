package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/pkg/errors"
	"medivu.co/auth/envs"
	"medivu.co/auth/internal/services"
	"medivu.co/auth/validator"
)

type TokenCtrl struct {
	tokenSvc    services.TokenSvc
	grantCodeSvc services.GrantCodeSvc
}

func NewTokenCtrl(tokenSvc services.TokenSvc, grantCodeSvc services.GrantCodeSvc) *TokenCtrl {
	return &TokenCtrl{
		tokenSvc:    tokenSvc,
		grantCodeSvc: grantCodeSvc,
	}
}

func (ctrl *TokenCtrl) ExchangeCodeToToken(c *fiber.Ctx) error {
	// TODO: Implementation of the token endpoint
	type reqBody struct {
		Code         string `json:"code"`
		CodeVerifier string `json:"code_verifier"`
		ClientID    string `json:"client_id"`
		RedirectURI string `json:"redirect_uri"`
	}
	// Parse and validate request body
	var req reqBody
	if err := c.BodyParser(&req); err != nil {
		return fiber.NewError(400, "failed to parse token request body")
	}
	if err := validator.ValidateStruct(req); err != nil {
		return fiber.NewError(400, "invalid token request: "+err.Error())
	}
	userID, err := ctrl.grantCodeSvc.GetUserIDFromCode(req.Code, req.CodeVerifier, req.ClientID, req.RedirectURI)
	if err != nil {
		if errors.Cause(err) == services.ErrInvalidGrantCode {
			fmt.Println(err)
			return fiber.NewError(401, "invalid grant code")
		}
		return fiber.NewError(500, err.Error())
	}
	
	accessToken, refreshToken, err := ctrl.tokenSvc.GenerateTokens(userID, c.Get("User-Agent"))
	// TODO: Handle various errors from token generation
	if err != nil {
		return fiber.NewError(500, err.Error())
	}

	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",
		Value:    refreshToken.Token,
		HTTPOnly: true,
		Secure:  true,
		SameSite: "strict",
	})
	return c.JSON(fiber.Map{
		"access_token": accessToken.Token,
		"token_type":   "bearer",
		"expires_in":   envs.JWTExpirationSec(),
		"expires_at":   accessToken.ExpireAt.Unix(),
	})
}