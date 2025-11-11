package main

import (
	"errors"

	"github.com/gofiber/contrib/fiberzap/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"medivu.co/auth/envs"
	"medivu.co/auth/internal/routers"
	"medivu.co/auth/logger"
	"medivu.co/auth/postgres"
	"medivu.co/auth/rdb"
	"medivu.co/auth/validator"
)

func main() {
	// Load environment variables
	envs.Load()

	// Initialize logger
	logger.Init()
	defer logger.Sync()

	// Connect to Redis
	rdb.Connect()
	defer rdb.Close()

	// Connect to Postgres
	postgres.Connect()
	defer postgres.Close()

	// Initialize validator
	validator.Init()

	// Initialize Fiber app
	app := fiber.New(fiber.Config{
		ErrorHandler: func(c *fiber.Ctx, err error) error {
			// Status code defaults to 500
			code := fiber.StatusInternalServerError

			// Retrieve the custom status code if it's a *fiber.Error
			var e *fiber.Error
			if errors.As(err, &e) {
				code = e.Code
			}
			// log the error message
			if code >= 500 {
				logger.Get().Error("Internal Server Error: " + err.Error())
				return c.Status(code).JSON(fiber.Map{
					"error": "internal_server_error",
				})
			}
			return c.Status(code).JSON(fiber.Map{
				"error": e.Message,
			})
		},
	})
	// CORS Middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://medivu.co, https://*.medivu.co, https://localhost:5173",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
		AllowCredentials: true,
	}))

	// Middleware(Logger)
	app.Use(fiberzap.New(fiberzap.Config{
		Logger: logger.Get(),
	}))

	// API route
	routers.RegisterAPIRouters(app)

	// SPA static files
	app.Static("/", "./public")
	app.Static("*", "./public/index.html")


	logger.Get().Fatal(app.Listen(":3000").Error())
}