package main

import (
	"./database"
	"./routes"
	"github.com/gofiber/fiber"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.SetupRoute(app)

	app.Listen(":8080")
}
