package routers

import (
	"github.com/gofiber/fiber/v2"
	"medivu.co/auth/internal/controllers"
	"medivu.co/auth/internal/services"
)

type authorizeRouter struct {
	prefix      string
	authorizeSvc services.GrantCodeSvc
	userSvc     services.UserSvc
	clientSvc   services.ClientSvc
}

func (r *authorizeRouter) Register(rootRouter fiber.Router) {
	authorizeGroup := rootRouter.Group(r.prefix)
	authorizeController := controllers.NewAuthorizeCtrl(r.authorizeSvc, r.userSvc, r.clientSvc)

	authorizeGroup.Post("/", func(c *fiber.Ctx) error {
		return authorizeController.Authorize(c)
	})
}
