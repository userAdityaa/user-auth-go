package controllers

import (
	"log"

	"github.com/gofiber/fiber/v2"

	"github.com/userAdityaa/go-user-auth/configs"
	"github.com/userAdityaa/go-user-auth/models"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var collection *mongo.Collection = configs.GetCollection("users")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CompareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(password), []byte(hash))
	return err == nil
}

func SignUp(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)

	if err != nil {
		log.Fatal(err)
	}

	user.Password, err = HashPassword(user.Password)

	if err != nil {
		log.Fatal(err)
	}

	_, err = collection.InsertOne(c.Context(), user)

	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusAccepted).JSON("User Created")
}
