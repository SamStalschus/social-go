package usecase

import (
	"social-go/cmd/api/core/model"
	"social-go/cmd/api/persistence"
	apierrors "social-go/cmd/api/utils/err"
)

type CreateUserUseCase interface {
	Execute(user string) (string, apierrors.ApiError)
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

	userCreated, err := createUser.userDao.Save(&user)

	if err != nil {
		return nil, err
	}

	return userCreated, nil
}
