package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/habibbushira/goblog/Controller"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
}
