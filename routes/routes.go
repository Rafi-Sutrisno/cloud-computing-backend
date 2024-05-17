package routes

import (
	"mods/controller"
	"mods/service"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, userController controller.UserController, diseaseController controller.DiseaseController, predictionController controller.PredictionController, jwtService service.JWTService) {
	userPublic := router.Group("/user")
	{
		// public can access
		userPublic.POST("/add", userController.AddUser)
		userPublic.GET("", userController.GetAllUser)
		userPublic.DELETE("/:id", userController.DeleteUser)
		userPublic.POST("/login", userController.UserLoginToken)
	}

	diseasePublic := router.Group("/disease")
	{
		diseasePublic.POST("/add", diseaseController.AddDisease)
		diseasePublic.GET("", diseaseController.GetAllDisease)
		diseasePublic.DELETE("/:id", diseaseController.DeleteDisease)
	}

	predictionPublic := router.Group("/prediction")
	{
		predictionPublic.POST("", predictionController.AddPrediction)
	}

}
