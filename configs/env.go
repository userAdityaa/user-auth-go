package configs

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func EnvMongoURL() string {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal(err)
	}

	return os.Getenv("MONGO_URI")
}
