package baserouter

import (
	"github.com/gofiber/fiber/v2"
)

func BaseRouter(server *fiber.App) {
	routes := server.Group("/")

	routes.Get("/", func(context *fiber.Ctx) error {
		return context.Render("index", fiber.Map{
			"Blabla1Danis": "Bla Bla 1'den merhaba!",
		})
	})

	routes.Get("/about-us", func(context *fiber.Ctx) error {
		return context.Render("about-us", fiber.Map{
			"Blabla2Danis": "Bla Bla 2'den merhaba!",
		})
	})
}
