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
	app.Get("/api/posts", controller.Posts)
	app.Get("/api/post/:id", controller.Post)
	app.Put("/api/post/:id", controller.UpdatePost)
	app.Get("/api/my_posts", controller.MyPosts)
	app.Delete("/api/post/:id", controller.DeletePost)

	app.Post("/api/image/", controller.Upload)
	app.Static("/api/uploads", "./media")
}
