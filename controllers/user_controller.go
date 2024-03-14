package controllers

import (
	"log"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"

	"github.com/userAdityaa/go-user-auth/configs"
	"github.com/userAdityaa/go-user-auth/models"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var collection *mongo.Collection = configs.GetCollection("users")

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 12)
	return string(bytes), err
}

func CompareHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func createToken(username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,
		jwt.MapClaims{
			"username": username,
			"exp":      time.Now().Add(time.Hour * 3).Unix(),
		})

	tokenString, err := token.SignedString([]byte(configs.EnvSecretKey()))

	if err != nil {
		log.Fatal(err)
	}

	return tokenString, nil
}

// func verifyToken(tokenString string) error {
// 	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
// 		return configs.EnvSecretKey(), nil
// 	})

// 	if err != nil {
// 		return err
// 	}

// 	if !token.Valid {
// 		return errors.New("invalid token")
// 	}

// 	return nil
// }

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

func Login(c *fiber.Ctx) error {
	var user models.User

	err := c.BodyParser(&user)

	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"email": user.Email}

	var dbUser bson.M

	err = collection.FindOne(c.Context(), filter).Decode(&dbUser)

	if err != nil {
		log.Fatal(err)
	}

	dbPass, ok := dbUser["password"].(string)

	if !ok {
		log.Fatal(err)
	}

	if !CompareHash(user.Password, dbPass) {
		return c.Status(fiber.StatusBadRequest).JSON("User password is not correct.")
	}

	// user Hain:

	key, err := createToken(user.Username)

	if err != nil {
		log.Fatal(err)
	}

	return c.Status(fiber.StatusAccepted).JSON(key)
}
