package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
	"medivu.co/auth/envs"
)

func main() {
	envs.Load()

	app := fiber.New()

	// API route
	app.Post("/api/authorize", func(c *fiber.Ctx) error {
		// parse form data
		type Request struct {
			Email               string `form:"email"`
			Password            string `form:"password"`
			ResponseType        string `form:"response_type"`
			ClientID            string `form:"client_id"`
			RedirectURI         string `form:"redirect_uri"`
			State               string `form:"state"`
			Scope               string `form:"scope"`
			CodeChallenge       string `form:"code_challenge"`
			CodeChallengeMethod string `form:"code_challenge_method"`
		}
		var req Request
		if err := c.BodyParser(&req); err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				"error": "invalid_request",
			})
		}
		fmt.Println(req)
		return c.Redirect(fmt.Sprintf("%s?code=%s&state=%s", req.RedirectURI, "codeasdf", req.State))

	})

	// SPA static files
	app.Static("/", "./public")
	app.Static("*", "./public/index.html")


	log.Fatal(app.Listen(":3000"))
}