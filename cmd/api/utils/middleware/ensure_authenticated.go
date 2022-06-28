package middleware

import (
	"fmt"
	"net/http"
	"os"
	apierrors "social-go/cmd/api/utils/err"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

type IEnsureAuthenticated interface {
	Execute(c *gin.Context)
}

type EnsureAuthenticated struct {
}

func NewEnsureAuthenticated() *EnsureAuthenticated {
	return &EnsureAuthenticated{}
}

func (ensureAuthenticated *EnsureAuthenticated) Execute(c *gin.Context) {
	var token string

	if len(c.Request.Header["Authorization"]) > 0 {
		token = strings.Split(c.Request.Header["Authorization"][0], " ")[1]
	}

	if token == "" {
		c.JSON(http.StatusUnauthorized, apierrors.NewApiError("Token missing", http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized))
		return
	}

	tkn, err := jwt.ParseWithClaims(token, jwt.MapClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})

	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			c.JSON(http.StatusUnauthorized, apierrors.NewApiError("Unauthorized", http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized))
			return
		}
		c.JSON(http.StatusUnauthorized, apierrors.NewApiError("Bad Request", http.StatusText(http.StatusBadRequest), http.StatusBadRequest))
		return
	}
	if tkn.Valid && err == nil {
		fmt.Println("Chamando next")
		c.Next()
	}

	c.JSON(http.StatusUnauthorized, apierrors.NewApiError("Unauthorized", http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized))
}
