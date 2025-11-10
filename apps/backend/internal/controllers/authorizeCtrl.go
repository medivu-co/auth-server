package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/pkg/errors"
	"medivu.co/auth/internal/services"
	"medivu.co/auth/validator"
)

type AuthorizeCtrl struct {
	grantCodeSvc services.GrantCodeSvc
	userSvc services.UserSvc
	clientSvc services.ClientSvc
}

func NewAuthorizeCtrl(authorizeSvc services.GrantCodeSvc, userSvc services.UserSvc, clientSvc services.ClientSvc) *AuthorizeCtrl {
	return &AuthorizeCtrl{
		grantCodeSvc: authorizeSvc,
		userSvc: userSvc,
		clientSvc: clientSvc,
	}
}

func (ctrl *AuthorizeCtrl) Authorize(c *fiber.Ctx) error {
	// TODO: Implementation of the authorize endpoint
	type reqBody struct {
		Email               string `form:"email" validate:"required,email"`
		Password            string `form:"password" validate:"required,printascii"`
		ResponseType        string `form:"response_type" validate:"required,eq=code"`
		ClientID            string `form:"client_id" validate:"required,uuid4_rfc4122"`
		RedirectURI         string `form:"redirect_uri" validate:"required,url"`
		State               string `form:"state" validate:"required,hexadecimal,len=64"`
		Scope               string `form:"scope" validate:"required"`
		CodeChallenge       string `form:"code_challenge" validate:"required,hexadecimal,len=64"`
		CodeChallengeMethod string `form:"code_challenge_method" validate:"required,eq=s256"`
	}
	var req reqBody
	if err := c.BodyParser(&req); err != nil {
		return c.Redirect("/error?error=invalid_request&error_description=failed to parse authorize request body", 302) // invalid request
	}
	if err := validator.ValidateStruct(req); err != nil {
		return c.Redirect("/error?error=invalid_request&error_description="+err.Error(), 302) // invalid request
	}
	clientID, _ := uuid.Parse(req.ClientID)
	isValidClient, err := ctrl.clientSvc.IsClientValid(clientID, req.RedirectURI)
	if err != nil {
		return c.Redirect("/error?error=server_error&error_description=Internal Server Error", 302) // server error
	}
	if !isValidClient {
		return c.Redirect("/error?error=unauthorized_client&error_description=Invalid client or redirect URI", 302) // invalid client
	}

	user, err := ctrl.userSvc.BasicAuthenticate(req.Email, req.Password)
	if err != nil {
		if errors.Cause(err) == services.ErrInvalidCredentials {
			return c.Redirect("/error?error=access_denied&error_description=Invalid credentials", 302) // invalid credentials
		} else {
			return c.Redirect("/error?error=server_error&error_description=Internal Server Error", 302) // server error
		}
	}

	grantCodeStr, err := ctrl.grantCodeSvc.GenerateGrantCode(user, clientID, req.RedirectURI, req.Scope, req.CodeChallenge)
	if err != nil {
		return c.Redirect("/error?error=server_error&error_description=Internal Server Error", 302)
	}
	return c.Redirect(req.RedirectURI + "?code=" + grantCodeStr + "&state=" + req.State)
}
