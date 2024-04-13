package main

import (
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/api/routes"
	"github.com/robert-tyssen/go-auth-jwt-demo/internal/data/db"
)

const maxTimeoutSeconds = 2

func main() {

	// Find .env file and load contents
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}

	// Print a greeting
	greeting := os.Getenv("GREETING")
	log.Println(greeting)

	// Get the MongoDB client
	mongoClient, err := db.GetMongoClient(); 
	if err != nil {
		log.Fatal("Failed to connect to MongoDB")
	}

	log.Print("Successfully connected to MongoDB")

	// Set up the router and start the server
	timeout := maxTimeoutSeconds * time.Second
	r := routes.InitRouter(timeout, mongoClient)
	r.Run()
}
