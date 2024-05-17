package routes

import (
	"mods/controller"
	"mods/service"

	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine, carController controller.UserController, jwtService service.JWTService) {
	userPublic := router.Group("/user")
	{
		// public can access
		userPublic.POST("/add", carController.AddUser)
		userPublic.GET("", carController.GetAllUser)
		userPublic.DELETE("/:id", carController.DeleteUser)
	}

}
