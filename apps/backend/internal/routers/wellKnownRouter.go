package routers

import (
	"github.com/gofiber/fiber/v2"
	"medivu.co/auth/internal/controllers"
)

type wellKnownRouter struct {
	prefix       string
}

func (r *wellKnownRouter) Register(rootRouter fiber.Router) {
	wellKnownGroup := rootRouter.Group(r.prefix)
	wellKnownController := controllers.NewWellKnownCtrl()

	wellKnownGroup.Get("/openid-configuration", func(c *fiber.Ctx) error {
		return wellKnownController.GetOpenIDConfig(c)
	})

	wellKnownGroup.Get("/jwks.json", func(c *fiber.Ctx) error {
		return wellKnownController.GetJWKS(c)
	})
}