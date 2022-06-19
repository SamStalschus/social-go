package app

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"github.com/shaj13/go-guardian/auth"
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

func ValidateUser(ctx context.Context, r *http.Request, userName, password string) (auth.Info, error) {
	if userName == "medium" && password == "medium" {
		return auth.NewDefaultUser("medium", "1", nil, nil), nil
	}
	return nil, fmt.Errorf("Invalid credentials")
}
func VerifyToken(ctx context.Context, r *http.Request, tokenString string) (auth.Info, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte("secret"), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user := auth.NewDefaultUser(claims["medium"].(string), "", nil, nil)
		return user, nil
	}
	return nil, fmt.Errorf("Invaled token")
}
