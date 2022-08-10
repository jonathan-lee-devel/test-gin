package userControllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	userDtos "test-gin/users/dtos"
	"test-gin/users/enums"
	"test-gin/users/services"
)

type RegisterController interface {
	PostRegister(ginContext *gin.Context)
}

type registerController struct {
	registerService services.RegisterService
}

func NewRegisterController(registerService services.RegisterService) RegisterController {
	return &registerController{
		registerService: registerService,
	}
}

func (registerController *registerController) PostRegister(ginContext *gin.Context) {
	var input userDtos.RegisterInput
	if err := ginContext.ShouldBindJSON(&input); err != nil {
		ginContext.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	registrationStatus, err := registerController.registerService.RegisterUser(&input)
	if err != nil {
		ginContext.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	switch registrationStatus {
	case enums.UserAlreadyExists:
		ginContext.JSON(http.StatusConflict, gin.H{"message": "User already exists"})
		return
	case enums.Failure:
		ginContext.JSON(http.StatusInternalServerError, gin.H{"message": "Failed to register user"})
		return
	case enums.Success:
		ginContext.JSON(http.StatusCreated, gin.H{"message": "Successfully validated"})
		return
	}
}
