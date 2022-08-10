package services

import (
	"errors"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"test-gin/mocks"
	userDtos "test-gin/users/dtos"
	"test-gin/users/enums"
	userModels "test-gin/users/models"
	"testing"
)

type UnitTestSuite struct {
	suite.Suite
	registerService    RegisterService
	userRepositoryMock *mocks.UserRepository
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (unitTestSuite *UnitTestSuite) SetupTest() {
	userRepositoryMock := mocks.UserRepository{}
	registerService := NewRegisterService(&userRepositoryMock)

	unitTestSuite.registerService = registerService
	unitTestSuite.userRepositoryMock = &userRepositoryMock
}

func (unitTestSuite *UnitTestSuite) TestRegisterUser_ValidUser() {
	unitTestSuite.userRepositoryMock.On("FindOne", mock.Anything).Return(userModels.User{
		ID:       primitive.ObjectID{},
		Username: "",
		Password: "",
	}, nil)

	unitTestSuite.userRepositoryMock.On("InsertOne", mock.Anything).Return(&mongo.InsertOneResult{InsertedID: 1}, nil)

	input := userDtos.RegisterInput{
		Username: "test@mail.com",
		Password: "password",
	}

	registrationStatus, err := unitTestSuite.registerService.RegisterUser(&input)

	unitTestSuite.Assert().Nil(err)
	unitTestSuite.Assert().Equal(enums.Success, registrationStatus)
}

func (unitTestSuite *UnitTestSuite) TestRegisterUser_UserExists() {
	const username = "test@mail.com"
	const password = "password"
	unitTestSuite.userRepositoryMock.On("FindOne", mock.Anything).Return(userModels.User{
		ID:       primitive.ObjectID{},
		Username: username,
		Password: password,
	}, nil)

	unitTestSuite.userRepositoryMock.On("InsertOne", mock.Anything).Return(&mongo.InsertOneResult{InsertedID: 0}, nil)

	input := userDtos.RegisterInput{
		Username: username,
		Password: password,
	}

	registrationStatus, err := unitTestSuite.registerService.RegisterUser(&input)

	unitTestSuite.Assert().Nil(err)
	unitTestSuite.Assert().Equal(enums.UserAlreadyExists, registrationStatus)
}

func (unitTestSuite *UnitTestSuite) TestRegisterUser_FindOneError() {
	const username = "test@mail.com"
	const password = "password"
	findOneError := errors.New("test error")
	unitTestSuite.userRepositoryMock.On("FindOne", mock.Anything).Return(userModels.User{
		ID:       primitive.ObjectID{},
		Username: username,
		Password: password,
	}, findOneError)

	unitTestSuite.userRepositoryMock.On("InsertOne", mock.Anything).Return(&mongo.InsertOneResult{InsertedID: 0}, nil)

	input := userDtos.RegisterInput{
		Username: username,
		Password: password,
	}

	registrationStatus, err := unitTestSuite.registerService.RegisterUser(&input)

	unitTestSuite.Assert().Equal(findOneError, err)
	unitTestSuite.Assert().Equal(enums.Failure, registrationStatus)
}

func (unitTestSuite *UnitTestSuite) TestRegisterUser_InsertOneError() {
	unitTestSuite.userRepositoryMock.On("FindOne", mock.Anything).Return(userModels.User{
		ID:       primitive.ObjectID{},
		Username: "test@mail.com",
		Password: "password",
	}, nil)

	insertOneError := errors.New("test error")
	unitTestSuite.userRepositoryMock.On("InsertOne", mock.Anything).Return(&mongo.InsertOneResult{InsertedID: 0}, insertOneError)

	input := userDtos.RegisterInput{
		Username: "another_test@mail.com",
		Password: "password",
	}

	registrationStatus, err := unitTestSuite.registerService.RegisterUser(&input)

	unitTestSuite.Assert().Equal(errors.New("error connecting to database"), err)
	unitTestSuite.Assert().Equal(enums.Failure, registrationStatus)
}
