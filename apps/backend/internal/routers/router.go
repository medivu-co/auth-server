package routers

import (
	"github.com/gofiber/fiber/v2"
	"medivu.co/auth/internal/repositories"
	"medivu.co/auth/internal/services"
	"medivu.co/auth/postgres"
	"medivu.co/auth/rdb"
)

type Router interface {
	Register(app fiber.Router)
}

func RegisterAPIRouters(app *fiber.App) {
	apiGroup := app.Group("/api")

	apiRouters := []Router{
		&authorizeRouter{
			prefix:     "/authorize",
			authorizeSvc: services.NewGrantCodeSvc(
				repositories.NewGrantRecordRepo(rdb.Client),
			),
			userSvc: services.NewUserSvc(
				repositories.NewUserRepo(postgres.Conn),
			),
			clientSvc: services.NewClientSvc(
				repositories.NewClientRepo(postgres.Conn),
			),
		},
		&tokenRouter{
			prefix: "/token",
			tokenSvc: services.NewTokenSvc(
				repositories.NewRefreshTokenRepo(rdb.Client),
				repositories.NewUserRepo(postgres.Conn),
			),
			grantCodeSvc: services.NewGrantCodeSvc(
				repositories.NewGrantRecordRepo(rdb.Client),
			),
		},
	}

	for _, router := range apiRouters {
		router.Register(apiGroup)
	}

	wellKnownRouter := &wellKnownRouter{
		prefix: "/.well-known",
	}
	wellKnownRouter.Register(app)
}