package apiRoutes

import (
	apiHandler "fiber-starter/handler"
	"github.com/gofiber/fiber/v2"
)

func ApiRouter(app fiber.Router) {
	app.Get("/foo", apiHandler.Foo())
	app.Post("/bar", apiHandler.Bar())
}
