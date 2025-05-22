package api

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func InitializeFiber() {
	app := fiber.New(fiber.Config{
		AppName: "Mentoring Api",
	})

	// Configura o middleware CORS
	app.Use(cors.New(cors.Config{
		AllowOrigins: "*", // Permite requisições de qualquer origem
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Origin, Content-Type, Accept, Authorization",
	}))

	InitializeRoutes(app)

	app.Listen(":3000")

}
