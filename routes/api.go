package routes

import (
	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"
	"go-fiber-bootstrap/controller"
)

func Setup(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World!")
	})

	app.Get("/swagger/*", swagger.HandlerDefault) // default

	app.Post("/signup", controller.SignUp)

}
