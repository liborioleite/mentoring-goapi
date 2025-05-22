package api

import (
	jwtware "github.com/gofiber/contrib/jwt"
	"github.com/gofiber/fiber/v2"
	"github.com/liborioleite/mentoring-goapi/controllers/auth"
	controllers "github.com/liborioleite/mentoring-goapi/controllers/user"
)

func InitializeRoutes(app *fiber.App) {

	app.Post("/register/mentor", auth.RegisterMentor)
	app.Post("/register/mentee", auth.RegisterMentee)
	app.Post("/login/mentor", auth.LoginMentor)
	app.Post("/login/mentee", auth.LoginMentee)

	// JWT Middleware
	app.Use(jwtware.New(jwtware.Config{
		SigningKey: jwtware.SigningKey{Key: []byte("secret")},
	}))

	// Mentor Routes
	app.Get("/mentees", controllers.GetMentees)
	app.Get("/mentee/:id", controllers.GetMentee)
	app.Put("/me/update/:id", controllers.UpdateMe)
	app.Put("/me/status-account/:id", controllers.ChangeMentorStatus) //REVER LOGICA

	// Mentee Routes
	app.Get("/mentors", controllers.GetMentors)
	app.Get("/mentee/:id", controllers.GetMentee)
	app.Put("/mentee/update/:id", controllers.UpdateMentee)
	app.Put("/mentee/status-account/:id", controllers.ChangeMenteeStatus) //REVER LOGICA
}
