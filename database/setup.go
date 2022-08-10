package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
	"test-gin/users/models"
)

var Connection *mongo.Database

const name = "test"

func ConnectToDatabase(ctx context.Context) *mongo.Client {

	uri := os.Getenv("DATABASE_URI")

	mongoClient, err := mongo.Connect(ctx, options.Client().ApplyURI(uri))
	if err != nil {
		log.Fatalf("Error connecting to MongoDB database: %s", err.Error())
	}

	if err := mongoClient.Ping(ctx, readpref.Primary()); err != nil {
		log.Fatalf("Error has occurred while pinging MongoDB: %s", err.Error())
	}

	log.Printf("Successfully connected to and pinged MongoDB database.")
	Connection = mongoClient.Database(name)

	return mongoClient
}

func InitializeModels(databaseConnection *mongo.Database) {
	userModels.Init(databaseConnection.Collection(userModels.CollectionName))
}
