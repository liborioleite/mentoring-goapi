package api

import "github.com/gofiber/fiber/v2"

func InitializeRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/home", func(c *fiber.Ctx) error {
		return c.SendString("Home Page")
	})
}
