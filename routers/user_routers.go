package routers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/userAdityaa/go-user-auth/controllers"
	"github.com/userAdityaa/go-user-auth/validator"
)

func UserRouters(app *fiber.App) {
	app.Post("/signup", validator.ValidateUser, controllers.SignUp)
}
