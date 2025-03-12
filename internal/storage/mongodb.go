package storage

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo/options"

	"go.mongodb.org/mongo-driver/mongo"
)

// EnvMongoURI gets MongoDB connection URI from .env file
func EnvMongoURI() string {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	return os.Getenv("DATABASE_URL")
}

// ConnectDB establishes connection to MongoDB
func ConnectDB() *mongo.Client {
	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()

	// mongo.Connect returns mongo.Client method
	uri := EnvMongoURI()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}

	// ping the database
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal("Error: " + err.Error())
	}
	log.Println("Connected to MongoDB")
	return client
}
