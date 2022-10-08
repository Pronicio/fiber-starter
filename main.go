package main

import (
	apiRoutes "fiber-starter/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/limiter"
	"time"
)

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "http://127.0.0.1:3000, https://pronicio.dev",
	}))

	app.Use(limiter.New(limiter.Config{
		Expiration: 10 * time.Second,
		Max:        7,
	}))

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	api := app.Group("/api")
	apiRoutes.ApiRouter(api)

	app.Listen(":3000")
}
