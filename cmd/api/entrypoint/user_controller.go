package entrypoint

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"social-go/cmd/api/core/model"
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

func NewUserController(createUser usecase.CreateUser) *UserController {
	return &UserController{
		createUser: createUser,
	}
}

func (controller *UserController) CreateUser(c *gin.Context) {

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		// TODO: Handler api errors
	}

	var user model.User

	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		// TODO: Handler api errors
	}

	response, err := controller.createUser.Execute(user)

	if err != nil {
		// TODO: Handler api errors
		// c.JSON(err.Status())
	}
	c.JSON(http.StatusCreated, response)
}

func (controller *UserController) GetUser(c *gin.Context) {
	controller.getUser.Execute()
}

func (controller *UserController) GetUsers(c *gin.Context) {
	controller.getUsers.Execute()
}

func (controller *UserController) UpdateUser(c *gin.Context) {
	controller.updateUser.Execute()
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	controller.deleteUser.Execute()
}
