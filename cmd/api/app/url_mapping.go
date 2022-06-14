package app

import "github.com/gin-gonic/gin"

func MapRoutes(r *gin.Engine, controller *Controllers) {

	// Users
	r.POST("/user", controller.User.CreateUser)
	r.GET("/user/{id}", controller.User.GetUser)
	r.GET("/user", controller.User.GetUsers)
	r.PUT("/user", controller.User.UpdateUser)
	r.DELETE("/user/{id}", controller.User.DeleteUser)
}
