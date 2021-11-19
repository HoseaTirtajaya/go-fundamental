package routes

import (
	"../controllers"
	"github.com/gofiber/fiber"
)

func SetupRoute(app *fiber.App) {
	app.Get("/", controllers.Register())

	app.Listen(":8080")
}
