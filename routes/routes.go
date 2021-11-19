package routes

import (
	"github.com/HoseaTirtajaya/go-fundamental/controllers"
	"github.com/gofiber/fiber/v2"
)

func SetupRoute(app *fiber.App) {
	app.Get("/", controllers.Register)

}
