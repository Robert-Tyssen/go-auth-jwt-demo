package db

import (
	"context"
	"os"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Get the MongoDB client using the connection string from .env
// Returns a pointer to a mongo.Client if successful, or an error
func GetMongoClient() (*mongo.Client, error) {

	// Get the connection string from the environment
	mongoUrl := os.Getenv("MONGO_DB_CONNECTION")

	// Create a context with a timeout
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	defer cancel()

	// Connect to the MongoDB server
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(mongoUrl))
	if err != nil {
		return nil, err
	}

	return client, nil

}
