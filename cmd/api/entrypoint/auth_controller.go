package entrypoint

import (
	"net/http"
	"social-go/cmd/api/core/usecase"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	genToken usecase.GenToken
}

func NewAuthController() *AuthController {
	return &AuthController{}
}

func (controller *AuthController) GenToken(c *gin.Context) {
	ctx := c.Request.Context()
	token, err := controller.genToken.Execute(ctx)

	if err != nil {
		c.JSON(int(err.Status()), err)
	}

	c.JSON(http.StatusCreated, token)
}
