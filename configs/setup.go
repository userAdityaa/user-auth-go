package configs

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func init() {
	clientOptions := options.Client().ApplyURI(EnvMongoURL())

	var err error
	client, err = mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB")
}

func GetClient() *mongo.Client {
	return client
}

func GetCollection(collectionName string) *mongo.Collection {
	collection := client.Database("go-user").Collection(collectionName)
	return collection
}
