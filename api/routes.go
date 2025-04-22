package api

import (
	"github.com/gofiber/fiber/v2"
	controllers "github.com/liborioleite/mentoring-goapi/controllers/user"
)

func InitializeRoutes(app *fiber.App) {

	app.Post("/register", controllers.RegisterUser)

}
