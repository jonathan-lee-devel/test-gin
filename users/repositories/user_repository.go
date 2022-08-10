package repositories

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	userModels "test-gin/users/models"
)

type UserRepository interface {
	FindOne(filter interface{}) (userModels.User, error)
	InsertOne(filter interface{}) (*mongo.InsertOneResult, error)
}

type userRepository struct {
	collection *mongo.Collection
}

func NewUserRepository(collection *mongo.Collection) UserRepository {
	return &userRepository{
		collection: collection,
	}
}

func (userRepository *userRepository) FindOne(filter interface{}) (userModels.User, error) {
	var user userModels.User
	err := userRepository.collection.FindOne(context.Background(), filter).Decode(&user)
	if err != nil {
		return user, nil // Only error is ideally that no documents were found to be decoded
	}
	return user, nil
}

func (userRepository *userRepository) InsertOne(filter interface{}) (*mongo.InsertOneResult, error) {
	return userRepository.collection.InsertOne(context.Background(), filter)
}
