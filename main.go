package main

import (
	"context"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
	"test-gin/database"
	"test-gin/users/controllers"
	userModels "test-gin/users/models"
	"test-gin/users/repositories"
	"test-gin/users/services"
	"time"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	mongoClient := database.ConnectToDatabase(ctx)
	defer func(client *mongo.Client, ctx context.Context) {
		if err := client.Disconnect(ctx); err != nil {
			log.Printf("Error disconnected from database %s\n", err.Error())
			return
		}
		log.Println("Disconnected from database")
	}(mongoClient, ctx)

	database.InitializeModels(database.Connection)

	userRepository := repositories.NewUserRepository(database.Connection.Collection(userModels.CollectionName))
	registerService := services.NewRegisterService(userRepository)
	registerController := userControllers.NewRegisterController(registerService)

	router := gin.Default()

	publicRouterGroup := router.Group("/api")
	publicRouterGroup.POST("/register", registerController.PostRegister)

	if err := router.Run(":8080"); err != nil {
		log.Fatalf("An error has occurred: %s\n", err.Error())
		return
	}
}
