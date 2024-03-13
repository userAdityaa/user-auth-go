package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/userAdityaa/go-user-auth/routers"
)

func main() {
	app := fiber.New()

	routers.UserRouters(app)

	app.Listen(":3000")
}
