package routes

import (
	"github.com/gofiber/fiber/v2"
	controller "github.com/habibbushira/goblog/Controller"
	"github.com/habibbushira/goblog/middleware"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)

	app.Use(middleware.IsAuthenticated)
	app.Post("/api/post", controller.CreatePost)
}
