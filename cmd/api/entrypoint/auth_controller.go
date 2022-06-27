package entrypoint

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"social-go/cmd/api/core/model"
	"social-go/cmd/api/core/usecase"
	apierrors "social-go/cmd/api/utils/err"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	genToken usecase.GenTokenUseCase
}

func NewAuthController(genToken usecase.GenTokenUseCase) *AuthController {
	return &AuthController{
		genToken: genToken,
	}
}

func (controller *AuthController) GenToken(c *gin.Context) {
	jsonData, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			apierrors.NewBadRequestApiError("Bad Request - Failed read to request body"))
		return
	}

	var authUser model.AuthUser

	err = json.Unmarshal(jsonData, &authUser)
	if err != nil {
		c.JSON(http.StatusBadRequest,
			apierrors.NewBadRequestApiError("Bad Request - Failed to unmarshal request body"))
		return
	}
	ctx := c.Request.Context()
	token, apiErr := controller.genToken.Execute(ctx, authUser)

	if apiErr != nil {
		c.JSON(int(apiErr.Status()), apiErr)
	}

	c.JSON(http.StatusCreated, token)
}
