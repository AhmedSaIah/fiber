package user

import (
	"github.com/gofiber/fiber/v2"
	_ "github.com/gofiber/fiber/v2"
)

type SignUpData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type SignUpController struct{}

func (c *SignUpController) Signup(ctx *fiber.Ctx) error {
	data := new(SignUpData)
	if err := ctx.BodyParser(data); err != nil {
		return fiber.NewError(fiber.StatusBadRequest, "Invalid request format")
	}

	if data.Email == "" || data.Password == "" {
		return fiber.NewError(fiber.StatusBadRequest, "Please provide Email and Password")
	}

	// Save user to database

	return ctx.Status(fiber.StatusCreated).JSON(fiber.Map{"message": "User registered successfully"})
}
