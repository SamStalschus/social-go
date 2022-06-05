package entrypoint

import (
	"social-go/cmd/api/core/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUser usecase.CreateUser
	getUser    usecase.GetUser
	getUsers   usecase.GetUsers
	updateUser usecase.UpdateUser
	deleteUser usecase.DeleteUser
}

func NewUserController() *UserController {
	return &UserController{}
}

func (controller *UserController) CreateUser(c *gin.Context) {
	controller.createUser.Execute("Samuca")
}

func (controller *UserController) GetUser(c *gin.Context) {
	controller.getUser.Execute()
}

func (controller *UserController) GetUsers(c *gin.Context) {
	controller.getUser.Execute()
}

func (controller *UserController) UpdateUser(c *gin.Context) {
	controller.getUser.Execute()
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	controller.getUser.Execute()
}
