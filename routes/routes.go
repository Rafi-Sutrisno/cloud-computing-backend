package routes

import (
	"mods/controller"
	"mods/middleware"
	"mods/service"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, userController controller.UserController, diseaseController controller.DiseaseController, predictionController controller.PredictionController, jwtService service.JWTService) {
	inscurePublic := router.Group("/inscure")
	{
		// public can access
		inscurePublic.POST("/add", userController.AddUser)
		inscurePublic.POST("/login", userController.UserLoginToken)
	}

	userPrivate := router.Group("/user").Use(middleware.Authenticate())
	{
		userPrivate.GET("/me", userController.Me)
		userPrivate.GET("", userController.GetAllUser)
		userPrivate.GET("/doctor", userController.GetAllDoctor)
		userPrivate.DELETE("/:id", userController.DeleteUser)
		userPrivate.PUT("/update", userController.UpdateUser)
	}

	diseasePublic := router.Group("/disease")
	{
		diseasePublic.POST("/add", diseaseController.AddDisease)
		diseasePublic.GET("", diseaseController.GetAllDisease)
		diseasePublic.DELETE("/:id", diseaseController.DeleteDisease)
		diseasePublic.GET("/:id", diseaseController.DiseaseByID)
	}

	predictionPublic := router.Group("/prediction").Use(middleware.Authenticate())
	{
		predictionPublic.POST("", predictionController.AddPrediction)
		predictionPublic.GET("/list", predictionController.GetPredictionByUserID)
		predictionPublic.GET("/:p_id", predictionController.GetPredictionByPredictionID)
	}

}
