package controllers

import (
	"github.com/HoseaTirtajaya/go-fundamental/models"
	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *fiber.Ctx) error {
	var data map[string]string

	err := c.BodyParser(&data)
	print(data["email"])

	if err != nil {
		return err
	}

	hashPass, _ := bcrypt.GenerateFromPassword([]byte(data["password"]), 14)

	user := models.User{
		Name:     data["name"],
		Email:    data["email"],
		Password: hashPass,
	}

	return c.JSON(user)
}
