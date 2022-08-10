package services

import (
	"errors"
	"go.mongodb.org/mongo-driver/bson"
	"log"
	userDtos "test-gin/users/dtos"
	"test-gin/users/enums"
	userModels "test-gin/users/models"
	"test-gin/users/repositories"
)

type RegisterService interface {
	RegisterUser(input *userDtos.RegisterInput) (enums.RegistrationStatus, error)
}

type registerService struct {
	userRepository repositories.UserRepository
}

func NewRegisterService(userRepository repositories.UserRepository) RegisterService {
	return &registerService{
		userRepository: userRepository,
	}
}

func (registerService *registerService) RegisterUser(input *userDtos.RegisterInput) (enums.RegistrationStatus, error) {
	exists, err := registerService.handleUserAlreadyExists(input.Username)
	if err != nil {
		return enums.Failure, err
	}
	if exists {
		return enums.UserAlreadyExists, nil
	}

	newUser := userModels.User{Username: input.Username, Password: input.Password}

	if _, err := registerService.userRepository.InsertOne(newUser); err != nil {
		log.Printf("Error occurred while writing to db: %s\n", err.Error())
		return enums.Failure, errors.New("error connecting to database")
	}

	return enums.Success, nil
}

func (registerService *registerService) handleUserAlreadyExists(username string) (bool, error) {
	user, err := registerService.userRepository.FindOne(bson.D{{"username", username}})
	return user.Username == username, err
}
