package userControllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
	"io"
	"net/http"
	"net/http/httptest"
	"test-gin/mocks"
	userDtos "test-gin/users/dtos"
	"test-gin/users/enums"
	"testing"
)

type UnitTestSuite struct {
	suite.Suite
	registerController  RegisterController
	registerServiceMock *mocks.RegisterService
}

func TestUnitTestSuite(t *testing.T) {
	suite.Run(t, &UnitTestSuite{})
}

func (unitTestSuite *UnitTestSuite) SetupTest() {
	registerServiceMock := mocks.RegisterService{}
	registerController := NewRegisterController(&registerServiceMock)

	unitTestSuite.registerController = registerController
	unitTestSuite.registerServiceMock = &registerServiceMock
}

func (unitTestSuite *UnitTestSuite) TestPostRegister_EmptyBody() {
	unitTestSuite.registerServiceMock.On("RegisterUser", mock.Anything).Return(enums.Success, nil)

	gin.SetMode(gin.TestMode)
	httpTestRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpTestRecorder)

	unitTestSuite.registerController.PostRegister(ginContext)

	unitTestSuite.Assert().Equal(http.StatusBadRequest, httpTestRecorder.Code)
	unitTestSuite.registerServiceMock.AssertNotCalled(unitTestSuite.T(), "RegisterUser", mock.Anything)
}

func (unitTestSuite *UnitTestSuite) TestPostRegister_UserExists() {
	unitTestSuite.registerServiceMock.On("RegisterUser", mock.Anything).Return(enums.UserAlreadyExists, nil)

	gin.SetMode(gin.TestMode)
	httpTestRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpTestRecorder)

	ginContext.Request = &http.Request{}

	user := userDtos.RegisterInput{
		Username: "username",
		Password: "password",
	}
	jsonValue, _ := json.Marshal(user)
	ginContext.Request.Body = io.NopCloser(bytes.NewReader(jsonValue))

	unitTestSuite.registerController.PostRegister(ginContext)

	unitTestSuite.Assert().Equal(http.StatusConflict, httpTestRecorder.Code)
	unitTestSuite.registerServiceMock.AssertExpectations(unitTestSuite.T())
}

func (unitTestSuite *UnitTestSuite) TestPostRegister_ValidBody_RegisterUserSuccess() {
	unitTestSuite.registerServiceMock.On("RegisterUser", mock.Anything).Return(enums.Success, nil)

	gin.SetMode(gin.TestMode)
	httpTestRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpTestRecorder)

	ginContext.Request = &http.Request{}

	user := userDtos.RegisterInput{
		Username: "username",
		Password: "password",
	}
	jsonValue, _ := json.Marshal(user)
	ginContext.Request.Body = io.NopCloser(bytes.NewReader(jsonValue))

	unitTestSuite.registerController.PostRegister(ginContext)

	unitTestSuite.Assert().Equal(http.StatusCreated, httpTestRecorder.Code)
	unitTestSuite.registerServiceMock.AssertExpectations(unitTestSuite.T())
}

func (unitTestSuite *UnitTestSuite) TestPostRegister_ValidBody_RegisterUserFailureError() {
	unitTestSuite.registerServiceMock.On("RegisterUser", mock.Anything).Return(enums.Failure, errors.New("error"))

	gin.SetMode(gin.TestMode)
	httpTestRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpTestRecorder)

	ginContext.Request = &http.Request{}

	user := userDtos.RegisterInput{
		Username: "username",
		Password: "password",
	}
	jsonValue, _ := json.Marshal(user)
	ginContext.Request.Body = io.NopCloser(bytes.NewReader(jsonValue))

	unitTestSuite.registerController.PostRegister(ginContext)

	unitTestSuite.Assert().Equal(http.StatusInternalServerError, httpTestRecorder.Code)
	unitTestSuite.registerServiceMock.AssertExpectations(unitTestSuite.T())
}

func (unitTestSuite *UnitTestSuite) TestPostRegister_ValidBody_RegisterUserFailureNoError() {
	unitTestSuite.registerServiceMock.On("RegisterUser", mock.Anything).Return(enums.Failure, nil)

	gin.SetMode(gin.TestMode)
	httpTestRecorder := httptest.NewRecorder()
	ginContext, _ := gin.CreateTestContext(httpTestRecorder)

	ginContext.Request = &http.Request{}

	user := userDtos.RegisterInput{
		Username: "username",
		Password: "password",
	}
	jsonValue, _ := json.Marshal(user)
	ginContext.Request.Body = io.NopCloser(bytes.NewReader(jsonValue))

	unitTestSuite.registerController.PostRegister(ginContext)

	unitTestSuite.Assert().Equal(http.StatusInternalServerError, httpTestRecorder.Code)
	unitTestSuite.registerServiceMock.AssertExpectations(unitTestSuite.T())
}
