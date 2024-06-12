package main

import (
	"log"
	"mods/config"
	"mods/controller"
	"mods/middleware"
	"mods/repository"
	"mods/routes"
	"mods/service"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println(err)
	}

	db := config.SetupDatabaseConnection()

	jwtService := service.NewJWTService()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controller.NewUserController(userService, jwtService)

	diseaseRepository := repository.NewDiseaseRepository(db)
	diseaseService := service.NewDiseaseService(diseaseRepository)
	diseaseController := controller.NewDiseaseController(diseaseService, jwtService)

	predictionRepository := repository.NewPredictionRepository(db)
	predictionService := service.NewPredictionService(predictionRepository, diseaseRepository)
	predictionController := controller.NewPredictionController(predictionService, jwtService)

	chatroomRepository := repository.NewChatroomRepository(db)
	chatroomService := service.NewChatRoomService(chatroomRepository)
	chatroomController := controller.NewChatroomController(chatroomService, jwtService)

	defer config.CloseDatabaseConnection(db)

	server := gin.Default()
	server.Use(middleware.CORSMiddleware())

	routes.Routes(server, userController, diseaseController, predictionController, chatroomController, jwtService)

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	server.Run(":" + port)
	// testing test
}
