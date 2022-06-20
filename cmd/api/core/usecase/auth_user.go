package usecase

import (
	"context"
	"os"
	"social-go/cmd/api/core/model"
	"social-go/cmd/api/persistence"
	apierrors "social-go/cmd/api/utils/err"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
)

type GenTokenUseCase interface {
	Execute(ctx context.Context, authUser model.AuthUser) (string, apierrors.ApiError)
}

type GenToken struct {
	userDao persistence.UserDao
}

func NewGenTokenUseCase(userDao persistence.UserDao) *GenToken {
	return &GenToken{
		userDao: userDao,
	}
}

func (genToken *GenToken) Execute(ctx context.Context, authUser model.AuthUser) (string, apierrors.ApiError) {

	if authUser.Username == "" || authUser.Password == "" {
		return "", apierrors.NewBadRequestApiError("Username or password incorrect.")
	}

	user, err := genToken.userDao.GetByUsernameWithPassword(authUser.Username)

	if err != nil {
		return "", err
	}

	if !passwordIsCorrect(authUser.Password, user.Password) {
		return "", apierrors.NewBadRequestApiError("Username or password incorrect.")
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"iss": "social-go",
		"sub": user,
		"aud": "any",
		"exp": time.Now().Add(time.Minute * 5).Unix(),
	})

	jwtToken, errJwt := token.SignedString([]byte(os.Getenv("SECRET_KEY")))

	if errJwt != nil {
		return "", apierrors.NewInternalServerError()
	}

	return jwtToken, nil
}

func passwordIsCorrect(passwordProvided, password string) bool {
	if err := bcrypt.CompareHashAndPassword([]byte(password), []byte(passwordProvided)); err != nil {
		return false
	}

	return true
}
