package usecase

import (
	"social-go/cmd/api/core/model"
	"social-go/cmd/api/persistence"
	apierrors "social-go/cmd/api/utils/err"
)

type GetUserUseCase interface {
	Execute(id int) (*model.User, apierrors.ApiError)
}

type GetUser struct {
	userDao persistence.UserDao
}

func NewGetUser(userDao persistence.UserDao) *GetUser {
	return &GetUser{
		userDao: userDao,
	}
}

func (getUser *GetUser) Execute(id int) (*model.User, apierrors.ApiError) {

	user, err := getUser.userDao.Get(id)

	if err != nil {
		return nil, err
	}

	return user, nil
}
