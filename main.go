package main

import (
	"github.com/HoseaTirtajaya/go-fundamental/database"
	"github.com/HoseaTirtajaya/go-fundamental/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

func main() {
	database.Connect()

	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowCredentials: true,
	}))

	routes.SetupRoute(app)

	app.Listen(":8080")
}
