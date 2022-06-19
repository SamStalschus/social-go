package usecase

import (
	"context"
	"os"
	apierrors "social-go/cmd/api/utils/err"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type GenTokenUseCase interface {
	Execute(ctx context.Context) (string, apierrors.ApiError)
}

type GenToken struct {
}

func NewGenTokenUseCase() *GenToken {
	return &GenToken{}
}

func (genToken *GenToken) Execute(ctx context.Context) (string, apierrors.ApiError) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "social-go",
		"sub": "user_id",
		"aud": "any",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	jwtToken, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if err != nil {
		return "", apierrors.NewInternalServerError()
	}

	return jwtToken, nil
}
