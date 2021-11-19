package controllers

import "github.com/gofiber/fiber"

func Register(c *fiber.Ctx) error {
	return c.SendString("Hello, World 👋!")
}
