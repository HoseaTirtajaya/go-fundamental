package routes

import (
	"github.com/HoseaTirtajaya/go-fundamental/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	app.Post("/api/v1/user/register", controllers.Register)
	app.Post("/api/v1/user/login", controllers.Login)
	app.Get("/api/v1/user/readme", controllers.ReadUser)
	app.Post("api/v1/user/logout", controllers.Logout)
}
