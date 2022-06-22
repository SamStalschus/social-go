package usecase

import (
	"social-go/cmd/api/core/model"
	"social-go/cmd/api/persistence"
	apierrors "social-go/cmd/api/utils/err"

	"golang.org/x/crypto/bcrypt"
)

type CreateUserUseCase interface {
	Execute(user model.User) (*model.User, apierrors.ApiError)
}

type CreateUser struct {
	userDao persistence.UserDao
}

func NewCreateUser(userDao persistence.UserDao) *CreateUser {
	return &CreateUser{
		userDao: userDao,
	}
}

func (createUser *CreateUser) Execute(user model.User) (*model.User, apierrors.ApiError) {
	passwordHash, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, apierrors.NewInternalServerError()
	}

	user.Password = string(passwordHash)

	userCreated, apiErr := createUser.userDao.Save(&user)

	if apiErr != nil {
		return nil, apiErr
	}

	return userCreated, nil
}
