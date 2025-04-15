package api

import (
	"github.com/gofiber/fiber/v2"
)

func InitializeFiber() {
	app := fiber.New(fiber.Config{
		AppName: "Mentoring Api",
	})

	InitializeRoutes(app)

	app.Listen(":3000")

}
