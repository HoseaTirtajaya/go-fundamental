package main

import (
	"github.com/HoseaTirtajaya/go-fundamental/database"
	"github.com/HoseaTirtajaya/go-fundamental/routes"
	"github.com/gofiber/fiber/v2"
)

func main() {
	database.Connect()

	app := fiber.New()

	routes.SetupRoute(app)

	app.Listen(":8080")
}
