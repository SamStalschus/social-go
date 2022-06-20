package app

import (
	"github.com/gin-gonic/gin"
)

func MapRoutes(r *gin.Engine, controller *Controllers) {

	// Users
	userGroup := r.Group("/v1/user")
	{
		userGroup.POST("", controller.User.CreateUser)
		userGroup.GET("/:id", controller.User.GetUser)
		userGroup.GET("", controller.User.GetUsers)
		userGroup.PUT("", controller.User.UpdateUser)
		userGroup.DELETE("/:id", controller.User.DeleteUser)
	}

	authGroup := r.Group("v1/auth")
	{
		authGroup.POST("/gen-token", controller.Auth.GenToken)
	}

}
