package db

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get the MongoDB client using the connection string from .env
// Returns a pointer to a mongo.Client if successful, or an error
func GetMongoClient() (*mongo.Client, error) {

	// Get the connection string from the environment
	mongoUrl := os.Getenv("MONGO_DB_CONNECTION")

	// Connect to the MongoDB server
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}

	// Ping the MongoDB server to check if the connection is successful
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return nil, err
	}

	return client, nil

}
