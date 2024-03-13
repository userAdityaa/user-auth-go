package validator

import (
	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/userAdityaa/go-user-auth/models"
)

type IError struct {
	Field string
	Tag   string
	Value string
}

var Validator = validator.New()

func ValidateUser(c *fiber.Ctx) error {
	var errors []*IError
	body := new(models.User)
	err := c.BodyParser(&body)

	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			"error": "Cannot parse JSON",
		})
	}

	err = Validator.Struct(body)

	if err != nil {
		for _, err := range err.(validator.ValidationErrors) {
			var el IError
			el.Field = err.Field()
			el.Tag = err.Tag()
			el.Value = err.Param()
			errors = append(errors, &el)
		}
		return c.Status(fiber.StatusBadRequest).JSON(errors)
	}

	return c.Next()
}
