package routes

import (
	"github.com/HoseaTirtajaya/go-fundamental/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	app.Post("/api/v1/users/register", controllers.Register)
	app.Post("/api/v1/users/login", controllers.Login)
}
