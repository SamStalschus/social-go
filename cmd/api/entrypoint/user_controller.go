package entrypoint

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"social-go/cmd/api/core/model"
	"social-go/cmd/api/core/usecase"
	apierrors "social-go/cmd/api/utils/err"
	"strconv"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	createUser usecase.CreateUserUseCase
	getUser    usecase.GetUserUseCase
	getUsers   usecase.GetUsersUseCase
	updateUser usecase.UpdateUserUseCase
	deleteUser usecase.DeleteUserUseCase
}

func NewUserController(
	createUser usecase.CreateUserUseCase,
	getUser usecase.GetUserUseCase,
	getUsers usecase.GetUsersUseCase,
	updateUser usecase.UpdateUserUseCase,
	deleteUser usecase.DeleteUserUseCase) *UserController {

	return &UserController{
		createUser: createUser,
		getUser:    getUser,
		updateUser: updateUser,
		getUsers:   getUsers,
		deleteUser: deleteUser,
	}
}

func (controller *UserController) CreateUser(c *gin.Context) {
	ctx := c.Request.Context()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			apierrors.NewBadRequestApiError("Bad Request - Failed read to request body"))
		return
	}

	var user model.User

	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			apierrors.NewBadRequestApiError("Bad Request - Failed to unmarshal request body"))
		return
	}

	response, apiError := controller.createUser.Execute(ctx, user)

	if apiError != nil {
		c.JSON(int(apiError.Status()),
			apiError)
		return
	}
	c.JSON(http.StatusCreated, response)
}

func (controller *UserController) GetUser(c *gin.Context) {
	ctx := c.Request.Context()

	receivedId := c.Param("id")

	id, err := strconv.Atoi(receivedId)
	if err != nil {
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequestApiError("Error to get id of user"))
	}

	user, apiError := controller.getUser.Execute(ctx, id)

	if apiError != nil {
		c.JSON(int(apiError.Status()),
			apiError)
		return
	}
	c.JSON(http.StatusOK, user)
}

func (controller *UserController) UpdateUser(c *gin.Context) {
	ctx := c.Request.Context()

	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			apierrors.NewBadRequestApiError("Bad Request - Failed read to request body"))
		return
	}

	var user model.User

	err = json.Unmarshal(jsonData, &user)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			apierrors.NewBadRequestApiError("Bad Request - Failed to unmarshal request body"))
		return
	}

	response, apiError := controller.updateUser.Execute(ctx, &user)

	if apiError != nil {
		c.JSON(int(apiError.Status()),
			apiError)
		return
	}
	c.JSON(http.StatusNoContent, response)
}

func (controller *UserController) GetUsers(c *gin.Context) {
	ctx := c.Request.Context()

	response, err := controller.getUsers.Execute(ctx)

	if err != nil {
		c.JSON(int(err.Status()),
			err)
		return
	}
	c.JSON(http.StatusOK, response)
}

func (controller *UserController) DeleteUser(c *gin.Context) {
	ctx := c.Request.Context()

	receivedId := c.Param("id")

	id, err := strconv.Atoi(receivedId)
	if err != nil {
		c.JSON(http.StatusBadRequest, apierrors.NewBadRequestApiError("Error to get id of user"))
	}

	apiError := controller.deleteUser.Execute(ctx, id)

	if apiError != nil {
		c.JSON(int(apiError.Status()),
			apiError)
		return
	}
	c.Status(http.StatusNoContent)
}
