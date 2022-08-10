package userModels

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,unique,omitempty"`
	Username string             `bson:"username,unique,omitempty"`
	Password string             `bson:"password,omitempty"`
}

const CollectionName = "users"

func Init(collection *mongo.Collection) {
	_, err := collection.Indexes().CreateOne(
		context.TODO(),
		mongo.IndexModel{
			Keys:    bson.D{{Key: "username", Value: 1}},
			Options: options.Index().SetUnique(true),
		})
	if err != nil {
		log.Fatalf("Unable to create index for User model: %s\n", err.Error())
	}
}
