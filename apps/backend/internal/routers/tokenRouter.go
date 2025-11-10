package routers

import (
	"github.com/gofiber/fiber/v2"
	"medivu.co/auth/internal/controllers"
	"medivu.co/auth/internal/services"
)

type tokenRouter struct {
	prefix       string
	tokenSvc     services.TokenSvc
	grantCodeSvc services.GrantCodeSvc
}

func (r *tokenRouter) Register(rootRouter fiber.Router) {
	tokenGroup := rootRouter.Group(r.prefix)
	tokenController := controllers.NewTokenCtrl(r.tokenSvc, r.grantCodeSvc)
	
	tokenGroup.Post("/", func(c *fiber.Ctx) error {
		return tokenController.ExchangeCodeToToken(c)
	})
}