package api

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/liborioleite/mentoring-goapi/controllers/auth"
	controllers "github.com/liborioleite/mentoring-goapi/controllers/user"
)

func InitializeRoutes(app *fiber.App) {

	app.Post("/register", auth.Register)
	app.Post("/login", auth.Login)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	// Mentor Routes
	app.Get("/mentees", controllers.GetMentees)
	app.Get("/mentee/:id", controllers.GetMentee)
}
